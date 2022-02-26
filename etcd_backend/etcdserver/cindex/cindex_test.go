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

package cindex

import (
	"math/rand"
	"testing"
	"time"

	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc/backend"
	betesting "github.com/ls-2018/etcd_cn/etcd_backend/mvcc/backend/testing"
	"github.com/stretchr/testify/assert"
)

// TestConsistentIndex ensures that LoadConsistentIndex/Save/ConsistentIndex and backend.BatchTx can work well together.
func TestConsistentIndex(t *testing.T) {

	be, tmpPath := betesting.NewTmpBackend(t, time.Microsecond, 10)
	ci := NewConsistentIndex(be)

	tx := be.BatchTx()
	if tx == nil {
		t.Fatal("batch tx is nil")
	}
	tx.Lock()

	UnsafeCreateMetaBucket(tx)
	tx.Unlock()
	be.ForceCommit()
	r := uint64(7890123)
	term := uint64(234)
	ci.SetConsistentIndex(r, term)
	index := ci.ConsistentIndex()
	if index != r {
		t.Errorf("expected %d,got %d", r, index)
	}
	tx.Lock()
	ci.UnsafeSave(tx)
	tx.Unlock()
	be.ForceCommit()
	be.Close()

	b := backend.NewDefaultBackend(tmpPath)
	defer b.Close()
	ci.SetBackend(b)
	index = ci.ConsistentIndex()
	assert.Equal(t, r, index)

	ci = NewConsistentIndex(b)
	index = ci.ConsistentIndex()
	assert.Equal(t, r, index)
}

func TestFakeConsistentIndex(t *testing.T) {

	r := rand.Uint64()
	ci := NewFakeConsistentIndex(r)
	index := ci.ConsistentIndex()
	if index != r {
		t.Errorf("expected %d,got %d", r, index)
	}
	r = rand.Uint64()
	ci.SetConsistentIndex(r, 5)
	index = ci.ConsistentIndex()
	if index != r {
		t.Errorf("expected %d,got %d", r, index)
	}

}
