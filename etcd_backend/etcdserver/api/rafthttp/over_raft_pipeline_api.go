package rafthttp

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
	pioutil "github.com/ls-2018/etcd_cn/pkg/ioutil"
	"github.com/ls-2018/etcd_cn/raft/raftpb"
	"go.uber.org/zap"
)

// newPipelineHandler
func newPipelineHandler(t *Transport, r Raft, cid types.ID) http.Handler {
	h := &pipelineHandler{
		lg:      t.Logger,
		localID: t.ID,
		tr:      t,
		r:       r,
		cid:     cid,
	}
	if h.lg == nil {
		h.lg = zap.NewNop()
	}
	return h
}

type pipelineHandler struct {
	lg      *zap.Logger
	localID types.ID
	tr      Transporter
	r       Raft
	cid     types.ID
}

func (h *pipelineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	limitedr := pioutil.NewLimitedBufferReader(r.Body, connReadLimitByte) // 限制返回的数据大小  64K
	b, err := ioutil.ReadAll(limitedr)
	if err != nil {
		h.lg.Warn("读取raft消息失败", zap.String("local-member-id", h.localID.String()), zap.Error(err))
		http.Error(w, "读取raft消息失败", http.StatusBadRequest)
		return
	}

	var m raftpb.Message
	if err := m.Unmarshal(b); err != nil {
		h.lg.Warn("发序列化raft消息失败", zap.String("local-member-id", h.localID.String()), zap.Error(err))
		http.Error(w, "发序列化raft消息失败", http.StatusBadRequest)
		return
	}

	if err := h.r.Process(context.TODO(), m); err != nil {
		switch v := err.(type) {
		case writerToResponse:
			v.WriteTo(w)
		default:
			h.lg.Warn("处理raft消息错误", zap.String("local-member-id", h.localID.String()), zap.Error(err))
			http.Error(w, "处理raft消息错误", http.StatusInternalServerError)
			w.(http.Flusher).Flush()
			// 断开http流的连接
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
