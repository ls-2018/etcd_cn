// Copyright 2016 The etcd Authors
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

package command

import (
	"fmt"
	"os"

	"github.com/ls-2018/etcd_cn/etcdutl/etcdutl"
	"github.com/ls-2018/etcd_cn/pkg/cobrautl"
	"github.com/spf13/cobra"
)

var defragDataDir string

func NewDefragCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "defrag",
		Short: "对给定端点的etcd成员的存储进行碎片整理",
		Run:   defragCommandFunc,
	}
	cmd.PersistentFlags().BoolVar(&epClusterEndpoints, "cluster", false, "使用集群成员列表中的所有端点")
	cmd.Flags().StringVar(&defragDataDir, "data-dir", "", "可选的.如果存在,对etcd不使用的数据目录进行碎片整理.")
	return cmd
}

func defragCommandFunc(cmd *cobra.Command, args []string) {
	if len(defragDataDir) > 0 {
		fmt.Fprintf(os.Stderr, "Use `etcdutl defrag` instead. The --data-dir is going to be decomissioned in v3.6.\n\n")
		err := etcdutl.DefragData(defragDataDir)
		if err != nil {
			cobrautl.ExitWithError(cobrautl.ExitError, err)
		}
	}

	failures := 0
	c := mustClientFromCmd(cmd)
	for _, ep := range endpointsFromCluster(cmd) {
		ctx, cancel := commandCtx(cmd)
		_, err := c.Defragment(ctx, ep)
		cancel()
		if err != nil {
			fmt.Fprintf(os.Stderr, "整理etcd成员失败 [%s] (%v)\n", ep, err)
			failures++
		} else {
			fmt.Printf("整理etcd成员完成[%s]\n", ep)
		}
	}

	if failures != 0 {
		os.Exit(cobrautl.ExitError)
	}
}
