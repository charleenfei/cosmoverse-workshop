// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/v1/swap.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type MsgSwap struct {
	Sender string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Offer  types.Coin `protobuf:"bytes,2,opt,name=offer,proto3" json:"offer"`
	MinAsk types.Coin `protobuf:"bytes,3,opt,name=minAsk,proto3" json:"minAsk"`
	// information for how to send the tokens to intended receiver
	PortId    string `protobuf:"bytes,4,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	ChannelId string `protobuf:"bytes,5,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Receiver  string `protobuf:"bytes,6,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (m *MsgSwap) Reset()         { *m = MsgSwap{} }
func (m *MsgSwap) String() string { return proto.CompactTextString(m) }
func (*MsgSwap) ProtoMessage()    {}
func (*MsgSwap) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6c642f0d1337204, []int{0}
}
func (m *MsgSwap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwap.Merge(m, src)
}
func (m *MsgSwap) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwap) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwap.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwap proto.InternalMessageInfo

func (m *MsgSwap) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgSwap) GetOffer() types.Coin {
	if m != nil {
		return m.Offer
	}
	return types.Coin{}
}

func (m *MsgSwap) GetMinAsk() types.Coin {
	if m != nil {
		return m.MinAsk
	}
	return types.Coin{}
}

func (m *MsgSwap) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *MsgSwap) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *MsgSwap) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

type MsgSwapResponse struct {
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgSwapResponse) Reset()         { *m = MsgSwapResponse{} }
func (m *MsgSwapResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapResponse) ProtoMessage()    {}
func (*MsgSwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6c642f0d1337204, []int{1}
}
func (m *MsgSwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapResponse.Merge(m, src)
}
func (m *MsgSwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapResponse proto.InternalMessageInfo

func (m *MsgSwapResponse) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgSwap)(nil), "eightball.v1.MsgSwap")
	proto.RegisterType((*MsgSwapResponse)(nil), "eightball.v1.MsgSwapResponse")
}

func init() { proto.RegisterFile("eightball/v1/swap.proto", fileDescriptor_a6c642f0d1337204) }

var fileDescriptor_a6c642f0d1337204 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4f, 0xeb, 0x30,
	0x14, 0xc5, 0x93, 0xf7, 0xda, 0xf4, 0xd5, 0x0f, 0x09, 0x29, 0x42, 0x34, 0x54, 0x22, 0x54, 0x9d,
	0xba, 0x34, 0x56, 0x40, 0x88, 0x85, 0x85, 0x32, 0x75, 0x60, 0x09, 0x12, 0x03, 0x0b, 0x4a, 0x9c,
	0xdb, 0xc4, 0x6a, 0xea, 0x6b, 0xec, 0x34, 0x85, 0x6f, 0xc1, 0xc7, 0xea, 0xd8, 0x91, 0x09, 0xa1,
	0x56, 0x7c, 0x0f, 0x94, 0x3f, 0x2a, 0x2b, 0x9b, 0x8f, 0x7f, 0xc7, 0x47, 0x3e, 0xf7, 0x92, 0x1e,
	0xf0, 0x24, 0xcd, 0xa3, 0x30, 0xcb, 0x68, 0xe1, 0x53, 0xbd, 0x0a, 0xa5, 0x27, 0x15, 0xe6, 0x68,
	0x1f, 0xec, 0x81, 0x57, 0xf8, 0xfd, 0xa3, 0x04, 0x13, 0xac, 0x00, 0x2d, 0x4f, 0xb5, 0xa7, 0xef,
	0x32, 0xd4, 0x0b, 0xd4, 0x34, 0x0a, 0x35, 0xd0, 0xc2, 0x8f, 0x20, 0x0f, 0x7d, 0xca, 0x90, 0x8b,
	0x9a, 0x0f, 0xbf, 0x4c, 0xd2, 0xb9, 0xd3, 0xc9, 0xfd, 0x2a, 0x94, 0xf6, 0x31, 0xb1, 0x34, 0x88,
	0x18, 0x94, 0x63, 0x0e, 0xcc, 0x51, 0x37, 0x68, 0x94, 0x7d, 0x49, 0xda, 0x38, 0x9b, 0x81, 0x72,
	0xfe, 0x0c, 0xcc, 0xd1, 0xff, 0xf3, 0x13, 0xaf, 0xce, 0xf4, 0xca, 0x4c, 0xaf, 0xc9, 0xf4, 0x6e,
	0x91, 0x8b, 0x49, 0x6b, 0xfd, 0x71, 0x66, 0x04, 0xb5, 0xdb, 0xbe, 0x22, 0xd6, 0x82, 0x8b, 0x1b,
	0x3d, 0x77, 0xfe, 0xfe, 0xee, 0x5d, 0x63, 0xb7, 0x7b, 0xa4, 0x23, 0x51, 0xe5, 0x4f, 0x3c, 0x76,
	0x5a, 0xf5, 0x47, 0x4a, 0x39, 0x8d, 0xed, 0x53, 0x42, 0x58, 0x1a, 0x0a, 0x01, 0x59, 0xc9, 0xda,
	0x15, 0xeb, 0x36, 0x37, 0xd3, 0xd8, 0xee, 0x93, 0x7f, 0x0a, 0x18, 0xf0, 0x02, 0x94, 0x63, 0x55,
	0x70, 0xaf, 0x87, 0x63, 0x72, 0xd8, 0xd4, 0x0c, 0x40, 0x4b, 0x14, 0x1a, 0x4a, 0xbb, 0x86, 0xe7,
	0x25, 0x08, 0x06, 0x55, 0xe1, 0x56, 0xb0, 0xd7, 0x93, 0x87, 0xf5, 0xd6, 0x35, 0x37, 0x5b, 0xd7,
	0xfc, 0xdc, 0xba, 0xe6, 0xdb, 0xce, 0x35, 0x36, 0x3b, 0xd7, 0x78, 0xdf, 0xb9, 0xc6, 0xe3, 0x75,
	0xc2, 0xf3, 0x74, 0x19, 0x79, 0x0c, 0x17, 0x94, 0xa5, 0xa1, 0xca, 0x00, 0xc4, 0x0c, 0x38, 0xad,
	0xba, 0x15, 0xa0, 0x34, 0x8c, 0x57, 0xa8, 0xe6, 0x3a, 0x45, 0x49, 0x5f, 0xe8, 0xcf, 0xea, 0xf2,
	0x57, 0x09, 0x3a, 0xb2, 0xaa, 0xa9, 0x5f, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x33, 0xc1, 0x98,
	0x66, 0xd4, 0x01, 0x00, 0x00,
}

func (m *MsgSwap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.MinAsk.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Offer.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintSwap(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovSwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSwap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = m.Offer.Size()
	n += 1 + l + sovSwap(uint64(l))
	l = m.MinAsk.Size()
	n += 1 + l + sovSwap(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	return n
}

func (m *MsgSwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovSwap(uint64(m.Sequence))
	}
	return n
}

func sovSwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSwap(x uint64) (n int) {
	return sovSwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
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
			return fmt.Errorf("proto: MsgSwap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Offer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAsk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinAsk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
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
func (m *MsgSwapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
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
			return fmt.Errorf("proto: MsgSwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
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
func skipSwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSwap
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
					return 0, ErrIntOverflowSwap
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
					return 0, ErrIntOverflowSwap
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
				return 0, ErrInvalidLengthSwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSwap = fmt.Errorf("proto: unexpected end of group")
)