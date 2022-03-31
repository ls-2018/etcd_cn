// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"testing"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/testutil"
)

func TestMain(m *testing.M) {
	testutil.MustTestMainWithLeakDetection(m)
}
