package walpb

import (
	"testing"

	"github.com/golang/protobuf/descriptor"
	"github.com/ls-2018/etcd_cn/raft/raftpb"
)

func TestSnapshotMetadataCompatibility(t *testing.T) {
	_, snapshotMetadataMd := descriptor.ForMessage(&raftpb.SnapshotMetadata{})
	_, snapshotMd := descriptor.ForMessage(&Snapshot{})
	if len(snapshotMetadataMd.GetField()) != len(snapshotMd.GetField()) {
		t.Errorf("Different number of fields in raftpb.SnapshotMetadata vs. walpb.Snapshot. " +
			"They are supposed to be in sync.")
	}
}
