package raftpb

import (
	"encoding/json"
)

//func (m *ConfChange) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *ConfState) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *HardState) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *Message) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *Snapshot) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *SnapshotMetadata) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *Entry) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *ConfChangeSingle) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *ConfChangeV2) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}

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

// func (m *Entry) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
//			}
//			m.Type = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Type |= EntryType(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
//			}
//			m.Term = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Term |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 3:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
//			}
//			m.Index = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Index |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 4:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
//			}
//			var byteLen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				byteLen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if byteLen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + byteLen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
//			if m.Data == nil {
//				m.Data = []byte{}
//			}
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//
//func (m *SnapshotMetadata) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: SnapshotMetadata: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: SnapshotMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field ConfState", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if err := m.ConfState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 2:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
//			}
//			m.Index = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Index |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 3:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
//			}
//			m.Term = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Term |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//
//func (m *Snapshot) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
//			}
//			var byteLen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				byteLen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if byteLen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + byteLen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
//			if m.Data == nil {
//				m.Data = []byte{}
//			}
//			iNdEx = postIndex
//		case 2:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//
//func (m *Message) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: Message: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
//			}
//			m.Type = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Type |= MessageType(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
//			}
//			m.To = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.To |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 3:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
//			}
//			m.From = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.From |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 4:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
//			}
//			m.Term = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Term |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 5:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field LogTerm", wireType)
//			}
//			m.LogTerm = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.LogTerm |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 6:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
//			}
//			m.Index = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Index |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 7:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Entries = append(m.Entries, Entry{})
//			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 8:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
//			}
//			m.Commit = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Commit |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 9:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if err := m.Snapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 10:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Reject", wireType)
//			}
//			var v int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				v |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			m.Reject = bool(v != 0)
//		case 11:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field RejectHint", wireType)
//			}
//			m.RejectHint = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.RejectHint |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 12:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
//			}
//			var byteLen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				byteLen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if byteLen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + byteLen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Context = append(m.Context[:0], dAtA[iNdEx:postIndex]...)
//			if m.Context == nil {
//				m.Context = []byte{}
//			}
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//
//func (m *HardState) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: HardState: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: HardState: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
//			}
//			m.Term = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Term |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
//			}
//			m.Vote = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Vote |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 3:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
//			}
//			m.Commit = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Commit |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//
//func (m *ConfState) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: ConfState: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: ConfState: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType == 0 {
//				var v uint64
//				for shift := uint(0); ; shift += 7 {
//					if shift >= 64 {
//						return ErrIntOverflowRaft
//					}
//					if iNdEx >= l {
//						return io.ErrUnexpectedEOF
//					}
//					b := dAtA[iNdEx]
//					iNdEx++
//					v |= uint64(b&0x7F) << shift
//					if b < 0x80 {
//						break
//					}
//				}
//				m.Voters = append(m.Voters, v)
//			} else if wireType == 2 {
//				var packedLen int
//				for shift := uint(0); ; shift += 7 {
//					if shift >= 64 {
//						return ErrIntOverflowRaft
//					}
//					if iNdEx >= l {
//						return io.ErrUnexpectedEOF
//					}
//					b := dAtA[iNdEx]
//					iNdEx++
//					packedLen |= int(b&0x7F) << shift
//					if b < 0x80 {
//						break
//					}
//				}
//				if packedLen < 0 {
//					return ErrInvalidLengthRaft
//				}
//				postIndex := iNdEx + packedLen
//				if postIndex < 0 {
//					return ErrInvalidLengthRaft
//				}
//				if postIndex > l {
//					return io.ErrUnexpectedEOF
//				}
//				var elementCount int
//				var count int
//				for _, integer := range dAtA[iNdEx:postIndex] {
//					if integer < 128 {
//						count++
//					}
//				}
//				elementCount = count
//				if elementCount != 0 && len(m.Voters) == 0 {
//					m.Voters = make([]uint64, 0, elementCount)
//				}
//				for iNdEx < postIndex {
//					var v uint64
//					for shift := uint(0); ; shift += 7 {
//						if shift >= 64 {
//							return ErrIntOverflowRaft
//						}
//						if iNdEx >= l {
//							return io.ErrUnexpectedEOF
//						}
//						b := dAtA[iNdEx]
//						iNdEx++
//						v |= uint64(b&0x7F) << shift
//						if b < 0x80 {
//							break
//						}
//					}
//					m.Voters = append(m.Voters, v)
//				}
//			} else {
//				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
//			}
//		case 2:
//			if wireType == 0 {
//				var v uint64
//				for shift := uint(0); ; shift += 7 {
//					if shift >= 64 {
//						return ErrIntOverflowRaft
//					}
//					if iNdEx >= l {
//						return io.ErrUnexpectedEOF
//					}
//					b := dAtA[iNdEx]
//					iNdEx++
//					v |= uint64(b&0x7F) << shift
//					if b < 0x80 {
//						break
//					}
//				}
//				m.Learners = append(m.Learners, v)
//			} else if wireType == 2 {
//				var packedLen int
//				for shift := uint(0); ; shift += 7 {
//					if shift >= 64 {
//						return ErrIntOverflowRaft
//					}
//					if iNdEx >= l {
//						return io.ErrUnexpectedEOF
//					}
//					b := dAtA[iNdEx]
//					iNdEx++
//					packedLen |= int(b&0x7F) << shift
//					if b < 0x80 {
//						break
//					}
//				}
//				if packedLen < 0 {
//					return ErrInvalidLengthRaft
//				}
//				postIndex := iNdEx + packedLen
//				if postIndex < 0 {
//					return ErrInvalidLengthRaft
//				}
//				if postIndex > l {
//					return io.ErrUnexpectedEOF
//				}
//				var elementCount int
//				var count int
//				for _, integer := range dAtA[iNdEx:postIndex] {
//					if integer < 128 {
//						count++
//					}
//				}
//				elementCount = count
//				if elementCount != 0 && len(m.Learners) == 0 {
//					m.Learners = make([]uint64, 0, elementCount)
//				}
//				for iNdEx < postIndex {
//					var v uint64
//					for shift := uint(0); ; shift += 7 {
//						if shift >= 64 {
//							return ErrIntOverflowRaft
//						}
//						if iNdEx >= l {
//							return io.ErrUnexpectedEOF
//						}
//						b := dAtA[iNdEx]
//						iNdEx++
//						v |= uint64(b&0x7F) << shift
//						if b < 0x80 {
//							break
//						}
//					}
//					m.Learners = append(m.Learners, v)
//				}
//			} else {
//				return fmt.Errorf("proto: wrong wireType = %d for field Learners", wireType)
//			}
//		case 3:
//			if wireType == 0 {
//				var v uint64
//				for shift := uint(0); ; shift += 7 {
//					if shift >= 64 {
//						return ErrIntOverflowRaft
//					}
//					if iNdEx >= l {
//						return io.ErrUnexpectedEOF
//					}
//					b := dAtA[iNdEx]
//					iNdEx++
//					v |= uint64(b&0x7F) << shift
//					if b < 0x80 {
//						break
//					}
//				}
//				m.VotersOutgoing = append(m.VotersOutgoing, v)
//			} else if wireType == 2 {
//				var packedLen int
//				for shift := uint(0); ; shift += 7 {
//					if shift >= 64 {
//						return ErrIntOverflowRaft
//					}
//					if iNdEx >= l {
//						return io.ErrUnexpectedEOF
//					}
//					b := dAtA[iNdEx]
//					iNdEx++
//					packedLen |= int(b&0x7F) << shift
//					if b < 0x80 {
//						break
//					}
//				}
//				if packedLen < 0 {
//					return ErrInvalidLengthRaft
//				}
//				postIndex := iNdEx + packedLen
//				if postIndex < 0 {
//					return ErrInvalidLengthRaft
//				}
//				if postIndex > l {
//					return io.ErrUnexpectedEOF
//				}
//				var elementCount int
//				var count int
//				for _, integer := range dAtA[iNdEx:postIndex] {
//					if integer < 128 {
//						count++
//					}
//				}
//				elementCount = count
//				if elementCount != 0 && len(m.VotersOutgoing) == 0 {
//					m.VotersOutgoing = make([]uint64, 0, elementCount)
//				}
//				for iNdEx < postIndex {
//					var v uint64
//					for shift := uint(0); ; shift += 7 {
//						if shift >= 64 {
//							return ErrIntOverflowRaft
//						}
//						if iNdEx >= l {
//							return io.ErrUnexpectedEOF
//						}
//						b := dAtA[iNdEx]
//						iNdEx++
//						v |= uint64(b&0x7F) << shift
//						if b < 0x80 {
//							break
//						}
//					}
//					m.VotersOutgoing = append(m.VotersOutgoing, v)
//				}
//			} else {
//				return fmt.Errorf("proto: wrong wireType = %d for field VotersOutgoing", wireType)
//			}
//		case 4:
//			if wireType == 0 {
//				var v uint64
//				for shift := uint(0); ; shift += 7 {
//					if shift >= 64 {
//						return ErrIntOverflowRaft
//					}
//					if iNdEx >= l {
//						return io.ErrUnexpectedEOF
//					}
//					b := dAtA[iNdEx]
//					iNdEx++
//					v |= uint64(b&0x7F) << shift
//					if b < 0x80 {
//						break
//					}
//				}
//				m.LearnersNext = append(m.LearnersNext, v)
//			} else if wireType == 2 {
//				var packedLen int
//				for shift := uint(0); ; shift += 7 {
//					if shift >= 64 {
//						return ErrIntOverflowRaft
//					}
//					if iNdEx >= l {
//						return io.ErrUnexpectedEOF
//					}
//					b := dAtA[iNdEx]
//					iNdEx++
//					packedLen |= int(b&0x7F) << shift
//					if b < 0x80 {
//						break
//					}
//				}
//				if packedLen < 0 {
//					return ErrInvalidLengthRaft
//				}
//				postIndex := iNdEx + packedLen
//				if postIndex < 0 {
//					return ErrInvalidLengthRaft
//				}
//				if postIndex > l {
//					return io.ErrUnexpectedEOF
//				}
//				var elementCount int
//				var count int
//				for _, integer := range dAtA[iNdEx:postIndex] {
//					if integer < 128 {
//						count++
//					}
//				}
//				elementCount = count
//				if elementCount != 0 && len(m.LearnersNext) == 0 {
//					m.LearnersNext = make([]uint64, 0, elementCount)
//				}
//				for iNdEx < postIndex {
//					var v uint64
//					for shift := uint(0); ; shift += 7 {
//						if shift >= 64 {
//							return ErrIntOverflowRaft
//						}
//						if iNdEx >= l {
//							return io.ErrUnexpectedEOF
//						}
//						b := dAtA[iNdEx]
//						iNdEx++
//						v |= uint64(b&0x7F) << shift
//						if b < 0x80 {
//							break
//						}
//					}
//					m.LearnersNext = append(m.LearnersNext, v)
//				}
//			} else {
//				return fmt.Errorf("proto: wrong wireType = %d for field LearnersNext", wireType)
//			}
//		case 5:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AutoLeave", wireType)
//			}
//			var v int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				v |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			m.AutoLeave = bool(v != 0)
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//
//func (m *ConfChange) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: ConfChange: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: ConfChange: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
//			}
//			m.ID = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.ID |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
//			}
//			m.Type = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Type |= ConfChangeType(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 3:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
//			}
//			m.NodeID = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.NodeID |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 4:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
//			}
//			var byteLen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				byteLen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if byteLen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + byteLen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Context = append(m.Context[:0], dAtA[iNdEx:postIndex]...)
//			if m.Context == nil {
//				m.Context = []byte{}
//			}
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//
//func (m *ConfChangeSingle) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: ConfChangeSingle: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: ConfChangeSingle: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
//			}
//			m.Type = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Type |= ConfChangeType(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
//			}
//			m.NodeID = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.NodeID |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//
//func (m *ConfChangeV2) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaft
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: ConfChangeV2: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: ConfChangeV2: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Transition", wireType)
//			}
//			m.Transition = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Transition |= ConfChangeTransition(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Changes = append(m.Changes, ConfChangeSingle{})
//			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 3:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
//			}
//			var byteLen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaft
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				byteLen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if byteLen < 0 {
//				return ErrInvalidLengthRaft
//			}
//			postIndex := iNdEx + byteLen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.Context = append(m.Context[:0], dAtA[iNdEx:postIndex]...)
//			if m.Context == nil {
//				m.Context = []byte{}
//			}
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaft(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaft
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}

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
