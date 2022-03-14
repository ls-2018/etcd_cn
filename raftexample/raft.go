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

package main

import (
	"context"
	"fmt"
	"github.com/ls-2018/etcd_cn/raft"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/fileutil"
	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
	"github.com/ls-2018/etcd_cn/etcd_backend/etcdserver/api/rafthttp"
	"github.com/ls-2018/etcd_cn/etcd_backend/etcdserver/api/snap"
	stats "github.com/ls-2018/etcd_cn/etcd_backend/etcdserver/api/v2stats"
	"github.com/ls-2018/etcd_cn/etcd_backend/wal"
	"github.com/ls-2018/etcd_cn/etcd_backend/wal/walpb"
	"github.com/ls-2018/etcd_cn/raft/raftpb"

	"go.uber.org/zap"
)

// 一个批次一个批次交给上层应用来处理
type commit struct {
	data       []string
	applyDoneC chan<- struct{}
}

// A key-value stream backed by raft
type raftNode struct {
	proposeC    <-chan string            // proposed messages (k,v)
	confChangeC <-chan raftpb.ConfChange // proposed cluster config changes
	commitC     chan<- *commit           // entries committed to log (k,v)
	errorC      chan<- error             // errors from raft session

	id          int      // client ID for raft session
	peers       []string // raft peer URLs
	join        bool     // node is joining an existing cluster
	waldir      string   // wal存储路径
	snapdir     string   // 存储快照的路径
	getSnapshot func() ([]byte, error)

	confState     raftpb.ConfState
	snapshotIndex uint64
	appliedIndex  uint64

	// raft backing for the commit/error channel
	node        raft.RaftNodeInterFace
	raftStorage *raft.MemoryStorage
	wal         *wal.WAL

	snapshotter      *snap.Snapshotter      // 快照管理者
	snapshotterReady chan *snap.Snapshotter // 快照管理者就绪的信号

	snapCount uint64
	transport *rafthttp.Transport // 负责 raft 节点之间的网络通信服务
	stopc     chan struct{}       // signals proposal channel closed
	httpstopc chan struct{}       // signals http etcd to shutdown
	httpdonec chan struct{}       // signals http etcd shutdown complete

	logger *zap.Logger
}

func (rc *raftNode) saveSnap(snap raftpb.Snapshot) error {
	walSnap := walpb.Snapshot{
		Index:     snap.Metadata.Index,
		Term:      snap.Metadata.Term,
		ConfState: &snap.Metadata.ConfState,
	}
	//在把快照写到wal之前保存快照文件。这使得快照文件有可能成为孤儿，但可以防止一个WAL快照条目没有相应的快照文件。
	if err := rc.snapshotter.SaveSnap(snap); err != nil {
		return err
	}
	if err := rc.wal.SaveSnapshot(walSnap); err != nil {
		return err
	}
	return rc.wal.ReleaseLockTo(snap.Metadata.Index)
}

func (rc *raftNode) entriesToApply(ents []raftpb.Entry) (nents []raftpb.Entry) {
	if len(ents) == 0 {
		return ents
	}
	firstIdx := ents[0].Index
	if firstIdx > rc.appliedIndex+1 {
		log.Fatalf("first index of committed entry[%d] should <= progress.appliedIndex[%d]+1", firstIdx, rc.appliedIndex)
	}
	if rc.appliedIndex-firstIdx+1 < uint64(len(ents)) {
		nents = ents[rc.appliedIndex-firstIdx+1:]
	}
	return nents
}

// publishEntries writes committed log entries to commit channel and returns
// whether all entries could be published.
func (rc *raftNode) publishEntries(ents []raftpb.Entry) (<-chan struct{}, bool) {
	if len(ents) == 0 {
		return nil, true
	}

	data := make([]string, 0, len(ents))
	for i := range ents {
		switch ents[i].Type {
		case raftpb.EntryNormal:
			if len(ents[i].Data) == 0 {
				// ignore empty messages
				break
			}
			s := string(ents[i].Data)
			data = append(data, s)
		case raftpb.EntryConfChange:
			var cc raftpb.ConfChange
			cc.Unmarshal(ents[i].Data)
			rc.confState = *rc.node.ApplyConfChange(cc)
			switch cc.Type {
			case raftpb.ConfChangeAddNode:
				if len(cc.Context) > 0 {
					rc.transport.AddPeer(types.ID(cc.NodeID), []string{string(cc.Context)})
				}
			case raftpb.ConfChangeRemoveNode:
				if cc.NodeID == uint64(rc.id) {
					log.Println("I've been removed from the cluster! Shutting down.")
					return nil, false
				}
				rc.transport.RemovePeer(types.ID(cc.NodeID))
			}
		}
	}

	var applyDoneC chan struct{}

	if len(data) > 0 {
		applyDoneC = make(chan struct{}, 1)
		select {
		case rc.commitC <- &commit{data, applyDoneC}:
		case <-rc.stopc:
			return nil, false
		}
	}

	// after commit, update appliedIndex
	rc.appliedIndex = ents[len(ents)-1].Index

	return applyDoneC, true
}

// openWAL returns a WAL ready for reading.
func (rc *raftNode) openWAL(snapshot *raftpb.Snapshot) *wal.WAL {
	if !wal.Exist(rc.waldir) {
		if err := os.Mkdir(rc.waldir, 0750); err != nil {
			log.Fatalf("raftexample: cannot create dir for wal (%v)", err)
		}

		w, err := wal.Create(zap.NewExample(), rc.waldir, nil)
		if err != nil {
			log.Fatalf("raftexample: create wal error (%v)", err)
		}
		w.Close()
	}

	walsnap := walpb.Snapshot{}
	if snapshot != nil {
		walsnap.Index, walsnap.Term = snapshot.Metadata.Index, snapshot.Metadata.Term
	}
	log.Printf("loading WAL at term %d and index %d", walsnap.Term, walsnap.Index)
	w, err := wal.Open(zap.NewExample(), rc.waldir, walsnap)
	if err != nil {
		log.Fatalf("raftexample: error loading wal (%v)", err)
	}

	return w
}

// replayWAL 重放wal日志
func (rc *raftNode) replayWAL() *wal.WAL {
	log.Printf("replaying WAL of member %d", rc.id)
	snapshot := rc.loadSnapshot()
	w := rc.openWAL(snapshot)
	_, st, ents, err := w.ReadAll()
	if err != nil {
		log.Fatalf("raftexample: failed to read WAL (%v)", err)
	}
	rc.raftStorage = raft.NewMemoryStorage()
	if snapshot != nil {
		rc.raftStorage.ApplySnapshot(*snapshot) // 从持久化的内存存储中恢复出快照
	}
	rc.raftStorage.SetHardState(st) // 从持久化的内存存储中恢复出状态

	// append to storage so raft starts at the right place in log
	rc.raftStorage.Append(ents) // 从持久化的内存存储中恢复出日志

	return w
}

func (rc *raftNode) writeError(err error) {
	rc.stopHTTP()
	close(rc.commitC)
	rc.errorC <- err
	close(rc.errorC)
	rc.node.Stop()
}

func (rc *raftNode) startRaft() {
	if !fileutil.Exist(rc.snapdir) {
		if err := os.Mkdir(rc.snapdir, 0750); err != nil {
			log.Fatalf("raftexample: 无法创建快照目录 (%v)", err)
		}
	}
	rc.snapshotter = snap.New(zap.NewExample(), rc.snapdir)
	// 创建 WAL 实例,然后加载快照并回放 WAL 日志
	oldwal := wal.Exist(rc.waldir)
	//raftNode.replayWAL() 方法首先会读取快照数据,
	//在快照数据中记录了该快照包含的最后一条Entry记录的 Term 值 和 索引值.
	//然后根据 Term 值 和 索引值确定读取 WAL 日志文件的位置, 并进行日志记录的读取.
	rc.wal = rc.replayWAL()
	rc.snapshotterReady <- rc.snapshotter

	rpeers := make([]raft.Peer, len(rc.peers))
	for i := range rpeers {
		rpeers[i] = raft.Peer{ID: uint64(i + 1)}
	}
	// 创建 raft.Config 实例
	c := &raft.Config{
		ID:                        uint64(rc.id),
		ElectionTick:              10, // 返回选举权检查对应多少次tick触发次数
		HeartbeatTick:             1,  // 返回心跳检查对应多少次tick触发次数
		Storage:                   rc.raftStorage,
		MaxSizePerMsg:             1024 * 1024,
		MaxInflightMsgs:           256,
		MaxUncommittedEntriesSize: 1 << 30,
	}
	// 初始化底层的 etcd-raft 模块,这里会根据 WAL 日志的回放情况,
	// 判断当前节点是首次启动还是重新启动
	if oldwal || rc.join {
		rc.node = raft.RestartNode(c)
	} else {
		rc.node = raft.StartNode(c, rpeers)
	}
	// 创建 Transport 实例并启动,他负责 raft 节点之间的网络通信服务
	rc.transport = &rafthttp.Transport{
		Logger:      rc.logger,
		ID:          types.ID(rc.id),
		ClusterID:   0x1000,
		Raft:        rc,
		ServerStats: stats.NewServerStats("", ""),
		LeaderStats: stats.NewLeaderStats(zap.NewExample(), strconv.Itoa(rc.id)),
		ErrorC:      make(chan error),
	}
	// 启动网络服务相关组件
	rc.transport.Start()
	// 建立与集群中其他各个节点的连接
	for i := range rc.peers {
		if i+1 != rc.id {
			rc.transport.AddPeer(types.ID(i+1), []string{rc.peers[i]})
		}
	}
	// 启动一个goroutine,其中会监听当前节点与集群中其他节点之间的网络连接
	go rc.serveRaft()
	// 启动后台 goroutine 处理上层应用与底层 etcd-raft 模块的交互
	go rc.serveChannels()
}

var snapshotCatchUpEntriesN uint64 = 10000

// -----------------------------------------------------------------------------------------------------------------

// 判断是否应该创建快照,每次apply就会调用
func (rc *raftNode) maybeTriggerSnapshot(applyDoneC <-chan struct{}) {
	if rc.appliedIndex-rc.snapshotIndex <= rc.snapCount {
		// 日志条数 阈值
		return
	}

	// 等待所有提交的条目被应用(或etcd被关闭)
	if applyDoneC != nil {
		select {
		case <-applyDoneC:
		case <-rc.stopc:
			return
		}
	}

	log.Printf("开始打快照 [applied index: %d | last snapshot index: %d]", rc.appliedIndex, rc.snapshotIndex)
	data, err := rc.getSnapshot()
	if err != nil {
		log.Panic(err)
	}
	snap, err := rc.raftStorage.CreateSnapshot(rc.appliedIndex, &rc.confState, data)
	if err != nil {
		panic(err)
	}
	if err := rc.saveSnap(snap); err != nil {
		panic(err)
	}

	compactIndex := uint64(1)
	if rc.appliedIndex > snapshotCatchUpEntriesN {
		compactIndex = rc.appliedIndex - snapshotCatchUpEntriesN
	}
	if err := rc.raftStorage.Compact(compactIndex); err != nil {
		panic(err)
	}

	log.Printf("压缩 [0 ,%d]的日志索引 ", compactIndex)
	rc.snapshotIndex = rc.appliedIndex
}

// 会单独启动一个后台 goroutine来负责上层模块 传递给 etcd-ra企 模块的数据,
// 主要 处理前面介绍的 proposeC、 confChangeC 两个通道
func (rc *raftNode) serveChannels() {
	// 这里是获取快照数据和快照的元数据
	snap, err := rc.raftStorage.Snapshot()
	if err != nil {
		panic(err)
	}
	rc.confState = snap.Metadata.ConfState
	rc.snapshotIndex = snap.Metadata.Index
	rc.appliedIndex = snap.Metadata.Index

	defer rc.wal.Close()

	// 创建一个每隔 lOOms 触发一次的定时器,那么在逻辑上,lOOms 即是 etcd-raft 组件的最小时间单位 ,
	// 该定时器每触发一次,则逻辑时钟推进一次
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// 单独启 动一个 goroutine 负责将 proposeC、 confChangeC 远远上接收到
	// 的数据传递给 etcd-raft 组件进行处理
	go func() {
		confChangeCount := uint64(0)

		for rc.proposeC != nil && rc.confChangeC != nil {
			select {
			case prop, ok := <-rc.proposeC:
				if !ok {
					// 发生异常将proposeC置空
					rc.proposeC = nil
				} else {
					// 阻塞直到消息被处理
					rc.node.Propose(context.TODO(), []byte(prop))
				}
				// 收到上层应用通过 confChangeC远远传递过来的数据
			case cc, ok := <-rc.confChangeC:
				if !ok {
					// 如果发生异常将confChangeC置空
					rc.confChangeC = nil
				} else {
					confChangeCount++
					cc.ID = confChangeCount
					rc.node.ProposeConfChange(context.TODO(), cc)
				}
			}
		}
		// 关闭 stopc 通道,触发 rafeNode.stop() 方法的调用
		close(rc.stopc)
	}()

	// 处理 etcd-raft 模块返回给上层模块的数据及其他相关的操作
	for {
		select {
		case <-ticker.C:
			// 上述 ticker 定时器触发一次
			rc.node.Tick()

		// 读取 node.readyc 通道
		// 该通道是 etcd-raft 组件与上层应用交互的主要channel之一
		// 其中传递的 Ready 实例也封装了很多信息
		case rd := <-rc.node.Ready(): // demo
			// 将当前 etcd raft 组件的状态信息,以及待持久化的 Entry 记录先记录到 WAL 日志文件中,
			// 即使之后宕机,这些信息也可以在节点下次启动时,通过前面回放 WAL 日志的方式进行恢复
			rc.wal.Save(rd.HardState, rd.Entries)
			// 检测到 etcd-raft 组件生成了新的快照数据
			if !raft.IsEmptySnap(rd.Snapshot) {
				// 将新的快照数据写入快照文件中
				rc.saveSnap(rd.Snapshot)
				rc.raftStorage.ApplySnapshot(rd.Snapshot) // 从持久化的内存存储中恢复出快照
				rc.publishSnapshot(rd.Snapshot)
			}
			// 将待持久化的 Entry 记录追加到 raftStorage 中完成持久化
			rc.raftStorage.Append(rd.Entries)
			// 将待发送的消息发送到指定节点
			rc.transport.Send(rd.Messages)
			// 将已提交、待应用的 Entry 记录应用到上层应用的状态机中
			applyDoneC, ok := rc.publishEntries(rc.entriesToApply(rd.CommittedEntries))
			if !ok {
				rc.stop()
				return
			}

			// 随着节点的运行, WAL 日志量和 raftLog.storage 中的 Entry 记录会不断增加 ,
			// 所以节点每处理 10000 条(默认值) Entry 记录,就会触发一次创建快照的过程,
			// 同时 WAL 会释放一些日志文件的句柄,raftLog.storage 也会压缩其保存的 Entry 记录
			rc.maybeTriggerSnapshot(applyDoneC)
			// 上层应用处理完该 Ready 实例,通知 etcd-raft 纽件准备返回下一个 Ready 实例
			rc.node.Advance()

		case err := <-rc.transport.ErrorC:
			rc.writeError(err)
			return

		case <-rc.stopc:
			rc.stop()
			return
		}
	}
}

func (rc *raftNode) serveRaft() {
	url, err := url.Parse(rc.peers[rc.id-1])
	if err != nil {
		log.Fatalf("raftexample: Failed parsing URL (%v)", err)
	}

	ln, err := newStoppableListener(url.Host, rc.httpstopc)
	if err != nil {
		log.Fatalf("raftexample: Failed to listen rafthttp (%v)", err)
	}

	err = (&http.Server{Handler: rc.transport.Handler()}).Serve(ln)
	select {
	case <-rc.httpstopc:
	default:
		log.Fatalf("raftexample: Failed to serve rafthttp (%v)", err)
	}
	close(rc.httpdonec)
}

func (rc *raftNode) Process(ctx context.Context, m raftpb.Message) error {
	return rc.node.Step(ctx, m)
}

func (rc *raftNode) IsIDRemoved(id uint64) bool                           { return false }
func (rc *raftNode) ReportUnreachable(id uint64)                          {}
func (rc *raftNode) ReportSnapshot(id uint64, status raft.SnapshotStatus) {}

// 主要完成了raftNode的初始化
// 使用上层模块传入的配置信息来创建raftNode实例,同时创建commitC 通道和errorC通道返回给上层模块使用
// 上层的应用通过这几个channel就能和raftNode进行交互
func newRaftNode(id int, peers []string, join bool, getSnapshot func() ([]byte, error),
	proposeC <-chan string,
	confChangeC <-chan raftpb.ConfChange) (<-chan *commit, <-chan error, <-chan *snap.Snapshotter) {

	// channel,主要传输Entry记录
	// raftNode会将etcd-raft模块返回的待apply Entry封装在 Ready实例中然后 写入commitC通道,
	// 另一方面,kvstore会从commitC通道中读取这些待应用的 Entry 记录井保存其中的键值对信息.
	commitC := make(chan *commit)
	errorC := make(chan error)

	rc := &raftNode{
		proposeC:    proposeC,
		confChangeC: confChangeC,
		commitC:     commitC,
		errorC:      errorC,
		id:          id,
		peers:       peers,
		join:        join,
		// 初始化存放 WAL 日志和 Snapshot 文件的的目录
		waldir:           fmt.Sprintf("./raftexample/db/raftexample-%d", id),
		snapdir:          fmt.Sprintf("./raftexample/db/raftexample-%d-snap", id),
		getSnapshot:      getSnapshot,
		snapCount:        10000,
		stopc:            make(chan struct{}),
		httpstopc:        make(chan struct{}),
		httpdonec:        make(chan struct{}),
		logger:           zap.NewExample(),
		snapshotterReady: make(chan *snap.Snapshotter, 1),
		// 其余结构在WAL重放后填充
	}
	// 启动一个goroutine,完成剩余的初始化工作
	go rc.startRaft()
	return commitC, errorC, rc.snapshotterReady
}

// stop closes http, closes all channels, and stops raft.
func (rc *raftNode) stop() {
	rc.stopHTTP()
	close(rc.commitC)
	close(rc.errorC)
	rc.node.Stop()
}

func (rc *raftNode) stopHTTP() {
	rc.transport.Stop()
	close(rc.httpstopc)
	<-rc.httpdonec
}

func (rc *raftNode) loadSnapshot() *raftpb.Snapshot {
	if wal.Exist(rc.waldir) {
		walSnaps, err := wal.ValidSnapshotEntries(rc.logger, rc.waldir)
		if err != nil {
			log.Fatalf("raftexample: error listing snapshots (%v)", err)
		}
		snapshot, err := rc.snapshotter.LoadNewestAvailable(walSnaps)
		if err != nil && err != snap.ErrNoSnapshot {
			log.Fatalf("raftexample: error loading snapshot (%v)", err)
		}
		return snapshot
	}
	return &raftpb.Snapshot{}
}

// 接收到了快照
func (rc *raftNode) publishSnapshot(snapshotToSave raftpb.Snapshot) {
	if raft.IsEmptySnap(snapshotToSave) {
		return
	}

	log.Printf("publishing snapshot at index %d", rc.snapshotIndex)
	defer log.Printf("finished publishing snapshot at index %d", rc.snapshotIndex)

	if snapshotToSave.Metadata.Index <= rc.appliedIndex {
		log.Fatalf("snapshot index [%d] should > progress.appliedIndex [%d]", snapshotToSave.Metadata.Index, rc.appliedIndex)
	}
	rc.commitC <- nil // 通知应用加载快照
	rc.confState = snapshotToSave.Metadata.ConfState
	rc.snapshotIndex = snapshotToSave.Metadata.Index
	rc.appliedIndex = snapshotToSave.Metadata.Index
}
