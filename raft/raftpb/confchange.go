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

package raftpb

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
)

// ConfChangeI 配置变更
type ConfChangeI interface {
	AsV2() ConfChangeV2
	AsV1() (ConfChange, bool)
}

func MarshalConfChange(c ConfChangeI) (EntryType, []byte, error) {
	var typ EntryType
	var ccdata []byte
	var err error
	if ccv1, ok := c.AsV1(); ok {
		typ = EntryConfChange
		ccdata, err = ccv1.Marshal()
	} else {
		ccv2 := c.AsV2()
		typ = EntryConfChangeV2
		ccdata, err = ccv2.Marshal()
	}
	return typ, ccdata, err
}

func (c ConfChange) AsV2() ConfChangeV2 {
	return ConfChangeV2{
		Changes: []ConfChangeSingle{{
			Type:   c.Type,
			NodeID: c.NodeID,
		}},
		Context: c.Context,
	}
}

func (c ConfChange) AsV1() (ConfChange, bool) {
	return c, true
}

func (c ConfChangeV2) AsV2() ConfChangeV2 { return c }

func (c ConfChangeV2) AsV1() (ConfChange, bool) { return ConfChange{}, false }

// EnterJoint returns two bools. The second bool is true if and only if this
// config change will use Joint Consensus, which is the case if it contains more
// than one change or if the use of Joint Consensus was requested explicitly.
// The first bool can only be true if second one is, and indicates whether the
// Joint State will be left automatically.
func (c ConfChangeV2) EnterJoint() (autoLeave bool, ok bool) {
	// NB: in theory, more config changes could qualify for the "simple"
	// protocol but it depends on the config on top of which the changes apply.
	// For example, adding two learners is not OK if both nodes are part of the
	// base config (i.e. two voters are turned into learners in the process of
	// applying the conf change). In practice, these distinctions should not
	// matter, so we keep it simple and use Joint Consensus liberally.
	if c.Transition != ConfChangeTransitionAuto || len(c.Changes) > 1 {
		// Use Joint Consensus.
		var autoLeave bool
		switch c.Transition {
		case ConfChangeTransitionAuto:
			autoLeave = true
		case ConfChangeTransitionJointImplicit:
			autoLeave = true
		case ConfChangeTransitionJointExplicit:
		default:
			panic(fmt.Sprintf("unknown transition: %+v", c))
		}
		return autoLeave, true
	}
	return false, false
}

// LeaveJoint is true if the configuration change leaves a joint configuration.
// This is the case if the ConfChangeV2 is zero, with the possible exception of
// the Context field.
func (c ConfChangeV2) LeaveJoint() bool {
	// NB: c is already a copy.
	c.Context = nil
	return proto.Equal(&c, &ConfChangeV2{})
}

// ConfChangesToString is the inverse to ConfChangesFromString.
func ConfChangesToString(ccs []ConfChangeSingle) string {
	var buf strings.Builder
	for i, cc := range ccs {
		if i > 0 {
			buf.WriteByte(' ')
		}
		switch cc.Type {
		case ConfChangeAddNode:
			buf.WriteByte('v')
		case ConfChangeAddLearnerNode:
			buf.WriteByte('l')
		case ConfChangeRemoveNode:
			buf.WriteByte('r')
		case ConfChangeUpdateNode:
			buf.WriteByte('u')
		default:
			buf.WriteString("unknown")
		}
		fmt.Fprintf(&buf, "%d", cc.NodeID)
	}
	return buf.String()
}
