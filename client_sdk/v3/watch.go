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

package clientv3

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ls-2018/etcd_cn/offical/api/v3/mvccpb"
	v3rpc "github.com/ls-2018/etcd_cn/offical/api/v3/v3rpc/rpctypes"
	pb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	EventTypeDelete = mvccpb.DELETE
	EventTypePut    = mvccpb.PUT

	closeSendErrTimeout = 250 * time.Millisecond
)

type Event mvccpb.Event

type WatchChan <-chan WatchResponse

type Watcher interface {
	Watch(ctx context.Context, key string, opts ...OpOption) WatchChan
	// RequestProgress requests a progress notify response be sent in all watch channels.
	RequestProgress(ctx context.Context) error
	// Close closes the watcher and cancels all watch requests.
	Close() error
}

type WatchResponse struct {
	Header pb.ResponseHeader
	Events []*Event

	// CompactRevision is the minimum revision the watcher may receive.
	CompactRevision int64

	CreatedRevision int64

	// Canceled is used to indicate watch failure.
	// If the watch failed and the stream was about to close, before the channel is closed,
	// the channel sends a final response that has Canceled set to true with a non-nil Err().
	Canceled bool

	// Created is used to indicate the creation of the watcher.
	Created bool

	closeErr error

	// cancelReason is a reason of canceling watch
	cancelReason string
}

// IsCreate returns true if the event tells that the key is newly created.
func (e *Event) IsCreate() bool {
	return e.Type == EventTypePut && e.Kv.CreateRevision == e.Kv.ModRevision
}

// IsModify returns true if the event tells that a new value is put on existing key.
func (e *Event) IsModify() bool {
	return e.Type == EventTypePut && e.Kv.CreateRevision != e.Kv.ModRevision
}

// Err is the error value if this WatchResponse holds an error.
func (wr *WatchResponse) Err() error {
	switch {
	case wr.closeErr != nil:
		return v3rpc.Error(wr.closeErr)
	case wr.CompactRevision != 0:
		return v3rpc.ErrCompacted
	case wr.Canceled:
		if len(wr.cancelReason) != 0 {
			return v3rpc.Error(status.Error(codes.FailedPrecondition, wr.cancelReason))
		}
		return v3rpc.ErrFutureRev
	}
	return nil
}

// IsProgressNotify returns true if the WatchResponse is progress notification.
func (wr *WatchResponse) IsProgressNotify() bool {
	return len(wr.Events) == 0 && !wr.Canceled && !wr.Created && wr.CompactRevision == 0 && wr.Header.Revision != 0
}

type watcher struct {
	remote   pb.WatchClient              // ?????????????????????????????????
	callOpts []grpc.CallOption           //
	mu       sync.Mutex                  //
	streams  map[string]*watchGrpcStream // ??????CTX ???????????????????????????GRPC???.
	lg       *zap.Logger                 //
}

// watchGrpcStream tracks all watch resources attached to a single grpc stream.
type watchGrpcStream struct {
	owner      *watcher
	remote     pb.WatchClient
	callOpts   []grpc.CallOption
	ctx        context.Context //  remote.Watch requests
	ctxKey     string          // ctxKey ??????????????????????????????
	cancel     context.CancelFunc
	substreams map[int64]*watcherStream // ????????? grpc ????????????????????????watchers
	resuming   []*watcherStream         // ??????????????? grpc ???????????????????????????????????????
	reqc       chan watchStreamRequest  // reqc ??? Watch() ??????????????????????????????
	respc      chan *pb.WatchResponse   // respc ??? watch ?????????????????????
	donec      chan struct{}            // donec ????????????????????????
	errc       chan error
	closingc   chan *watcherStream // ????????????????????????????????????
	wg         sync.WaitGroup      // ??????????????? goroutine ????????????,wg ??????
	resumec    chan struct{}       // resumec ?????????????????????????????????????????????
	closeErr   error               // closeErr ???????????????????????????
	lg         *zap.Logger
}

// watchStreamRequest is a union of the supported watch request operation types
type watchStreamRequest interface {
	toPB() *pb.WatchRequest
}

type watchRequest struct {
	ctx            context.Context
	key            string
	end            string
	rev            int64
	createdNotify  bool // ??????????????????true,??????????????????????????????
	progressNotify bool // ????????????
	fragment       bool // ??????????????????,??????????????????
	filters        []pb.WatchCreateRequest_FilterType
	prevKV         bool
	retc           chan chan WatchResponse
}

// progressRequest is issued by the subscriber to request watch progress
type progressRequest struct{}

// watcherStream ????????????????????????
// watch()???,??????watchgrpcstream????????????watcherStream,??????????????????watch rpc??????,??????????????????key,??????key????????????,??????????????????.
type watcherStream struct {
	initReq watchRequest        // initReq ??????????????????????????????
	outc    chan WatchResponse  // outc ??????????????????watch??????
	recvc   chan *WatchResponse // recvc buffers watch responses before publishing
	donec   chan struct{}       // ??? watcherStream goroutine ????????? donec ??????
	closing bool                // ???????????????????????????,closures ????????? true.
	id      int64               // id ?????? grpc ??????????????? watch id
	buf     []*WatchResponse    // buf ????????? etcd ????????????????????????????????????????????????
}

func NewWatcher(c *Client) Watcher {
	return NewWatchFromWatchClient(pb.NewWatchClient(c.conn), c)
}

// NewWatchFromWatchClient watch?????????,??????????????????
func NewWatchFromWatchClient(wc pb.WatchClient, c *Client) Watcher {
	w := &watcher{
		remote:  wc,
		streams: make(map[string]*watchGrpcStream),
	}
	if c != nil {
		w.callOpts = c.callOpts
		w.lg = c.lg
	}
	return w
}

// never closes
var valCtxCh = make(chan struct{})
var zeroTime = time.Unix(0, 0)

// ctx with only the values; never Done
type valCtx struct{ context.Context }

func (vc *valCtx) Deadline() (time.Time, bool) { return zeroTime, false }
func (vc *valCtx) Done() <-chan struct{}       { return valCtxCh }
func (vc *valCtx) Err() error                  { return nil }

// ??????????????????  gRPC??????,????????????serverWatchStream.recvLoop()
func (w *watcher) newWatcherGrpcStream(inctx context.Context) *watchGrpcStream {
	ctx, cancel := context.WithCancel(&valCtx{inctx})
	wgs := &watchGrpcStream{
		owner:      w,
		remote:     w.remote,
		callOpts:   w.callOpts,
		ctx:        ctx,
		ctxKey:     streamKeyFromCtx(inctx),
		cancel:     cancel,
		substreams: make(map[int64]*watcherStream),
		respc:      make(chan *pb.WatchResponse),
		reqc:       make(chan watchStreamRequest),
		donec:      make(chan struct{}),
		errc:       make(chan error, 1),
		closingc:   make(chan *watcherStream),
		resumec:    make(chan struct{}),
		lg:         w.lg,
	}
	go wgs.run()
	return wgs
}

// Watch ??????watch??????,??????????????????
func (w *watcher) Watch(ctx context.Context, key string, opts ...OpOption) WatchChan {
	ow := opWatch(key, opts...) // ??????watch??????

	var filters []pb.WatchCreateRequest_FilterType
	if ow.filterPut {
		filters = append(filters, pb.WatchCreateRequest_NOPUT)
	}
	if ow.filterDelete {
		filters = append(filters, pb.WatchCreateRequest_NODELETE)
	}

	wr := &watchRequest{
		ctx:            ctx,
		createdNotify:  ow.createdNotify,
		key:            ow.key,
		end:            ow.end,
		rev:            ow.rev,
		progressNotify: ow.progressNotify,
		fragment:       ow.fragment,
		filters:        filters,
		prevKV:         ow.prevKV,
		retc:           make(chan chan WatchResponse, 1),
	}

	ok := false
	ctxKey := streamKeyFromCtx(ctx) // map[hasleader:[true]]

	var closeCh chan WatchResponse
	for {
		// ????????????????????????GRPC??????       ????????????
		w.mu.Lock()
		if w.streams == nil { // ??????????????????
			// closed
			w.mu.Unlock()
			ch := make(chan WatchResponse)
			close(ch)
			return ch
		}
		// streams?????????map,??????????????? ctx ?????????????????? grpc ???
		// ?????????????????????????????????,?????????
		wgs := w.streams[ctxKey]
		if wgs == nil {
			// newWatcherGrpcStream new??????watch grpc stream?????????watch??????
			// ??????goroutine???????????????key???watch????????????
			wgs = w.newWatcherGrpcStream(ctx) // ???????????????watch???
			w.streams[ctxKey] = wgs
		}
		donec := wgs.donec
		reqc := wgs.reqc
		w.mu.Unlock()

		// couldn't create channel; return closed channel
		if closeCh == nil {
			closeCh = make(chan WatchResponse, 1)
		}

		// submit request
		select {
		case reqc <- wr: // reqc ??? Watch() ??????????????????????????????
			ok = true
		case <-wr.ctx.Done():
			ok = false
		case <-donec:
			ok = false
			if wgs.closeErr != nil {
				closeCh <- WatchResponse{Canceled: true, closeErr: wgs.closeErr}
				break
			}
			// retry; may have dropped stream from no ctxs
			continue
		}

		// receive channel
		if ok {
			select {
			case ret := <-wr.retc:
				return ret
			case <-ctx.Done():
			case <-donec:
				if wgs.closeErr != nil {
					closeCh <- WatchResponse{Canceled: true, closeErr: wgs.closeErr}
					break
				}
				// retry; may have dropped stream from no ctxs
				continue
			}
		}
		break
	}

	close(closeCh)
	return closeCh
}

func (w *watcher) Close() (err error) {
	w.mu.Lock()
	streams := w.streams
	w.streams = nil
	w.mu.Unlock()
	for _, wgs := range streams {
		if werr := wgs.close(); werr != nil {
			err = werr
		}
	}
	// Consider context.Canceled as a successful close
	if err == context.Canceled {
		err = nil
	}
	return err
}

// RequestProgress requests a progress notify response be sent in all watch channels.
func (w *watcher) RequestProgress(ctx context.Context) (err error) {
	ctxKey := streamKeyFromCtx(ctx)

	w.mu.Lock()
	if w.streams == nil {
		w.mu.Unlock()
		return fmt.Errorf("no stream found for context")
	}
	wgs := w.streams[ctxKey]
	if wgs == nil {
		wgs = w.newWatcherGrpcStream(ctx) // ???????????????watch???
		w.streams[ctxKey] = wgs
	}
	donec := wgs.donec
	reqc := wgs.reqc
	w.mu.Unlock()

	pr := &progressRequest{}

	select {
	case reqc <- pr:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-donec:
		if wgs.closeErr != nil {
			return wgs.closeErr
		}
		// retry; may have dropped stream from no ctxs
		return w.RequestProgress(ctx)
	}
}

func (w *watchGrpcStream) close() (err error) {
	w.cancel()
	<-w.donec
	select {
	case err = <-w.errc:
	default:
	}
	return toErr(w.ctx, err)
}

func (w *watcher) closeStream(wgs *watchGrpcStream) {
	w.mu.Lock()
	close(wgs.donec)
	wgs.cancel()
	if w.streams != nil {
		delete(w.streams, wgs.ctxKey)
	}
	w.mu.Unlock()
}

func (w *watchGrpcStream) addSubstream(resp *pb.WatchResponse, ws *watcherStream) {
	// check watch ID for backward compatibility (<= v3.3)
	if resp.WatchId == -1 || (resp.Canceled && resp.CancelReason != "") {
		w.closeErr = v3rpc.Error(errors.New(resp.CancelReason))
		// failed; no channel
		close(ws.recvc)
		return
	}
	ws.id = resp.WatchId
	w.substreams[ws.id] = ws
}

func (w *watchGrpcStream) sendCloseSubstream(ws *watcherStream, resp *WatchResponse) {
	select {
	case ws.outc <- *resp:
	case <-ws.initReq.ctx.Done():
	case <-time.After(closeSendErrTimeout):
	}
	close(ws.outc)
}

func (w *watchGrpcStream) closeSubstream(ws *watcherStream) {
	// send channel response in case stream was never established
	select {
	case ws.initReq.retc <- ws.outc:
	default:
	}
	// close subscriber's channel
	if closeErr := w.closeErr; closeErr != nil && ws.initReq.ctx.Err() == nil {
		go w.sendCloseSubstream(ws, &WatchResponse{Canceled: true, closeErr: w.closeErr})
	} else if ws.outc != nil {
		close(ws.outc)
	}
	if ws.id != -1 {
		delete(w.substreams, ws.id)
		return
	}
	for i := range w.resuming {
		if w.resuming[i] == ws {
			w.resuming[i] = nil
			return
		}
	}
}

// run ????????????watch client
func (w *watchGrpcStream) run() {
	var wc pb.Watch_WatchClient
	var closeErr error

	// ?????????????????????,???goroutine????????????;?????????????????????recvc???GRPC?????????
	closing := make(map[*watcherStream]struct{})

	defer func() {
		w.closeErr = closeErr
		// ???????????????????????????
		for _, ws := range w.substreams {
			if _, ok := closing[ws]; !ok {
				close(ws.recvc)
				closing[ws] = struct{}{}
			}
		}
		for _, ws := range w.resuming {
			if _, ok := closing[ws]; ws != nil && !ok {
				close(ws.recvc)
				closing[ws] = struct{}{}
			}
		}
		w.joinSubstreams()
		for range closing {
			w.closeSubstream(<-w.closingc)
		}
		w.wg.Wait()
		w.owner.closeStream(w)
	}()

	// ???etcd??????grpc???
	if wc, closeErr = w.newWatchClient(); closeErr != nil {
		return
	}

	cancelSet := make(map[int64]struct{})

	var cur *pb.WatchResponse
	for {
		select {
		// Watch() requested
		case req := <-w.reqc:
			switch wreq := req.(type) {
			case *watchRequest:
				outc := make(chan WatchResponse, 1)
				// TODO: pass custom watch ID?
				ws := &watcherStream{
					initReq: *wreq,
					id:      -1,
					outc:    outc,
					// unbuffered so resumes won't cause repeat events
					recvc: make(chan *WatchResponse),
				}

				ws.donec = make(chan struct{})
				w.wg.Add(1)
				go w.serveSubstream(ws, w.resumec)

				// queue up for watcher creation/resume
				w.resuming = append(w.resuming, ws)
				if len(w.resuming) == 1 {
					// head of resume queue, can register a new watcher
					if err := wc.Send(ws.initReq.toPB()); err != nil {
						w.lg.Debug("error when sending request", zap.Error(err))
					}
				}
			case *progressRequest:
				if err := wc.Send(wreq.toPB()); err != nil {
					w.lg.Debug("error when sending request", zap.Error(err))
				}
			}

		case pbresp := <-w.respc: // ??????watch client????????????
			if cur == nil || pbresp.Created || pbresp.Canceled {
				cur = pbresp
			} else if cur != nil && cur.WatchId == pbresp.WatchId {
				// ???????????????
				cur.Events = append(cur.Events, pbresp.Events...)
				// update "Fragment" field; last response with "Fragment" == false
				cur.Fragment = pbresp.Fragment
			}

			switch {
			case pbresp.Created: // ????????????????????????
				// response to head of queue creation
				if len(w.resuming) != 0 {
					if ws := w.resuming[0]; ws != nil {
						w.addSubstream(pbresp, ws)
						w.dispatchEvent(pbresp)
						w.resuming[0] = nil
					}
				}

				if ws := w.nextResume(); ws != nil {
					if err := wc.Send(ws.initReq.toPB()); err != nil {
						w.lg.Debug("error when sending request", zap.Error(err))
					}
				}

				// reset for next iteration
				cur = nil

			case pbresp.Canceled && pbresp.CompactRevision == 0:
				delete(cancelSet, pbresp.WatchId)
				if ws, ok := w.substreams[pbresp.WatchId]; ok {
					// signal to stream goroutine to update closingc
					close(ws.recvc)
					closing[ws] = struct{}{}
				}

				// reset for next iteration
				cur = nil

			case cur.Fragment: // ???????????????????????????,????????????????????????,??????????????????????????????
				continue

			default:
				// dispatch to appropriate watch stream
				ok := w.dispatchEvent(cur)

				// reset for next iteration
				cur = nil

				if ok {
					break
				}

				// watch response on unexpected watch id; cancel id
				if _, ok := cancelSet[pbresp.WatchId]; ok {
					break
				}

				cancelSet[pbresp.WatchId] = struct{}{}
				cr := &pb.WatchRequest_CancelRequest{
					CancelRequest: &pb.WatchCancelRequest{
						WatchId: pbresp.WatchId,
					},
				}
				req := &pb.WatchRequest{WatchRequest_CancelRequest: cr}
				w.lg.Debug("sending watch cancel request for failed dispatch", zap.Int64("watch-id", pbresp.WatchId))
				if err := wc.Send(req); err != nil {
					w.lg.Debug("failed to send watch cancel request", zap.Int64("watch-id", pbresp.WatchId), zap.Error(err))
				}
			}

		// ??????client Recv??????.????????????,???????????????,??????????????????watch??????
		// ????????????watch????????????,?????????watch client??????????????????

		case err := <-w.errc:
			if isHaltErr(w.ctx, err) || toErr(w.ctx, err) == v3rpc.ErrNoLeader {
				closeErr = err
				return
			}
			if wc, closeErr = w.newWatchClient(); closeErr != nil {
				return
			}
			if ws := w.nextResume(); ws != nil {
				if err := wc.Send(ws.initReq.toPB()); err != nil {
					w.lg.Debug("error when sending request", zap.Error(err))
				}
			}
			cancelSet = make(map[int64]struct{})

		case <-w.ctx.Done():
			return

		case ws := <-w.closingc:
			w.closeSubstream(ws)
			delete(closing, ws)
			// no more watchers on this stream, shutdown, skip cancellation
			if len(w.substreams)+len(w.resuming) == 0 {
				return
			}
			if ws.id != -1 {
				// client is closing an established watch; close it on the etcd proactively instead of waiting
				// to close when the next message arrives
				cancelSet[ws.id] = struct{}{}
				cr := &pb.WatchRequest_CancelRequest{
					CancelRequest: &pb.WatchCancelRequest{
						WatchId: ws.id,
					},
				}
				req := &pb.WatchRequest{
					WatchRequest_CancelRequest: cr,
				}
				w.lg.Debug("sending watch cancel request for closed watcher", zap.Int64("watch-id", ws.id))
				if err := wc.Send(req); err != nil {
					w.lg.Debug("failed to send watch cancel request", zap.Int64("watch-id", ws.id), zap.Error(err))
				}
			}
		}
	}
}

// nextResume chooses the next resuming to register with the grpc stream. Abandoned
// streams are marked as nil in the queue since the head must wait for its inflight registration.
func (w *watchGrpcStream) nextResume() *watcherStream {
	for len(w.resuming) != 0 {
		if w.resuming[0] != nil {
			return w.resuming[0]
		}
		w.resuming = w.resuming[1:len(w.resuming)]
	}
	return nil
}

// dispatchEvent sends a WatchResponse to the appropriate watcher stream
func (w *watchGrpcStream) dispatchEvent(pbresp *pb.WatchResponse) bool {
	events := make([]*Event, len(pbresp.Events))
	for i, ev := range pbresp.Events {
		events[i] = (*Event)(ev)
	}
	// TODO: return watch ID?
	wr := &WatchResponse{
		Header:          *pbresp.Header,
		Events:          events,
		CompactRevision: pbresp.CompactRevision,
		Created:         pbresp.Created,
		Canceled:        pbresp.Canceled,
		cancelReason:    pbresp.CancelReason,
	}

	// watch IDs are zero indexed, so request notify watch responses are assigned a watch ID of -1 to
	// indicate they should be broadcast.
	if wr.IsProgressNotify() && pbresp.WatchId == -1 {
		return w.broadcastResponse(wr)
	}

	return w.unicastResponse(wr, pbresp.WatchId)
}

// broadcastResponse send a watch response to all watch substreams.
func (w *watchGrpcStream) broadcastResponse(wr *WatchResponse) bool {
	for _, ws := range w.substreams {
		select {
		case ws.recvc <- wr:
		case <-ws.donec:
		}
	}
	return true
}

// unicastResponse sends a watch response to a specific watch substream.
func (w *watchGrpcStream) unicastResponse(wr *WatchResponse, watchId int64) bool {
	ws, ok := w.substreams[watchId]
	if !ok {
		return false
	}
	select {
	case ws.recvc <- wr:
	case <-ws.donec:
		return false
	}
	return true
}

// serveWatchClient forwards messages from the grpc stream to run()
func (w *watchGrpcStream) serveWatchClient(wc pb.Watch_WatchClient) {
	for {
		resp, err := wc.Recv()
		if err != nil {
			select {
			case w.errc <- err:
			case <-w.donec:
			}
			return
		}
		select {
		case w.respc <- resp:
		case <-w.donec:
			return
		}
	}
}

// serveSubstream forwards watch responses from run() to the subscriber
func (w *watchGrpcStream) serveSubstream(ws *watcherStream, resumec chan struct{}) {
	if ws.closing {
		panic("created substream goroutine but substream is closing")
	}

	// nextRev is the minimum expected next revision
	nextRev := ws.initReq.rev
	resuming := false
	defer func() {
		if !resuming {
			ws.closing = true
		}
		close(ws.donec)
		if !resuming {
			w.closingc <- ws
		}
		w.wg.Done()
	}()

	emptyWr := &WatchResponse{}
	for {
		curWr := emptyWr
		outc := ws.outc

		if len(ws.buf) > 0 {
			curWr = ws.buf[0]
		} else {
			outc = nil
		}
		select {
		case outc <- *curWr:
			if ws.buf[0].Err() != nil {
				return
			}
			ws.buf[0] = nil
			ws.buf = ws.buf[1:]
		case wr, ok := <-ws.recvc:
			if !ok {
				// shutdown from closeSubstream
				return
			}

			if wr.Created {
				if ws.initReq.retc != nil {
					ws.initReq.retc <- ws.outc
					// to prevent next write from taking the slot in buffered channel
					// and posting duplicate create events
					ws.initReq.retc = nil

					// send first creation event only if requested
					if ws.initReq.createdNotify {
						ws.outc <- *wr
					}
					// once the watch channel is returned, a current revision
					// watch must resume at the store revision. This is necessary
					// for the following case to work as expected:
					//	wch := m1.Watch("a")
					//	m2.Put("a", "b")
					//	<-wch
					// If the revision is only bound on the first observed event,
					// if wch is disconnected before the Put is issued, then reconnects
					// after it is committed, it'll miss the Put.
					if ws.initReq.rev == 0 {
						nextRev = wr.Header.Revision
					}
				}
			} else {
				// current progress of watch; <= store revision
				nextRev = wr.Header.Revision
			}

			if len(wr.Events) > 0 {
				nextRev = wr.Events[len(wr.Events)-1].Kv.ModRevision + 1
			}
			ws.initReq.rev = nextRev

			// created event is already sent above,
			// watcher should not post duplicate events
			if wr.Created {
				continue
			}

			// TODO pause channel if buffer gets too large
			ws.buf = append(ws.buf, wr)
		case <-w.ctx.Done():
			return
		case <-ws.initReq.ctx.Done():
			return
		case <-resumec:
			resuming = true
			return
		}
	}
	// lazily send cancel message if events on missing id
}

func (w *watchGrpcStream) newWatchClient() (pb.Watch_WatchClient, error) {
	// ??????????????????????????????
	close(w.resumec)
	w.resumec = make(chan struct{})
	w.joinSubstreams()
	for _, ws := range w.substreams {
		ws.id = -1
		w.resuming = append(w.resuming, ws)
	}
	// strip out nils, if any
	var resuming []*watcherStream
	for _, ws := range w.resuming {
		if ws != nil {
			resuming = append(resuming, ws)
		}
	}
	w.resuming = resuming
	w.substreams = make(map[int64]*watcherStream)

	// connect to grpc stream while accepting watcher cancelation
	stopc := make(chan struct{})
	donec := w.waitCancelSubstreams(stopc)
	wc, err := w.openWatchClient()
	close(stopc)
	<-donec

	// serve all non-closing streams, even if there's a client error
	// so that the teardown path can shutdown the streams as expected.
	for _, ws := range w.resuming {
		if ws.closing {
			continue
		}
		ws.donec = make(chan struct{})
		w.wg.Add(1)
		go w.serveSubstream(ws, w.resumec)
	}

	if err != nil {
		return nil, v3rpc.Error(err)
	}

	// receive data from new grpc stream
	go w.serveWatchClient(wc)
	return wc, nil
}

func (w *watchGrpcStream) waitCancelSubstreams(stopc <-chan struct{}) <-chan struct{} {
	var wg sync.WaitGroup
	wg.Add(len(w.resuming))
	donec := make(chan struct{})
	for i := range w.resuming {
		go func(ws *watcherStream) {
			defer wg.Done()
			if ws.closing {
				if ws.initReq.ctx.Err() != nil && ws.outc != nil {
					close(ws.outc)
					ws.outc = nil
				}
				return
			}
			select {
			case <-ws.initReq.ctx.Done():
				// closed ws will be removed from resuming
				ws.closing = true
				close(ws.outc)
				ws.outc = nil
				w.wg.Add(1)
				go func() {
					defer w.wg.Done()
					w.closingc <- ws
				}()
			case <-stopc:
			}
		}(w.resuming[i])
	}
	go func() {
		defer close(donec)
		wg.Wait()
	}()
	return donec
}

// joinSubstreams ????????????sub stream??????
func (w *watchGrpcStream) joinSubstreams() {
	for _, ws := range w.substreams {
		<-ws.donec
	}
	for _, ws := range w.resuming {
		if ws != nil {
			<-ws.donec
		}
	}
}

var maxBackoff = 100 * time.Millisecond

// openWatchClient retries opening a watch client until success or halt.
// manually retry in case "ws==nil && err==nil"
// TODO: remove FailFast=false
func (w *watchGrpcStream) openWatchClient() (ws pb.Watch_WatchClient, err error) {
	backoff := time.Millisecond
	for {
		select {
		case <-w.ctx.Done():
			if err == nil {
				return nil, w.ctx.Err()
			}
			return nil, err
		default:
		}
		if ws, err = w.remote.Watch(w.ctx, w.callOpts...); ws != nil && err == nil {
			break
		}
		if isHaltErr(w.ctx, err) {
			return nil, v3rpc.Error(err)
		}
		if isUnavailableErr(w.ctx, err) {
			// retry, but backoff
			if backoff < maxBackoff {
				// 25% backoff factor
				backoff = backoff + backoff/4
				if backoff > maxBackoff {
					backoff = maxBackoff
				}
			}
			time.Sleep(backoff)
		}
	}
	return ws, nil
}

// toPB converts an internal watch request structure to its protobuf WatchRequest structure.
func (wr *watchRequest) toPB() *pb.WatchRequest {
	req := &pb.WatchCreateRequest{
		StartRevision:  wr.rev,
		Key:            string([]byte(wr.key)),
		RangeEnd:       string([]byte(wr.end)),
		ProgressNotify: wr.progressNotify,
		Filters:        wr.filters,
		PrevKv:         wr.prevKV,
		Fragment:       wr.fragment,
	}
	cr := &pb.WatchRequest_CreateRequest{CreateRequest: req}
	return &pb.WatchRequest{WatchRequest_CreateRequest: cr}
}

// toPB converts an internal progress request structure to its protobuf WatchRequest structure.
func (pr *progressRequest) toPB() *pb.WatchRequest {
	req := &pb.WatchProgressRequest{}
	cr := &pb.WatchRequest_ProgressRequest{ProgressRequest: req}
	return &pb.WatchRequest{WatchRequest_ProgressRequest: cr}
}

// ???ctx?????????str
func streamKeyFromCtx(ctx context.Context) string {
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		return fmt.Sprintf("%+v", md)
	}
	return ""
}
