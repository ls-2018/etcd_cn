package etcdserverpb

import (
	"encoding/json"
)

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *Metadata) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}
