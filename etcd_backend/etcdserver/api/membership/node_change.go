package membership

import (
	"errors"

	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc/backend"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
	"github.com/ls-2018/etcd_cn/etcd_backend/etcdserver/api/v2error"
	"go.uber.org/zap"
)

func (c *RaftCluster) IsReadyToAddVotingMember() bool {
	nmembers := 1
	nstarted := 0

	for _, member := range c.VotingMembers() {
		if member.IsStarted() {
			nstarted++
		}
		nmembers++
	}

	if nstarted == 1 && nmembers == 2 {
		// a case of adding a new node to 1-member cluster for restoring cluster data
		// https://github.com/etcd-io/website/blob/main/content/docs/v2/admin_guide.md#restoring-the-cluster
		c.lg.Debug("number of started member is 1; can accept add member request")
		return true
	}

	nquorum := nmembers/2 + 1
	if nstarted < nquorum {
		c.lg.Warn("rejecting member add; started member will be less than quorum", zap.Int("number-of-started-member", nstarted), zap.Int("quorum", nquorum), zap.String("cluster-id", c.cid.String()), zap.String("local-member-id", c.localID.String()))
		return false
	}

	return true
}

func (c *RaftCluster) IsReadyToRemoveVotingMember(id uint64) bool {
	nmembers := 0
	nstarted := 0

	for _, member := range c.VotingMembers() {
		if uint64(member.ID) == id {
			continue
		}

		if member.IsStarted() {
			nstarted++
		}
		nmembers++
	}

	nquorum := nmembers/2 + 1
	if nstarted < nquorum {
		c.lg.Warn(
			"rejecting member remove; started member will be less than quorum",
			zap.Int("number-of-started-member", nstarted),
			zap.Int("quorum", nquorum),
			zap.String("cluster-id", c.cid.String()),
			zap.String("local-member-id", c.localID.String()),
		)
		return false
	}

	return true
}

func (c *RaftCluster) IsReadyToPromoteMember(id uint64) bool {
	nmembers := 1 // We count the learner to be promoted for the future quorum
	nstarted := 1 // and we also count it as started.

	for _, member := range c.VotingMembers() {
		if member.IsStarted() {
			nstarted++
		}
		nmembers++
	}

	nquorum := nmembers/2 + 1
	if nstarted < nquorum {
		c.lg.Warn(
			"rejecting member promote; started member will be less than quorum",
			zap.Int("number-of-started-member", nstarted),
			zap.Int("quorum", nquorum),
			zap.String("cluster-id", c.cid.String()),
			zap.String("local-member-id", c.localID.String()),
		)
		return false
	}

	return true
}

// ------------------------------------------------ over ------------------------------------------------

// PromoteMember 将该成员的IsLearner属性标记为false。
func (c *RaftCluster) PromoteMember(id types.ID, shouldApplyV3 ShouldApplyV3) {
	c.Lock()
	defer c.Unlock()

	c.members[id].RaftAttributes.IsLearner = false
	if c.v2store != nil {
		// 内存里面的一个树形node结构
		mustUpdateMemberInStore(c.lg, c.v2store, c.members[id])
	}
	if c.be != nil && shouldApplyV3 {
		unsafeSaveMemberToBackend(c.lg, c.be, c.members[id])
	}

	c.lg.Info("成员角色提升", zap.String("cluster-id", c.cid.String()), zap.String("local-member-id", c.localID.String()))
}

// AddMember 在集群中添加一个新的成员，并将给定成员的raftAttributes保存到存储空间。给定的成员应该有空的属性。 一个具有匹配id的成员必须不存在。
func (c *RaftCluster) AddMember(m *Member, shouldApplyV3 ShouldApplyV3) {
	c.Lock()
	defer c.Unlock()

	var v2Err, beErr error
	if c.v2store != nil {
		v2Err = unsafeSaveMemberToStore(c.lg, c.v2store, m)
		if v2Err != nil {
			if e, ok := v2Err.(*v2error.Error); !ok || e.ErrorCode != v2error.EcodeNodeExist {
				c.lg.Panic("保存member到v2store失败", zap.String("member-id", m.ID.String()), zap.Error(v2Err))
			}
		}
	}
	_ = backend.MyBackend{}
	if c.be != nil && shouldApplyV3 {
		beErr = unsafeSaveMemberToBackend(c.lg, c.be, m) // 保存到bolt.db     members
		if beErr != nil && !errors.Is(beErr, errMemberAlreadyExist) {
			c.lg.Panic("保存member到backend失败", zap.String("member-id", m.ID.String()), zap.Error(beErr))
		}
	}
	if v2Err != nil && (beErr != nil || c.be == nil) {
		c.lg.Panic("保存member到store失败", zap.String("member-id", m.ID.String()), zap.Error(v2Err))
	}

	c.members[m.ID] = m

	c.lg.Info("添加成员", zap.String("cluster-id", c.cid.String()), zap.String("local-member-id", c.localID.String()), zap.String("added-peer-id", m.ID.String()), zap.Strings("added-peer-peer-urls", m.PeerURLs))
}

// RemoveMember  store中必须存在该ID，否则会panic
func (c *RaftCluster) RemoveMember(id types.ID, shouldApplyV3 ShouldApplyV3) {
	c.Lock()
	defer c.Unlock()
	var v2Err, beErr error
	if c.v2store != nil {
		v2Err = unsafeDeleteMemberFromStore(c.v2store, id)
		if v2Err != nil {
			if e, ok := v2Err.(*v2error.Error); !ok || e.ErrorCode != v2error.EcodeKeyNotFound {
				c.lg.Panic("从v2store删除节点失败", zap.String("member-id", id.String()), zap.Error(v2Err))
			}
		}
	}
	if c.be != nil && shouldApplyV3 {
		beErr = unsafeDeleteMemberFromBackend(c.be, id)
		if beErr != nil && !errors.Is(beErr, errMemberNotFound) {
			c.lg.Panic("从backend  bolt 删除节点失败", zap.String("member-id", id.String()), zap.Error(beErr))
		}
	}
	if v2Err != nil && (beErr != nil || c.be == nil) {
		c.lg.Panic("从store中删除节点失败", zap.String("member-id", id.String()), zap.Error(v2Err))
	}

	m, ok := c.members[id]
	delete(c.members, id)
	c.removed[id] = true

	if ok {
		c.lg.Info("移除成员", zap.String("cluster-id", c.cid.String()), zap.String("local-member-id", c.localID.String()), zap.String("removed-remote-peer-id", id.String()), zap.Strings("removed-remote-peer-urls", m.PeerURLs))
	} else {
		c.lg.Warn("该成员已经移除", zap.String("cluster-id", c.cid.String()), zap.String("local-member-id", c.localID.String()), zap.String("removed-remote-peer-id", id.String()))
	}
}
