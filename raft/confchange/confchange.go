// Copyright 2019 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package confchange

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ls-2018/etcd_cn/raft/quorum"
	pb "github.com/ls-2018/etcd_cn/raft/raftpb"
	"github.com/ls-2018/etcd_cn/raft/tracker"
)

// Changer facilitates configuration changes. It exposes methods to handle
// simple and joint consensus while performing the proper validation that allows
// refusing invalid configuration changes before they affect the active
// configuration.
type Changer struct {
	Tracker   tracker.ProgressTracker
	LastIndex uint64
}

// EnterJoint verifies that the outgoing (=right) majority config of the joint
// config is empty and initializes it with a copy of the incoming (=left)
// majority config. That is, it transitions from
//
//     (1 2 3)&&()
// to
//     (1 2 3)&&(1 2 3).
//
// The supplied changes are then applied to the incoming majority config,
// resulting in a joint configuration that in terms of the Raft thesis[1]
// (Section 4.3) corresponds to `C_{new,old}`.
//
// [1]: https://github.com/ongardie/dissertation/blob/master/online-trim.pdf
func (c Changer) EnterJoint(autoLeave bool, ccs ...pb.ConfChangeSingle) (tracker.Config, tracker.ProgressMap, error) {
	cfg, prs, err := c.checkAndCopy()
	if err != nil {
		return c.err(err)
	}
	if joint(cfg) {
		err := errors.New("config is already joint")
		return c.err(err)
	}
	if len(incoming(cfg.Voters)) == 0 {
		// We allow adding nodes to an empty config for convenience (testing and
		// bootstrap), but you can't enter a joint state.
		err := errors.New("can't make a zero-voter config joint")
		return c.err(err)
	}
	// Clear the outgoing config.
	*outgoingPtr(&cfg.Voters) = quorum.MajorityConfig{}
	// Copy incoming to outgoing.
	for id := range incoming(cfg.Voters) {
		outgoing(cfg.Voters)[id] = struct{}{}
	}

	if err := c.apply(&cfg, prs, ccs...); err != nil {
		return c.err(err)
	}
	cfg.AutoLeave = autoLeave
	return checkAndReturn(cfg, prs)
}

// LeaveJoint transitions out of a joint configuration. It is an error to call
// this method if the configuration is not joint, i.e. if the outgoing majority
// config Voters[1] is empty.
//
// The outgoing majority config of the joint configuration will be removed,
// that is, the incoming config is promoted as the sole decision maker. In the
// notation of the Raft thesis[1] (Section 4.3), this method transitions from
// `C_{new,old}` into `C_new`.
//
// At the same time, any staged learners (LearnersNext) the addition of which
// was held back by an overlapping voter in the former outgoing config will be
// inserted into Learners.
//
// [1]: https://github.com/ongardie/dissertation/blob/master/online-trim.pdf
func (c Changer) LeaveJoint() (tracker.Config, tracker.ProgressMap, error) {
	cfg, prs, err := c.checkAndCopy()
	if err != nil {
		return c.err(err)
	}
	if !joint(cfg) {
		err := errors.New("can't leave a non-joint config")
		return c.err(err)
	}
	if len(outgoing(cfg.Voters)) == 0 {
		err := fmt.Errorf("configuration is not joint: %v", cfg)
		return c.err(err)
	}
	for id := range cfg.LearnersNext {
		nilAwareAdd(&cfg.Learners, id)
		prs[id].IsLearner = true
	}
	cfg.LearnersNext = nil

	for id := range outgoing(cfg.Voters) {
		_, isVoter := incoming(cfg.Voters)[id]
		_, isLearner := cfg.Learners[id]

		if !isVoter && !isLearner {
			delete(prs, id)
		}
	}
	*outgoingPtr(&cfg.Voters) = nil
	cfg.AutoLeave = false

	return checkAndReturn(cfg, prs)
}

// ------------------------------------  OVER  --------------------------------------------------

// ???????????????????????????
func joint(cfg tracker.Config) bool {
	return len(outgoing(cfg.Voters)) > 0
}

// ????????????????????????,?????????????????????JointConfig[0]???incoming?????????????????????.
// ???EnterJoint???,?????????????????????outgoing???,?????????????????????incoming???.
// LeaveJoint???,?????????????????????,?????????????????????????????????incoming???,????????????????????????.
func incoming(voters quorum.JointConfig) quorum.MajorityConfig      { return voters[0] }
func outgoing(voters quorum.JointConfig) quorum.MajorityConfig      { return voters[1] }
func outgoingPtr(voters *quorum.JointConfig) *quorum.MajorityConfig { return &voters[1] }

// Describe ??????????????????
func Describe(ccs ...pb.ConfChangeSingle) string {
	var buf strings.Builder
	for _, cc := range ccs {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%s(%d)", cc.Type, cc.NodeID)
	}
	return buf.String()
}

// checkInvariants ???????????????????????????????????????.
func checkInvariants(cfg tracker.Config, prs tracker.ProgressMap) error {
	// cfg????????????,prs????????????'
	for _, ids := range []map[uint64]struct{}{
		cfg.Voters.IDs(), // JointConfig??????????????????ID
		cfg.Learners,     //
		cfg.LearnersNext, // ????????????learner?????????
	} {
		for id := range ids {
			if _, ok := prs[id]; !ok {
				return fmt.Errorf("?????????????????????????????? %d", id)
			}
		}
	}

	// LearnersNext ??????????????????,  ?????????leader???follower ?????? ; ??????????????????.
	for id := range cfg.LearnersNext {
		// ????????????????????????????????????
		if _, ok := outgoing(cfg.Voters)[id]; !ok {
			return fmt.Errorf("%d ??????LearnersNext???,????????????Voters???.[1]", id)
		}
		// ??????????????????????????????learner
		if prs[id].IsLearner {
			return fmt.Errorf("%d ??????LearnersNext???,?????????????????????learner.", id)
		}
	}
	// ??????,learner??????????????????????????????.
	for id := range cfg.Learners {
		if _, ok := outgoing(cfg.Voters)[id]; ok {
			return fmt.Errorf("%d ??? Learners ??? Voters[1]", id)
		}
		if _, ok := incoming(cfg.Voters)[id]; ok {
			return fmt.Errorf("%d ??? Learners ??? Voters[0]", id)
		}
		if !prs[id].IsLearner {
			return fmt.Errorf("%d ??? Learners, ?????????????????????learner", id)
		}
	}

	if !joint(cfg) { // ????????????????????????
		// ??????????????????,???map???nil?????????0.
		if outgoing(cfg.Voters) != nil {
			return fmt.Errorf("cfg.Voters[1]?????????nil ???????????????????????????")
		}
		if cfg.LearnersNext != nil {
			return fmt.Errorf("cfg.LearnersNext?????????nil ???????????????????????????")
		}
		if cfg.AutoLeave {
			return fmt.Errorf("AutoLeave?????????false ???????????????????????????")
		}
	}
	return nil
}

// checkAndCopy ?????????????????????????????????Map,?????????????????????.
func (c Changer) checkAndCopy() (tracker.Config, tracker.ProgressMap, error) {
	cfg := c.Tracker.Config.Clone() // ?????????????????????
	var _ tracker.ProgressMap = c.Tracker.Progress
	prs := tracker.ProgressMap{} // ??????

	for id, pr := range c.Tracker.Progress {
		// pr???????????????
		// ??????????????????????????????,??????????????????Learner??????????????????.
		ppr := *pr
		prs[id] = &ppr
	}
	return checkAndReturn(cfg, prs) // cfg????????????,prs????????????;???????????????????????????????????????
}

// checkAndReturn ??????????????????checkInvariants,??????????????????????????????.
func checkAndReturn(cfg tracker.Config, prs tracker.ProgressMap) (tracker.Config, tracker.ProgressMap, error) {
	// cfg????????????,prs????????????
	if err := checkInvariants(cfg, prs); err != nil { // ???????????????????????????????????????
		return tracker.Config{}, tracker.ProgressMap{}, err
	}
	return cfg, prs, nil
}

// initProgress ???????????????Follower???Learner???Progress,??????????????????????????????????????????,????????????.
// ID???Peer???ID,match???next???????????????Progress???.
func (c Changer) initProgress(cfg *tracker.Config, prs tracker.ProgressMap, id uint64, isLearner bool) {
	if !isLearner {
		// ?????????  cfg.Voters[0][id] = struct{}{}
		incoming(cfg.Voters)[id] = struct{}{}
	} else {
		nilAwareAdd(&cfg.Learners, id)
	}
	// Follower???????????????????????????,Learner?????????,?????????????????????????????????.?????????Follower??????
	// Learner???????????????Progress,???????????????????????????????????????.
	prs[id] = &tracker.Progress{
		// ?????????Progress????????????Next???Match???Inflights?????????????????????learner,???????????????
		// ????????????????????????,raft???????????????????????????Match=0,Next=1.
		Next:      c.LastIndex,
		Match:     0,
		Inflights: tracker.NewInflights(c.Tracker.MaxInflight),
		IsLearner: isLearner,
		// ????????????????????????????????????,???????????????????????????????????????.??????,??????CheckQuorum??????????????????,?????????????????????????????????.
		// ????????????????????????????????????????????????????????????,???????????????????????????.
		RecentActive: true,
	}
}

// nilAwareAdd ????????????map??????,??????????????????,??????map
func nilAwareAdd(m *map[uint64]struct{}, id uint64) {
	if *m == nil {
		*m = map[uint64]struct{}{}
	}
	(*m)[id] = struct{}{}
}

// nilAwareDelete ?????????map?????????,????????????map?????????,???????????????.
func nilAwareDelete(m *map[uint64]struct{}, id uint64) {
	if *m == nil {
		return
	}
	delete(*m, id)
	if len(*m) == 0 {
		*m = nil
	}
}

// err ???????????????????????????.
func (c Changer) err(err error) (tracker.Config, tracker.ProgressMap, error) {
	return tracker.Config{}, nil, err
}

// makeVoter ????????????????????????ID,????????????????????????????????????????????????.
func (c Changer) makeVoter(cfg *tracker.Config, prs tracker.ProgressMap, id uint64) {
	// cfg????????????,prs????????????;?????????????????????,???????????????????????????????????????
	pr := prs[id]
	if pr == nil {
		// ????????????
		c.initProgress(cfg, prs, id, false /* isLearner */)
		return
	}
	// ???????????? [learner->follower]
	pr.IsLearner = false
	nilAwareDelete(&cfg.Learners, id)
	nilAwareDelete(&cfg.LearnersNext, id)
	incoming(cfg.Voters)[id] = struct{}{}
}

// remove ??? voter[0]???learner???learnersNext ?????????
func (c Changer) remove(cfg *tracker.Config, prs tracker.ProgressMap, id uint64) {
	if _, ok := prs[id]; !ok {
		return
	}

	delete(incoming(cfg.Voters), id)
	nilAwareDelete(&cfg.Learners, id)
	nilAwareDelete(&cfg.LearnersNext, id)

	// ????????????vote[1]??? ??????;?????????vote[1]???,?????????leaveJoint??????
	if _, onRight := outgoing(cfg.Voters)[id]; !onRight {
		// cfg.Voters[1][id]
		delete(prs, id)
	}
}

// makeLearner ??????learner
func (c Changer) makeLearner(cfg *tracker.Config, prs tracker.ProgressMap, id uint64) {
	pr := prs[id]
	if pr == nil {
		// ???????????????learner
		c.initProgress(cfg, prs, id, true /* isLearner */)
		return
	}
	if pr.IsLearner {
		return
	}
	// // ??? voters???learner???learnersNext ?????????
	c.remove(cfg, prs, id) // ??? voter[0]???learner???learnersNext ?????????
	prs[id] = pr
	//  ???????????????????????????learner?????????learner???,????????????,???peer???veto[1]???,???????????????
	//  ????????? LearnersNext.
	//  ???LeaveJoint()???,LearnersNext?????????????????????learner.
	//  ??????,???????????????????????????learner.
	if _, onRight := outgoing(cfg.Voters)[id]; onRight {
		// ??????
		nilAwareAdd(&cfg.LearnersNext, id)
	} else {
		// ??????
		pr.IsLearner = true
		nilAwareAdd(&cfg.Learners, id)
	}
}

// apply ????????????????????????.????????????,???voter?????????????????? Voters[0] ??????.
//  [??????????????????,???????????????] ??? [?????????nil]
//  Voters[1]???????????????,???????????????????????????????????????????????????.
func (c Changer) apply(cfg *tracker.Config, prs tracker.ProgressMap, ccs ...pb.ConfChangeSingle) error {
	// cfg????????????,prs????????????;?????????????????????,???????????????????????????????????????
	for _, cc := range ccs {
		if cc.NodeID == 0 {
			continue
		}
		// ????????? ?????? id?????????;   peer url ?????????
		// ??????cfg prs ??????
		switch cc.Type {
		case pb.ConfChangeAddNode:
			c.makeVoter(cfg, prs, cc.NodeID)
		case pb.ConfChangeAddLearnerNode:
			c.makeLearner(cfg, prs, cc.NodeID)
		case pb.ConfChangeRemoveNode:
			c.remove(cfg, prs, cc.NodeID)
		case pb.ConfChangeUpdateNode:
		default:
			return fmt.Errorf("?????????conf type %d", cc.Type)
		}
	}
	if len(incoming(cfg.Voters)) == 0 {
		return errors.New("?????????????????????")
	}
	return nil
}

// symdiff ????????????????????????  len(a-b)+len(b-a):?????????????????????
func symdiff(l, r map[uint64]struct{}) int {
	var n int
	// [[1,2],[1,2]]
	pairs := [][2]quorum.MajorityConfig{
		{l, r}, // count elems in l but not in r
		{r, l}, // count elems in r but not in l
	}
	for _, p := range pairs {
		for id := range p[0] {
			if _, ok := p[1][id]; !ok {
				n++
			}
		}
	}
	return n
}

// Simple ??????????????????????????????,(????????????)????????????????????????Voters[0]??????????????????.
// ??????????????????,???????????????:quorum??????,????????????????????????????????????(?????????????????????????????????),?????????????????????????????????.
func (c Changer) Simple(ccs ...pb.ConfChangeSingle) (tracker.Config, tracker.ProgressMap, error) {
	cfg, prs, err := c.checkAndCopy() // cfg????????????,prs????????????;?????????????????????,???????????????????????????????????????
	if err != nil {
		return c.err(err)
	}
	if joint(cfg) { // ???????????????????????????
		err := errors.New("???????????????????????????????????????????????????")
		return c.err(err)
	}
	if err := c.apply(&cfg, prs, ccs...); err != nil { // ?????????cfg???prs??????
		return c.err(err)
	}
	// incoming voters[0]
	// prs = c.Tracker.Progress cfg = c.Tracker.Config
	_ = c.Tracker.Voters //  quorum.JointConfig
	// c.Tracker.Config.Voters   ????????????  Config??????????????????  c.Tracker.Voters

	if n := symdiff(incoming(c.Tracker.Voters), incoming(cfg.Voters)); n > 1 {
		// [12,123] ????????? 1
		// ?????????????????????
		return tracker.Config{}, nil, errors.New("???????????????????????????????????????????????????????????????")
	}

	return checkAndReturn(cfg, prs) // ????????????
}
