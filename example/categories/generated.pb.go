// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/zimbabao/proteus/example/protos/github.com/zimbabao/proteus/example/categories/generated.proto

/*
	Package categories is a generated protocol buffer package.

	It is generated from these files:
		github.com/zimbabao/proteus/example/protos/github.com/zimbabao/proteus/example/categories/generated.proto

	It has these top-level messages:
		CategoryOptions
*/
package categories

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func (m *CategoryOptions) Reset()                    { *m = CategoryOptions{} }
func (m *CategoryOptions) String() string            { return proto.CompactTextString(m) }
func (*CategoryOptions) ProtoMessage()               {}
func (*CategoryOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func init() {
	proto.RegisterType((*CategoryOptions)(nil), "gopkg.in.srcd.proteus.v1.example.categories.CategoryOptions")
}
func (m *CategoryOptions) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CategoryOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ShowPrices {
		dAtA[i] = 0x8
		i++
		if m.ShowPrices {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CanBuy {
		dAtA[i] = 0x10
		i++
		if m.CanBuy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CategoryOptions) ProtoSize() (n int) {
	var l int
	_ = l
	if m.ShowPrices {
		n += 2
	}
	if m.CanBuy {
		n += 2
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CategoryOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CategoryOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CategoryOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShowPrices", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ShowPrices = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanBuy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanBuy = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/zimbabao/proteus/example/protos/github.com/zimbabao/proteus/example/categories/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x4a, 0x4d, 0xcf, 0x2f, 0xc8,
	0x4e, 0xd7, 0xcb, 0xcc, 0xd3, 0x2f, 0x2e, 0x4a, 0xd6, 0x4d, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0x49,
	0x2d, 0x2d, 0xd6, 0x2b, 0x33, 0xd4, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x05, 0x0b, 0xe5,
	0x17, 0xeb, 0x13, 0x56, 0x98, 0x9c, 0x58, 0x92, 0x9a, 0x9e, 0x5f, 0x94, 0x99, 0x5a, 0xac, 0x9f,
	0x9e, 0x9a, 0x97, 0x5a, 0x94, 0x58, 0x92, 0x9a, 0xa2, 0x07, 0xd6, 0x2d, 0xa4, 0x0d, 0xd3, 0xad,
	0x57, 0x5c, 0x94, 0x0c, 0x11, 0x84, 0x68, 0xd6, 0x83, 0x6a, 0xd6, 0x43, 0x68, 0x96, 0xd2, 0x4d,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0x87, 0xb8,
	0x20, 0xa9, 0x34, 0x0d, 0xcc, 0x03, 0x73, 0xc0, 0x2c, 0x88, 0xd9, 0x4a, 0xa1, 0x5c, 0xfc, 0xce,
	0x10, 0xcd, 0x95, 0xfe, 0x05, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x42, 0xf2, 0x5c, 0xdc, 0xc5, 0x19,
	0xf9, 0xe5, 0xf1, 0x05, 0x45, 0x99, 0xc9, 0xa9, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41,
	0x5c, 0x20, 0xa1, 0x00, 0xb0, 0x88, 0x90, 0x38, 0x17, 0x7b, 0x72, 0x62, 0x5e, 0x7c, 0x52, 0x69,
	0xa5, 0x04, 0x13, 0x58, 0x92, 0x2d, 0x39, 0x31, 0xcf, 0xa9, 0xb4, 0xd2, 0x8a, 0xa3, 0x63, 0x81,
	0x3c, 0xc3, 0x87, 0x85, 0xf2, 0x0c, 0x4e, 0x0a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xc3, 0x8c, 0xc7, 0x72, 0x0c, 0x0b, 0x1e, 0xcb, 0x31, 0x46, 0x71, 0x21,
	0xdc, 0x99, 0xc4, 0x06, 0xb6, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x87, 0x07, 0x58, 0x2e,
	0x44, 0x01, 0x00, 0x00,
}
