#!/bin/bash

infra1IP=192.168.1.100
infra2IP=192.168.1.101
infra3IP=192.168.1.102

/usr/local/etcd/etcd -name infra1 \
-initial-advertise-peer-urls http://$(infra1IP):2380 \
-listen-peer-urls http://$(infra1IP):2380 \
-listen-client-urls http://$(infra1IP):2379,http://127.0.0.1:2379 \
-advertise-client-urls http://$(infra1IP):2379 \
-initial-cluster-token etcd-cluster-1 \
-initial-cluster infra1=http://$(infra1IP):2380,infra2=http://$(infra2IP):2380,infra3=http://$(infra3IP):2380 \
-initial-cluster-state new >>/var/log/etcd_log/etcd.log 2>&1 &

/usr/local/etcd/etcd -name infra2 \
-initial-advertise-peer-urls http://$(infra2IP):2380 \
-listen-peer-urls http://$(infra2IP):2380 \
-listen-client-urls http://$(infra2IP):2379,http://127.0.0.1:2379 \
-advertise-client-urls http://$(infra2IP):2379 \
-initial-cluster-token etcd-cluster-1 \
-initial-cluster infra1=http://$(infra1IP):2380,infra2=http://$(infra2IP):2380,infra3=http://$(infra3IP):2380 \
-initial-cluster-state new >>/var/log/etcd_log/etcd.log 2>&1 &

/usr/local/etcd/etcd -name infra3 \
-initial-advertise-peer-urls http://$(infra3IP):2380 \
-listen-peer-urls http://$(infra3IP):2380 \
-listen-client-urls http://$(infra3IP):2379,http://127.0.0.1:2379 \
-advertise-client-urls http://$(infra3IP):2379 \
-initial-cluster-token etcd-cluster-1 \
-initial-cluster infra1=http://$(infra1IP):2380,infra2=http://$(infra2IP):2380,infra3=http://$(infra3IP):2380 \
-initial-cluster-state new >>/var/log/etcd_log/etcd.log 2>&1 &

iptables -A INPUT -p tcp -m state --state NEW -m tcp --dport 2379 -j ACCEPT
iptables -A INPUT -p tcp -m state --state NEW -m tcp --dport 2380 -j ACCEPT

echo -e "frontend etcd
    bind 10.10.0.14:2379
    mode tcp
    option tcplog
    default_backend etcd
    log 127.0.0.1 local3
    backend etcd
    balance roundrobin
    fullconn 1024
    server etcd1 $(infra1IP):2379 check port 2379 inter 300 fall 3
    server etcd1 $(infra2IP):2379 check port 2379 inter 300 fall 3
    server etcd1 $(infra3IP):2379 check port 2379 inter 300 fall 3
" >/etc/haproxy/haproxy.cfg

curl http://10.10.0.14:2379/v2/members
