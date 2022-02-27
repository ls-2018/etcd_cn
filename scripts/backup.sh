#!/bin/bash

# shellcheck disable=SC2006
date_time=`date +%Y%m%d`
etcdctl backup --data-dir /usr/local/etcd/niub3.etcd/ --backup-dir /niub/etcd_backup/"${date_time}"

find /niub/etcd_backup/ -ctime +7 -exec rm -r {} \;
