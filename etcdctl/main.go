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

// Etcdctl是一个控制etcd的命令行应用程序.
package main

import (
	"fmt"
	"os"

	"github.com/ls-2018/etcd_cn/etcdctl/ctlv3"
)

const (
	apiEnv = "ETCDCTL_API"
)

func main() {
	apiv := os.Getenv(apiEnv)

	os.Unsetenv(apiEnv)
	if len(apiv) == 0 || apiv == "3" {
		ctlv3.MustStart()
		return
	}

	fmt.Fprintf(os.Stderr, "unsupported API version: %v\n", apiv)
	os.Exit(1)
}
