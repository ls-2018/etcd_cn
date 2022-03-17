# etcd

![etcd Logo](logos/etcd-horizontal-color.svg)
v3.5.2
Etcd是分布式系统中最关键的数据的可靠的分布式键值存储,其重点是:

自己看源码 用  

### 配置

```
peer-cert-allowed-cn    允许的客户端证书CommonName    your name or your server's hostname
```

## 空间占用整理

```模拟 
设置etcd存储大小
etcd --quota-backend-bytes=$((16*1024*1024))

写爆磁盘
while [ 1 ]; do dd if=/dev/urandom bs=1024 count=1024 | ETCDCTL_API=3 etcdctl put key || break;done

查看endpoint状态
ETCDCTL_API=3 etcdctl --write-out=table endpoint status

查看alarm
ETCDCTL_API=3 etcdctl alarm list

清理碎片
ETCDCTL_API=3 etcdctl defrag

清理alarm
ETCDCTL_API=3 etcdctl alarm disarm
```

```
//--auto-compaction-mode=revision --auto-compaction-retention=1000 每5分钟自动压缩"latest revision" - 1000；
//--auto-compaction-mode=periodic --auto-compaction-retention=12h 每1小时自动压缩并保留12小时窗口.
👁etcd_backend/embed/config_test.go:TestAutoCompactionModeParse

- 只保存一个小时的历史版本`etcd --auto-compaction-retention=1`
- 只保留最近的3个版本`etcdctl compact 3`
- 碎片整理`etcdctl defrag`
```

### URL

```
http://127.0.0.1:2379/members

```

### msgType

| 消息类型 | 处理方 | 描述 |
| :--- | :--- | :--: |
| MsgHup | 节点支持 | 本地：开启选举,---->会触发vote或pre-vote |
| MsgBeat | Leader | 本地：心跳,---->给peers发送Msghearbeat |
| MsgProp | Leader、Candidate、Follower | 本地：Propose -----> MsgApp |
| MsgApp | Candidate、Follower | 非本地：操作日志【复制、配置变更 req】 |
| MsgAppResp | Leader | 非本地：操作日志【复制 res】 |
| MsgVote | 节点支持 | 非本地：投票请求 |
| MsgVoteResp | Candidate | 非本地：投票相应 |
| MsgPreVote | 节点支持 | 非本地：预投票请求 |
| MsgPreVoteResp | Candidate | 非本地：预投票相应 |
| MsgSnap | Candidate、Follower | 非本地：leader向follower拷贝快照,响应是MsgAppResp,告诉leader继续复制之后的值 |
| MsgHeartbeat | Candidate、Follower | |
| MsgHeartbeatResp | Leader | |
| MsgUnreachable | Leader | 非本地：etcdserver通过这个消息告诉raft状态机某个follower不可达,让其发送消息的方式由pipeline切成ping-pong模式 |
| MsgSnapStatus | Leader | 非本地：etcdserver通过这个消息告诉raft状态机快照发送成功还是失败 |
| MsgCheckQuorum | Leader | |
| MsgTransferLeader | Leader、Follower | 非本地： |
| MsgTimeoutNow | Candidate、Follower | 非本地： |
| MsgReadIndex | Leader、Follower | 非本地： |
| MsgReadIndexResp | Follower | 非本地： |

### issue

- 1、CertFile与ClientCertFile KeyFile与ClientKeyFile的区别
  ```
  在运行的过程中是配置的相同的;
  一般情况下,client与server是使用相同的ca进行的签发,   所有server端可以使用自己的私钥与证书验证client证书
  但如果不是同一个ca签发的; 那么就需要一个与client相同ca签发的证书文件与key
  
  ```
- 2、url
  ```
  
  	ErrUnsetAdvertiseClientURLsFlag = fmt.Errorf("--advertise-client-urls is required when --listen-client-urls is set explicitly")
	ErrLogRotationInvalidLogOutput  = fmt.Errorf("--log-outputs requires a single file path when --log-rotate-config-json is defined")
  
    --data-dir 指定节点的数据存储目录,这些数据包括节点ID,集群ID,集群初始化配置,Snapshot文件,若未指定—wal-dir,还会存储WAL文件;
    --wal-dir 指定节点的was文件的存储目录,若指定了该参数,wal文件会和其他数据文件分开存储.
  # member  
    这个参数是etcd服务器自己监听时用的,也就是说,监听本机上的哪个网卡,哪个端口
    --listen-client-urls        DefaultListenClientURLs = "http://192.168.1.100:2379"
    和成员之间通信的地址.用于监听其他etcd member的url
    --listen-peer-urls          DefaultListenPeerURLs   = "http://192.168.1.100:2380"
  
  # cluster
    就是客户端(etcdctl/curl等)跟etcd服务进行交互时请求的url
    --advertise-client-urls             http://127.0.0.1:2379,http://192.168.1.100:2379,http://10.10.10.10:2379      
    集群成员的 URL地址.且会通告群集的其余成员节点.  
    --initial-advertise-peer-urls       http://127.0.0.1:12380       告知集群其他节点url.           
    # 集群中所有节点的信息
    --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' 
  
  
    请求流程:
    etcdctl endpoints=http://192.168.1.100：2379 --debug ls
    首先与endpoints建立链接, 获取配置在advertise-client-urls的参数
    然后依次与每一个地址建立链接,直到操作成功
  
  
      --advertise-client-urls=https://192.168.1.100:2379
      --cert-file=/etc/kubernetes/pki/etcd/server.crt
      --client-cert-auth=true
  
      --initial-advertise-peer-urls=https://192.168.1.100:2380
      --initial-cluster=k8s-master01=https://192.168.1.100:2380
  
      --key-file=/etc/kubernetes/pki/etcd/server.key
      --listen-client-urls=https://127.0.0.1:2379,https://192.168.1.100:2379
      --listen-metrics-urls=http://127.0.0.1:2381
      --listen-peer-urls=https://192.168.1.100:2380
  
      --name=k8s-master01
  
      --peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt
      --peer-client-cert-auth=true
      --peer-key-file=/etc/kubernetes/pki/etcd/peer.key
  
      --peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
      --trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
    initial-advertise-peer-urls与initial-cluster要都包含
  
  ```
- 3 JournalLogOutput 日志
  ```
  systemd-journal是syslog 的补充,收集来自内核、启动过程早期阶段、标准输出、系统日志、守护进程启动和运行期间错误的信息,
  它会默认把日志记录到/run/log/journal中,仅保留一个月的日志,且系统重启后也会消失.
  但是当新建 /var/log/journal 目录后,它又会把日志记录到这个目录中,永久保存.
  ```


- checkquorum  过半机制：
  ```
  每隔一段时间,leader节点会尝试连接集群中的节点(发送心跳),如果发现自己可以连接到的节点个数没有超过半数,则主动切换成follower状态.
  这样在网络分区的情况下,旧的leader节点可以很快的知道自己已经过期了.
  ```


- PreVote优化
  ```
  当follower节点准备发起选举时候,先连接其他节点,并询问它们是否愿意参与选举(其他人是否能正常收到leader节点的信息),当有半数以上节点响应并参与则可以发起新一轮选举.
  解决分区之后节点重新恢复但term过大导致的leader选举问题
  ```
- WAL
  ```
  WAL全称是Write Ahead Log,是数据库中常用的持久化数据的方法.比如我们更新数据库的一条数据,如果直接找到这条数据并更新,
  可能会耗费比较长的时间.更快更安全的方式是先写一条Log数据到文件中,然后由后台线程来完成最终数据的更新,这条log中通常包含的是一条指令.
  ```
- 发送心跳消息的时候leader是怎么设置各个follower的commit？

- leader收到follower的心跳响应之后会怎么去修改对应的follower元数据呢？

- 快照 follower 当数据远落后于leader , leader会将快照发送过来 但由于网络原因,这一过程很慢 ,但是leader又生成了新的快照,wal没有旧的数据, 这时follower同步完,leader将最新新消息
  发送follower , follower reject ,但是此时wal已经没有对应的wal 又会发送新的快照, 这就会陷入死循环.....how? 看完源码再说吧
  ![](./images/MsgReadIndex.png)
- JointConfig 为什么是两个

### Ref

- https://blog.csdn.net/cuichongxin/article/details/118678009
- https://blog.csdn.net/crazyj4/category_10585293.html
- https://zhuanlan.zhihu.com/p/113149149
- https://blog.csdn.net/lkree/article/details/99085339
- https://blog.csdn.net/xxb249/category_8693355.html
- https://blog.csdn.net/luo222/article/details/98849114
- https://www.cnblogs.com/ricklz/category/2004842.html
- https://blog.csdn.net/skh2015java/category_9284671.html
- https://mp.weixin.qq.com/s/o_g5z77VZbImgTqjNBSktA
- https://www.jianshu.com/p/089a4c464c95
- https://www.coder55.com/article/10608
- https://www.freesion.com/article/93891147362/
- https://www.cnblogs.com/huaweiyuncce/p/10130522.html
- https://www.cnblogs.com/myd620/p/13189604.html
- https://blog.csdn.net/weixin_42663840/article/details/101039942
- https://www.jianshu.com/p/ff6aaa66ea0f
- https://www.jianshu.com/p/518f75c5bf55
- https://www.cnblogs.com/ricklz/p/15155095.html
- https://www.lixueduan.com/categories/
- https://zhuanlan.zhihu.com/p/452483457
- https://so.csdn.net/so/search?q=etcd&t=blog&u=devclouds
- https://blog.csdn.net/weixin_42663840/article/details/100005978
- https://www.jianshu.com/p/267e1d626c22
- https://blog.csdn.net/qq_40504851/category_10905363.html
- https://blog.csdn.net/cyq6239075/category_9756361.html             ☆
-【一文搞懂Raft算法】https://www.cnblogs.com/xybaby/p/10124083.html    
-【寻找一种易于理解的一致性算法（扩展版）】https://github.com/maemual/raft-zh_cn/blob/master/raft-zh_cn.md  
-【raft演示动画】https://raft.github.io/raftscope/index.html    
-【理解 raft 算法】https://sanyuesha.com/2019/04/18/raft/  
-【理解Raft一致性算法—一篇学术论文总结】https://mp.weixin.qq.com/s/RkMeYyUck1WQPjNiGvahKQ  
-【Raft协议原理详解】https://zhuanlan.zhihu.com/p/91288179  
-【Raft算法详解】https://zhuanlan.zhihu.com/p/32052223
-【etcd技术内幕】一本关于etcd不错的书籍  
-【高可用分布式存储 etcd 的实现原理】https://draveness.me/etcd-introduction/  
-【Raft 在 etcd 中的实现】https://blog.betacat.io/post/raft-implementation-in-etcd/  tickHeartbeart 会同时推进两个计数器  heartbeatElapsed 和 electionElapsed .
-【etcd Raft库解析】https://www.codedump.info/post/20180922-etcd-raft/  
-【etcd raft 设计与实现《一》】https://zhuanlan.zhihu.com/p/51063866    (1) heartbeatElapsed
-【raftexample 源码解读】https://zhuanlan.zhihu.com/p/91314329  
-【etcd实现-全流程分析】https://zhuanlan.zhihu.com/p/135891186    当 heartbeatElapsed 超时,发送 MsgBeat 消息给当前节点,当前节点收到消息之后会广播心跳消息(bcastHeartbeat)给其他节点 MsgHeartbeat 消息.
-【线性一致性和Raft】https://pingcap.com/zh/blog/linearizability-and-raft  
-【etcd raft 设计与实现《二》】https://zhuanlan.zhihu.com/p/51065416  当 Follower 或者 Candidate 收到 MsgHeartbeat 消息会重置 electionElapsed 为 0,同时会响应 MsgHeartbeatResp 消息.
-【《深入浅出etcd》part 3 – 解析etcd的日志同步机制】https://mp.weixin.qq.com/s/o_g5z77VZbImgTqjNBSktA    
-【一文入门ETCD】https://juejin.cn/post/6844904031186321416   
-【etcd：从应用场景到实现原理的全方位解读】https://www.infoq.cn/article/etcd-interpretation-application-scenario-implement-principle   
-【Etcd 架构与实现解析】http://jolestar.com/etcd-architecture/   
-【linux单节点和集群的etcd】https://www.jianshu.com/p/07ca88b6ff67   
-【软负载均衡与硬负载均衡、4层与7层负载均衡】https://cloud.tencent.com/developer/article/1446391   
-【Etcd Lock详解】https://tangxusc.github.io/blog/2019/05/etcd-lock%E8%AF%A6%E8%A7%A3/   
-【etcd基础与使用】https://zhuyasen.com/post/etcd.html   
-【ETCD核心机制解析】https://www.cnblogs.com/FG123/p/13632095.html      
-【etcd watch机制】http://liangjf.top/2019/12/31/110.etcd-watch%E6%9C%BA%E5%88%B6%E5%88%86%E6%9E%90/   
-【ETCD 源码学习--Watch(server)】https://www.codeleading.com/article/15455457381/   
-【etcdV3—watcher服务端源码解析】https://blog.csdn.net/stayfoolish_yj/article/details/104497233
- https://www.jianshu.com/p/f0a63762ac13
- https://www.zhihu.com/question/63995014/answer/2251115833
- https://blog.csdn.net/weixin_42017400/article/details/123174473
- https://www.jianshu.com/p/1621360b0b8e



(2) electionElapsed

当 electionElapsed 超时,发送 MsgCheckQuorum 给当前节点,当前节点收到消息之后,进行自我检查,判断是否能继续维持 Leader 状态,如果不能切换为Follower.同时如果节点正在进行 Leader 切换(切换其他节点为Leader),当 electionElapsed 超时,说明 Leader 节点转移超时,会终止切换.


```
curl    --------http--------->    gateway ------------> etcd grpc server 2379
                            将http转换成了grpc



127.0.0.1:2379
1、HTTP2
2、HTTP1     

```

### module

- github.com/soheilhy/cmux 可以在同一个listener上监听不同协议的请求
-


```
etcdServer 会单独处理 Propose消息, 其余消息交给raft.step 来处理 [该函数,会随着节点角色的改变而发生改变] [会首先判断任期、索引,在判断消息类型]

StartEtcd
  1、etcdserver.NewServer -> 
      heartbeat := time.Duration(cfg.TickMs) * time.Millisecond
      MySelfStartRaft
        newRaftNode
          r.ticker = time.NewTicker(r.heartbeat)                创建定时器、心跳
        startNode -> 
          raft.StartNode -> 
            go n.run()
              rd = n.rn.readyWithoutAccept()                      获取待发送消息,会获取到r.msgs
              readyc = n.readyc                                   待发送消息channel 
              - case pm := <-propc                                网络发来的消息、除Propose消息
              - case m := <-n.recvc                               G 处理来自peer的消息 
              - case cc := <-n.confc
              - case <-n.tickc                                    F取出数据
                  n.rn.Tick()
                    rn.raft.tick()                                根据角色调用自己的函数
                      - r.tickElection 
                          r.Step(pb.Message{From: r.id, Type: pb.MsgHup}) 该函数是处理所有到来消息的入口
                            r.send(pb.Message
                              r.msgs = append(r.msgs, m)          放入要发送的消息
                      - r.tickHeartbeat
                          r.Step(pb.Message{From: r.id, Type: pb.MsgCheckQuorum})
              - case readyc <- rd                                                   A放入数据
              - case <-advancec: 
              - case c := <-n.status:
              - case <-n.stop: 
      tr.AddPeer  
        startPeer                                                 与每个peer都建立一个链接
          r.Process
            s.Process
              s.r.Step(ctx, m)
                n.step
                  stepWithWaitOption
                    case n.recvc <- m                             G 接收来自peer的消息 
        
  2、e.Server.Start ->
    EtcdServer.strat ->
      s.start()
        go s.run() 
            --> | # s.r=raftNode
            --> | s.r.start(rh)
            --> |   go func()
            --> |     - case <-r.ticker.C:                          接收定时器信号
            --> |       r.tick() 
            --> |         r.Tick()
            --> |           case n.tickc <- struct{}{}              F放入数据、不会阻塞,有size
            --> |     - case rd := <-r.Ready()                      获取可以发送的数据     A取出数据  
                          case r.applyc <- ap                                           B放入数据
                          r.transport.Send(msgs)                        发出响应数据
            --> |     - case <-r.stopped:
            
          - case ap := <-s.r.apply()                                                   B取出数据
              读取applyc的数据,封装为JOB,放入调度器
          - case leases := <-expiredLeaseC
              处理过期租约
          - case err := <-s.errorc
              处理运行过程中出现的err,直接退出
          - getSyncC
          - case <-s.stop:
              启动过程中失败
    
  3、e.servePeers
  4、e.serveClients
  5、e.serveMetrics
    
```

```
快照 + storage + unstable 的区别
compacted <--- compacted <--- applied <--- committed <--- stable <--- unstable
WAL 日志 
11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46
-------------------------------------------------------------------|--MemoryStorage|file--|-----------------
-----压缩---------|---------------压缩---------------------|-------------------------------------------------
----快照----------------- |----------------快照--------------------- | storage: 落盘的       | unstable 内存中的  
----快照----------------- |----------------快照--------------------- |         | 在没有被持久化之前如果遇到了换届选举,这个日志可能会被相同索引值的新日志覆盖


每次用户提交日志,该日志会保存到 MemoryStorage 以及wal里，每当raft发送给上层程序一批已经commited日志，就会触发maybeTriggerSnapshot,当用户apply以后
判断是否进行触发 MemoryStorage 打快照,当打了快照以后，会把当前快照点10000条以前的记录从 MemoryStorage.ents去除掉 【俗称压缩】



MemoryStorage并不是持久存储啊,其实日志是被双写了,文件和MemoryStorage各存储了一份,
而raft包只能访问MemoryStorage中的内容.这样设计的目的是用内存缓冲文件中的日志,在频繁操作日志的时候性能会更高.
此处需要注意的是,MemoryStorage中的日志仅仅代表日志是可靠的,与提交和应用没有任何关系.
```

### WAl数据日志数据

```
type Record struct {
    Type                 int64   
    Crc                  uint32  
    Data                 []byte   
}
- metadataType  :1  元数据类型,元数据会保存当前的node id和cluster id.
    type Metadata struct {
        NodeID    uint64
        ClusterID uint64
    }
- entryType     :2  日志条目
    type Entry struct {
        Term  uint64
        Index uint64
        Type  EntryType  
                EntryNormal
                  # msgType
                EntryConfChange   
                EntryConfChangeV2 
        Data  []byte 
              msg 👆🏻
    }
- stateType     :3  当前Term,当前竞选者、当前已经commit的日志.
- crcType       :4  存放crc校验字段 Data为nil
- snapshotType  :5  快照的、日志的Index和Term
    type Snapshot struct {
        Index     uint64
        Term      uint64
        ConfState *raftpb.ConfState
    }
```

```
raft commit->apply  的数据 封装在ready结构体里   <-r.Ready()
    raftNode拿到该ready做一些处理,过滤出操作日志   publishEntries
        上层应用拿到过滤后的,将其应用到kvstore
```


集群节点变更：
1、先检查是否有待应用的变更
2、将变更信息放入raft unstable 等待发送----->发送,等到apply
3、apply 该变更
- ap := <-s.r.apply()
  - s.applyAll(&ep, &ap)
    - s.applyEntries(ep, apply)
      - s.apply(ents, &ep.confState)
        - case raftpb.EntryConfChange:
          - s.applyConfChange(cc, confState, shouldApplyV3)
