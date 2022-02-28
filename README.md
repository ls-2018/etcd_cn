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
//--auto-compaction-mode=periodic --auto-compaction-retention=12h 每1小时自动压缩并保留12小时窗口。
👁etcd_backend/embed/config_test.go:TestAutoCompactionModeParse

- 只保存一个小时的历史版本```etcd --auto-compaction-retention=1```
- 只保留最近的3个版本```etcdctl compact 3```
- 碎片整理```etcdctl defrag```
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
  
  ```
- 3 JournalLogOutput 日志
  ```
  systemd-journal是syslog 的补充,收集来自内核、启动过程早期阶段、标准输出、系统日志、守护进程启动和运行期间错误的信息,
  它会默认把日志记录到/run/log/journal中,仅保留一个月的日志,且系统重启后也会消失.
  但是当新建 /var/log/journal 目录后,它又会把日志记录到这个目录中,永久保存.
  ```
