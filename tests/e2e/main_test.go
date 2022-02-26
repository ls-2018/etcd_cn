// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package e2e

import (
	"flag"
	"os"
	"runtime"
	"testing"

	"github.com/ls-2018/etcd_cn/client/pkg/testutil"
	"github.com/ls-2018/etcd_cn/tests/integration"
)

var (
	binDir  string
	certDir string

	certPath       string
	privateKeyPath string
	caPath         string

	certPath2       string
	privateKeyPath2 string

	certPath3       string
	privateKeyPath3 string

	crlPath               string
	revokedCertPath       string
	revokedPrivateKeyPath string
)

func TestMain(m *testing.M) {
	os.Setenv("ETCD_UNSUPPORTED_ARCH", runtime.GOARCH)
	os.Unsetenv("ETCDCTL_API")

	binDirDef := integration.MustAbsPath("../../bin")
	certDirDef := fixturesDir

	flag.StringVar(&binDir, "bin-dir", binDirDef, "The directory for store etcd and etcdctl binaries.")
	flag.StringVar(&certDir, "cert-dir", certDirDef, "The directory for store certificate files.")
	flag.Parse()

	binPath = binDir + "/etcd"
	ctlBinPath = binDir + "/etcdctl"
	utlBinPath = binDir + "/etcdutl"
	certPath = certDir + "/etcd.crt"
	privateKeyPath = certDir + "/etcd.key.insecure"
	caPath = certDir + "/ca.crt"
	revokedCertPath = certDir + "/etcd-revoked.crt"
	revokedPrivateKeyPath = certDir + "/etcd-revoked.key.insecure"
	crlPath = certDir + "/revoke.crl"

	certPath2 = certDir + "/server2.crt"
	privateKeyPath2 = certDir + "/server2.key.insecure"

	certPath3 = certDir + "/server3.crt"
	privateKeyPath3 = certDir + "/server3.key.insecure"

	v := m.Run()
	if v == 0 && testutil.CheckLeakedGoroutine() {
		os.Exit(1)
	}
	os.Exit(v)
}
