package walpb

import (
	"encoding/json"
	"fmt"
)

//
//func (m *Record) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}

type Temp struct {
	Type int64
	Crc  uint32
	Data string
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	if m.Crc == 1285409499 {
		fmt.Println(123)
	}
	return json.Marshal(Temp{
		Type: m.Type,
		Crc:  m.Crc,
		Data: string(m.Data),
	})
}

//
//func (m *Record) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRecord
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
//			return fmt.Errorf("proto: Record: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
//			}
//			m.Type = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRecord
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Type |= int64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Crc", wireType)
//			}
//			m.Crc = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRecord
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.Crc |= uint32(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 3:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
//			}
//			var byteLen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRecord
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
//				return ErrInvalidLengthRecord
//			}
//			postIndex := iNdEx + byteLen
//			if postIndex < 0 {
//				return ErrInvalidLengthRecord
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
//			skippy, err := skipRecord(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRecord
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}

func (m *Record) Unmarshal(dAtA []byte) error {
	a := Temp{}
	err := json.Unmarshal(dAtA, &a)
	if err != nil {
		return err
	}
	m.Type = a.Type
	m.Crc = a.Crc
	m.Data = []byte(a.Data)
	return nil
}
