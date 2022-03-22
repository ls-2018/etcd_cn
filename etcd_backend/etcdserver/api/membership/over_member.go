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

package membership

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
)

// RaftAttributes  与raft相关的etcd成员属性
type RaftAttributes struct {
	PeerURLs  []string `json:"peerURLs"`            // 是raft集群中的对等体列表.
	IsLearner bool     `json:"isLearner,omitempty"` // 表示该成员是否是raft Learner.
}

// Attributes 代表一个etcd成员的所有非raft的相关属性.
type Attributes struct {
	Name       string   `json:"name,omitempty"`       // 节点创建时设置的name   默认default
	ClientURLs []string `json:"clientURLs,omitempty"` // 当接受到来自该Name的请求时,会
}

type Member struct {
	ID             types.ID `json:"id"` // hash得到的, 本节点ID
	RaftAttributes          // 与raft相关的etcd成员属性
	Attributes              // 代表一个etcd成员的所有非raft的相关属性
}

// NewMember 创建一个没有ID的成员,并根据集群名称、peer的URLS 和时间生成一个ID.这是用来引导/添加新成员的.
func NewMember(name string, peerURLs types.URLs, clusterName string, now *time.Time) *Member {
	memberId := computeMemberId(peerURLs, clusterName, now)
	return newMember(name, peerURLs, memberId, false)
}

// NewMemberAsLearner 创建一个没有ID的成员,并根据集群名称、peer的URLS 和时间生成一个ID.这是用来引导新learner成员的.
func NewMemberAsLearner(name string, peerURLs types.URLs, clusterName string, now *time.Time) *Member {
	memberId := computeMemberId(peerURLs, clusterName, now)
	return newMember(name, peerURLs, memberId, true)
}

// 计算成员ID
func computeMemberId(peerURLs types.URLs, clusterName string, now *time.Time) types.ID {
	peerURLstrs := peerURLs.StringSlice()
	sort.Strings(peerURLstrs)
	joinedPeerUrls := strings.Join(peerURLstrs, "")
	b := []byte(joinedPeerUrls)

	b = append(b, []byte(clusterName)...)
	if now != nil {
		b = append(b, []byte(fmt.Sprintf("%d", now.Unix()))...)
	}

	hash := sha1.Sum(b)
	return types.ID(binary.BigEndian.Uint64(hash[:8]))
}

func newMember(name string, peerURLs types.URLs, memberId types.ID, isLearner bool) *Member {
	m := &Member{
		RaftAttributes: RaftAttributes{
			PeerURLs:  peerURLs.StringSlice(),
			IsLearner: isLearner,
		},
		Attributes: Attributes{Name: name},
		ID:         memberId,
	}
	return m
}

// PickPeerURL chooses a random address from a given Member's PeerURLs.
// It will panic if there is no PeerURLs available in Member.
func (m *Member) PickPeerURL() string {
	if len(m.PeerURLs) == 0 {
		panic("member should always have some peer url")
	}
	return m.PeerURLs[rand.Intn(len(m.PeerURLs))]
}

// Clone 返回member deepcopy
func (m *Member) Clone() *Member {
	if m == nil {
		return nil
	}
	mm := &Member{
		ID: m.ID,
		RaftAttributes: RaftAttributes{
			IsLearner: m.IsLearner,
		},
		Attributes: Attributes{
			Name: m.Name,
		},
	}
	if m.PeerURLs != nil {
		mm.PeerURLs = make([]string, len(m.PeerURLs))
		copy(mm.PeerURLs, m.PeerURLs)
	}
	if m.ClientURLs != nil {
		mm.ClientURLs = make([]string, len(m.ClientURLs))
		copy(mm.ClientURLs, m.ClientURLs)
	}
	return mm
}

func (m *Member) IsStarted() bool {
	return len(m.Name) != 0
}

type MembersByID []*Member

func (ms MembersByID) Len() int           { return len(ms) }
func (ms MembersByID) Less(i, j int) bool { return ms[i].ID < ms[j].ID }
func (ms MembersByID) Swap(i, j int)      { ms[i], ms[j] = ms[j], ms[i] }

type MembersByPeerURLs []*Member

func (ms MembersByPeerURLs) Len() int { return len(ms) }
func (ms MembersByPeerURLs) Less(i, j int) bool {
	return ms[i].PeerURLs[0] < ms[j].PeerURLs[0]
}
func (ms MembersByPeerURLs) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }
