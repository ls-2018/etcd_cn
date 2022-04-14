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

package v3rpc

import (
	"context"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/ls-2018/etcd_cn/etcd/auth"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver"
	"github.com/ls-2018/etcd_cn/etcd/mvcc"
	"github.com/ls-2018/etcd_cn/offical/api/v3/mvccpb"
	"github.com/ls-2018/etcd_cn/offical/api/v3/v3rpc/rpctypes"
	pb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"

	"go.uber.org/zap"
)

const minWatchProgressInterval = 100 * time.Millisecond

type watchServer struct {
	lg              *zap.Logger
	clusterID       int64
	memberID        int64
	maxRequestBytes int
	sg              etcdserver.RaftStatusGetter
	watchable       mvcc.WatchableKV
	ag              AuthGetter
}

var (
	progressReportInterval   = 10 * time.Minute
	progressReportIntervalMu sync.RWMutex
)

// SetProgressReportInterval 更新进度汇报间隔
func SetProgressReportInterval(newTimeout time.Duration) {
	progressReportIntervalMu.Lock()
	progressReportInterval = newTimeout
	progressReportIntervalMu.Unlock()
}

// We send ctrl response inside the read loop. We do not want
// send to block read, but we still want ctrl response we sent to
// be serialized. Thus we use a buffered chan to solve the problem.
// A small buffer should be OK for most cases, since we expect the
// ctrl requests are infrequent.
const ctrlStreamBufLen = 16

// serverWatchStream is an etcd etcd side stream. It receives requests
// from client side gRPC stream. It receives watch events from mvcc.WatchStream,
// and creates responses that forwarded to gRPC stream.
// It also forwards control message like watch created and canceled.
type serverWatchStream struct {
	lg *zap.Logger

	clusterID int64
	memberID  int64

	maxRequestBytes int

	sg        etcdserver.RaftStatusGetter
	watchable mvcc.WatchableKV
	ag        AuthGetter

	gRPCStream  pb.Watch_WatchServer
	watchStream mvcc.WatchStream
	ctrlStream  chan *pb.WatchResponse // 用来发送控制响应的Chan，比如watcher创建和取消。

	// mu protects progress, prevKV, fragment
	mu sync.RWMutex
	// tracks the watchID that stream might need to send progress to
	// TODO: combine progress and prevKV into a single struct?
	progress map[mvcc.WatchID]bool // 储存watcher
	// record watch IDs that need return previous key-value pair
	prevKV map[mvcc.WatchID]bool
	// records fragmented watch IDs
	fragment map[mvcc.WatchID]bool

	closec chan struct{}
	wg     sync.WaitGroup // 等待send loop 完成
}

// Watch 接收到watch请求
func (ws *watchServer) Watch(stream pb.Watch_WatchServer) (err error) {
	sws := serverWatchStream{
		lg: ws.lg,

		clusterID: ws.clusterID,
		memberID:  ws.memberID,

		maxRequestBytes: ws.maxRequestBytes,

		sg:        ws.sg,
		watchable: ws.watchable,
		ag:        ws.ag,

		gRPCStream:  stream,
		watchStream: ws.watchable.NewWatchStream(),
		// 用来发送控制响应的Chan，比如watcher创建和取消。
		ctrlStream: make(chan *pb.WatchResponse, ctrlStreamBufLen),

		progress: make(map[mvcc.WatchID]bool),
		prevKV:   make(map[mvcc.WatchID]bool),
		fragment: make(map[mvcc.WatchID]bool),

		closec: make(chan struct{}),
	}

	sws.wg.Add(1)
	go func() {
		sws.sendLoop() // 回复变更事件、阻塞
		sws.wg.Done()
	}()

	errc := make(chan error, 1)
	// 理想情况下，recvLoop也会使用sws.wg。当stream. context (). done()被关闭时，流的recv可能会继续阻塞，因为它使用了不同的上下文，导致调用sws.close()时死锁。
	go func() {
		if rerr := sws.recvLoop(); rerr != nil {
			if isClientCtxErr(stream.Context().Err(), rerr) {
				sws.lg.Debug("从gRPC流接收watch请求失败", zap.Error(rerr))
			} else {
				sws.lg.Warn("从gRPC流接收watch请求失败", zap.Error(rerr))
			}
			errc <- rerr
		}
	}()
	select {
	case err = <-errc:
		if err == context.Canceled {
			err = rpctypes.ErrGRPCWatchCanceled
		}
		close(sws.ctrlStream)
	case <-stream.Context().Done():
		err = stream.Context().Err()
		if err == context.Canceled {
			err = rpctypes.ErrGRPCWatchCanceled
		}
	}

	sws.close()
	return err
}

func (sws *serverWatchStream) isWatchPermitted(wcr *pb.WatchCreateRequest) bool {
	authInfo, err := sws.ag.AuthInfoFromCtx(sws.gRPCStream.Context())
	if err != nil {
		return false
	}
	if authInfo == nil {
		// if auth is enabled, IsRangePermitted() can cause an error
		authInfo = &auth.AuthInfo{}
	}
	return sws.ag.AuthStore().IsRangePermitted(authInfo, []byte(wcr.Key), []byte(wcr.RangeEnd)) == nil
}

// 接收watch请求，可以是创建、取消、和
func (sws *serverWatchStream) recvLoop() error {
	for {
		// 同一个连接，可以接收多次不同的消息
		req, err := sws.gRPCStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if req.WatchRequest_CreateRequest != nil {
			uv := &pb.WatchRequest_CreateRequest{}
			uv = req.WatchRequest_CreateRequest
			if uv.CreateRequest == nil {
				continue
			}

			creq := uv.CreateRequest
			if len(creq.Key) == 0 { // a
				// \x00 is the smallest key
				creq.Key = string([]byte{0})
			}
			if len(creq.RangeEnd) == 0 {
				// force nil since watchstream.Watch distinguishes
				// between nil and []byte{} for single key / >=
				creq.RangeEnd = ""
			}
			if len(creq.RangeEnd) == 1 && creq.RangeEnd[0] == 0 {
				// support  >= key queries
				creq.RangeEnd = string([]byte{})
			}
			// 权限校验
			if !sws.isWatchPermitted(creq) { // 当前请求 权限不允许
				wr := &pb.WatchResponse{
					Header:       sws.newResponseHeader(sws.watchStream.Rev()),
					WatchId:      creq.WatchId,
					Canceled:     true,
					Created:      true,
					CancelReason: rpctypes.ErrGRPCPermissionDenied.Error(),
				}

				select {
				case sws.ctrlStream <- wr:
					continue
				case <-sws.closec:
					return nil
				}
			}

			filters := FiltersFromRequest(creq) // server端  从watch请求中 获取一些过滤调价

			wsrev := sws.watchStream.Rev() // 获取当前kv的修订版本
			rev := creq.StartRevision      // 监听从哪个修订版本之后的变更，没穿就是当前
			if rev == 0 {
				rev = wsrev + 1
			}
			id, err := sws.watchStream.Watch(mvcc.WatchID(creq.WatchId), []byte(creq.Key), []byte(creq.RangeEnd), rev, filters...)
			if err == nil {
				sws.mu.Lock()
				if creq.ProgressNotify { // 默认FALSE
					sws.progress[id] = true
				}
				if creq.PrevKv { // 默认FALSE
					sws.prevKV[id] = true
				}
				if creq.Fragment { // 拆分大的事件
					sws.fragment[id] = true
				}
				sws.mu.Unlock()
			}
			wr := &pb.WatchResponse{
				Header:   sws.newResponseHeader(wsrev), //
				WatchId:  int64(id),
				Created:  true,
				Canceled: err != nil,
			}
			if err != nil {
				wr.CancelReason = err.Error()
			}
			select {
			case sws.ctrlStream <- wr:
			case <-sws.closec:
				return nil
			}
		}
		if req.WatchRequest_CancelRequest != nil {
			uv := &pb.WatchRequest_CancelRequest{}
			uv = req.WatchRequest_CancelRequest
			if uv.CancelRequest != nil {
				id := uv.CancelRequest.WatchId
				err := sws.watchStream.Cancel(mvcc.WatchID(id))
				if err == nil {
					sws.ctrlStream <- &pb.WatchResponse{
						Header:   sws.newResponseHeader(sws.watchStream.Rev()),
						WatchId:  id,
						Canceled: true,
					}
					sws.mu.Lock()
					delete(sws.progress, mvcc.WatchID(id))
					delete(sws.prevKV, mvcc.WatchID(id))
					delete(sws.fragment, mvcc.WatchID(id))
					sws.mu.Unlock()
				}
			}
		}
		if req.WatchRequest_ProgressRequest != nil {
			uv := &pb.WatchRequest_ProgressRequest{}
			uv = req.WatchRequest_ProgressRequest
			if uv.ProgressRequest != nil {
				sws.ctrlStream <- &pb.WatchResponse{
					Header:  sws.newResponseHeader(sws.watchStream.Rev()),
					WatchId: -1, // response is not associated with any WatchId and will be broadcast to all watch channels
				}
			}
		}
	}
}

// 往watch stream 发送消息
func (sws *serverWatchStream) sendLoop() {
	// watch ids that are currently active
	// 监视当前活动的id
	ids := make(map[mvcc.WatchID]struct{})
	// watch responses pending on a watch id creation message
	// 等待在一个手表id创建消息上的手表响应
	pending := make(map[mvcc.WatchID][]*pb.WatchResponse)

	interval := GetProgressReportInterval()
	progressTicker := time.NewTicker(interval)

	defer func() {
		progressTicker.Stop()
	}()

	for {
		select {
		case wresp, ok := <-sws.watchStream.Chan(): // //watchStream Channel中提取event发送
			if !ok {
				return
			}

			// TODO: evs is []mvccpb.Event type
			// either return []*mvccpb.Event from the mvcc package
			// or define protocol buffer with []mvccpb.Event.
			evs := wresp.Events
			events := make([]*mvccpb.Event, len(evs))
			sws.mu.RLock()
			needPrevKV := sws.prevKV[wresp.WatchID]
			sws.mu.RUnlock()
			for i := range evs {
				events[i] = &evs[i]
				if needPrevKV && !IsCreateEvent(evs[i]) {
					opt := mvcc.RangeOptions{Rev: evs[i].Kv.ModRevision - 1}
					r, err := sws.watchable.Range(context.TODO(), []byte(evs[i].Kv.Key), nil, opt)
					if err == nil && len(r.KVs) != 0 {
						events[i].PrevKv = &(r.KVs[0])
					}
				}
			}

			canceled := wresp.CompactRevision != 0
			wr := &pb.WatchResponse{
				Header:          sws.newResponseHeader(wresp.Revision),
				WatchId:         int64(wresp.WatchID),
				Events:          events,
				CompactRevision: wresp.CompactRevision,
				Canceled:        canceled,
			}

			if _, okID := ids[wresp.WatchID]; !okID {
				// buffer if id not yet announced
				wrs := append(pending[wresp.WatchID], wr)
				pending[wresp.WatchID] = wrs
				continue
			}

			sws.mu.RLock()
			fragmented, ok := sws.fragment[wresp.WatchID]
			sws.mu.RUnlock()

			var serr error
			if !fragmented && !ok {
				serr = sws.gRPCStream.Send(wr)
			} else {
				serr = sendFragments(wr, sws.maxRequestBytes, sws.gRPCStream.Send)
			}

			if serr != nil {
				if isClientCtxErr(sws.gRPCStream.Context().Err(), serr) {
					sws.lg.Debug("failed to send watch response to gRPC stream", zap.Error(serr))
				} else {
					sws.lg.Warn("failed to send watch response to gRPC stream", zap.Error(serr))
				}
				return
			}

			sws.mu.Lock()
			if len(evs) > 0 && sws.progress[wresp.WatchID] {
				// elide next progress update if sent a key update
				sws.progress[wresp.WatchID] = false
			}
			sws.mu.Unlock()

		case c, ok := <-sws.ctrlStream:
			if !ok {
				return
			}

			if err := sws.gRPCStream.Send(c); err != nil {
				if isClientCtxErr(sws.gRPCStream.Context().Err(), err) {
					sws.lg.Debug("failed to send watch control response to gRPC stream", zap.Error(err))
				} else {
					sws.lg.Warn("failed to send watch control response to gRPC stream", zap.Error(err))
				}
				return
			}

			// track id creation
			wid := mvcc.WatchID(c.WatchId)
			if c.Canceled {
				delete(ids, wid)
				continue
			}
			if c.Created {
				// flush buffered events
				ids[wid] = struct{}{}
				for _, v := range pending[wid] {
					if err := sws.gRPCStream.Send(v); err != nil {
						if isClientCtxErr(sws.gRPCStream.Context().Err(), err) {
							sws.lg.Debug("failed to send pending watch response to gRPC stream", zap.Error(err))
						} else {
							sws.lg.Warn("failed to send pending watch response to gRPC stream", zap.Error(err))
						}
						return
					}
				}
				delete(pending, wid)
			}

		case <-progressTicker.C:
			sws.mu.Lock()
			for id, ok := range sws.progress {
				if ok {
					sws.watchStream.RequestProgress(id)
				}
				sws.progress[id] = true
			}
			sws.mu.Unlock()

		case <-sws.closec:
			return
		}
	}
}

func IsCreateEvent(e mvccpb.Event) bool {
	return e.Type == mvccpb.PUT && e.Kv.CreateRevision == e.Kv.ModRevision
}

func sendFragments(wr *pb.WatchResponse, maxRequestBytes int, sendFunc func(*pb.WatchResponse) error) error {
	// no need to fragment if total request size is smaller
	// than max request limit or response contains only one event
	if wr.Size() < maxRequestBytes || len(wr.Events) < 2 {
		return sendFunc(wr)
	}

	ow := *wr
	ow.Events = make([]*mvccpb.Event, 0)
	ow.Fragment = true

	var idx int
	for {
		cur := ow
		for _, ev := range wr.Events[idx:] {
			cur.Events = append(cur.Events, ev)
			if len(cur.Events) > 1 && cur.Size() >= maxRequestBytes {
				cur.Events = cur.Events[:len(cur.Events)-1]
				break
			}
			idx++
		}
		if idx == len(wr.Events) {
			// last response has no more fragment
			cur.Fragment = false
		}
		if err := sendFunc(&cur); err != nil {
			return err
		}
		if !cur.Fragment {
			break
		}
	}
	return nil
}

func (sws *serverWatchStream) close() {
	sws.watchStream.Close()
	close(sws.closec)
	sws.wg.Wait()
}

func filterNoDelete(e mvccpb.Event) bool {
	return e.Type == mvccpb.DELETE
}

func filterNoPut(e mvccpb.Event) bool {
	return e.Type == mvccpb.PUT
}

// ----------------------------------------------  OVER ---------------------------------------------------------------------

// NewWatchServer OK
func NewWatchServer(s *etcdserver.EtcdServer) pb.WatchServer {
	srv := &watchServer{
		lg:              s.Cfg.Logger,
		clusterID:       int64(s.Cluster().ID()),
		memberID:        int64(s.ID()),
		maxRequestBytes: int(s.Cfg.MaxRequestBytes + grpcOverheadBytes),
		sg:              s,
		watchable:       s.Watchable(),
		ag:              s,
	}
	if srv.lg == nil {
		srv.lg = zap.NewNop()
	}
	if s.Cfg.WatchProgressNotifyInterval > 0 {
		if s.Cfg.WatchProgressNotifyInterval < minWatchProgressInterval {
			srv.lg.Warn("将watch 进度通知时间间隔调整为最小周期", zap.Duration("min-watch-progress-notify-interval", minWatchProgressInterval))
			s.Cfg.WatchProgressNotifyInterval = minWatchProgressInterval
		}
		SetProgressReportInterval(s.Cfg.WatchProgressNotifyInterval)
	}
	return srv
}

// GetProgressReportInterval OK
func GetProgressReportInterval() time.Duration {
	progressReportIntervalMu.RLock()
	interval := progressReportInterval
	progressReportIntervalMu.RUnlock()

	// add rand(1/10*progressReportInterval) as jitter so that etcdserver will not
	// send progress notifications to watchers around the same time even when watchers
	// are created around the same time (which is common when a client restarts itself).
	jitter := time.Duration(rand.Int63n(int64(interval) / 10))

	return interval + jitter
}

func FiltersFromRequest(creq *pb.WatchCreateRequest) []mvcc.FilterFunc {
	filters := make([]mvcc.FilterFunc, 0, len(creq.Filters))
	for _, ft := range creq.Filters {
		switch ft {
		case pb.WatchCreateRequest_NOPUT:
			filters = append(filters, filterNoPut)
		case pb.WatchCreateRequest_NODELETE:
			filters = append(filters, filterNoDelete)
		default:
		}
	}
	return filters
}

// 当前的修订版本
func (sws *serverWatchStream) newResponseHeader(rev int64) *pb.ResponseHeader {
	return &pb.ResponseHeader{
		ClusterId: uint64(sws.clusterID),
		MemberId:  uint64(sws.memberID),
		Revision:  rev,
		RaftTerm:  sws.sg.Term(),
	}
}
