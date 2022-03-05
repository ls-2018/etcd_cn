# etcd

![etcd Logo](logos/etcd-horizontal-color.svg)

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

- 只保存一个小时的历史版本```etcd --auto-compaction-retention=1```
- 只保留最近的3个版本```etcdctl compact 3```
- 碎片整理```etcdctl defrag```
```

### URL

```
http://127.0.0.1:2379/members


```

### msg

```
MsgHup                      本地：开启选举,---->会触发vote或pre-vote
MsgBeat                     本地：心跳，---->给peers发送Msghearbeat
MsgProp                     本地：Propose -----> MsgApp
MsgApp                      非本地：操作日志【复制、配置变更 req】
MsgAppResp                  非本地：操作日志【复制 res】
MsgVote                     非本地：投票请求
MsgVoteResp                 非本地：投票相应
MsgPreVote                  非本地：预投票请求
MsgPreVoteResp              非本地：预投票相应
MsgSnap                     非本地：leader向follower拷贝快照,响应是MsgAppResp,告诉leader继续复制之后的值
MsgHeartbeat                
MsgHeartbeatResp            
MsgUnreachable              非本地：etcdserver通过这个消息告诉raft状态机某个follower不可达,让其发送消息的方式由pipeline切成ping-pong模式
MsgSnapStatus               非本地：etcdserver通过这个消息告诉raft状态机快照发送成功还是失败
MsgCheckQuorum               
MsgTransferLeader           非本地：
MsgTimeoutNow               非本地：
MsgReadIndex                非本地：
MsgReadIndexResp            非本地：
```

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


- checkquorum机制：
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

![](./images/MsgReadIndex.png)

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

```
tickHeartbeart 会同时推进两个计数器  heartbeatElapsed 和 electionElapsed .

(1) heartbeatElapsed

当 heartbeatElapsed 超时,发送 MsgBeat 消息给当前节点,当前节点收到消息之后会广播心跳消息(bcastHeartbeat)给其他节点 MsgHeartbeat 消息.

当 Follower 或者 Candidate 收到 MsgHeartbeat 消息会重置 electionElapsed 为 0,同时会响应 MsgHeartbeatResp 消息.

当 Leader 收到 MsgHeartbeatResp 消息,会更新对应节点的状态(存活、日志复制状态等)

(2) electionElapsed

当 electionElapsed 超时,发送 MsgCheckQuorum 给当前节点,当前节点收到消息之后,进行自我检查,判断是否能继续维持 Leader 状态,如果不能切换为Follower.同时如果节点正在进行 Leader 切换(切换其他节点为Leader),当 electionElapsed 超时,说明 Leader 节点转移超时,会终止切换.

```

```
grpc  client   --------grpc--------->    gateway ------------> etcd http server 2379
                            将grpc转换成了http 



127.0.0.1:2379
1、HTTP2
2、HTTP1     

```

### module

- github.com/soheilhy/cmux 可以在同一个listener上监听不同协议的请求
- 

