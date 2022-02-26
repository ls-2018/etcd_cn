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

// Every change should be reflected on help.go as well.

package etcdmain

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/ls-2018/client/pkg/logutil"
	cconfig "github.com/ls-2018/etcd/config"
	"github.com/ls-2018/etcd/embed"
	"github.com/ls-2018/etcd/etcdserver/api/rafthttp"
	"github.com/ls-2018/pkg/flags"
	"go.etcd.io/etcd/api/v3/version"

	"go.uber.org/zap"
	"sigs.k8s.io/yaml"
)

var (
	proxyFlagOff      = "off"
	proxyFlagReadonly = "readonly"
	proxyFlagOn       = "on"

	fallbackFlagExit  = "exit"
	fallbackFlagProxy = "proxy"

	ignored = []string{
		"cluster-active-size",
		"cluster-remove-delay",
		"cluster-sync-interval",
		"config",
		"force",
		"max-result-buffer",
		"max-retry-attempts",
		"peer-heartbeat-interval",
		"peer-election-timeout",
		"retry-interval",
		"snapshot",
		"v",
		"vv",
		// for coverage testing
		"test.coverprofile",
		"test.outputdir",
	}
)

type configProxy struct {
	ProxyFailureWaitMs     uint `json:"proxy-failure-wait"`
	ProxyRefreshIntervalMs uint `json:"proxy-refresh-interval"`
	ProxyDialTimeoutMs     uint `json:"proxy-dial-timeout"`
	ProxyWriteTimeoutMs    uint `json:"proxy-write-timeout"`
	ProxyReadTimeoutMs     uint `json:"proxy-read-timeout"`
	Fallback               string
	Proxy                  string
	ProxyJSON              string `json:"proxy"`
	FallbackJSON           string `json:"discovery-fallback"`
}

// config 保存etcd命令行调用的配置
type config struct {
	ec           embed.Config
	cp           configProxy
	cf           configFlags
	configFile   string // 从文件加载服务器配置。
	printVersion bool
	ignored      []string
}

// configFlags 是否有一组标志用于命令行解析配置
type configFlags struct {
	flagSet       *flag.FlagSet
	clusterState  *flags.SelectiveStringValue
	fallback      *flags.SelectiveStringValue
	proxy         *flags.SelectiveStringValue
	v2deprecation *flags.SelectiveStringsValue
}

func newConfig() *config {
	cfg := &config{
		ec: *embed.NewConfig(),
		cp: configProxy{
			Proxy:                  proxyFlagOff,
			ProxyFailureWaitMs:     5000,
			ProxyRefreshIntervalMs: 30000,
			ProxyDialTimeoutMs:     1000,
			ProxyWriteTimeoutMs:    5000,
		},
		ignored: ignored,
	}
	cfg.cf = configFlags{
		flagSet: flag.NewFlagSet("etcd", flag.ContinueOnError),
		clusterState: flags.NewSelectiveStringValue(
			embed.ClusterStateFlagNew,
			embed.ClusterStateFlagExisting,
		),
		fallback: flags.NewSelectiveStringValue(
			fallbackFlagProxy,
			fallbackFlagExit,
		),
		proxy: flags.NewSelectiveStringValue(
			proxyFlagOff,
			proxyFlagReadonly,
			proxyFlagOn,
		),
		v2deprecation: flags.NewSelectiveStringsValue(
			string(cconfig.V2_DEPR_0_NOT_YET),
			string(cconfig.V2_DEPR_1_WRITE_ONLY),
			string(cconfig.V2_DEPR_1_WRITE_ONLY_DROP),
			string(cconfig.V2_DEPR_2_GONE)),
	}

	fs := cfg.cf.flagSet
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, usageline)
	}

	fs.StringVar(&cfg.configFile, "config-file", "", "从文件加载服务器配置。")

	// member
	fs.StringVar(&cfg.ec.Dir, "data-dir", cfg.ec.Dir, "Path to the data directory.")
	fs.StringVar(&cfg.ec.WalDir, "wal-dir", cfg.ec.WalDir, "Path to the dedicated wal directory.")
	fs.Var(
		flags.NewUniqueURLsWithExceptions(embed.DefaultListenPeerURLs, ""),
		"listen-peer-urls",
		"List of URLs to listen on for peer traffic.",
	)
	fs.Var(
		flags.NewUniqueURLsWithExceptions(embed.DefaultListenClientURLs, ""), "listen-client-urls",
		"List of URLs to listen on for client traffic.",
	)
	fs.Var(flags.NewUniqueURLsWithExceptions("", ""), "listen-metrics-urls", "要监听指标和运行状况端点的url列表。")
	fs.UintVar(&cfg.ec.MaxSnapFiles, "max-snapshots", cfg.ec.MaxSnapFiles, "Maximum number of snapshot files to retain (0 is unlimited).")
	fs.UintVar(&cfg.ec.MaxWalFiles, "max-wals", cfg.ec.MaxWalFiles, "Maximum number of wal files to retain (0 is unlimited).")
	fs.StringVar(&cfg.ec.Name, "name", cfg.ec.Name, "Human-readable name for this member.")
	fs.Uint64Var(&cfg.ec.SnapshotCount, "snapshot-count", cfg.ec.SnapshotCount, "Number of committed transactions to trigger a snapshot to disk.")
	fs.UintVar(&cfg.ec.TickMs, "heartbeat-interval", cfg.ec.TickMs, "Time (in milliseconds) of a heartbeat interval.")
	fs.UintVar(&cfg.ec.ElectionMs, "election-timeout", cfg.ec.ElectionMs, "Time (in milliseconds) for an election to timeout.")
	fs.BoolVar(&cfg.ec.InitialElectionTickAdvance, "initial-election-tick-advance", cfg.ec.InitialElectionTickAdvance, "Whether to fast-forward initial election ticks on boot for faster election.")
	fs.Int64Var(&cfg.ec.QuotaBackendBytes, "quota-backend-bytes", cfg.ec.QuotaBackendBytes, "Raise alarms when backend size exceeds the given quota. 0 means use the default quota.")
	fs.StringVar(&cfg.ec.BackendFreelistType, "backend-bbolt-freelist-type", cfg.ec.BackendFreelistType, "BackendFreelistType specifies the type of freelist that boltdb backend uses(array and map are supported types)")
	fs.DurationVar(&cfg.ec.BackendBatchInterval, "backend-batch-interval", cfg.ec.BackendBatchInterval, "BackendBatchInterval is the maximum time before commit the backend transaction.")
	fs.IntVar(&cfg.ec.BackendBatchLimit, "backend-batch-limit", cfg.ec.BackendBatchLimit, "BackendBatchLimit is the maximum operations before commit the backend transaction.")
	fs.UintVar(&cfg.ec.MaxTxnOps, "max-txn-ops", cfg.ec.MaxTxnOps, "Maximum number of operations permitted in a transaction.")
	fs.UintVar(&cfg.ec.MaxRequestBytes, "max-request-bytes", cfg.ec.MaxRequestBytes, "Maximum client request size in bytes the etcd will accept.")
	fs.DurationVar(&cfg.ec.GRPCKeepAliveMinTime, "grpc-keepalive-min-time", cfg.ec.GRPCKeepAliveMinTime, "Minimum interval duration that a client should wait before pinging etcd.")
	fs.DurationVar(&cfg.ec.GRPCKeepAliveInterval, "grpc-keepalive-interval", cfg.ec.GRPCKeepAliveInterval, "Frequency duration of etcd-to-client ping to check if a connection is alive (0 to disable).")
	fs.DurationVar(&cfg.ec.GRPCKeepAliveTimeout, "grpc-keepalive-timeout", cfg.ec.GRPCKeepAliveTimeout, "Additional duration of wait before closing a non-responsive connection (0 to disable).")
	fs.BoolVar(&cfg.ec.SocketOpts.ReusePort, "socket-reuse-port", cfg.ec.SocketOpts.ReusePort, "Enable to set socket option SO_REUSEPORT on listeners allowing rebinding of a port already in use.")
	fs.BoolVar(&cfg.ec.SocketOpts.ReuseAddress, "socket-reuse-address", cfg.ec.SocketOpts.ReuseAddress, "Enable to set socket option SO_REUSEADDR on listeners allowing binding to an address in `TIME_WAIT` state.")

	// raft 连接超时
	fs.DurationVar(&rafthttp.ConnReadTimeout, "raft-read-timeout", rafthttp.DefaultConnReadTimeout, "Read timeout set on each rafthttp connection")
	fs.DurationVar(&rafthttp.ConnWriteTimeout, "raft-write-timeout", rafthttp.DefaultConnWriteTimeout, "Write timeout set on each rafthttp connection")

	// 集群
	fs.Var(
		flags.NewUniqueURLsWithExceptions(embed.DefaultInitialAdvertisePeerURLs, ""),
		"initial-advertise-peer-urls",
		"List of this member's peer URLs to advertise to the rest of the cluster.",
	)
	fs.Var(
		flags.NewUniqueURLsWithExceptions(embed.DefaultAdvertiseClientURLs, ""),
		"advertise-client-urls",
		"List of this member's client URLs to advertise to the public.",
	)
	fs.StringVar(&cfg.ec.Durl, "discovery", cfg.ec.Durl, "Discovery URL used to bootstrap the cluster.")
	fs.Var(cfg.cf.fallback, "discovery-fallback", fmt.Sprintf("Valid values include %q", cfg.cf.fallback.Valids()))

	fs.StringVar(&cfg.ec.Dproxy, "discovery-proxy", cfg.ec.Dproxy, "HTTP proxy to use for traffic to discovery service.")
	fs.StringVar(&cfg.ec.DNSCluster, "discovery-srv", cfg.ec.DNSCluster, "DNS domain used to bootstrap initial cluster.")
	fs.StringVar(&cfg.ec.DNSClusterServiceName, "discovery-srv-name", cfg.ec.DNSClusterServiceName, "Service name to query when using DNS discovery.")
	fs.StringVar(&cfg.ec.InitialCluster, "initial-cluster", cfg.ec.InitialCluster, "Initial cluster configuration for bootstrapping.")
	fs.StringVar(&cfg.ec.InitialClusterToken, "initial-cluster-token", cfg.ec.InitialClusterToken, "Initial cluster token for the etcd cluster during bootstrap.")
	fs.Var(cfg.cf.clusterState, "initial-cluster-state", "Initial cluster state ('new' or 'existing').")

	fs.BoolVar(&cfg.ec.StrictReconfigCheck, "strict-reconfig-check", cfg.ec.StrictReconfigCheck, "Reject reconfiguration requests that would cause quorum loss.")

	fs.BoolVar(&cfg.ec.PreVote, "pre-vote", cfg.ec.PreVote, "Enable to run an additional Raft election phase.")

	fs.BoolVar(&cfg.ec.EnableV2, "enable-v2", cfg.ec.EnableV2, "Accept etcd V2 client requests. Deprecated in v3.5. Will be decommission in v3.6.")
	fs.StringVar(&cfg.ec.ExperimentalEnableV2V3, "experimental-enable-v2v3", cfg.ec.ExperimentalEnableV2V3, "v3 prefix for serving emulated v2 state. Deprecated in 3.5. Will be decomissioned in 3.6.")
	fs.Var(cfg.cf.v2deprecation, "v2-deprecation", fmt.Sprintf("v2store deprecation stage: %q. ", cfg.cf.proxy.Valids()))

	// proxy
	fs.Var(cfg.cf.proxy, "proxy", fmt.Sprintf("Valid values include %q", cfg.cf.proxy.Valids()))
	fs.UintVar(&cfg.cp.ProxyFailureWaitMs, "proxy-failure-wait", cfg.cp.ProxyFailureWaitMs, "Time (in milliseconds) an endpoint will be held in a failed state.")
	fs.UintVar(&cfg.cp.ProxyRefreshIntervalMs, "proxy-refresh-interval", cfg.cp.ProxyRefreshIntervalMs, "Time (in milliseconds) of the endpoints refresh interval.")
	fs.UintVar(&cfg.cp.ProxyDialTimeoutMs, "proxy-dial-timeout", cfg.cp.ProxyDialTimeoutMs, "Time (in milliseconds) for a dial to timeout.")
	fs.UintVar(&cfg.cp.ProxyWriteTimeoutMs, "proxy-write-timeout", cfg.cp.ProxyWriteTimeoutMs, "Time (in milliseconds) for a write to timeout.")
	fs.UintVar(&cfg.cp.ProxyReadTimeoutMs, "proxy-read-timeout", cfg.cp.ProxyReadTimeoutMs, "Time (in milliseconds) for a read to timeout.")

	// etcdctl通信的证书配置
	fs.StringVar(&cfg.ec.ClientTLSInfo.CertFile, "cert-file", "", "客户端证书")
	fs.StringVar(&cfg.ec.ClientTLSInfo.KeyFile, "key-file", "", "客户端私钥")

	fs.StringVar(&cfg.ec.ClientTLSInfo.ClientCertFile, "client-cert-file", "", "验证client客户端时使用的 证书文件路径,否则在需要客户认证时将使用cert-file文件")
	fs.StringVar(&cfg.ec.ClientTLSInfo.ClientKeyFile, "client-key-file", "", "验证client客户端时使用的 密钥文件路径,否则在需要客户认证时将使用key-file文件。")
	fs.BoolVar(&cfg.ec.ClientTLSInfo.ClientCertAuth, "client-cert-auth", false, "启用客户端证书验证;默认false")
	fs.StringVar(&cfg.ec.ClientTLSInfo.CRLFile, "client-crl-file", "", "客户端证书吊销列表文件的路径。")
	fs.StringVar(&cfg.ec.ClientTLSInfo.AllowedHostname, "client-cert-allowed-hostname", "", "允许客户端证书认证使用TLS主机名。")
	fs.StringVar(&cfg.ec.ClientTLSInfo.TrustedCAFile, "trusted-ca-file", "", "客户端etcd通信 的可信CA证书文件")
	fs.BoolVar(&cfg.ec.ClientAutoTLS, "auto-tls", false, "客户端TLS使用生成的证书")
	// etcd通信之间的证书配置
	fs.StringVar(&cfg.ec.PeerTLSInfo.CertFile, "peer-cert-file", "", "证书路径")
	fs.StringVar(&cfg.ec.PeerTLSInfo.KeyFile, "peer-key-file", "", "私钥路径")

	fs.StringVar(&cfg.ec.PeerTLSInfo.ClientCertFile, "peer-client-cert-file", "", "验证server客户端时使用的 证书文件路径,否则在需要客户认证时将使用cert-file文件")
	fs.StringVar(&cfg.ec.PeerTLSInfo.ClientKeyFile, "peer-client-key-file", "", "验证server客户端时使用的 密钥文件路径,否则在需要客户认证时将使用key-file文件。")

	fs.BoolVar(&cfg.ec.PeerTLSInfo.ClientCertAuth, "peer-client-cert-auth", false, "启用server客户端证书验证;默认false")
	fs.StringVar(&cfg.ec.PeerTLSInfo.TrustedCAFile, "peer-trusted-ca-file", "", "服务器端ca证书")
	fs.BoolVar(&cfg.ec.PeerAutoTLS, "peer-auto-tls", false, "节点之间使用生成的证书通信;默认false")
	fs.UintVar(&cfg.ec.SelfSignedCertValidity, "self-signed-cert-validity", 1, "客户端证书和同级证书的有效期,单位为年")
	fs.StringVar(&cfg.ec.PeerTLSInfo.CRLFile, "peer-crl-file", "", "服务端证书吊销列表文件的路径。")
	fs.StringVar(&cfg.ec.PeerTLSInfo.AllowedCN, "peer-cert-allowed-cn", "", "允许的server客户端证书CommonName")
	fs.StringVar(&cfg.ec.PeerTLSInfo.AllowedHostname, "peer-cert-allowed-hostname", "", "允许的server客户端证书hostname")
	fs.Var(flags.NewStringsValue(""), "cipher-suites", "客户端/etcds之间支持的TLS加密套件的逗号分隔列表(空将由Go自动填充)。")
	fs.BoolVar(&cfg.ec.PeerTLSInfo.SkipClientSANVerify, "experimental-peer-skip-client-san-verification", false, "跳过server 客户端证书中SAN字段的验证。默认false")

	fs.Var(flags.NewUniqueURLsWithExceptions("*", "*"), "cors", "逗号分隔的CORS白名单，或跨来源资源共享，(空或*表示允许所有)")
	fs.Var(flags.NewUniqueStringsValue("*"), "host-whitelist", "如果etcd是不安全的(空意味着允许所有)，用逗号分隔HTTP客户端请求中的可接受主机名。")

	// 日志
	fs.StringVar(&cfg.ec.Logger, "logger", "zap", "当前只支持zap,结构化数据")
	fs.Var(flags.NewUniqueStringsValue(embed.DefaultLogOutput), "log-outputs", "指定'stdout'或'stderr'以跳过日志记录,即使在systemd或逗号分隔的输出目标列表下运行也是如此。")
	fs.StringVar(&cfg.ec.LogLevel, "log-level", logutil.DefaultLogLevel, "日志等级,只支持 debug, info, warn, error, panic, or fatal. Default 'info'.")
	fs.BoolVar(&cfg.ec.EnableLogRotation, "enable-log-rotation", false, "Enable log rotation of a single log-outputs file target.")
	fs.StringVar(&cfg.ec.LogRotationConfigJSON, "log-rotation-config-json", embed.DefaultLogRotationConfig, "Configures log rotation if enabled with a JSON logger config. Default: MaxSize=100(MB), MaxAge=0(days,no limit), MaxBackups=0(no limit), LocalTime=false(UTC), Compress=false(gzip)")

	// 版本
	fs.BoolVar(&cfg.printVersion, "version", false, "打印版本并退出。")

	fs.StringVar(&cfg.ec.AutoCompactionRetention, "auto-compaction-retention", "0", "Auto compaction retention for mvcc key value store. 0 means disable auto compaction.")
	fs.StringVar(&cfg.ec.AutoCompactionMode, "auto-compaction-mode", "periodic", "interpret 'auto-compaction-retention' one of: periodic|revision. 'periodic' for duration based retention, defaulting to hours if no time unit is provided (e.g. '5m'). 'revision' for revision number based retention.")

	// 性能分析器 通过 HTTP
	fs.BoolVar(&cfg.ec.EnablePprof, "enable-pprof", false, "通过HTTP服务器启用运行时分析数据。地址位于客户端URL +“/ debug / pprof /”")

	// additional metrics
	fs.StringVar(&cfg.ec.Metrics, "metrics", cfg.ec.Metrics, "设置导出的指标的详细程度,指定“扩展”以包括直方图指标")

	// experimental distributed tracing
	fs.BoolVar(&cfg.ec.ExperimentalEnableDistributedTracing, "experimental-enable-distributed-tracing", false, "Enable experimental distributed  tracing using OpenTelemetry Tracing.")
	fs.StringVar(&cfg.ec.ExperimentalDistributedTracingAddress, "experimental-distributed-tracing-address", embed.ExperimentalDistributedTracingAddress, "Address for distributed tracing used for OpenTelemetry Tracing (if enabled with experimental-enable-distributed-tracing flag).")
	fs.StringVar(&cfg.ec.ExperimentalDistributedTracingServiceName, "experimental-distributed-tracing-service-name", embed.ExperimentalDistributedTracingServiceName, "Configures service name for distributed tracing to be used to define service name for OpenTelemetry Tracing (if enabled with experimental-enable-distributed-tracing flag). 'etcd' is the default service name. Use the same service name for all instances of etcd.")
	fs.StringVar(&cfg.ec.ExperimentalDistributedTracingServiceInstanceID, "experimental-distributed-tracing-instance-id", "", "Configures service instance ID for distributed tracing to be used to define service instance ID key for OpenTelemetry Tracing (if enabled with experimental-enable-distributed-tracing flag). There is no default value set. This ID must be unique per etcd instance.")

	// auth
	fs.StringVar(&cfg.ec.AuthToken, "auth-token", cfg.ec.AuthToken, "Specify auth token specific options.")
	fs.UintVar(&cfg.ec.BcryptCost, "bcrypt-cost", cfg.ec.BcryptCost, "Specify bcrypt algorithm cost factor for auth password hashing.")
	fs.UintVar(&cfg.ec.AuthTokenTTL, "auth-token-ttl", cfg.ec.AuthTokenTTL, "The lifetime in seconds of the auth token.")

	// gateway
	fs.BoolVar(&cfg.ec.EnableGRPCGateway, "enable-grpc-gateway", cfg.ec.EnableGRPCGateway, "Enable GRPC gateway.")

	// experimental
	fs.BoolVar(&cfg.ec.ExperimentalInitialCorruptCheck, "experimental-initial-corrupt-check", cfg.ec.ExperimentalInitialCorruptCheck, "Enable to check data corruption before serving any client/peer traffic.")
	fs.DurationVar(&cfg.ec.ExperimentalCorruptCheckTime, "experimental-corrupt-check-time", cfg.ec.ExperimentalCorruptCheckTime, "Duration of time between cluster corruption check passes.")

	fs.BoolVar(&cfg.ec.ExperimentalEnableLeaseCheckpoint, "experimental-enable-lease-checkpoint", false, "Enable leader to send regular checkpoints to other members to prevent reset of remaining TTL on leader change.")
	// TODO: delete in v3.7
	fs.BoolVar(&cfg.ec.ExperimentalEnableLeaseCheckpointPersist, "experimental-enable-lease-checkpoint-persist", false, "Enable persisting remainingTTL to prevent indefinite auto-renewal of long lived leases. Always enabled in v3.6. Should be used to ensure smooth upgrade from v3.5 clusters with this feature enabled. Requires experimental-enable-lease-checkpoint to be enabled.")
	fs.IntVar(&cfg.ec.ExperimentalCompactionBatchLimit, "experimental-compaction-batch-limit", cfg.ec.ExperimentalCompactionBatchLimit, "Sets the maximum revisions deleted in each compaction batch.")
	fs.DurationVar(&cfg.ec.ExperimentalWatchProgressNotifyInterval, "experimental-watch-progress-notify-interval", cfg.ec.ExperimentalWatchProgressNotifyInterval, "Duration of periodic watch progress notifications.")
	fs.DurationVar(&cfg.ec.ExperimentalDowngradeCheckTime, "experimental-downgrade-check-time", cfg.ec.ExperimentalDowngradeCheckTime, "Duration of time between two downgrade status check.")
	fs.DurationVar(&cfg.ec.ExperimentalWarningApplyDuration, "experimental-warning-apply-duration", cfg.ec.ExperimentalWarningApplyDuration, "Time duration after which a warning is generated if request takes more time.")
	fs.BoolVar(&cfg.ec.ExperimentalMemoryMlock, "experimental-memory-mlock", cfg.ec.ExperimentalMemoryMlock, "Enable to enforce etcd pages (in particular bbolt) to stay in RAM.")
	fs.BoolVar(&cfg.ec.ExperimentalTxnModeWriteWithSharedBuffer, "experimental-txn-mode-write-with-shared-buffer", true, "Enable the write transaction to use a shared buffer in its readonly check operations.")
	fs.UintVar(&cfg.ec.ExperimentalBootstrapDefragThresholdMegabytes, "experimental-bootstrap-defrag-threshold-megabytes", 0, "Enable the defrag during etcd etcd bootstrap on condition that it will free at least the provided threshold of disk space. Needs to be set to non-zero value to take effect.")

	// 非安全
	fs.BoolVar(&cfg.ec.UnsafeNoFsync, "unsafe-no-fsync", false, "Disables fsync, unsafe, will cause data loss.")
	fs.BoolVar(&cfg.ec.ForceNewCluster, "force-new-cluster", false, "强制创建新的单成员群集。它提交配置更改,强制删除集群中的所有现有成员并添加自身。需要将其设置为还原备份。")

	// ignored
	for _, f := range cfg.ignored {
		fs.Var(&flags.IgnoredFlag{Name: f}, f, "")
	}
	return cfg
}

func (cfg *config) parse(arguments []string) error {
	perr := cfg.cf.flagSet.Parse(arguments)
	switch perr {
	case nil:
	case flag.ErrHelp:
		fmt.Println(flagsline)
		os.Exit(0)
	default:
		os.Exit(2)
	}
	if len(cfg.cf.flagSet.Args()) != 0 {
		return fmt.Errorf("'%s' is not a valid flag", cfg.cf.flagSet.Arg(0))
	}

	if cfg.printVersion {
		fmt.Printf("etcd Version: %s\n", version.Version)
		fmt.Printf("Git SHA: %s\n", version.GitSHA)
		fmt.Printf("Go Version: %s\n", runtime.Version())
		fmt.Printf("Go OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	var err error

	// This env variable must be parsed separately
	// because we need to determine whether to use or
	// ignore the env variables based on if the config file is set.
	if cfg.configFile == "" {
		cfg.configFile = os.Getenv(flags.FlagToEnv("ETCD", "config-file"))
	}

	if cfg.configFile != "" {
		err = cfg.configFromFile(cfg.configFile)
		if lg := cfg.ec.GetLogger(); lg != nil {
			lg.Info(
				"loaded etcd configuration, other configuration command line flags and environment variables will be ignored if provided",
				zap.String("path", cfg.configFile),
			)
		}
	} else {
		err = cfg.configFromCmdLine()
	}

	if cfg.ec.V2Deprecation == "" {
		cfg.ec.V2Deprecation = cconfig.V2_DEPR_DEFAULT
	}

	// now logger is set up
	return err
}

func (cfg *config) configFromCmdLine() error {
	// user-specified logger is not setup yet, use this logger during flag parsing
	lg, err := zap.NewProduction()
	if err != nil {
		return err
	}

	verKey := "ETCD_VERSION"
	if verVal := os.Getenv(verKey); verVal != "" {
		// unset to avoid any possible side-effect.
		os.Unsetenv(verKey)

		lg.Warn(
			"cannot set special environment variable",
			zap.String("key", verKey),
			zap.String("value", verVal),
		)
	}

	err = flags.SetFlagsFromEnv(lg, "ETCD", cfg.cf.flagSet)
	if err != nil {
		return err
	}

	if rafthttp.ConnReadTimeout < rafthttp.DefaultConnReadTimeout {
		rafthttp.ConnReadTimeout = rafthttp.DefaultConnReadTimeout
		lg.Info(fmt.Sprintf("raft-read-timeout increased to minimum value: %v", rafthttp.DefaultConnReadTimeout))
	}
	if rafthttp.ConnWriteTimeout < rafthttp.DefaultConnWriteTimeout {
		rafthttp.ConnWriteTimeout = rafthttp.DefaultConnWriteTimeout
		lg.Info(fmt.Sprintf("raft-write-timeout increased to minimum value: %v", rafthttp.DefaultConnWriteTimeout))
	}

	cfg.ec.LPUrls = flags.UniqueURLsFromFlag(cfg.cf.flagSet, "listen-peer-urls")
	cfg.ec.APUrls = flags.UniqueURLsFromFlag(cfg.cf.flagSet, "initial-advertise-peer-urls")
	cfg.ec.LCUrls = flags.UniqueURLsFromFlag(cfg.cf.flagSet, "listen-client-urls")
	cfg.ec.ACUrls = flags.UniqueURLsFromFlag(cfg.cf.flagSet, "advertise-client-urls")
	cfg.ec.ListenMetricsUrls = flags.UniqueURLsFromFlag(cfg.cf.flagSet, "listen-metrics-urls")

	cfg.ec.CORS = flags.UniqueURLsMapFromFlag(cfg.cf.flagSet, "cors")
	cfg.ec.HostWhitelist = flags.UniqueStringsMapFromFlag(cfg.cf.flagSet, "host-whitelist")

	cfg.ec.CipherSuites = flags.StringsFromFlag(cfg.cf.flagSet, "cipher-suites")

	cfg.ec.LogOutputs = flags.UniqueStringsFromFlag(cfg.cf.flagSet, "log-outputs")

	cfg.ec.ClusterState = cfg.cf.clusterState.String()
	cfg.cp.Fallback = cfg.cf.fallback.String()
	cfg.cp.Proxy = cfg.cf.proxy.String()

	cfg.ec.V2Deprecation = cconfig.V2DeprecationEnum(cfg.cf.v2deprecation.String())

	// disable default advertise-client-urls if lcurls is set
	missingAC := flags.IsSet(cfg.cf.flagSet, "listen-client-urls") && !flags.IsSet(cfg.cf.flagSet, "advertise-client-urls")
	if !cfg.mayBeProxy() && missingAC {
		cfg.ec.ACUrls = nil
	}

	// disable default initial-cluster if discovery is set
	if (cfg.ec.Durl != "" || cfg.ec.DNSCluster != "" || cfg.ec.DNSClusterServiceName != "") && !flags.IsSet(cfg.cf.flagSet, "initial-cluster") {
		cfg.ec.InitialCluster = ""
	}

	return cfg.validate()
}

func (cfg *config) configFromFile(path string) error {
	eCfg, err := embed.ConfigFromFile(path)
	if err != nil {
		return err
	}
	cfg.ec = *eCfg

	// load extra config information
	b, rerr := ioutil.ReadFile(path)
	if rerr != nil {
		return rerr
	}
	if yerr := yaml.Unmarshal(b, &cfg.cp); yerr != nil {
		return yerr
	}

	if cfg.cp.FallbackJSON != "" {
		if err := cfg.cf.fallback.Set(cfg.cp.FallbackJSON); err != nil {
			log.Fatalf("unexpected error setting up discovery-fallback flag: %v", err)
		}
		cfg.cp.Fallback = cfg.cf.fallback.String()
	}

	if cfg.cp.ProxyJSON != "" {
		if err := cfg.cf.proxy.Set(cfg.cp.ProxyJSON); err != nil {
			log.Fatalf("unexpected error setting up proxyFlag: %v", err)
		}
		cfg.cp.Proxy = cfg.cf.proxy.String()
	}
	return nil
}

func (cfg *config) mayBeProxy() bool {
	mayFallbackToProxy := cfg.ec.Durl != "" && cfg.cp.Fallback == fallbackFlagProxy
	return cfg.cp.Proxy != proxyFlagOff || mayFallbackToProxy
}

func (cfg *config) validate() error {
	err := cfg.ec.Validate()
	// TODO(yichengq): check this for joining through discovery service case
	if err == embed.ErrUnsetAdvertiseClientURLsFlag && cfg.mayBeProxy() {
		return nil
	}
	return err
}

func (cfg config) isProxy() bool               { return cfg.cf.proxy.String() != proxyFlagOff }
func (cfg config) isReadonlyProxy() bool       { return cfg.cf.proxy.String() == proxyFlagReadonly }
func (cfg config) shouldFallbackToProxy() bool { return cfg.cf.fallback.String() == fallbackFlagProxy }
