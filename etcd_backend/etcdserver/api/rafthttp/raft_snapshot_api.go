package rafthttp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
	"github.com/ls-2018/etcd_cn/etcd_backend/etcdserver/api/snap"
	"github.com/ls-2018/etcd_cn/raft/raftpb"
	"go.uber.org/zap"
)

type snapshotHandler struct {
	lg          *zap.Logger
	tr          Transporter
	r           Raft
	snapshotter *snap.Snapshotter

	localID types.ID
	cid     types.ID
}

func newSnapshotHandler(t *Transport, r Raft, snapshotter *snap.Snapshotter, cid types.ID) http.Handler {
	h := &snapshotHandler{
		lg:          t.Logger,
		tr:          t,
		r:           r,
		snapshotter: snapshotter,
		localID:     t.ID,
		cid:         cid,
	}
	if h.lg == nil {
		h.lg = zap.NewNop()
	}
	return h
}

// ServeHTTP serves HTTP request to receive and process snapshot message.
//
// If request sender dies without closing underlying TCP connection,
// the handler will keep waiting for the request body until TCP keepalive
// finds out that the connection is broken after several minutes.
// This is acceptable because
// 1. snapshot messages sent through other TCP connections could still be
// received and processed.
// 2. this case should happen rarely, so no further optimization is done.
func (h *snapshotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("X-Etcd-Cluster-ID", h.cid.String())

	if err := checkClusterCompatibilityFromHeader(h.lg, h.localID, r.Header, h.cid); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	addRemoteFromRequest(h.tr, r)

	dec := &messageDecoder{r: r.Body}
	// let snapshots be very large since they can exceed 512MB for large installations
	m, err := dec.decodeLimit(snapshotLimitByte)
	from := types.ID(m.From).String()
	if err != nil {
		msg := fmt.Sprintf("failed to decode raft message (%v)", err)
		h.lg.Warn(
			"failed to decode Raft message",
			zap.String("local-member-id", h.localID.String()),
			zap.String("remote-snapshot-sender-id", from),
			zap.Error(err),
		)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	msgSize := m.Size()

	if m.Type != raftpb.MsgSnap {
		h.lg.Warn(
			"unexpected Raft message type",
			zap.String("local-member-id", h.localID.String()),
			zap.String("remote-snapshot-sender-id", from),
			zap.String("message-type", m.Type.String()),
		)
		http.Error(w, "wrong raft message type", http.StatusBadRequest)
		return
	}

	h.lg.Info(
		"receiving database snapshot",
		zap.String("local-member-id", h.localID.String()),
		zap.String("remote-snapshot-sender-id", from),
		zap.Uint64("incoming-snapshot-index", m.Snapshot.Metadata.Index),
		zap.Int("incoming-snapshot-message-size-bytes", msgSize),
		zap.String("incoming-snapshot-message-size", humanize.Bytes(uint64(msgSize))),
	)

	// save incoming database snapshot.

	n, err := h.snapshotter.SaveDBFrom(r.Body, m.Snapshot.Metadata.Index)
	if err != nil {
		msg := fmt.Sprintf("failed to save KV snapshot (%v)", err)
		h.lg.Warn(
			"failed to save incoming database snapshot",
			zap.String("local-member-id", h.localID.String()),
			zap.String("remote-snapshot-sender-id", from),
			zap.Uint64("incoming-snapshot-index", m.Snapshot.Metadata.Index),
			zap.Error(err),
		)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	downloadTook := time.Since(start)
	h.lg.Info(
		"接受并保存数据库快照",
		zap.String("local-member-id", h.localID.String()),
		zap.String("remote-snapshot-sender-id", from),
		zap.Uint64("incoming-snapshot-index", m.Snapshot.Metadata.Index),
		zap.Int64("incoming-snapshot-size-bytes", n),
		zap.String("incoming-snapshot-size", humanize.Bytes(uint64(n))),
		zap.String("download-took", downloadTook.String()),
	)

	if err := h.r.Process(context.TODO(), m); err != nil {
		switch v := err.(type) {
		// Process may return writerToResponse error when doing some
		// additional checks before calling raft.RaftNodeInterFace.Step.
		case writerToResponse:
			v.WriteTo(w)
		default:
			msg := fmt.Sprintf("failed to process raft message (%v)", err)
			h.lg.Warn(
				"failed to process Raft message",
				zap.String("local-member-id", h.localID.String()),
				zap.String("remote-snapshot-sender-id", from),
				zap.Error(err),
			)
			http.Error(w, msg, http.StatusInternalServerError)
		}
		return
	}

	// Write StatusNoContent header after the message has been processed by
	// raft, which facilitates the client to report MsgSnap status.
	w.WriteHeader(http.StatusNoContent)
}
