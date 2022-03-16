package raftpb

import (
	"encoding/json"
)

func (m *ConfChange) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *ConfState) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *HardState) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

type temp struct {
	Data     string
	Metadata SnapshotMetadata
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	t := temp{
		Data:     string(m.Data),
		Metadata: m.Metadata,
	}
	return json.Marshal(t)
}

func (m *SnapshotMetadata) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *ConfChangeSingle) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *ConfChangeV2) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *Entry) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *SnapshotMetadata) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *Snapshot) Unmarshal(dAtA []byte) error {
	t := temp{}
	err := json.Unmarshal(dAtA, &t)
	if err == nil {
		m.Data = []byte(t.Data)
		m.Metadata = t.Metadata
	}
	return err
}

func (m *Message) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *HardState) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *ConfState) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *ConfChange) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *ConfChangeSingle) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *ConfChangeV2) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}
