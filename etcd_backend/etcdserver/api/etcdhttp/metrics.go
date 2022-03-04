// Copyright 2017 The etcd Authors
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

package etcdhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ls-2018/etcd_cn/etcd_backend/auth"
	"github.com/ls-2018/etcd_cn/etcd_backend/etcdserver"
	"github.com/ls-2018/etcd_cn/raft"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"go.uber.org/zap"
	"net/http"
)

const (
	PathMetrics      = "/metrics"
	PathHealth       = "/health"
	PathProxyMetrics = "/proxy/metrics"
	PathProxyHealth  = "/proxy/health"
)

// HandleMetricsHealthForV3 注册指标、健康检查
func HandleMetricsHealthForV3(lg *zap.Logger, mux *http.ServeMux, srv *etcdserver.EtcdServer) {
	mux.Handle(PathMetrics, promhttp.Handler())
	mux.Handle(PathHealth, NewHealthHandler(lg, func(excludedAlarms AlarmSet) Health { return checkV3Health(lg, srv, excludedAlarms) }))
}

// HandlePrometheus 注册指标
func HandlePrometheus(mux *http.ServeMux) {
	mux.Handle(PathMetrics, promhttp.Handler())
}

// NewHealthHandler OK
func NewHealthHandler(lg *zap.Logger, hfunc func(excludedAlarms AlarmSet) Health) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			lg.Warn("/health error", zap.Int("status-code", http.StatusMethodNotAllowed))
			return
		}
		excludedAlarms := getExcludedAlarms(r) // 排除的警告
		h := hfunc(excludedAlarms)
		defer func() {
			if h.Health == "true" {
				healthSuccess.Inc()
			} else {
				healthFailed.Inc()
			}
		}()
		d, _ := json.Marshal(h)
		if h.Health != "true" {
			http.Error(w, string(d), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(d)
	}
}

var (
	healthSuccess = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "etcd",
		Subsystem: "etcd",
		Name:      "health_success",
		Help:      "调用健康检查成功次数",
	})
	healthFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "etcd",
		Subsystem: "etcd",
		Name:      "health_failures",
		Help:      "调用健康检查失败次数",
	})
)

func init() {
	prometheus.MustRegister(healthSuccess)
	prometheus.MustRegister(healthFailed)
}

// Health 健康检查的状态
type Health struct {
	Health string `json:"health"`
	Reason string `json:"reason"`
}

type AlarmSet map[string]struct{}

// 排除特定警告
func getExcludedAlarms(r *http.Request) (alarms AlarmSet) {
	alarms = make(map[string]struct{}, 2)
	alms, found := r.URL.Query()["exclude"]
	if found {
		for _, alm := range alms {
			if len(alm) == 0 {
				continue
			}
			alarms[alm] = struct{}{}
		}
	}
	return alarms
}

func checkHealth(lg *zap.Logger, srv etcdserver.ServerV2, excludedAlarms AlarmSet) Health {
	h := Health{}
	h.Health = "true"
	as := srv.Alarms()
	if len(as) > 0 {
		for _, v := range as {
			alarmName := v.Alarm.String()
			if _, found := excludedAlarms[alarmName]; found {
				lg.Debug("/health 忽略警告", zap.String("alarm", v.String()))
				continue
			}
			h.Health = "false"
			switch v.Alarm {
			case etcdserverpb.AlarmType_NOSPACE:
				h.Reason = "警报 NOSPACE"
			case etcdserverpb.AlarmType_CORRUPT:
				h.Reason = "警报 CORRUPT"
			default:
				h.Reason = "警报 UNKNOWN"
			}
			lg.Warn("/health 不健康;警报", zap.String("alarm", v.String()))
			return h
		}
	}

	if uint64(srv.Leader()) == raft.None {
		h.Health = "false"
		h.Reason = "RAFT没有leader"
		lg.Warn("/health不健康;没有leader")
		return h
	}
	return h
}

func checkV3Health(lg *zap.Logger, srv *etcdserver.EtcdServer, excludedAlarms AlarmSet) (h Health) {
	if h = checkHealth(lg, srv, excludedAlarms); h.Health != "true" {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), srv.Cfg.ReqTimeout())
	_, err := srv.Range(ctx, &etcdserverpb.RangeRequest{KeysOnly: true, Limit: 1})
	cancel()
	if err != nil && err != auth.ErrUserEmpty && err != auth.ErrPermissionDenied {
		h.Health = "false"
		h.Reason = fmt.Sprintf("RANGE 失败:%s", err)
		return
	}
	return
}
