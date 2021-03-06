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

package etcdserver

import (
	"context"
	"encoding/json"
	"expvar"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ls-2018/etcd_cn/raft"

	"github.com/coreos/go-semver/semver"
	humanize "github.com/dustin/go-humanize"
	"github.com/ls-2018/etcd_cn/etcd/config"
	"go.uber.org/zap"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/fileutil"
	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
	"github.com/ls-2018/etcd_cn/etcd/auth"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/membership"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/rafthttp"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/snap"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/v2discovery"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/v2http/httptypes"
	stats "github.com/ls-2018/etcd_cn/etcd/etcdserver/api/v2stats"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/v2store"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/v3alarm"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/v3compactor"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/cindex"
	"github.com/ls-2018/etcd_cn/etcd/lease"
	"github.com/ls-2018/etcd_cn/etcd/mvcc"
	"github.com/ls-2018/etcd_cn/etcd/mvcc/backend"
	"github.com/ls-2018/etcd_cn/etcd/wal"
	"github.com/ls-2018/etcd_cn/offical/api/v3/version"
	pb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"
	"github.com/ls-2018/etcd_cn/pkg/idutil"
	"github.com/ls-2018/etcd_cn/pkg/pbutil"
	"github.com/ls-2018/etcd_cn/pkg/runtime"
	"github.com/ls-2018/etcd_cn/pkg/schedule"
	"github.com/ls-2018/etcd_cn/pkg/traceutil"
	"github.com/ls-2018/etcd_cn/pkg/wait"
	"github.com/ls-2018/etcd_cn/raft/raftpb"
)

const (
	DefaultSnapshotCount = 100000

	// DefaultSnapshotCatchUpEntries is the number of entries for a slow follower
	// to catch-up after compacting the raft storage entries.
	// We expect the follower has a millisecond level latency with the leader.
	// The max throughput is around 10K. Keep a 5K entries is enough for helping
	// follower to catch up.
	DefaultSnapshotCatchUpEntries uint64 = 5000

	StoreClusterPrefix = "/0"
	StoreKeysPrefix    = "/1"

	// HealthInterval is the minimum time the cluster should backend healthy
	// before accepting add member requests.
	HealthInterval = 5 * time.Second

	purgeFileInterval = 30 * time.Second

	// max number of in-flight snapshot messages etcdserver allows to have
	// This number is more than enough for most clusters with 5 machines.
	maxInFlightMsgSnap = 16

	releaseDelayAfterSnapshot = 30 * time.Second

	// maxPendingRevokes is the maximum number of outstanding expired lease revocations.
	maxPendingRevokes = 16

	recommendedMaxRequestBytes = 10 * 1024 * 1024 // 10M

	readyPercent = 0.9

	DowngradeEnabledPath = "/downgrade/enabled"
)

var (
	// monitorVersionInterval should backend smaller than the timeout
	// on the connection. Or we will not backend able to reuse the connection
	// (since it will timeout).
	monitorVersionInterval = rafthttp.ConnWriteTimeout - time.Second

	recommendedMaxRequestBytesString = humanize.Bytes(uint64(recommendedMaxRequestBytes))
	storeMemberAttributeRegexp       = regexp.MustCompile(path.Join(membership.StoreMembersPrefix, "[[:xdigit:]]{1,16}", "attributes"))
)

func init() {
	rand.Seed(time.Now().UnixNano())

	expvar.Publish(
		"file_descriptor_limit",
		expvar.Func(
			func() interface{} {
				n, _ := runtime.FDLimit()
				return n
			},
		),
	)
}

type Response struct {
	Term    uint64
	Index   uint64
	Event   *v2store.Event
	Watcher v2store.Watcher
	Err     error
}

type ServerV2 interface {
	Server
	Leader() types.ID
	// Do takes a V2 request and attempts to fulfill it, returning a Response.
	Do(ctx context.Context, r pb.Request) (Response, error)
	stats.Stats
	ClientCertAuthEnabled() bool
}

type ServerV3 interface {
	Server
	RaftStatusGetter
}

func (s *EtcdServer) ClientCertAuthEnabled() bool { return s.Cfg.ClientCertAuthEnabled }

type Server interface {
	AddMember(ctx context.Context, memb membership.Member) ([]*membership.Member, error)          // http  ????????????
	RemoveMember(ctx context.Context, id uint64) ([]*membership.Member, error)                    // http  ????????????
	UpdateMember(ctx context.Context, updateMemb membership.Member) ([]*membership.Member, error) // http  ????????????
	PromoteMember(ctx context.Context, id uint64) ([]*membership.Member, error)                   // http  ????????????
	ClusterVersion() *semver.Version                                                              //
	Cluster() api.Cluster                                                                         // ??????????????????cluster ?????????
	Alarms() []*pb.AlarmMember                                                                    //
	LeaderChangedNotify() <-chan struct{}                                                         // ?????????????????????
	// 1. ???????????????????????????,???????????????????????????.
	// 2. ??????,??????????????????????????????????????????.
	// 3. ?????????????????????????????????API????????????????????????????????????.
}

// EtcdServer ??????etcd????????????????????????,??????etcd?????????????????????????????????????????????.
type EtcdServer struct {
	inflightSnapshots int64  // ?????????????????????snapshot??????
	appliedIndex      uint64 // ??????apply?????????????????????index
	committedIndex    uint64 // ?????????????????????index,?????????leader??????????????????????????????????????????index
	term              uint64
	lead              uint64
	consistIndex      cindex.ConsistentIndexer // ??????????????????kvstore???index
	r                 raftNode                 // ?????????????????????,?????????raft??????????????????.
	readych           chan struct{}            // ?????????????????????????????????cluster,??????????????????.
	Cfg               config.ServerConfig      // ?????????
	lgMu              *sync.RWMutex
	lg                *zap.Logger
	w                 wait.Wait     // ??????????????????????????????????????????????????????????????????.
	readMu            sync.RWMutex  // ??????3???????????????????????????linearizable ????????????
	readwaitc         chan struct{} // ?????????readwaitC?????????????????????????????????etcd??????????????????????????????
	readNotifier      *notifier     // ????????????????????????read goroutine ??????????????????

	stop            chan struct{}           // ????????????
	stopping        chan struct{}           // ???????????????????????????
	done            chan struct{}           // etcd???start????????????????????????,?????????????????????
	leaderChanged   chan struct{}           // leader????????? ??????linearizable read loop   drop??????????????????
	leaderChangedMu sync.RWMutex            //
	errorc          chan error              // ????????????,?????????????????????????????????,??????raft?????????.
	id              types.ID                // etcd??????id
	attributes      membership.Attributes   // etcd????????????
	cluster         *membership.RaftCluster // ????????????
	v2store         v2store.Store           // v2???kv??????
	snapshotter     *snap.Snapshotter       // ??????snapshot
	applyV2         ApplierV2               // v2???applier,?????????commited index apply???raft?????????
	applyV3         applierV3               // v3???applier,?????????commited index apply???raft?????????
	applyV3Base     applierV3               // ?????????????????????????????????applyV3
	applyV3Internal applierV3Internal       // v3?????????applier
	applyWait       wait.WaitTime           // apply???????????????,????????????index?????????apply??????
	kv              mvcc.WatchableKV        // v3??????kv??????
	lessor          lease.Lessor            // v3???,???????????????????????????
	backendLock     sync.Mutex              // ????????????????????????,????????????????????????????????????????????????
	backend         backend.Backend         // ????????????  bolt.db
	beHooks         *backendHooks           // ????????????
	authStore       auth.AuthStore          // ??????????????????
	alarmStore      *v3alarm.AlarmStore     // ??????????????????
	stats           *stats.ServerStats      // ??????????????????
	lstats          *stats.LeaderStats      // leader??????
	SyncTicker      *time.Ticker            // v2???,??????ttl???????????????
	compactor       v3compactor.Compactor   // ???????????????????????????
	peerRt          http.RoundTripper       // ????????????????????????
	reqIDGen        *idutil.Generator       // ??????????????????id
	// wgMu blocks concurrent waitgroup mutation while etcd stopping
	wgMu sync.RWMutex
	// wg is used to wait for the goroutines that depends on the etcd state
	// to exit when stopping the etcd.
	wg                  sync.WaitGroup
	ctx                 context.Context // ?????????etcd??????????????????????????????????????????etcd????????????????????????.
	cancel              context.CancelFunc
	leadTimeMu          sync.RWMutex
	leadElectedTime     time.Time
	firstCommitInTermMu sync.RWMutex
	firstCommitInTermC  chan struct{} // ?????????????????????commit????????????
	*AccessController
}

// ??????????????????
type backendHooks struct {
	indexer   cindex.ConsistentIndexer // ????????????????????????
	lg        *zap.Logger
	confState raftpb.ConfState // ???????????????????????????
	// first write changes it to 'dirty'. false by default, so
	// not initialized `confState` is meaningless.
	confStateDirty bool
	confStateLock  sync.Mutex
}

func (bh *backendHooks) OnPreCommitUnsafe(tx backend.BatchTx) {
	bh.indexer.UnsafeSave(tx)
	bh.confStateLock.Lock()
	defer bh.confStateLock.Unlock()
	if bh.confStateDirty {
		membership.MustUnsafeSaveConfStateToBackend(bh.lg, tx, &bh.confState)
		// save bh.confState
		bh.confStateDirty = false
	}
}

func (bh *backendHooks) SetConfState(confState *raftpb.ConfState) {
	bh.confStateLock.Lock()
	defer bh.confStateLock.Unlock()
	bh.confState = *confState
	bh.confStateDirty = true
}

type Temp struct {
	Bepath   string
	W        *wal.WAL
	N        raft.RaftNodeInterFace
	S        *raft.MemoryStorage
	ID       types.ID
	CL       *membership.RaftCluster
	Remotes  []*membership.Member
	Snapshot *raftpb.Snapshot
	Prt      http.RoundTripper
	SS       *snap.Snapshotter
	ST       v2store.Store
	CI       cindex.ConsistentIndexer
	BeExist  bool
	BeHooks  *backendHooks
	BE       backend.Backend
}

func MySelfStartRaft(cfg config.ServerConfig) (temp *Temp, err error) {
	temp = &Temp{}
	temp.ST = v2store.New(StoreClusterPrefix, StoreKeysPrefix) // ???????????????store?????????   /0 /1

	if cfg.MaxRequestBytes > recommendedMaxRequestBytes { // 10M
		cfg.Logger.Warn(
			"??????????????????????????????",
			zap.Uint("max-request-bytes", cfg.MaxRequestBytes),
			zap.String("max-request-size", humanize.Bytes(uint64(cfg.MaxRequestBytes))),
			zap.Int("recommended-request-bytes", recommendedMaxRequestBytes),
			zap.String("recommended-request-size", recommendedMaxRequestBytesString),
		)
	}
	// ???????????????
	if terr := fileutil.TouchDirAll(cfg.DataDir); terr != nil {
		return nil, fmt.Errorf("????????????????????????: %v", terr)
	}

	haveWAL := wal.Exist(cfg.WALDir()) // default.etcd/member/wal
	// default.etcd/member/snap
	if err = fileutil.TouchDirAll(cfg.SnapDir()); err != nil {
		cfg.Logger.Fatal(
			"????????????????????????",
			zap.String("path", cfg.SnapDir()),
			zap.Error(err),
		)
	}
	// ???????????????????????????
	if err = fileutil.RemoveMatchFile(cfg.Logger, cfg.SnapDir(), func(fileName string) bool {
		return strings.HasPrefix(fileName, "tmp")
	}); err != nil {
		cfg.Logger.Error(
			"????????????????????????????????????",
			zap.String("path", cfg.SnapDir()),
			zap.Error(err),
		)
	}
	// ????????????struct
	temp.SS = snap.New(cfg.Logger, cfg.SnapDir())

	temp.Bepath = cfg.BackendPath() // default.etcd/member/snap/db
	temp.BeExist = fileutil.Exist(temp.Bepath)

	temp.CI = cindex.NewConsistentIndex(nil) // pointer
	temp.BeHooks = &backendHooks{lg: cfg.Logger, indexer: temp.CI}
	temp.BE = openBackend(cfg, temp.BeHooks)
	temp.CI.SetBackend(temp.BE)
	cindex.CreateMetaBucket(temp.BE.BatchTx())

	// ?????????,?????????????????????????????????
	if cfg.ExperimentalBootstrapDefragThresholdMegabytes != 0 {
		err := maybeDefragBackend(cfg, temp.BE)
		if err != nil {
			return nil, err
		}
	}

	defer func() {
		if err != nil {
			temp.BE.Close()
		}
	}()
	// ????????????
	temp.Prt, err = rafthttp.NewRoundTripper(cfg.PeerTLSInfo, cfg.PeerDialTimeout())
	if err != nil {
		return nil, err
	}

	switch {
	case !haveWAL && !cfg.NewCluster: // false true   ?????????????????????
		if err = cfg.VerifyJoinExisting(); err != nil {
			return nil, err
		}
		temp.CL, err = membership.NewClusterFromURLsMap(cfg.Logger, cfg.InitialClusterToken, cfg.InitialPeerURLsMap)
		if err != nil {
			return nil, err
		}
		existingCluster, gerr := GetClusterFromRemotePeers(cfg.Logger, getRemotePeerURLs(temp.CL, cfg.Name), temp.Prt)
		if gerr != nil {
			return nil, fmt.Errorf("cannot fetch cluster info from peer urls: %v", gerr)
		}
		if err = membership.ValidateClusterAndAssignIDs(cfg.Logger, temp.CL, existingCluster); err != nil {
			return nil, fmt.Errorf("error validating peerURLs %s: %v", existingCluster, err)
		}
		if !isCompatibleWithCluster(cfg.Logger, temp.CL, temp.CL.MemberByName(cfg.Name).ID, temp.Prt) {
			return nil, fmt.Errorf("incompatible with current running cluster")
		}

		temp.Remotes = existingCluster.Members()
		temp.CL.SetID(types.ID(0), existingCluster.ID())
		temp.CL.SetStore(temp.ST)
		temp.CL.SetBackend(temp.BE)
		temp.ID, temp.N, temp.S, temp.W = startNode(cfg, temp.CL, nil)
		temp.CL.SetID(temp.ID, existingCluster.ID())

	case !haveWAL && cfg.NewCluster: // false true   ???????????????
		if err = cfg.VerifyBootstrap(); err != nil { // ??????peer ???????????????--initial-advertise-peer-urls" and "--initial-cluster
			return nil, err
		}
		// ??????RaftCluster
		temp.CL, err = membership.NewClusterFromURLsMap(cfg.Logger, cfg.InitialClusterToken, cfg.InitialPeerURLsMap)
		if err != nil {
			return nil, err
		}
		m := temp.CL.MemberByName(cfg.Name) // ????????????????????????
		if isMemberBootstrapped(cfg.Logger, temp.CL, cfg.Name, temp.Prt, cfg.BootstrapTimeoutEffective()) {
			return nil, fmt.Errorf("?????? %s ???????????????", m.ID)
		}
		// TODO ????????????discovery ??????????????????
		if cfg.ShouldDiscover() {
			var str string
			str, err = v2discovery.JoinCluster(cfg.Logger, cfg.DiscoveryURL, cfg.DiscoveryProxy, m.ID, cfg.InitialPeerURLsMap.String())
			if err != nil {
				return nil, &DiscoveryError{Op: "join", Err: err}
			}
			var urlsmap types.URLsMap
			urlsmap, err = types.NewURLsMap(str)
			if err != nil {
				return nil, err
			}
			if config.CheckDuplicateURL(urlsmap) {
				return nil, fmt.Errorf("discovery cluster %s has duplicate url", urlsmap)
			}
			if temp.CL, err = membership.NewClusterFromURLsMap(cfg.Logger, cfg.InitialClusterToken, urlsmap); err != nil {
				return nil, err
			}
		}
		temp.CL.SetStore(temp.ST) // ?????????
		temp.CL.SetBackend(temp.BE)
		// ????????????
		temp.ID, temp.N, temp.S, temp.W = startNode(cfg, temp.CL, temp.CL.MemberIDs()) // ????????? ????????????????
		temp.CL.SetID(temp.ID, temp.CL.ID())

	case haveWAL:
		if err = fileutil.IsDirWriteable(cfg.MemberDir()); err != nil {
			return nil, fmt.Errorf("cannot write to member directory: %v", err)
		}

		if err = fileutil.IsDirWriteable(cfg.WALDir()); err != nil {
			return nil, fmt.Errorf("cannot write to WAL directory: %v", err)
		}

		if cfg.ShouldDiscover() {
			cfg.Logger.Warn(
				"discovery token is ignored since cluster already initialized; valid logs are found",
				zap.String("wal-dir", cfg.WALDir()),
			)
		}

		// Find a snapshot to start/restart a raft node
		walSnaps, err := wal.ValidSnapshotEntries(cfg.Logger, cfg.WALDir())
		if err != nil {
			return nil, err
		}
		// snapshot files can backend orphaned if etcd crashes after writing them but before writing the corresponding
		// wal log entries
		temp.Snapshot, err = temp.SS.LoadNewestAvailable(walSnaps)
		if err != nil && err != snap.ErrNoSnapshot {
			return nil, err
		}

		if temp.Snapshot != nil {
			if err = temp.ST.Recovery(temp.Snapshot.Data); err != nil {
				cfg.Logger.Panic("failed to recover from snapshot", zap.Error(err))
			}

			if err = assertNoV2StoreContent(cfg.Logger, temp.ST, cfg.V2Deprecation); err != nil {
				cfg.Logger.Error("illegal v2store content", zap.Error(err))
				return nil, err
			}

			cfg.Logger.Info(
				"recovered v2 store from snapshot",
				zap.Uint64("snapshot-index", temp.Snapshot.Metadata.Index),
				zap.String("snapshot-size", humanize.Bytes(uint64(temp.Snapshot.Size()))),
			)

			if temp.BE, err = recoverSnapshotBackend(cfg, temp.BE, *temp.Snapshot, temp.BeExist, temp.BeHooks); err != nil {
				cfg.Logger.Panic("failed to recover v3 backend from snapshot", zap.Error(err))
			}
			// A snapshot db may have already been recovered, and the old db should have
			// already been closed in this case, so we should set the backend again.
			temp.CI.SetBackend(temp.BE)
			s1, s2 := temp.BE.Size(), temp.BE.SizeInUse()
			cfg.Logger.Info(
				"recovered v3 backend from snapshot",
				zap.Int64("backend-size-bytes", s1),
				zap.String("backend-size", humanize.Bytes(uint64(s1))),
				zap.Int64("backend-size-in-use-bytes", s2),
				zap.String("backend-size-in-use", humanize.Bytes(uint64(s2))),
			)
		} else {
			cfg.Logger.Info("No snapshot found. Recovering WAL from scratch!")
		}

		if !cfg.ForceNewCluster {
			temp.ID, temp.CL, temp.N, temp.S, temp.W = restartNode(cfg, temp.Snapshot)
		} else {
			temp.ID, temp.CL, temp.N, temp.S, temp.W = restartAsStandaloneNode(cfg, temp.Snapshot)
		}

		temp.CL.SetStore(temp.ST)
		temp.CL.SetBackend(temp.BE)
		temp.CL.Recover(api.UpdateCapability)
		if temp.CL.Version() != nil && !temp.CL.Version().LessThan(semver.Version{Major: 3}) && !temp.BeExist {
			os.RemoveAll(temp.Bepath)
			return nil, fmt.Errorf("database file (%v) of the backend is missing", temp.Bepath)
		}

	default:
		return nil, fmt.Errorf("????????????????????????")
	}

	if terr := fileutil.TouchDirAll(cfg.MemberDir()); terr != nil {
		return nil, fmt.Errorf("????????????????????????: %v", terr)
	}

	return
}

// NewServer ???????????????????????????????????????EtcdServer.???EtcdServer??????????????????,??????????????????????????????.
func NewServer(cfg config.ServerConfig) (srv *EtcdServer, err error) {
	temp := &Temp{}
	temp, err = MySelfStartRaft(cfg) // ?????????????????????
	serverStats := stats.NewServerStats(cfg.Name, temp.ID.String())
	leaderStats := stats.NewLeaderStats(cfg.Logger, temp.ID.String())

	heartbeat := time.Duration(cfg.TickMs) * time.Millisecond
	srv = &EtcdServer{
		readych:     make(chan struct{}),
		Cfg:         cfg,
		lgMu:        new(sync.RWMutex),
		lg:          cfg.Logger,
		errorc:      make(chan error, 1),
		v2store:     temp.ST,
		snapshotter: temp.SS,
		r: *newRaftNode(
			raftNodeConfig{
				lg:                cfg.Logger,
				isIDRemoved:       func(id uint64) bool { return temp.CL.IsIDRemoved(types.ID(id)) },
				RaftNodeInterFace: temp.N,
				heartbeat:         heartbeat,
				raftStorage:       temp.S,
				storage:           NewStorage(temp.W, temp.SS),
			},
		),
		id:                 temp.ID,
		attributes:         membership.Attributes{Name: cfg.Name, ClientURLs: cfg.ClientURLs.StringSlice()},
		cluster:            temp.CL,
		stats:              serverStats,
		lstats:             leaderStats,
		SyncTicker:         time.NewTicker(500 * time.Millisecond),
		peerRt:             temp.Prt,
		reqIDGen:           idutil.NewGenerator(uint16(temp.ID), time.Now()),
		AccessController:   &AccessController{CORS: cfg.CORS, HostWhitelist: cfg.HostWhitelist},
		consistIndex:       temp.CI,
		firstCommitInTermC: make(chan struct{}),
	}
	srv.applyV2 = NewApplierV2(cfg.Logger, srv.v2store, srv.cluster)

	srv.backend = temp.BE
	srv.beHooks = temp.BeHooks
	// ????????????????????????leader?????????,lease????????????,??????ttl????????????????????????,?????????
	minTTL := time.Duration((3*cfg.ElectionTicks)/2) * heartbeat
	// ???????????????????????????2s,

	// ?????????KV?????????????????????.???????????????mvcc.KV???,?????????????????????????????????????????????.?????????????????????mvcc.KV,?????????????????????????????????????????????????????????.
	srv.lessor = lease.NewLessor(srv.Logger(), srv.backend, srv.cluster, lease.LessorConfig{
		MinLeaseTTL:                int64(math.Ceil(minTTL.Seconds())),
		CheckpointInterval:         cfg.LeaseCheckpointInterval,
		CheckpointPersist:          cfg.LeaseCheckpointPersist,
		ExpiredLeasesRetryInterval: srv.Cfg.ReqTimeout(),
	})

	tp, err := auth.NewTokenProvider(cfg.Logger, cfg.AuthToken, // ????????????  simple???jwt
		func(index uint64) <-chan struct{} {
			return srv.applyWait.Wait(index)
		},
		time.Duration(cfg.TokenTTL)*time.Second,
	)
	if err != nil {
		cfg.Logger.Warn("??????????????????????????????", zap.Error(err))
		return nil, err
	}
	// watch | kv ...
	srv.kv = mvcc.New(srv.Logger(), srv.backend, srv.lessor, mvcc.StoreConfig{CompactionBatchLimit: cfg.CompactionBatchLimit})

	kvindex := temp.CI.ConsistentIndex()
	srv.lg.Debug("??????consistentIndex", zap.Uint64("index", kvindex))
	if temp.BeExist {
		// TODO: remove kvindex != 0 checking when we do not expect users to upgrade
		// etcd from pre-3.0 release.
		if temp.Snapshot != nil && kvindex < temp.Snapshot.Metadata.Index {
			if kvindex != 0 {
				return nil, fmt.Errorf("database file (%v index %d) does not match with snapshot (index %d)", temp.Bepath, kvindex, temp.Snapshot.Metadata.Index)
			}
			cfg.Logger.Warn(
				"consistent index was never saved",
				zap.Uint64("snapshot-index", temp.Snapshot.Metadata.Index),
			)
		}
	}

	srv.authStore = auth.NewAuthStore(srv.Logger(), srv.backend, tp, int(cfg.BcryptCost)) // BcryptCost ?????????????????????????????????bcrypt???????????????/????????????10

	newSrv := srv // since srv == nil in defer if srv is returned as nil
	defer func() {
		// closing backend without first closing kv can cause
		// resumed compactions to fail with closed tx errors
		if err != nil {
			newSrv.kv.Close()
		}
	}()
	if num := cfg.AutoCompactionRetention; num != 0 {
		srv.compactor, err = v3compactor.New(cfg.Logger, cfg.AutoCompactionMode, num, srv.kv, srv)
		if err != nil {
			return nil, err
		}
		srv.compactor.Run()
	}

	srv.applyV3Base = srv.newApplierV3Backend()
	srv.applyV3Internal = srv.newApplierV3Internal()
	// ???????????????????????????
	if err = srv.restoreAlarms(); err != nil {
		return nil, err
	}

	if srv.Cfg.EnableLeaseCheckpoint {
		// ????????????checkpointer???????????????????????????.
		srv.lessor.SetCheckpointer(func(ctx context.Context, cp *pb.LeaseCheckpointRequest) {
			// ?????????????????? Lease ????????? TTL ?????? Raft Log ????????? Follower ??????,Follower ???????????? CheckPoint ?????????,
			// ???????????????????????? LeaseMap ????????? TTL ??????.
			srv.raftRequestOnce(ctx, pb.InternalRaftRequest{LeaseCheckpoint: cp})
		})
	}

	// TODO: move transport initialization near the definition of remote
	tr := &rafthttp.Transport{
		Logger:      cfg.Logger,
		TLSInfo:     cfg.PeerTLSInfo,
		DialTimeout: cfg.PeerDialTimeout(),
		ID:          temp.ID,
		URLs:        cfg.PeerURLs,
		ClusterID:   temp.CL.ID(),
		Raft:        srv,
		Snapshotter: temp.SS,
		ServerStats: serverStats,
		LeaderStats: leaderStats,
		ErrorC:      srv.errorc,
	}
	if err = tr.Start(); err != nil {
		return nil, err
	}
	// add all remotes into transport
	for _, m := range temp.Remotes {
		if m.ID != temp.ID {
			tr.AddRemote(m.ID, m.PeerURLs)
		}
	}
	for _, m := range temp.CL.Members() {
		if m.ID != temp.ID {
			tr.AddPeer(m.ID, m.PeerURLs)
		}
	}
	srv.r.transport = tr

	return srv, nil
}

// assertNoV2StoreContent -> depending on the deprecation stage, warns or report an error
// if the v2store contains custom content.
func assertNoV2StoreContent(lg *zap.Logger, st v2store.Store, deprecationStage config.V2DeprecationEnum) error {
	metaOnly, err := membership.IsMetaStoreOnly(st)
	if err != nil {
		return err
	}
	if metaOnly {
		return nil
	}
	if deprecationStage.IsAtLeast(config.V2_DEPR_1_WRITE_ONLY) {
		return fmt.Errorf("detected disallowed custom content in v2store for stage --v2-deprecation=%s", deprecationStage)
	}
	lg.Warn("detected custom v2store content. Etcd v3.5 is the last version allowing to access it using API v2. Please remove the content.")
	return nil
}

func (s *EtcdServer) Logger() *zap.Logger {
	s.lgMu.RLock()
	l := s.lg
	s.lgMu.RUnlock()
	return l
}

func tickToDur(ticks int, tickMs uint) string {
	return fmt.Sprintf("%v", time.Duration(ticks)*time.Duration(tickMs)*time.Millisecond)
}

func (s *EtcdServer) adjustTicks() {
	lg := s.Logger()
	clusterN := len(s.cluster.Members())

	// single-node fresh start, or single-node recovers from snapshot
	if clusterN == 1 {
		ticks := s.Cfg.ElectionTicks - 1
		lg.Info(
			"started as single-node; fast-forwarding election ticks",
			zap.String("local-member-id", s.ID().String()),
			zap.Int("forward-ticks", ticks),
			zap.String("forward-duration", tickToDur(ticks, s.Cfg.TickMs)),
			zap.Int("election-ticks", s.Cfg.ElectionTicks),
			zap.String("election-timeout", tickToDur(s.Cfg.ElectionTicks, s.Cfg.TickMs)),
		)
		s.r.advanceTicks(ticks)
		return
	}

	if !s.Cfg.InitialElectionTickAdvance {
		lg.Info("skipping initial election tick advance", zap.Int("election-ticks", s.Cfg.ElectionTicks))
		return
	}
	lg.Info("starting initial election tick advance", zap.Int("election-ticks", s.Cfg.ElectionTicks))

	// retry up to "rafthttp.ConnReadTimeout", which is 5-sec
	// until peer connection reports; otherwise:
	// 1. all connections failed, or
	// 2. no active peers, or
	// 3. restarted single-node with no snapshot
	// then, do nothing, because advancing ticks would have no effect
	waitTime := rafthttp.ConnReadTimeout
	itv := 50 * time.Millisecond
	for i := int64(0); i < int64(waitTime/itv); i++ {
		select {
		case <-time.After(itv):
		case <-s.stopping:
			return
		}

		peerN := s.r.transport.ActivePeers()
		if peerN > 1 {
			// multi-node received peer connection reports
			// adjust ticks, in case slow leader message receive
			ticks := s.Cfg.ElectionTicks - 2

			lg.Info(
				"initialized peer connections; fast-forwarding election ticks",
				zap.String("local-member-id", s.ID().String()),
				zap.Int("forward-ticks", ticks),
				zap.String("forward-duration", tickToDur(ticks, s.Cfg.TickMs)),
				zap.Int("election-ticks", s.Cfg.ElectionTicks),
				zap.String("election-timeout", tickToDur(s.Cfg.ElectionTicks, s.Cfg.TickMs)),
				zap.Int("active-remote-members", peerN),
			)

			s.r.advanceTicks(ticks)
			return
		}
	}
}

func (s *EtcdServer) Start() {
	s.start()
	s.GoAttach(func() { s.adjustTicks() })
	s.GoAttach(func() { s.publish(s.Cfg.ReqTimeout()) })
	s.GoAttach(s.purgeFile)
	s.GoAttach(s.monitorVersions)
	s.GoAttach(s.linearizableReadLoop)
	s.GoAttach(s.monitorKVHash)
	s.GoAttach(s.monitorDowngrade)
}

func (s *EtcdServer) start() {
	lg := s.Logger()

	if s.Cfg.SnapshotCount == 0 { // ????????????????????????????????????????????????
		lg.Info("??????????????????????????????",
			zap.Uint64("given-snapshot-count", s.Cfg.SnapshotCount), // ????????????????????????????????????????????????
			zap.Uint64("updated-snapshot-count", DefaultSnapshotCount),
		)
		s.Cfg.SnapshotCount = DefaultSnapshotCount // ????????????????????????????????????????????????
	}
	if s.Cfg.SnapshotCatchUpEntries == 0 {
		lg.Info("??????????????????????????????????????????",
			zap.Uint64("given-snapshot-catchup-entries", s.Cfg.SnapshotCatchUpEntries),
			zap.Uint64("updated-snapshot-catchup-entries", DefaultSnapshotCatchUpEntries),
		)
		s.Cfg.SnapshotCatchUpEntries = DefaultSnapshotCatchUpEntries
	}

	s.w = wait.New()
	s.applyWait = wait.NewTimeList()
	s.done = make(chan struct{})
	s.stop = make(chan struct{})
	s.stopping = make(chan struct{}, 1)
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.readwaitc = make(chan struct{}, 1)
	s.readNotifier = newNotifier()
	s.leaderChanged = make(chan struct{})
	if s.ClusterVersion() != nil {
		lg.Info("??????etcd", zap.String("local-member-id", s.ID().String()),
			zap.String("local-etcd-version", version.Version),
			zap.String("cluster-id", s.Cluster().ID().String()),
			zap.String("cluster-version", version.Cluster(s.ClusterVersion().String())),
		)
	} else {
		lg.Info("??????etcd", zap.String("local-member-id", s.ID().String()),
			zap.String("local-etcd-version", version.Version), zap.String("cluster-version", "to_be_decided"))
	}

	go s.run()
}

func (s *EtcdServer) purgeFile() {
	lg := s.Logger()
	var dberrc, serrc, werrc <-chan error
	var dbdonec, sdonec, wdonec <-chan struct{}
	if s.Cfg.MaxSnapFiles > 0 {
		dbdonec, dberrc = fileutil.PurgeFileWithDoneNotify(lg, s.Cfg.SnapDir(), "snap.db", s.Cfg.MaxSnapFiles, purgeFileInterval, s.stopping)
		sdonec, serrc = fileutil.PurgeFileWithDoneNotify(lg, s.Cfg.SnapDir(), "snap", s.Cfg.MaxSnapFiles, purgeFileInterval, s.stopping)
	}
	if s.Cfg.MaxWALFiles > 0 {
		wdonec, werrc = fileutil.PurgeFileWithDoneNotify(lg, s.Cfg.WALDir(), "wal", s.Cfg.MaxWALFiles, purgeFileInterval, s.stopping)
	}

	select {
	case e := <-dberrc:
		lg.Fatal("failed to purge snap db file", zap.Error(e))
	case e := <-serrc:
		lg.Fatal("failed to purge snap file", zap.Error(e))
	case e := <-werrc:
		lg.Fatal("failed to purge wal file", zap.Error(e))
	case <-s.stopping:
		if dbdonec != nil {
			<-dbdonec
		}
		if sdonec != nil {
			<-sdonec
		}
		if wdonec != nil {
			<-wdonec
		}
		return
	}
}

type ServerPeer interface {
	ServerV2
	RaftHandler() http.Handler
	LeaseHandler() http.Handler
}

func (s *EtcdServer) RaftHandler() http.Handler {
	return s.r.transport.Handler()
}

type ServerPeerV2 interface {
	ServerPeer
	HashKVHandler() http.Handler
	DowngradeEnabledHandler() http.Handler
}

func (s *EtcdServer) ReportUnreachable(id uint64) {
	s.r.ReportUnreachable(id)
}

// ReportSnapshot reports snapshot sent status to the raft state machine,
// and clears the used snapshot from the snapshot store.
func (s *EtcdServer) ReportSnapshot(id uint64, status raft.SnapshotStatus) {
	s.r.ReportSnapshot(id, status)
}

type etcdProgress struct {
	confState raftpb.ConfState
	snapi     uint64
	appliedt  uint64
	appliedi  uint64
}

// raftReadyHandler contains a set of EtcdServer operations to backend called by raftNode,
// and helps decouple state machine logic from Raft algorithms.
// TODO: add a state machine interface to apply the commit entries and do snapshot/recover
type raftReadyHandler struct {
	getLead              func() (lead uint64)
	updateLead           func(lead uint64)
	updateLeadership     func(newLeader bool)
	updateCommittedIndex func(uint64)
}

func (s *EtcdServer) run() {
	lg := s.Logger()

	sn, err := s.r.raftStorage.Snapshot()
	if err != nil {
		lg.Panic("???Raft????????????????????????", zap.Error(err))
	}

	// asynchronously accept apply packets, dispatch progress in-order
	sched := schedule.NewFIFOScheduler()

	var (
		smu   sync.RWMutex
		syncC <-chan time.Time
	)
	setSyncC := func(ch <-chan time.Time) {
		smu.Lock()
		syncC = ch
		smu.Unlock()
	}
	getSyncC := func() (ch <-chan time.Time) {
		smu.RLock()
		ch = syncC
		smu.RUnlock()
		return
	}
	rh := &raftReadyHandler{
		getLead:    func() (lead uint64) { return s.getLead() },
		updateLead: func(lead uint64) { s.setLead(lead) },
		updateLeadership: func(newLeader bool) {
			if !s.isLeader() {
				// ????????????leader???
				if s.lessor != nil {
					s.lessor.Demote() // ?????????????????????
				}
				if s.compactor != nil {
					s.compactor.Pause()
				}
				setSyncC(nil)
			} else {
				if newLeader {
					t := time.Now()
					s.leadTimeMu.Lock()
					s.leadElectedTime = t
					s.leadTimeMu.Unlock()
				}
				setSyncC(s.SyncTicker.C)
				if s.compactor != nil {
					s.compactor.Resume()
				}
			}
			if newLeader {
				s.leaderChangedMu.Lock()
				lc := s.leaderChanged
				s.leaderChanged = make(chan struct{})
				close(lc)
				s.leaderChangedMu.Unlock()
			}
			// TODO: remove the nil checking
			// current test utility does not provide the stats
			if s.stats != nil {
				s.stats.BecomeLeader()
			}
		},
		updateCommittedIndex: func(ci uint64) {
			cci := s.getCommittedIndex()
			if ci > cci {
				s.setCommittedIndex(ci)
			}
		},
	}
	s.r.start(rh)

	ep := etcdProgress{
		confState: sn.Metadata.ConfState,
		snapi:     sn.Metadata.Index,
		appliedt:  sn.Metadata.Term,
		appliedi:  sn.Metadata.Index,
	}

	defer func() {
		s.wgMu.Lock() // block concurrent waitgroup adds in GoAttach while stopping
		close(s.stopping)
		s.wgMu.Unlock()
		s.cancel()
		sched.Stop()

		// wait for gouroutines before closing raft so wal stays open
		s.wg.Wait()

		s.SyncTicker.Stop()

		// must stop raft after scheduler-- etcdserver can leak rafthttp pipelines
		// by adding a peer after raft stops the transport
		s.r.stop()

		s.Cleanup()

		close(s.done)
	}()
	var expiredLeaseC <-chan []*lease.Lease // ???????????????????????????????????????CHAN.
	if s.lessor != nil {                    // v3???,???????????????????????????
		expiredLeaseC = s.lessor.ExpiredLeasesC()
	}

	for {
		select {
		case ap := <-s.r.apply():
			// ???????????????,??????apply????????????
			// index1:EntryConfChange {"Type":0,"NodeID":10276657743932975437,"Context":"{\"id\":10276657743932975437,\"peerURLs\":[\"http://localhost:2380\"],\"name\":\"default\"}","ID":0}
			// index2:EntryNormal nil    ????????????????????????commit
			// index3:EntryNormal {"ID":7587861549007417858,"Method":"PUT","Path":"/0/members/8e9e05c52164694d/attributes","Val":"{\"name\":\"default\",\"clientURLs\":[\"http://localhost:2379\"]}","Dir":false,"PrevValue":"","PrevIndex":0,"Expiration":0,"Wait":false,"Since":0,"Recursive":false,"Sorted":false,"Quorum":false,"Time":0,"Stream":false}
			// ?????? ??????applyc?????????
			f := func(context.Context) {
				s.applyAll(&ep, &ap)
			}
			sched.Schedule(f)
		case leases := <-expiredLeaseC:
			s.GoAttach(func() {
				// ?????????????????????????????????????????????????????????
				c := make(chan struct{}, maxPendingRevokes) // ???????????????  ????????????16
				for _, lease := range leases {
					select {
					case c <- struct{}{}:
					case <-s.stopping:
						return
					}
					lid := lease.ID
					s.GoAttach(func() {
						ctx := s.authStore.WithRoot(s.ctx)
						_, lerr := s.LeaseRevoke(ctx, &pb.LeaseRevokeRequest{ID: int64(lid)})
						if lerr == nil {
						} else {
							lg.Warn("??????????????????", zap.String("lease-id", fmt.Sprintf("%016x", lid)), zap.Error(lerr))
						}
						<-c
					})
				}
			})
		case err := <-s.errorc:
			lg.Warn("etcd error", zap.Error(err))
			lg.Warn("???????????????data-dir????????????")
			return
		case <-getSyncC():
			if s.v2store.HasTTLKeys() {
				s.sync(s.Cfg.ReqTimeout())
			}
		case <-s.stop:
			return
		}
	}
}

// Cleanup removes allocated objects by EtcdServer.NewServer in
// situation that EtcdServer::Start was not called (that takes care of cleanup).
func (s *EtcdServer) Cleanup() {
	// kv, lessor and backend can backend nil if running without v3 enabled
	// or running unit tests.
	if s.lessor != nil {
		s.lessor.Stop()
	}
	if s.kv != nil {
		s.kv.Close()
	}
	if s.authStore != nil {
		s.authStore.Close()
	}
	if s.backend != nil {
		s.backend.Close()
	}
	if s.compactor != nil {
		s.compactor.Stop()
	}
}

func (s *EtcdServer) applyAll(ep *etcdProgress, apply *apply) {
	s.applySnapshot(ep, apply) // ?????????????????????????????????????????????
	s.applyEntries(ep, apply)

	s.applyWait.Trigger(ep.appliedi)

	// wait for the raft routine to finish the disk writes before triggering a
	// snapshot. or applied index might backend greater than the last index in raft
	// storage, since the raft routine might backend slower than apply routine.
	<-apply.notifyc

	s.triggerSnapshot(ep)
	select {
	// snapshot requested via send()
	case m := <-s.r.msgSnapC:
		merged := s.createMergedSnapshotMessage(m, ep.appliedt, ep.appliedi, ep.confState)
		s.sendMergedSnap(merged)
	default:
	}
}

func (s *EtcdServer) applySnapshot(ep *etcdProgress, apply *apply) {
	if raft.IsEmptySnap(apply.snapshot) {
		return
	}

	lg := s.Logger()
	lg.Info("??????????????????",
		zap.Uint64("current-snapshot-index", ep.snapi),
		zap.Uint64("current-applied-index", ep.appliedi),
		zap.Uint64("incoming-leader-snapshot-index", apply.snapshot.Metadata.Index),
		zap.Uint64("incoming-leader-snapshot-term", apply.snapshot.Metadata.Term),
	)
	defer func() {
		lg.Info("???????????????",
			zap.Uint64("current-snapshot-index", ep.snapi),
			zap.Uint64("current-applied-index", ep.appliedi),
			zap.Uint64("incoming-leader-snapshot-index", apply.snapshot.Metadata.Index),
			zap.Uint64("incoming-leader-snapshot-term", apply.snapshot.Metadata.Term),
		)
	}()

	if apply.snapshot.Metadata.Index <= ep.appliedi {
		lg.Panic("???????????? ????????????????????????????????????",
			zap.Uint64("current-snapshot-index", ep.snapi),
			zap.Uint64("current-applied-index", ep.appliedi),
			zap.Uint64("incoming-leader-snapshot-index", apply.snapshot.Metadata.Index),
			zap.Uint64("incoming-leader-snapshot-term", apply.snapshot.Metadata.Term),
		)
	}

	// ??????raftnode???????????????????????????
	<-apply.notifyc

	newbe, err := openSnapshotBackend(s.Cfg, s.snapshotter, apply.snapshot, s.beHooks)
	if err != nil {
		lg.Panic("failed to open snapshot backend", zap.Error(err))
	}

	// always recover lessor before kv. When we recover the mvcc.KV it will reattach keys to its leases.
	// If we recover mvcc.KV first, it will attach the keys to the wrong lessor before it recovers.
	if s.lessor != nil {
		lg.Info("restoring lease store")

		s.lessor.Recover(newbe, func() lease.TxnDelete { return s.kv.Write(traceutil.TODO()) })

		lg.Info("restored lease store")
	}

	lg.Info("restoring mvcc store")

	if err := s.kv.Restore(newbe); err != nil {
		lg.Panic("failed to restore mvcc store", zap.Error(err))
	}

	s.consistIndex.SetBackend(newbe)
	lg.Info("restored mvcc store", zap.Uint64("consistent-index", s.consistIndex.ConsistentIndex()))

	// Closing old backend might block until all the txns
	// on the backend are finished.
	// We do not want to wait on closing the old backend.
	s.backendLock.Lock()
	oldbe := s.backend
	go func() {
		lg.Info("closing old backend file")
		defer func() {
			lg.Info("closed old backend file")
		}()
		if err := oldbe.Close(); err != nil {
			lg.Panic("failed to close old backend", zap.Error(err))
		}
	}()

	s.backend = newbe
	s.backendLock.Unlock()

	lg.Info("restoring alarm store")

	if err := s.restoreAlarms(); err != nil {
		lg.Panic("failed to restore alarm store", zap.Error(err))
	}

	lg.Info("restored alarm store")

	if s.authStore != nil {
		lg.Info("restoring auth store")

		s.authStore.Recover(newbe)

		lg.Info("restored auth store")
	}

	lg.Info("restoring v2 store")
	if err := s.v2store.Recovery(apply.snapshot.Data); err != nil {
		lg.Panic("failed to restore v2 store", zap.Error(err))
	}

	if err := assertNoV2StoreContent(lg, s.v2store, s.Cfg.V2Deprecation); err != nil {
		lg.Panic("illegal v2store content", zap.Error(err))
	}

	lg.Info("restored v2 store")

	s.cluster.SetBackend(newbe)

	lg.Info("restoring cluster configuration")

	s.cluster.Recover(api.UpdateCapability)

	lg.Info("restored cluster configuration")
	lg.Info("removing old peers from network")

	// recover raft transport
	s.r.transport.RemoveAllPeers()

	lg.Info("removed old peers from network")
	lg.Info("adding peers from new cluster configuration")

	for _, m := range s.cluster.Members() {
		if m.ID == s.ID() {
			continue
		}
		s.r.transport.AddPeer(m.ID, m.PeerURLs)
	}

	lg.Info("added peers from new cluster configuration")

	ep.appliedt = apply.snapshot.Metadata.Term
	ep.appliedi = apply.snapshot.Metadata.Index
	ep.snapi = ep.appliedi
	ep.confState = apply.snapshot.Metadata.ConfState
}

func (s *EtcdServer) applyEntries(ep *etcdProgress, apply *apply) {
	if len(apply.entries) == 0 {
		return
	}
	firsti := apply.entries[0].Index
	if firsti > ep.appliedi+1 {
		lg := s.Logger()
		lg.Panic("????????? ???????????????",
			zap.Uint64("current-applied-index", ep.appliedi),
			zap.Uint64("first-committed-entry-index", firsti),
		)
	}
	var ents []raftpb.Entry
	if ep.appliedi+1-firsti < uint64(len(apply.entries)) {
		ents = apply.entries[ep.appliedi+1-firsti:]
	}
	if len(ents) == 0 {
		return
	}
	var shouldstop bool
	if ep.appliedt, ep.appliedi, shouldstop = s.apply(ents, &ep.confState); shouldstop {
		go s.stopWithDelay(10*100*time.Millisecond, fmt.Errorf(""))
	}
}

func (s *EtcdServer) triggerSnapshot(ep *etcdProgress) {
	if ep.appliedi-ep.snapi <= s.Cfg.SnapshotCount { // ????????????????????????????????????????????????
		return
	}

	lg := s.Logger()
	lg.Info("???????????????",
		zap.String("local-member-id", s.ID().String()),
		zap.Uint64("local-member-applied-index", ep.appliedi),
		zap.Uint64("local-member-snapshot-index", ep.snapi),
		zap.Uint64("local-member-snapshot-count", s.Cfg.SnapshotCount),
	)

	s.snapshot(ep.appliedi, ep.confState)
	ep.snapi = ep.appliedi
}

func (s *EtcdServer) hasMultipleVotingMembers() bool {
	return s.cluster != nil && len(s.cluster.VotingMemberIDs()) > 1
}

func (s *EtcdServer) isLeader() bool {
	return uint64(s.ID()) == s.Lead()
}

// MoveLeader leader??????
func (s *EtcdServer) MoveLeader(ctx context.Context, lead, transferee uint64) error {
	if !s.cluster.IsMemberExist(types.ID(transferee)) || s.cluster.Member(types.ID(transferee)).IsLearner {
		return ErrBadLeaderTransferee
	}

	now := time.Now()
	interval := time.Duration(s.Cfg.TickMs) * time.Millisecond

	lg := s.Logger()
	lg.Info(
		"??????leader??????",
		zap.String("local-member-id", s.ID().String()),
		zap.String("current-leader-member-id", types.ID(lead).String()),
		zap.String("transferee-member-id", types.ID(transferee).String()),
	)

	s.r.TransferLeadership(ctx, lead, transferee) // ??????leader??????
	for s.Lead() != transferee {
		select {
		case <-ctx.Done(): // time out
			return ErrTimeoutLeaderTransfer
		case <-time.After(interval):
		}
	}

	// ??????????????????,????????????????????????leader?????????
	lg.Info(
		"leader????????????",
		zap.String("local-member-id", s.ID().String()),
		zap.String("old-leader-member-id", types.ID(lead).String()),
		zap.String("new-leader-member-id", types.ID(transferee).String()),
		zap.Duration("took", time.Since(now)),
	)
	return nil
}

func (s *EtcdServer) TransferLeadership() error {
	lg := s.Logger()
	if !s.isLeader() {
		lg.Info(
			"skipped leadership transfer; local etcd is not leader",
			zap.String("local-member-id", s.ID().String()),
			zap.String("current-leader-member-id", types.ID(s.Lead()).String()),
		)
		return nil
	}

	if !s.hasMultipleVotingMembers() {
		lg.Info(
			"skipped leadership transfer for single voting member cluster",
			zap.String("local-member-id", s.ID().String()),
			zap.String("current-leader-member-id", types.ID(s.Lead()).String()),
		)
		return nil
	}

	transferee, ok := longestConnected(s.r.transport, s.cluster.VotingMemberIDs())
	if !ok {
		return ErrUnhealthy
	}

	tm := s.Cfg.ReqTimeout()
	ctx, cancel := context.WithTimeout(s.ctx, tm)
	err := s.MoveLeader(ctx, s.Lead(), uint64(transferee))
	cancel()
	return err
}

// HardStop ??????????????????????????????????????????????????????etcd.
func (s *EtcdServer) HardStop() {
	select {
	case s.stop <- struct{}{}:
	case <-s.done:
		return
	}
	<-s.done
}

// Stop ?????????????????????, ?????????leader??????leader??????
func (s *EtcdServer) Stop() {
	lg := s.Logger()
	if err := s.TransferLeadership(); err != nil {
		lg.Warn("leader????????????", zap.String("local-member-id", s.ID().String()), zap.Error(err))
	}
	s.HardStop()
}

// ReadyNotify  ???etcd ????????????????????????,?????????ready ch
func (s *EtcdServer) ReadyNotify() <-chan struct{} { return s.readych }

func (s *EtcdServer) stopWithDelay(d time.Duration, err error) {
	select {
	case <-time.After(d):
	case <-s.done:
	}
	select {
	case s.errorc <- err:
	default:
	}
}

// StopNotify ???etcd?????????????????????channel?????? empty struct
// when the etcd is stopped.
func (s *EtcdServer) StopNotify() <-chan struct{} { return s.done }

// StoppingNotify returns a channel that receives a empty struct
// when the etcd is being stopped.
func (s *EtcdServer) StoppingNotify() <-chan struct{} { return s.stopping }

func (s *EtcdServer) SelfStats() []byte { return s.stats.JSON() }

func (s *EtcdServer) LeaderStats() []byte {
	lead := s.getLead()
	if lead != uint64(s.id) {
		return nil
	}
	return s.lstats.JSON()
}

func (s *EtcdServer) StoreStats() []byte { return s.v2store.JsonStats() }

// ???????????????????????????
func (s *EtcdServer) checkMembershipOperationPermission(ctx context.Context) error {
	_ = auth.NewAuthStore
	if s.authStore == nil {
		// ????????????etcd?????????,s.authStore??????????????????.???????????????????????????server_test.go????????????
		return nil
	}

	// ?????????,????????????????????????API????????????,??????TOCTOU?????????????????????????????????????????????:
	// ????????????A???????????????------??????A????????????------??????????????????????????????????????????
	// ??????,??????????????????????????????????????????????????????.??????????????????????????????????????????????????????.
	authInfo, err := s.AuthInfoFromCtx(ctx)
	if err != nil {
		return err
	}

	return s.AuthStore().IsAdminPermitted(authInfo)
}

// ??????learner???????????????leader
// ??????:????????????????????????????????????,???????????????????????????,????????????nil.
// ??????????????????????????????????????????????????????????????????.
func (s *EtcdServer) isLearnerReady(id uint64) error {
	rs := s.raftStatus()
	if rs.Progress == nil {
		return ErrNotLeader
	}

	var learnerMatch uint64
	isFound := false
	leaderID := rs.ID
	for memberID, progress := range rs.Progress {
		if id == memberID {
			learnerMatch = progress.Match
			isFound = true
			break
		}
	}

	if isFound {
		leaderMatch := rs.Progress[leaderID].Match
		// learner?????????????????????????????????
		if float64(learnerMatch) < float64(leaderMatch)*readyPercent {
			return ErrLearnerNotReady
		}
	}
	return nil
}

func (s *EtcdServer) mayRemoveMember(id types.ID) error {
	if !s.Cfg.StrictReconfigCheck { // ????????????????????????
		return nil
	}

	lg := s.Logger()
	isLearner := s.cluster.IsMemberExist(id) && s.cluster.Member(id).IsLearner
	// no need to check quorum when removing non-voting member
	if isLearner {
		return nil
	}

	if !s.cluster.IsReadyToRemoveVotingMember(uint64(id)) {
		lg.Warn(
			"rejecting member remove request; not enough healthy members",
			zap.String("local-member-id", s.ID().String()),
			zap.String("requested-member-remove-id", id.String()),
			zap.Error(ErrNotEnoughStartedMembers),
		)
		return ErrNotEnoughStartedMembers
	}

	// downed member is safe to remove since it's not part of the active quorum
	if t := s.r.transport.ActiveSince(id); id != s.ID() && t.IsZero() {
		return nil
	}

	// protect quorum if some members are down
	m := s.cluster.VotingMembers()
	active := numConnectedSince(s.r.transport, time.Now().Add(-HealthInterval), s.ID(), m)
	if (active - 1) < 1+((len(m)-1)/2) {
		lg.Warn(
			"rejecting member remove request; local member has not been connected to all peers, reconfigure breaks active quorum",
			zap.String("local-member-id", s.ID().String()),
			zap.String("requested-member-remove", id.String()),
			zap.Int("active-peers", active),
			zap.Error(ErrUnhealthy),
		)
		return ErrUnhealthy
	}

	return nil
}

// FirstCommitInTermNotify
// ??????????????????commit????????????channel????????????,?????????leader??????????????????????????????
// Leader??????????????????????????????,??????????????????????????????
func (s *EtcdServer) FirstCommitInTermNotify() <-chan struct{} {
	s.firstCommitInTermMu.RLock()
	defer s.firstCommitInTermMu.RUnlock()
	return s.firstCommitInTermC
}

type confChangeResponse struct {
	membs []*membership.Member
	err   error
}

// configureAndSendRaft ??????raft??????????????????,???????????????????????????etcd.????????????,?????????????????????????????????.
func (s *EtcdServer) configureAndSendRaft(ctx context.Context, cc raftpb.ConfChangeV1) ([]*membership.Member, error) {
	lg := s.Logger()
	cc.ID = s.reqIDGen.Next()
	ch := s.w.Register(cc.ID)

	start := time.Now()
	if err := s.r.ProposeConfChange(ctx, cc); err != nil {
		s.w.Trigger(cc.ID, nil)
		return nil, err
	}

	select {
	case x := <-ch:
		if x == nil {
			lg.Panic("????????????")
		}
		resp := x.(*confChangeResponse)
		lg.Info(
			"??????raft??????????????????",
			zap.String("local-member-id", s.ID().String()),
			zap.String("raft-conf-change", cc.Type.String()),
			zap.String("raft-conf-change-node-id", types.ID(cc.NodeID).String()),
		)
		return resp.membs, resp.err

	case <-ctx.Done():
		s.w.Trigger(cc.ID, nil) // GC wait
		return nil, s.parseProposeCtxErr(ctx.Err(), start)

	case <-s.stopping:
		return nil, ErrStopped
	}
}

// sync proposes a SYNC request and is non-blocking.
// This makes no guarantee that the request will backend proposed or performed.
// The request will backend canceled after the given timeout.
func (s *EtcdServer) sync(timeout time.Duration) {
	req := pb.Request{
		Method: "SYNC",
		ID:     s.reqIDGen.Next(),
		Time:   time.Now().UnixNano(),
	}
	data := pbutil.MustMarshal(&req)
	// There is no promise that node has leader when do SYNC request,
	// so it uses goroutine to propose.
	ctx, cancel := context.WithTimeout(s.ctx, timeout)
	s.GoAttach(func() {
		s.r.Propose(ctx, data)
		cancel()
	})
}

// publish registers etcd information into the cluster. The information
// is the JSON representation of this etcd's member struct, updated with the
// static clientURLs of the etcd.
// The function keeps attempting to register until it succeeds,
// or its etcd is stopped.
//
// Use v2 store to encode member attributes, and apply through Raft
// but does not go through v2 API endpoint, which means even with v2
// client handler disabled (e.g. --enable-v2=false), cluster can still
// process publish requests through rafthttp
// TODO: Remove in 3.6 (start using publishV3)
func (s *EtcdServer) publish(timeout time.Duration) {
	lg := s.Logger()
	b, err := json.Marshal(s.attributes)
	if err != nil {
		lg.Panic("failed to marshal JSON", zap.Error(err))
		return
	}
	req := pb.Request{
		Method: "PUT",
		Path:   membership.MemberAttributesStorePath(s.id),
		Val:    string(b),
	}

	for {
		ctx, cancel := context.WithTimeout(s.ctx, timeout)
		_, err := s.Do(ctx, req)
		cancel()
		switch err {
		case nil:
			close(s.readych)
			lg.Info(
				"published local member to cluster through raft",
				zap.String("local-member-id", s.ID().String()),
				zap.String("local-member-attributes", fmt.Sprintf("%+v", s.attributes)),
				zap.String("request-path", req.Path),
				zap.String("cluster-id", s.cluster.ID().String()),
				zap.Duration("publish-timeout", timeout),
			)
			return

		case ErrStopped:
			lg.Warn(
				"stopped publish because etcd is stopped",
				zap.String("local-member-id", s.ID().String()),
				zap.String("local-member-attributes", fmt.Sprintf("%+v", s.attributes)),
				zap.Duration("publish-timeout", timeout),
				zap.Error(err),
			)
			return

		default:
			lg.Warn(
				"??????raft????????????????????????",
				zap.String("local-member-id", s.ID().String()),
				zap.String("local-member-attributes", fmt.Sprintf("%+v", s.attributes)),
				zap.String("request-path", req.Path),
				zap.Duration("publish-timeout", timeout),
				zap.Error(err),
			)
		}
	}
}

func (s *EtcdServer) sendMergedSnap(merged snap.Message) {
	atomic.AddInt64(&s.inflightSnapshots, 1)

	lg := s.Logger()
	fields := []zap.Field{
		zap.String("from", s.ID().String()),
		zap.String("to", types.ID(merged.To).String()),
		zap.Int64("bytes", merged.TotalSize),
		zap.String("size", humanize.Bytes(uint64(merged.TotalSize))),
	}

	now := time.Now()
	s.r.transport.SendSnapshot(merged)
	lg.Info("sending merged snapshot", fields...)

	s.GoAttach(func() {
		select {
		case ok := <-merged.CloseNotify():
			// delay releasing inflight snapshot for another 30 seconds to
			// block log compaction.
			// If the follower still fails to catch up, it is probably just too slow
			// to catch up. We cannot avoid the snapshot cycle anyway.
			if ok {
				select {
				case <-time.After(releaseDelayAfterSnapshot):
				case <-s.stopping:
				}
			}

			atomic.AddInt64(&s.inflightSnapshots, -1)

			lg.Info("sent merged snapshot", append(fields, zap.Duration("took", time.Since(now)))...)

		case <-s.stopping:
			lg.Warn("canceled sending merged snapshot; etcd stopping", fields...)
			return
		}
	})
}

// apply ???raft ?????????committed ---> applying
func (s *EtcdServer) apply(es []raftpb.Entry, confState *raftpb.ConfState) (appliedt uint64, appliedi uint64, shouldStop bool) {
	// confState ?????????????????? ????????????
	s.lg.Debug("??????????????????", zap.Int("num-entries", len(es)))
	for i := range es {
		e := es[i]
		s.lg.Debug("??????????????????", zap.Uint64("index", e.Index), zap.Uint64("term", e.Term), zap.Stringer("type", e.Type))
		switch e.Type {
		case raftpb.EntryNormal:
			s.applyEntryNormal(&e)
			s.setAppliedIndex(e.Index)
			s.setTerm(e.Term)

		case raftpb.EntryConfChange:
			shouldApplyV3 := membership.ApplyV2storeOnly    // false
			if e.Index > s.consistIndex.ConsistentIndex() { // ?????? bolt.db meta ????????? 	consistent_index???term
				s.consistIndex.SetConsistentIndex(e.Index, e.Term) // ??????????????????
				shouldApplyV3 = membership.ApplyBoth               // true
			}

			var cc raftpb.ConfChangeV1
			_ = cc.Unmarshal
			pbutil.MustUnmarshal(&cc, e.Data)
			// ConfChangeAddNode {"id":10276657743932975437,"peerURLs":["http://localhost:2380"],"name":"default"}
			removedSelf, err := s.applyConfChange(cc, confState, shouldApplyV3)
			s.setAppliedIndex(e.Index)
			s.setTerm(e.Term)
			shouldStop = shouldStop || removedSelf
			s.w.Trigger(cc.ID, &confChangeResponse{s.cluster.Members(), err})

		default:
			lg := s.Logger()
			lg.Panic(
				"?????????????????????;????????? EntryNormal ??? EntryConfChange",
				zap.String("type", e.Type.String()),
			)
		}
		appliedi, appliedt = e.Index, e.Term
	}
	return appliedt, appliedi, shouldStop
}

// TODO: non-blocking snapshot
func (s *EtcdServer) snapshot(snapi uint64, confState raftpb.ConfState) {
	clone := s.v2store.Clone()
	// commit kv to write metadata (for example: consistent index) to disk.
	//
	// This guarantees that Backend's consistent_index is >= index of last snapshot.
	//
	// KV().commit() updates the consistent index in backend.
	// All operations that update consistent index?????????called sequentially
	// from applyAll function.
	// So KV().Commit() cannot run in parallel with apply. It has to backend called outside
	// the go routine created below.
	s.KV().Commit()

	s.GoAttach(func() {
		lg := s.Logger()

		d, err := clone.SaveNoCopy()
		// TODO: current store will never fail to do a snapshot
		// what should we do if the store might fail?
		if err != nil {
			lg.Panic("failed to save v2 store", zap.Error(err))
		}
		snap, err := s.r.raftStorage.CreateSnapshot(snapi, &confState, d)
		if err != nil {
			// the snapshot was done asynchronously with the progress of raft.
			// raft might have already got a newer snapshot.
			if err == raft.ErrSnapOutOfDate {
				return
			}
			lg.Panic("failed to create snapshot", zap.Error(err))
		}
		// SaveSnap saves the snapshot to file and appends the corresponding WAL entry.
		if err = s.r.storage.SaveSnap(snap); err != nil {
			lg.Panic("failed to save snapshot", zap.Error(err))
		}
		if err = s.r.storage.Release(snap); err != nil {
			lg.Panic("failed to release wal", zap.Error(err))
		}

		lg.Info(
			"saved snapshot",
			zap.Uint64("snapshot-index", snap.Metadata.Index),
		)

		// When sending a snapshot, etcd will pause compaction.
		// After receives a snapshot, the slow follower needs to get all the entries right after
		// the snapshot sent to catch up. If we do not pause compaction, the log entries right after
		// the snapshot sent might already backend compacted. It happens when the snapshot takes long time
		// to send and save. Pausing compaction avoids triggering a snapshot sending cycle.
		if atomic.LoadInt64(&s.inflightSnapshots) != 0 {
			lg.Info("skip compaction since there is an inflight snapshot")
			return
		}

		// keep some in memory log entries for slow followers.
		compacti := uint64(1)
		if snapi > s.Cfg.SnapshotCatchUpEntries {
			compacti = snapi - s.Cfg.SnapshotCatchUpEntries // ???????????????????????????,??????follower????????????
		}

		err = s.r.raftStorage.Compact(compacti)
		if err != nil {
			// the compaction was done asynchronously with the progress of raft.
			// raft log might already been compact.
			if err == raft.ErrCompacted {
				return
			}
			lg.Panic("failed to compact", zap.Error(err))
		}
		lg.Info(
			"compacted Raft logs",
			zap.Uint64("compact-index", compacti),
		)
	})
}

// CutPeer drops messages to the specified peer.
func (s *EtcdServer) CutPeer(id types.ID) {
	tr, ok := s.r.transport.(*rafthttp.Transport)
	if ok {
		tr.CutPeer(id)
	}
}

// MendPeer recovers the message dropping behavior of the given peer.
func (s *EtcdServer) MendPeer(id types.ID) {
	tr, ok := s.r.transport.(*rafthttp.Transport)
	if ok {
		tr.MendPeer(id)
	}
}

func (s *EtcdServer) PauseSending() { s.r.pauseSending() }

func (s *EtcdServer) ResumeSending() { s.r.resumeSending() }

// monitorVersions checks the member's version every monitorVersionInterval.
// It updates the cluster version if all members agrees on a higher one.
// It prints out log if there is a member with a higher version than the
// local version.
func (s *EtcdServer) monitorVersions() {
	for {
		select {
		case <-s.FirstCommitInTermNotify():
		case <-time.After(monitorVersionInterval):
		case <-s.stopping:
			return
		}

		if s.Leader() != s.ID() {
			continue
		}

		v := decideClusterVersion(s.Logger(), getVersions(s.Logger(), s.cluster, s.id, s.peerRt))
		if v != nil {
			// only keep major.minor version for comparison
			v = &semver.Version{
				Major: v.Major,
				Minor: v.Minor,
			}
		}

		// if the current version is nil:
		// 1. use the decided version if possible
		// 2. or use the min cluster version
		if s.cluster.Version() == nil {
			verStr := version.MinClusterVersion
			if v != nil {
				verStr = v.String()
			}
			s.GoAttach(func() { s.updateClusterVersionV2(verStr) })
			continue
		}

		if v != nil && membership.IsValidVersionChange(s.cluster.Version(), v) {
			s.GoAttach(func() { s.updateClusterVersionV2(v.String()) })
		}
	}
}

func (s *EtcdServer) updateClusterVersionV2(ver string) {
	lg := s.Logger()
	if s.cluster.Version() == nil {
		lg.Info("??????v2 API ????????????????????????", zap.String("cluster-version", version.Cluster(ver)))
	} else {
		lg.Info("??????v2 API ????????????????????????", zap.String("from", version.Cluster(s.cluster.Version().String())), zap.String("to", version.Cluster(ver)))
	}

	req := pb.Request{
		Method: "PUT",
		Path:   membership.StoreClusterVersionKey(), // /0/version
		Val:    ver,
	}

	ctx, cancel := context.WithTimeout(s.ctx, s.Cfg.ReqTimeout())
	fmt.Println("start", time.Now())
	_, err := s.Do(ctx, req)
	fmt.Println("end", time.Now())
	cancel()

	switch err {
	case nil:
		lg.Info("?????????????????????", zap.String("cluster-version", version.Cluster(ver)))
		return

	case ErrStopped:
		lg.Warn("????????????????????????;etcd????????????", zap.Error(err))
		return

	default:
		lg.Warn("????????????????????????", zap.Error(err))
	}
}

func (s *EtcdServer) monitorDowngrade() {
	t := s.Cfg.DowngradeCheckTime
	if t == 0 {
		return
	}
	lg := s.Logger()
	for {
		select {
		case <-time.After(t):
		case <-s.stopping:
			return
		}

		if !s.isLeader() {
			continue
		}

		d := s.cluster.DowngradeInfo()
		if !d.Enabled {
			continue
		}

		targetVersion := d.TargetVersion
		v := semver.Must(semver.NewVersion(targetVersion))
		if isMatchedVersions(s.Logger(), v, getVersions(s.Logger(), s.cluster, s.id, s.peerRt)) {
			lg.Info("the cluster has been downgraded", zap.String("cluster-version", targetVersion))
			ctx, cancel := context.WithTimeout(context.Background(), s.Cfg.ReqTimeout())
			if _, err := s.downgradeCancel(ctx); err != nil {
				lg.Warn("failed to cancel downgrade", zap.Error(err))
			}
			cancel()
		}
	}
}

func (s *EtcdServer) parseProposeCtxErr(err error, start time.Time) error {
	switch err {
	case context.Canceled:
		return ErrCanceled

	case context.DeadlineExceeded:
		s.leadTimeMu.RLock()
		curLeadElected := s.leadElectedTime
		s.leadTimeMu.RUnlock()
		prevLeadLost := curLeadElected.Add(-2 * time.Duration(s.Cfg.ElectionTicks) * time.Duration(s.Cfg.TickMs) * time.Millisecond)
		if start.After(prevLeadLost) && start.Before(curLeadElected) {
			return ErrTimeoutDueToLeaderFail
		}
		lead := types.ID(s.getLead())
		switch lead {
		case types.ID(raft.None):
			// ????????????leader
		case s.ID(): // leader?????????
			if !isConnectedToQuorumSince(s.r.transport, start, s.ID(), s.cluster.Members()) { // ??????????????????????????????????????????
				return ErrTimeoutDueToConnectionLost
			}
		default:
			// ?????????????????????????????????,????????????????????????
			if !isConnectedSince(s.r.transport, start, lead) {
				return ErrTimeoutDueToConnectionLost
			}
		}
		return ErrTimeout

	default:
		return err
	}
}

func (s *EtcdServer) Backend() backend.Backend {
	s.backendLock.Lock()
	defer s.backendLock.Unlock()
	return s.backend
}

func (s *EtcdServer) AuthStore() auth.AuthStore { return s.authStore }

// ???????????????????????????
func (s *EtcdServer) restoreAlarms() error {
	s.applyV3 = s.newApplierV3()
	as, err := v3alarm.NewAlarmStore(s.lg, s)
	if err != nil {
		return err
	}
	s.alarmStore = as
	// ???????????????????????????
	if len(as.Get(pb.AlarmType_NOSPACE)) > 0 {
		s.applyV3 = newApplierV3Capped(s.applyV3)
	}
	if len(as.Get(pb.AlarmType_CORRUPT)) > 0 {
		s.applyV3 = newApplierV3Corrupt(s.applyV3)
	}
	return nil
}

// ----------------------------------------- OVER  --------------------------------------------------------------

// GoAttach ????????????????????????
func (s *EtcdServer) GoAttach(f func()) {
	s.wgMu.RLock() // stopping ?????? ???????????????
	defer s.wgMu.RUnlock()
	select {
	case <-s.stopping:
		lg := s.Logger()
		lg.Warn("etcd ?????????; ?????? GoAttach")
		return
	default:
	}

	// ???????????????????????????????????????????????????????????????.
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		f()
	}()
}

// applyEntryNormal ??????????????????raft??????
func (s *EtcdServer) applyEntryNormal(e *raftpb.Entry) {
	shouldApplyV3 := membership.ApplyV2storeOnly
	index := s.consistIndex.ConsistentIndex()
	if e.Index > index {
		// ????????????entry??????????????????
		s.consistIndex.SetConsistentIndex(e.Index, e.Term)
		shouldApplyV3 = membership.ApplyBoth // v2store ???bolt.db ???????????????
	}
	s.lg.Debug("????????????", zap.Uint64("consistent-index", index),
		zap.Uint64("entry-index", e.Index),
		zap.Bool("should-applyV3", bool(shouldApplyV3)))

	// ???leader?????????raft????????????????????????noop??????. ?????????????????????????????????????????????????????????.
	if len(e.Data) == 0 {
		s.notifyAboutFirstCommitInTerm() // ???????????? ?????????commit??????channel
		// ??????????????????leader ????????????????????????????????????????????????follower.
		if s.isLeader() {
			// ??????leader???,????????????????????????
			s.lessor.Promote(s.Cfg.ElectionTimeout())
		}
		return
	}
	// e.Data ?????? pb.InternalRaftRequest??? ??????????????????
	var raftReq pb.InternalRaftRequest
	if pbutil.MaybeUnmarshal(&raftReq, e.Data) {
	} else {
		// ??????????????????????????????
		// {"ID":7587861231285799684,"Method":"PUT","Path":"/0/version","Val":"3.5.0","Dir":false,"PrevValue":"","PrevIndex":0,"Expiration":0,"Wait":false,"Since":0,"Recursive":false,"Sorted":false,"Quorum":false,"Time":0,"Stream":false}
		var r pb.Request
		rp := &r
		pbutil.MustUnmarshal(rp, e.Data)
		s.w.Trigger(r.ID, s.applyV2Request((*RequestV2)(rp), shouldApplyV3))
		fmt.Println("pbutil.MustUnmarshal return")
		return
	}
	// ?????????
	//{"header":{"ID":7587861231285799685},"put":{"key":"YQ==","value":"Yg=="}}
	if raftReq.V2 != nil {
		req := (*RequestV2)(raftReq.V2)
		s.w.Trigger(req.ID, s.applyV2Request(req, shouldApplyV3))
		return
	}

	id := raftReq.ID
	if id == 0 {
		id = raftReq.Header.ID
	}

	var ar *applyResult
	needResult := s.w.IsRegistered(id)
	if needResult || !noSideEffect(&raftReq) {
		if !needResult && raftReq.Txn != nil {
			removeNeedlessRangeReqs(raftReq.Txn)
		}
		ar = s.applyV3.Apply(&raftReq, shouldApplyV3)
	}

	if !shouldApplyV3 { //  ???????????????bolt.db
		return
	}

	if ar == nil {
		return
	}

	if ar.err != ErrNoSpace || len(s.alarmStore.Get(pb.AlarmType_NOSPACE)) > 0 {
		s.w.Trigger(id, ar)
		return
	}

	lg := s.Logger()
	lg.Warn("??????????????????????????????????????????", zap.Int64("quota-size-bytes", s.Cfg.QuotaBackendBytes),
		zap.String("quota-size", humanize.Bytes(uint64(s.Cfg.QuotaBackendBytes))),
		zap.Error(ar.err),
	)

	s.GoAttach(func() {
		a := &pb.AlarmRequest{
			MemberID: uint64(s.ID()),
			Action:   pb.AlarmRequest_ACTIVATE, // ???????????????, ????????????
			Alarm:    pb.AlarmType_NOSPACE,
		}
		s.raftRequest(s.ctx, pb.InternalRaftRequest{Alarm: a})
		s.w.Trigger(id, ar)
	})
}

// ?????????????????????????????????commit
func (s *EtcdServer) notifyAboutFirstCommitInTerm() {
	newNotifier := make(chan struct{})
	s.firstCommitInTermMu.Lock()
	notifierToClose := s.firstCommitInTermC
	// ???????????????????????????
	s.firstCommitInTermC = newNotifier
	s.firstCommitInTermMu.Unlock()
	close(notifierToClose)
}

// IsLearner ????????????????????? raft learner
func (s *EtcdServer) IsLearner() bool {
	return s.cluster.IsLocalMemberLearner()
}

// IsMemberExist returns if the member with the given id exists in cluster.
func (s *EtcdServer) IsMemberExist(id types.ID) bool {
	return s.cluster.IsMemberExist(id)
}

// raftStatus ?????????????????????raft??????
func (s *EtcdServer) raftStatus() raft.Status {
	return s.r.RaftNodeInterFace.Status()
}

// ????????????
func maybeDefragBackend(cfg config.ServerConfig, be backend.Backend) error {
	size := be.Size()
	sizeInUse := be.SizeInUse()
	freeableMemory := uint(size - sizeInUse) // ??????
	thresholdBytes := cfg.ExperimentalBootstrapDefragThresholdMegabytes * 1024 * 1024
	if freeableMemory < thresholdBytes {
		cfg.Logger.Info("??????????????????",
			zap.Int64("current-db-size-bytes", size),
			zap.String("current-db-size", humanize.Bytes(uint64(size))),
			zap.Int64("current-db-size-in-use-bytes", sizeInUse),
			zap.String("current-db-size-in-use", humanize.Bytes(uint64(sizeInUse))),
			zap.Uint("experimental-bootstrap-defrag-threshold-bytes", thresholdBytes),
			zap.String("experimental-bootstrap-defrag-threshold", humanize.Bytes(uint64(thresholdBytes))),
		)
		return nil
	}
	return be.Defrag()
}

// applyConfChange ?????????confChange???????????????raft,???????????????committed
func (s *EtcdServer) applyConfChange(cc raftpb.ConfChangeV1, confState *raftpb.ConfState, shouldApplyV3 membership.ShouldApplyV3) (bool, error) {
	if err := s.cluster.ValidateConfigurationChange(cc); err != nil {
		cc.NodeID = raft.None // ??????,???????????????
		s.r.ApplyConfChange(cc)
		return false, err
	}

	lg := s.Logger()
	*confState = *s.r.ApplyConfChange(cc) // ?????????????????????
	s.beHooks.SetConfState(confState)
	switch cc.Type {
	// ??????????????????quorum.JointConfig???peer??????????????????
	case raftpb.ConfChangeAddNode, raftpb.ConfChangeAddLearnerNode:
		confChangeContext := new(membership.ConfigChangeContext)
		if err := json.Unmarshal([]byte(cc.Context), confChangeContext); err != nil {
			lg.Panic("????????????????????????", zap.Error(err))
		}
		if cc.NodeID != uint64(confChangeContext.Member.ID) {
			lg.Panic("?????????????????????ID",
				zap.String("member-id-from-config-change-entry", types.ID(cc.NodeID).String()),
				zap.String("member-id-from-message", confChangeContext.Member.ID.String()),
			)
		}
		if confChangeContext.IsPromote { // ??????????????????
			s.cluster.PromoteMember(confChangeContext.Member.ID, shouldApplyV3)
		} else {
			s.cluster.AddMember(&confChangeContext.Member, shouldApplyV3) // ????????????  /0/members/8e9e05c52164694d
			if confChangeContext.Member.ID != s.id {                      // ???????????????
				s.r.transport.AddPeer(confChangeContext.Member.ID, confChangeContext.PeerURLs)
			}
		}

	case raftpb.ConfChangeRemoveNode:
		id := types.ID(cc.NodeID)
		s.cluster.RemoveMember(id, shouldApplyV3) // ???
		if id == s.id {
			return true, nil
		}
		s.r.transport.RemovePeer(id)

	case raftpb.ConfChangeUpdateNode:
		m := new(membership.Member)
		if err := json.Unmarshal([]byte(cc.Context), m); err != nil {
			lg.Panic("??????????????????", zap.Error(err))
		}
		if cc.NodeID != uint64(m.ID) {
			lg.Panic("????????????????????????ID",
				zap.String("member-id-from-config-change-entry", types.ID(cc.NodeID).String()),
				zap.String("member-id-from-message", m.ID.String()),
			)
		}
		s.cluster.UpdateRaftAttributes(m.ID, m.RaftAttributes, shouldApplyV3)
		if m.ID != s.id {
			s.r.transport.UpdatePeer(m.ID, m.PeerURLs)
		}
	}
	return false, nil
}

// Alarms ?????????????????????,
func (s *EtcdServer) Alarms() []*pb.AlarmMember {
	return s.alarmStore.Get(pb.AlarmType_NONE)
}

func (s *EtcdServer) setCommittedIndex(v uint64) {
	atomic.StoreUint64(&s.committedIndex, v)
}

func (s *EtcdServer) getCommittedIndex() uint64 {
	return atomic.LoadUint64(&s.committedIndex)
}

func (s *EtcdServer) setAppliedIndex(v uint64) {
	atomic.StoreUint64(&s.appliedIndex, v)
}

func (s *EtcdServer) getAppliedIndex() uint64 {
	return atomic.LoadUint64(&s.appliedIndex)
}

func (s *EtcdServer) setTerm(v uint64) {
	atomic.StoreUint64(&s.term, v)
}

func (s *EtcdServer) getTerm() uint64 {
	return atomic.LoadUint64(&s.term)
}

func (s *EtcdServer) setLead(v uint64) {
	atomic.StoreUint64(&s.lead, v)
}

func (s *EtcdServer) getLead() uint64 {
	return atomic.LoadUint64(&s.lead)
}

func (s *EtcdServer) IsIDRemoved(id uint64) bool { return s.cluster.IsIDRemoved(types.ID(id)) }

func (s *EtcdServer) Cluster() api.Cluster { return s.cluster }

func (s *EtcdServer) LeaderChangedNotify() <-chan struct{} {
	s.leaderChangedMu.RLock()
	defer s.leaderChangedMu.RUnlock()
	return s.leaderChanged
}

func (s *EtcdServer) KV() mvcc.WatchableKV { return s.kv }

// Process ????????????raft????????????????????????etcd???raft?????????,??????ctx?????????.
func (s *EtcdServer) Process(ctx context.Context, m raftpb.Message) error {
	lg := s.Logger()
	// ??????????????????????????????????????????
	if s.cluster.IsIDRemoved(types.ID(m.From)) {
		lg.Warn("?????????????????????????????????raft?????????",
			zap.String("local-member-id", s.ID().String()),
			zap.String("removed-member-id", types.ID(m.From).String()),
		)
		return httptypes.NewHTTPError(http.StatusForbidden, "??????????????????????????????????????????")
	}
	// ???????????????????????????????????? req???
	if m.Type == raftpb.MsgApp {
		s.stats.RecvAppendReq(types.ID(m.From).String(), m.Size())
	}
	var _ raft.RaftNodeInterFace = raftNode{}
	//_ = raftNode{}.Step
	return s.r.Step(ctx, m)
}

func (s *EtcdServer) ClusterVersion() *semver.Version {
	if s.cluster == nil {
		return nil
	}
	return s.cluster.Version()
}
