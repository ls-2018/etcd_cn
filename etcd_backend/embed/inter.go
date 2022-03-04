package embed

import (
	gw "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ls-2018/etcd_cn/etcd_backend/etcdserver"
	"net/http"
)

func mux() {
	var _ etcdserver.EtcdServer
	var _ gw.ServeMux
	var _ http.ServeMux
	var _ http.Handler // ServeHTTP方法
}
