// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/interchainquery/v1/interchainquery.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Query struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ChainId      string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	QueryType    string `protobuf:"bytes,4,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	Request      []byte `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	// change these to uint64 in v0.5.0
	Period       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=period,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"period"`
	LastHeight   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=last_height,json=lastHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_height"`
	CallbackId   string                                 `protobuf:"bytes,8,opt,name=callback_id,json=callbackId,proto3" json:"callback_id,omitempty"`
	Ttl          uint64                                 `protobuf:"varint,9,opt,name=ttl,proto3" json:"ttl,omitempty"`
	LastEmission github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=last_emission,json=lastEmission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_emission"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_e12f0828e1ddee43, []int{0}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Query.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return m.Size()
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Query) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *Query) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *Query) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Query) GetCallbackId() string {
	if m != nil {
		return m.CallbackId
	}
	return ""
}

func (m *Query) GetTtl() uint64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type DataPoint struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// change these to uint64 in v0.5.0
	RemoteHeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=remote_height,json=remoteHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"remote_height"`
	LocalHeight  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=local_height,json=localHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"local_height"`
	Value        []byte                                 `protobuf:"bytes,4,opt,name=value,proto3" json:"result,omitempty"`
}

func (m *DataPoint) Reset()         { *m = DataPoint{} }
func (m *DataPoint) String() string { return proto.CompactTextString(m) }
func (*DataPoint) ProtoMessage()    {}
func (*DataPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_e12f0828e1ddee43, []int{1}
}
func (m *DataPoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataPoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPoint.Merge(m, src)
}
func (m *DataPoint) XXX_Size() int {
	return m.Size()
}
func (m *DataPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPoint.DiscardUnknown(m)
}

var xxx_messageInfo_DataPoint proto.InternalMessageInfo

func (m *DataPoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataPoint) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Query)(nil), "quicksilver.interchainquery.v1.Query")
	proto.RegisterType((*DataPoint)(nil), "quicksilver.interchainquery.v1.DataPoint")
}

func init() {
	proto.RegisterFile("quicksilver/interchainquery/v1/interchainquery.proto", fileDescriptor_e12f0828e1ddee43)
}

var fileDescriptor_e12f0828e1ddee43 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x8b, 0xd3, 0x50,
	0x10, 0x6e, 0xda, 0x6d, 0xbb, 0x9d, 0xb6, 0xb2, 0x3c, 0xf6, 0x90, 0x5d, 0x30, 0x2d, 0x2b, 0x48,
	0x11, 0xdb, 0xb2, 0xe8, 0x51, 0x3c, 0x14, 0x05, 0x73, 0xd3, 0xb0, 0x07, 0x11, 0x24, 0xbc, 0x26,
	0x8f, 0xf4, 0xd1, 0x97, 0xf7, 0xd2, 0xbc, 0x49, 0x31, 0xbf, 0xc1, 0x8b, 0x3f, 0xc6, 0x1f, 0xb1,
	0xc7, 0xc5, 0x93, 0x28, 0x14, 0x69, 0x6f, 0xfe, 0x0a, 0xc9, 0x4b, 0x8a, 0x65, 0xf7, 0xda, 0x53,
	0x66, 0xe6, 0x9b, 0xf9, 0xe6, 0x9b, 0xc9, 0x1b, 0x78, 0xb9, 0xca, 0x78, 0xb0, 0xd4, 0x5c, 0xac,
	0x59, 0x3a, 0xe5, 0x12, 0x59, 0x1a, 0x2c, 0x28, 0x97, 0xab, 0x8c, 0xa5, 0xf9, 0x74, 0x7d, 0x7d,
	0x3f, 0x34, 0x49, 0x52, 0x85, 0x8a, 0x38, 0x07, 0x55, 0x93, 0xfb, 0x29, 0xeb, 0xeb, 0xcb, 0xf3,
	0x48, 0x45, 0xca, 0xa4, 0x4e, 0x0b, 0xab, 0xac, 0xba, 0xbc, 0x08, 0x94, 0x8e, 0x95, 0xf6, 0x4b,
	0xa0, 0x74, 0x4a, 0xe8, 0xea, 0x77, 0x03, 0x9a, 0x1f, 0x8a, 0x6a, 0xf2, 0x08, 0xea, 0x3c, 0xb4,
	0xad, 0xa1, 0x35, 0xea, 0x78, 0x75, 0x1e, 0x92, 0x27, 0xd0, 0x0f, 0x94, 0x94, 0x2c, 0x40, 0xae,
	0xa4, 0xcf, 0x43, 0xbb, 0x6e, 0xa0, 0xde, 0xff, 0xa0, 0x1b, 0x92, 0x0b, 0x38, 0x35, 0x02, 0x0a,
	0xbc, 0x61, 0xf0, 0xb6, 0xf1, 0xdd, 0x90, 0x3c, 0x06, 0x30, 0xb2, 0x7c, 0xcc, 0x13, 0x66, 0x9f,
	0x18, 0xb0, 0x63, 0x22, 0x37, 0x79, 0xc2, 0x88, 0x0d, 0xed, 0x94, 0xad, 0x32, 0xa6, 0xd1, 0x6e,
	0x0e, 0xad, 0x51, 0xcf, 0xdb, 0xbb, 0xe4, 0x06, 0x5a, 0x09, 0x4b, 0xb9, 0x0a, 0xed, 0x56, 0x51,
	0x34, 0x7b, 0x75, 0xbb, 0x19, 0xd4, 0x7e, 0x6d, 0x06, 0x4f, 0x23, 0x8e, 0x8b, 0x6c, 0x3e, 0x09,
	0x54, 0x5c, 0xcd, 0x50, 0x7d, 0xc6, 0x3a, 0x5c, 0x4e, 0x8b, 0x2e, 0x7a, 0xe2, 0x4a, 0xfc, 0xf1,
	0x7d, 0x0c, 0xd5, 0x88, 0xae, 0x44, 0xaf, 0xe2, 0x22, 0x9f, 0xa1, 0x2b, 0xa8, 0x46, 0x7f, 0xc1,
	0x78, 0xb4, 0x40, 0xbb, 0x7d, 0x04, 0x6a, 0x28, 0x08, 0xdf, 0x19, 0x3e, 0x32, 0x80, 0x6e, 0x40,
	0x85, 0x98, 0xd3, 0x60, 0x59, 0xec, 0xe2, 0xd4, 0x8c, 0x0b, 0xfb, 0x90, 0x1b, 0x92, 0x33, 0x68,
	0x20, 0x0a, 0xbb, 0x33, 0xb4, 0x46, 0x27, 0x5e, 0x61, 0x12, 0x0a, 0x7d, 0xa3, 0x88, 0xc5, 0x5c,
	0x6b, 0xae, 0xa4, 0x0d, 0x47, 0xd0, 0xd4, 0x2b, 0x28, 0xdf, 0x56, 0x8c, 0x57, 0x5f, 0xeb, 0xd0,
	0x79, 0x43, 0x91, 0xbe, 0x57, 0x5c, 0xe2, 0x83, 0x3f, 0x4c, 0xa1, 0x9f, 0xb2, 0x58, 0x21, 0xdb,
	0x2f, 0xa5, 0x7e, 0x0c, 0x01, 0x25, 0x65, 0xb5, 0x16, 0x1f, 0x7a, 0x42, 0x05, 0x54, 0xec, 0x3b,
	0x34, 0x8e, 0xd0, 0xa1, 0x6b, 0x18, 0xab, 0x06, 0xcf, 0xa0, 0xb9, 0xa6, 0x22, 0x2b, 0x1f, 0x58,
	0x6f, 0x76, 0xfe, 0x77, 0x33, 0x38, 0x4b, 0x99, 0xce, 0x04, 0x3e, 0x57, 0x31, 0x47, 0x16, 0x27,
	0x98, 0x7b, 0x65, 0xca, 0xec, 0xe3, 0xed, 0xd6, 0xb1, 0xee, 0xb6, 0x8e, 0xf5, 0x67, 0xeb, 0x58,
	0xdf, 0x76, 0x4e, 0xed, 0x6e, 0xe7, 0xd4, 0x7e, 0xee, 0x9c, 0xda, 0xa7, 0xd7, 0x07, 0x42, 0xb8,
	0x8c, 0x98, 0xcc, 0x38, 0xe6, 0xe3, 0x79, 0xc6, 0x45, 0x38, 0x3d, 0xbc, 0xd3, 0x2f, 0x0f, 0x2e,
	0xd5, 0x88, 0x9c, 0xb7, 0xcc, 0x31, 0xbd, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x48, 0xc3,
	0x47, 0xd5, 0x03, 0x00, 0x00,
}

func (m *Query) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Query) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Query) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LastEmission.Size()
		i -= size
		if _, err := m.LastEmission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterchainquery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.Ttl != 0 {
		i = encodeVarintInterchainquery(dAtA, i, uint64(m.Ttl))
		i--
		dAtA[i] = 0x48
	}
	if len(m.CallbackId) > 0 {
		i -= len(m.CallbackId)
		copy(dAtA[i:], m.CallbackId)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.CallbackId)))
		i--
		dAtA[i] = 0x42
	}
	{
		size := m.LastHeight.Size()
		i -= size
		if _, err := m.LastHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterchainquery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.Period.Size()
		i -= size
		if _, err := m.Period.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterchainquery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Request) > 0 {
		i -= len(m.Request)
		copy(dAtA[i:], m.Request)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.Request)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataPoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataPoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataPoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.LocalHeight.Size()
		i -= size
		if _, err := m.LocalHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterchainquery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RemoteHeight.Size()
		i -= size
		if _, err := m.RemoteHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterchainquery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterchainquery(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterchainquery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Query) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = len(m.Request)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = m.Period.Size()
	n += 1 + l + sovInterchainquery(uint64(l))
	l = m.LastHeight.Size()
	n += 1 + l + sovInterchainquery(uint64(l))
	l = len(m.CallbackId)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	if m.Ttl != 0 {
		n += 1 + sovInterchainquery(uint64(m.Ttl))
	}
	l = m.LastEmission.Size()
	n += 1 + l + sovInterchainquery(uint64(l))
	return n
}

func (m *DataPoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = m.RemoteHeight.Size()
	n += 1 + l + sovInterchainquery(uint64(l))
	l = m.LocalHeight.Size()
	n += 1 + l + sovInterchainquery(uint64(l))
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	return n
}

func sovInterchainquery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterchainquery(x uint64) (n int) {
	return sovInterchainquery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Query) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterchainquery
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
			return fmt.Errorf("proto: Query: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Query: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request[:0], dAtA[iNdEx:postIndex]...)
			if m.Request == nil {
				m.Request = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Period.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEmission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastEmission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterchainquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterchainquery
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
func (m *DataPoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterchainquery
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
			return fmt.Errorf("proto: DataPoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataPoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemoteHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LocalHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterchainquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterchainquery
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
func skipInterchainquery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterchainquery
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
					return 0, ErrIntOverflowInterchainquery
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
					return 0, ErrIntOverflowInterchainquery
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
				return 0, ErrInvalidLengthInterchainquery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterchainquery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterchainquery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterchainquery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterchainquery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterchainquery = fmt.Errorf("proto: unexpected end of group")
)