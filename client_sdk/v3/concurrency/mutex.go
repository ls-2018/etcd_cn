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

package concurrency

import (
	"context"
	"errors"
	"fmt"
	v3 "github.com/ls-2018/etcd_cn/client_sdk/v3"
	"sync"

	pb "go.etcd.io/etcd/api/v3/etcdserverpb"
)

// ErrLocked is returned by TryLock when Mutex is already locked by another session.
var ErrLocked = errors.New("mutex: Locked by another session")
var ErrSessionExpired = errors.New("mutex: session is expired")

// Mutex implements the sync Locker interface with etcd
//即前缀机制，也称目录机制，例如，一个名为 `/mylock` 的锁，两个争抢它的客户端进行写操作，
// 实际写入的Key分别为：`key1="/mylock/UUID1"`,`key2="/mylock/UUID2"`，
//其中，UUID表示全局唯一的ID，确保两个Key的唯一性。很显然，写操作都会成功，但返回的Revision不一样，
//那么，如何判断谁获得了锁呢？通过前缀`“/mylock”`查询，返回包含两个Key-Value对的Key-Value列表，
//同时也包含它们的Revision，通过Revision大小，客户端可以判断自己是否获得锁，如果抢锁失败，则等待锁释放（对应的 Key 被删除或者租约过期），
//然后再判断自己是否可以获得锁。
type Mutex struct {
	s *Session

	pfx   string // 前缀
	myKey string // key
	myRev int64  // 自增的Revision
	hdr   *pb.ResponseHeader
}

func NewMutex(s *Session, pfx string) *Mutex {
	return &Mutex{s, pfx + "/", "", -1, nil}
}

// TryLock locks the mutex if not already locked by another session.
// If lock is held by another session, return immediately after attempting necessary cleanup
// The ctx argument is used for the sending/receiving Txn RPC.
func (m *Mutex) TryLock(ctx context.Context) error {
	resp, err := m.tryAcquire(ctx)
	if err != nil {
		return err
	}
	// if no key on prefix / the minimum rev is key, already hold the lock
	ownerKey := resp.Responses[1].GetResponseRange().Kvs
	if len(ownerKey) == 0 || ownerKey[0].CreateRevision == m.myRev {
		m.hdr = resp.Header
		return nil
	}
	client := m.s.Client()
	// Cannot lock, so delete the key
	if _, err := client.Delete(ctx, m.myKey); err != nil {
		return err
	}
	m.myKey = "\x00"
	m.myRev = -1
	return ErrLocked
}

// Lock locks the mutex with a cancelable context. If the context is canceled
// while trying to acquire the lock, the mutex tries to clean its stale lock entry.
// Lock 使用可取消的context锁定互斥锁.如果context被取消
// 在尝试获取锁时,互斥锁会尝试清除其过时的锁条目.
func (m *Mutex) Lock(ctx context.Context) error {
	resp, err := m.tryAcquire(ctx)
	if err != nil {
		return err
	}
	// if no key on prefix / the minimum rev is key, already hold the lock
	ownerKey := resp.Responses[1].GetResponseRange().Kvs
	if len(ownerKey) == 0 || ownerKey[0].CreateRevision == m.myRev {
		m.hdr = resp.Header
		return nil
	}
	client := m.s.Client()
	// wait for deletion revisions prior to myKey
	// waitDeletes 有效地等待,直到所有键匹配前缀且不大于
	// 创建的version.
	// TODO: early termination if the session key is deleted before other session keys with smaller revisions.
	_, werr := waitDeletes(ctx, client, m.pfx, m.myRev-1)
	// release lock key if wait failed
	if werr != nil {
		m.Unlock(client.Ctx())
		return werr
	}

	// make sure the session is not expired, and the owner key still exists.
	gresp, werr := client.Get(ctx, m.myKey)
	if werr != nil {
		m.Unlock(client.Ctx())
		return werr
	}

	if len(gresp.Kvs) == 0 { // is the session key lost?
		return ErrSessionExpired
	}
	m.hdr = gresp.Header

	return nil
}

func (m *Mutex) tryAcquire(ctx context.Context) (*v3.TxnResponse, error) {
	s := m.s
	client := m.s.Client()
	// s.Lease()租约
	m.myKey = fmt.Sprintf("%s%x", m.pfx, s.Lease())
	// 比较Revision, 这里构建了一个比较表达式
	// 具体的比较逻辑在下面的client.Txn用到
	// 如果等于0,写入当前的key,并设置租约,
	// 否则获取这个key,重用租约中的锁(这里主要目的是在于重入)
	// 通过第二次获取锁,判断锁是否存在来支持重入
	// 所以只要租约一致,那么是可以重入的.
	cmp := v3.Compare(v3.CreateRevision(m.myKey), "=", 0)
	// 通过 myKey 将自己锁在waiters；最早的waiters将获得锁
	put := v3.OpPut(m.myKey, "", v3.WithLease(s.Lease()))
	// 获取已经拿到锁的key的信息
	get := v3.OpGet(m.myKey)
	// 仅使用一个 RPC 获取当前持有者以完成无竞争路径
	getOwner := v3.OpGet(m.pfx, v3.WithFirstCreate()...)
	// 这里是比较的逻辑,如果等于0,写入当前的key,否则则读取这个key
	// 大佬的代码写的就是奇妙
	resp, err := client.Txn(ctx).If(cmp).Then(put, getOwner).Else(get, getOwner).Commit()
	if err != nil {
		return nil, err
	}
	// 根据比较操作的结果写入Revision到m.myRev中
	m.myRev = resp.Header.Revision
	if !resp.Succeeded {
		m.myRev = resp.Responses[0].GetResponseRange().Kvs[0].CreateRevision
	}
	return resp, nil
}

func (m *Mutex) Unlock(ctx context.Context) error {
	client := m.s.Client()
	if _, err := client.Delete(ctx, m.myKey); err != nil {
		return err
	}
	m.myKey = "\x00"
	m.myRev = -1
	return nil
}

func (m *Mutex) IsOwner() v3.Cmp {
	return v3.Compare(v3.CreateRevision(m.myKey), "=", m.myRev)
}

func (m *Mutex) Key() string { return m.myKey }

// Header is the response header received from etcd on acquiring the lock.
func (m *Mutex) Header() *pb.ResponseHeader { return m.hdr }

type lockerMutex struct{ *Mutex }

func (lm *lockerMutex) Lock() {
	client := lm.s.Client()
	if err := lm.Mutex.Lock(client.Ctx()); err != nil {
		panic(err)
	}
}
func (lm *lockerMutex) Unlock() {
	client := lm.s.Client()
	if err := lm.Mutex.Unlock(client.Ctx()); err != nil {
		panic(err)
	}
}

// NewLocker creates a sync.Locker backed by an etcd mutex.
func NewLocker(s *Session, pfx string) sync.Locker {
	return &lockerMutex{NewMutex(s, pfx)}
}
