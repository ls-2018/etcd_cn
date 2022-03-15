package membership

import (
	"testing"

	"github.com/coreos/go-semver/semver"
	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
	betesting "github.com/ls-2018/etcd_cn/etcd_backend/mvcc/backend/testing"
	"github.com/stretchr/testify/assert"

	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc/backend"
	"go.uber.org/zap"
)

func TestAddRemoveMember(t *testing.T) {
	c := newTestCluster(t, nil)
	be, bepath := betesting.NewDefaultTmpBackend(t)
	c.SetBackend(be)
	c.AddMember(newTestMember(17, nil, "node17", nil), true)
	c.RemoveMember(17, true)
	c.AddMember(newTestMember(18, nil, "node18", nil), true)

	// Skipping removal of already removed member
	c.RemoveMember(17, true)
	err := be.Close()
	assert.NoError(t, err)

	be2 := backend.NewDefaultBackend(bepath)
	defer func() {
		assert.NoError(t, be2.Close())
	}()

	if false {
		// TODO: Enable this code when Recover is reading membership from the backend.
		c2 := newTestCluster(t, nil)
		c2.SetBackend(be2)
		c2.Recover(func(*zap.Logger, *semver.Version) {})
		assert.Equal(t, []*Member{{
			ID:         types.ID(18),
			Attributes: Attributes{Name: "node18"},
		}}, c2.Members())
		assert.Equal(t, true, c2.IsIDRemoved(17))
		assert.Equal(t, false, c2.IsIDRemoved(18))
	}
}
