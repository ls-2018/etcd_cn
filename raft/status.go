// Copyright 2015 The etcd Authors
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

package raft

import (
	"fmt"

	pb "github.com/ls-2018/etcd_cn/raft/raftpb"
	"github.com/ls-2018/etcd_cn/raft/tracker"
)

// Status contains information about this Raft peer and its view of the system.
// The Progress is only populated on the leader.
type Status struct {
	BasicStatus
	Config   tracker.Config
	Progress map[uint64]tracker.Progress // 如果是Leader,还有其他节点的进度
}

type BasicStatus struct {
	ID uint64
	pb.HardState
	SoftState
	Applied        uint64 // 应用索引,其实这个使用者自己也知道,因为Ready的回调里提交日志被应用都会有日志的索引.
	LeadTransferee uint64 // Leader转移ID,如果正处于Leader转移期间.
}

func getProgressCopy(r *raft) map[uint64]tracker.Progress {
	m := make(map[uint64]tracker.Progress)
	r.prs.Visit(func(id uint64, pr *tracker.Progress) {
		p := *pr
		p.Inflights = pr.Inflights.Clone()
		pr = nil

		m[id] = p
	})
	return m
}

func getBasicStatus(r *raft) BasicStatus {
	s := BasicStatus{
		ID:             r.id,
		LeadTransferee: r.leadTransferee,
	}
	s.HardState = r.hardState()
	s.SoftState = *r.softState()
	s.Applied = r.raftLog.applied
	return s
}

// getStatus gets a copy of the current raft status.
func getStatus(r *raft) Status {
	var s Status
	s.BasicStatus = getBasicStatus(r)
	if s.RaftState == StateLeader {
		s.Progress = getProgressCopy(r)
	}
	s.Config = r.prs.Config.Clone()
	return s
}

// MarshalJSON translates the raft status into JSON.
// TODO: try to simplify this by introducing ID type into raft
func (s Status) MarshalJSON() ([]byte, error) {
	j := fmt.Sprintf(`{"id":"%x","term":%d,"vote":"%x","commit":%d,"lead":"%x","raftState":%q,"applied":%d,"progress":{`,
		s.ID, s.Term, s.Vote, s.Commit, s.Lead, s.RaftState, s.Applied)

	if len(s.Progress) == 0 {
		j += "},"
	} else {
		for k, v := range s.Progress {
			subj := fmt.Sprintf(`"%x":{"match":%d,"next":%d,"state":%q},`, k, v.Match, v.Next, v.State)
			j += subj
		}
		// remove the trailing ","
		j = j[:len(j)-1] + "},"
	}

	j += fmt.Sprintf(`"leadtransferee":"%x"}`, s.LeadTransferee)
	return []byte(j), nil
}

func (s Status) String() string {
	b, err := s.MarshalJSON()
	if err != nil {
		getLogger().Panicf("unexpected error: %v", err)
	}
	return string(b)
}

// SoftState 提供对日志和调试有用的状态.该状态是不稳定的,不需要持久化到WAL中.
type SoftState struct {
	Lead      uint64    // 当前leader
	RaftState StateType // 节点状态
}

func (a *SoftState) equal(b *SoftState) bool {
	return a.Lead == b.Lead && a.RaftState == b.RaftState
}
