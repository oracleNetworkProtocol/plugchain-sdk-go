// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/bank/v1beta1/tx.proto

package bank

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_oracleNetworkProtocol_plugchain_sdk_go_types "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	types "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
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

// MsgSend represents a message to send coins from one account to another.
type MsgSend struct {
	FromAddress string                                          `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress   string                                          `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	Amount      github_com_oracleNetworkProtocol_plugchain_sdk_go_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/oracleNetworkProtocol/plugchain-sdk-go/types.Coins" json:"amount"`
}

func (m *MsgSend) Reset()         { *m = MsgSend{} }
func (m *MsgSend) String() string { return proto.CompactTextString(m) }
func (*MsgSend) ProtoMessage()    {}
func (*MsgSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8cb1613481f5b7, []int{0}
}
func (m *MsgSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSend.Merge(m, src)
}
func (m *MsgSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSend proto.InternalMessageInfo

// MsgMultiSend represents an arbitrary multi-in, multi-out send message.
type MsgMultiSend struct {
	Inputs  []Input  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs"`
	Outputs []Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs"`
}

func (m *MsgMultiSend) Reset()         { *m = MsgMultiSend{} }
func (m *MsgMultiSend) String() string { return proto.CompactTextString(m) }
func (*MsgMultiSend) ProtoMessage()    {}
func (*MsgMultiSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8cb1613481f5b7, []int{1}
}
func (m *MsgMultiSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMultiSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMultiSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMultiSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMultiSend.Merge(m, src)
}
func (m *MsgMultiSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgMultiSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMultiSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMultiSend proto.InternalMessageInfo

func (m *MsgMultiSend) GetInputs() []Input {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *MsgMultiSend) GetOutputs() []Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSend)(nil), "cosmos.bank.v1beta1.MsgSend")
	proto.RegisterType((*MsgMultiSend)(nil), "cosmos.bank.v1beta1.MsgMultiSend")
}

func init() { proto.RegisterFile("cosmos/bank/v1beta1/tx.proto", fileDescriptor_1d8cb1613481f5b7) }

var fileDescriptor_1d8cb1613481f5b7 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0xef, 0xd2, 0x40,
	0x18, 0xc7, 0x5b, 0x20, 0x20, 0x07, 0x8b, 0x45, 0x23, 0xa2, 0x69, 0x49, 0x27, 0x16, 0xee, 0x44,
	0x34, 0x9a, 0x3a, 0x59, 0x27, 0x4d, 0x88, 0x49, 0xdd, 0x5c, 0xcc, 0xf5, 0x8f, 0xa5, 0xd2, 0xf6,
	0x69, 0xb8, 0xab, 0x81, 0x77, 0x60, 0xe2, 0xe2, 0x4b, 0x60, 0xf6, 0x95, 0x30, 0x32, 0x3a, 0xa1,
	0x81, 0xc5, 0x55, 0x5e, 0x81, 0xb9, 0x6b, 0xa9, 0x24, 0x32, 0xfc, 0xb6, 0xbb, 0x7c, 0x9e, 0xcf,
	0xf3, 0x3c, 0xdf, 0x3c, 0xe8, 0xa1, 0x07, 0x2c, 0x01, 0x46, 0x5c, 0x9a, 0x2e, 0xc8, 0xe7, 0x89,
	0x1b, 0x70, 0x3a, 0x21, 0x7c, 0x85, 0xb3, 0x25, 0x70, 0xd0, 0x7a, 0x05, 0xc5, 0x82, 0xe2, 0x92,
	0x0e, 0xee, 0x84, 0x10, 0x82, 0xe4, 0x44, 0xbc, 0x8a, 0xd2, 0x81, 0x5e, 0x35, 0x62, 0x41, 0xd5,
	0xc8, 0x83, 0x28, 0xfd, 0x8f, 0x5f, 0x0c, 0x92, 0x7d, 0x25, 0x37, 0xff, 0xa8, 0xa8, 0x35, 0x63,
	0xe1, 0xbb, 0x20, 0xf5, 0x35, 0x0b, 0x75, 0x3f, 0x2e, 0x21, 0xf9, 0x40, 0x7d, 0x7f, 0x19, 0x30,
	0xd6, 0x57, 0x87, 0xea, 0xa8, 0x6d, 0xdf, 0x3b, 0xed, 0x8d, 0xde, 0x9a, 0x26, 0xb1, 0x65, 0x5e,
	0x52, 0xd3, 0xe9, 0x88, 0xef, 0xcb, 0xe2, 0xa7, 0x3d, 0x41, 0x88, 0x43, 0x65, 0xd6, 0xa4, 0x79,
	0xf7, 0xb4, 0x37, 0x6e, 0x17, 0xe6, 0x3f, 0x66, 0x3a, 0x6d, 0x0e, 0x67, 0xeb, 0x13, 0x6a, 0xd2,
	0x04, 0xf2, 0x94, 0xf7, 0xeb, 0xc3, 0xfa, 0xa8, 0xf3, 0xf8, 0x3e, 0xae, 0x92, 0xb3, 0xe0, 0x9c,
	0x1c, 0xbf, 0x82, 0x28, 0xb5, 0x9f, 0x6d, 0xf7, 0x86, 0xf2, 0xfd, 0xa7, 0x41, 0xc2, 0x88, 0xcf,
	0x73, 0x17, 0x7b, 0x90, 0x90, 0x55, 0x44, 0x17, 0xf4, 0xe9, 0x94, 0x64, 0x71, 0x1e, 0x7a, 0x73,
	0x1a, 0xa5, 0x63, 0xe6, 0x2f, 0xc6, 0x21, 0x10, 0xbe, 0xce, 0x02, 0x26, 0x3d, 0xe6, 0x94, 0x13,
	0xac, 0x5b, 0x5f, 0x36, 0x86, 0xf2, 0x7b, 0x63, 0x28, 0xe6, 0x57, 0x15, 0x75, 0x67, 0x2c, 0x9c,
	0xe5, 0x31, 0x8f, 0x64, 0xf0, 0xe7, 0xa8, 0x19, 0xa5, 0x59, 0xce, 0x45, 0x64, 0xb1, 0xc6, 0x00,
	0x5f, 0x39, 0x00, 0x7e, 0x2d, 0x4a, 0xec, 0x86, 0xd8, 0xc3, 0x29, 0xeb, 0xb5, 0x17, 0xa8, 0x05,
	0x39, 0x97, 0x6a, 0x4d, 0xaa, 0x0f, 0xae, 0xaa, 0x6f, 0x65, 0x4d, 0xe9, 0x9e, 0x0d, 0xab, 0x21,
	0xb6, 0xb1, 0xdf, 0x6c, 0x0f, 0xba, 0xba, 0x3b, 0xe8, 0xea, 0xaf, 0x83, 0xae, 0x7e, 0x3b, 0xea,
	0xca, 0xee, 0xa8, 0x2b, 0x3f, 0x8e, 0xba, 0xf2, 0xfe, 0xd1, 0x4d, 0xa2, 0x26, 0xe0, 0xe7, 0x71,
	0x50, 0x1c, 0xd8, 0x6d, 0xca, 0xa3, 0x4e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x0d, 0xd4,
	0x79, 0x5f, 0x02, 0x00, 0x00,
}

func (m *MsgSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMultiSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMultiSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMultiSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Outputs) > 0 {
		for iNdEx := len(m.Outputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Outputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Inputs) > 0 {
		for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgMultiSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Inputs) > 0 {
		for _, e := range m.Inputs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Outputs) > 0 {
		for _, e := range m.Outputs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgMultiSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMultiSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMultiSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inputs = append(m.Inputs, Input{})
			if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Outputs = append(m.Outputs, Output{})
			if err := m.Outputs[len(m.Outputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
