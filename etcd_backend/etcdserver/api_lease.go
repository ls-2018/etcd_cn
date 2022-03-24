package etcdserver

import (
	"github.com/ls-2018/etcd_cn/etcd_backend/lease/leasehttp"
	"net/http"
)

func (s *EtcdServer) LeaseHandler() http.Handler {
	if s.lessor == nil {
		return nil
	}
	return leasehttp.NewHandler(s.lessor, s.ApplyWait)
}
func (s *EtcdServer) ApplyWait() <-chan struct{} { return s.applyWait.Wait(s.getCommittedIndex()) }
