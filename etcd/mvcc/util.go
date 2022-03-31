// Copyright 2016 The etcd Authors
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

package mvcc

import (
	"fmt"

	"github.com/ls-2018/etcd_cn/etcd/mvcc/backend"
	"github.com/ls-2018/etcd_cn/etcd/mvcc/buckets"
	"github.com/ls-2018/etcd_cn/offical/api/v3/mvccpb"
)

func WriteKV(be backend.Backend, kv mvccpb.KeyValue) {
	indexBytes := newRevBytes()
	revToBytes(revision{main: kv.ModRevision}, indexBytes)

	d, err := kv.Marshal()
	if err != nil {
		panic(fmt.Errorf("cannot marshal event: %v", err))
	}

	be.BatchTx().Lock()
	be.BatchTx().UnsafePut(buckets.Key, indexBytes, d)
	be.BatchTx().Unlock()
}
