// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tibc/core/types/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	types1 "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the ibc module's genesis state.
type GenesisState struct {
	// ICS002 - Clients genesis state
	ClientGenesis types.GenesisState `protobuf:"bytes,1,opt,name=client_genesis,json=clientGenesis,proto3" json:"client_genesis" yaml:"client_genesis"`
	// ICS004 - Packet genesis state
	PacketGenesis types1.GenesisState `protobuf:"bytes,3,opt,name=packet_genesis,json=packetGenesis,proto3" json:"packet_genesis" yaml:"packet_genesis"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7379bcbeb8a4fd03, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetClientGenesis() types.GenesisState {
	if m != nil {
		return m.ClientGenesis
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetPacketGenesis() types1.GenesisState {
	if m != nil {
		return m.PacketGenesis
	}
	return types1.GenesisState{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tibc.core.types.v1.GenesisState")
}

func init() { proto.RegisterFile("tibc/core/types/v1/genesis.proto", fileDescriptor_7379bcbeb8a4fd03) }

var fileDescriptor_7379bcbeb8a4fd03 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xc9, 0x4c, 0x4a,
	0xd6, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x4f,
	0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0xa9,
	0xd0, 0x03, 0xa9, 0xd0, 0x03, 0xab, 0xd0, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x52, 0x8a, 0x08, 0xb3, 0x92, 0x73, 0x32, 0x53, 0xf3, 0x4a,
	0x30, 0x0c, 0x43, 0x56, 0x52, 0x90, 0x98, 0x9c, 0x9d, 0x8a, 0xa9, 0x44, 0xe9, 0x01, 0x23, 0x17,
	0x8f, 0x3b, 0x44, 0x24, 0xb8, 0x24, 0xb1, 0x24, 0x55, 0x28, 0x9d, 0x8b, 0x0f, 0x62, 0x5c, 0x3c,
	0x54, 0xa1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xa2, 0x1e, 0xc2, 0x65, 0x10, 0x05, 0x7a,
	0x65, 0x86, 0x7a, 0xc8, 0x5a, 0x9d, 0x64, 0x4f, 0xdc, 0x93, 0x67, 0xf8, 0x74, 0x4f, 0x5e, 0xb4,
	0x32, 0x31, 0x37, 0xc7, 0x4a, 0x09, 0xd5, 0x18, 0xa5, 0x20, 0x5e, 0x88, 0x00, 0x54, 0x0b, 0xc8,
	0x22, 0x88, 0xa3, 0xe0, 0x16, 0x31, 0x63, 0x58, 0x04, 0x51, 0x40, 0xc8, 0x22, 0x54, 0x63, 0x94,
	0x82, 0x78, 0x21, 0x02, 0x50, 0x2d, 0x4e, 0x7e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0x65, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x94,
	0x99, 0x98, 0x97, 0x95, 0x99, 0x9a, 0x98, 0xa9, 0x0f, 0xb2, 0x5e, 0x37, 0x3d, 0x5f, 0x3f, 0x37,
	0x3f, 0xa5, 0x34, 0x27, 0xb5, 0x58, 0x1f, 0x2d, 0xce, 0x92, 0xd8, 0xc0, 0x21, 0x67, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x89, 0xa0, 0xb1, 0x88, 0xcd, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PacketGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ClientGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClientGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.PacketGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClientGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
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
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
