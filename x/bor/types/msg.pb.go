// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/bor/v1beta1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgProposeSpan struct {
	SpanId     uint64 `protobuf:"varint,1,opt,name=span_id,json=spanId,proto3" json:"span_id" yaml:"span_id"`
	Proposer   string `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer" yaml:"proposer"`
	StartBlock uint64 `protobuf:"varint,3,opt,name=start_block,json=startBlock,proto3" json:"start_block" yaml:"start_block"`
	EndBlock   uint64 `protobuf:"varint,4,opt,name=end_block,json=endBlock,proto3" json:"end_block" yaml:"end_block"`
	BorChainId string `protobuf:"bytes,5,opt,name=bor_chain_id,json=borChainId,proto3" json:"bor_chain_id" yaml:"bor_chain_id"`
	Seed       string `protobuf:"bytes,6,opt,name=seed,proto3" json:"seed" yaml:"seed"`
}

func (m *MsgProposeSpan) Reset()         { *m = MsgProposeSpan{} }
func (m *MsgProposeSpan) String() string { return proto.CompactTextString(m) }
func (*MsgProposeSpan) ProtoMessage()    {}
func (*MsgProposeSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d99edf48c57200b, []int{0}
}
func (m *MsgProposeSpan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposeSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposeSpan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposeSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposeSpan.Merge(m, src)
}
func (m *MsgProposeSpan) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposeSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposeSpan.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposeSpan proto.InternalMessageInfo

func (m *MsgProposeSpan) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *MsgProposeSpan) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *MsgProposeSpan) GetStartBlock() uint64 {
	if m != nil {
		return m.StartBlock
	}
	return 0
}

func (m *MsgProposeSpan) GetEndBlock() uint64 {
	if m != nil {
		return m.EndBlock
	}
	return 0
}

func (m *MsgProposeSpan) GetBorChainId() string {
	if m != nil {
		return m.BorChainId
	}
	return ""
}

func (m *MsgProposeSpan) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

// MsgProposeSpanResponse defines the Msg/MsgProposeSpan response type.
type MsgProposeSpanResponse struct {
}

func (m *MsgProposeSpanResponse) Reset()         { *m = MsgProposeSpanResponse{} }
func (m *MsgProposeSpanResponse) String() string { return proto.CompactTextString(m) }
func (*MsgProposeSpanResponse) ProtoMessage()    {}
func (*MsgProposeSpanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d99edf48c57200b, []int{1}
}
func (m *MsgProposeSpanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposeSpanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposeSpanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposeSpanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposeSpanResponse.Merge(m, src)
}
func (m *MsgProposeSpanResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposeSpanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposeSpanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposeSpanResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgProposeSpan)(nil), "heimdall.bor.v1beta1.MsgProposeSpan")
	proto.RegisterType((*MsgProposeSpanResponse)(nil), "heimdall.bor.v1beta1.MsgProposeSpanResponse")
}

func init() { proto.RegisterFile("heimdall/bor/v1beta1/msg.proto", fileDescriptor_7d99edf48c57200b) }

var fileDescriptor_7d99edf48c57200b = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0xc0, 0x37, 0xdd, 0x75, 0x6d, 0xa7, 0x52, 0x25, 0x56, 0x0d, 0x8b, 0x66, 0xea, 0xa0, 0xb4,
	0x68, 0x9b, 0x50, 0x05, 0x0f, 0x15, 0x3c, 0xac, 0x50, 0xd8, 0x43, 0xa1, 0xa4, 0x9e, 0xbc, 0x2c,
	0x93, 0xcd, 0x90, 0x0d, 0x4d, 0xe6, 0x85, 0x99, 0x51, 0xdb, 0xab, 0xe0, 0x5d, 0xf0, 0xe6, 0xd1,
	0x4f, 0xe3, 0xb1, 0xe0, 0xc5, 0xd3, 0x20, 0xbb, 0x9e, 0x72, 0xcc, 0x27, 0x90, 0x4c, 0xb2, 0xff,
	0xa0, 0x07, 0x6f, 0xf3, 0x7e, 0xef, 0xcd, 0xef, 0x31, 0xf3, 0x1e, 0x72, 0xc7, 0x2c, 0xc9, 0x22,
	0x9a, 0xa6, 0x7e, 0x08, 0xc2, 0xff, 0x78, 0x18, 0x32, 0x45, 0x0f, 0xfd, 0x4c, 0xc6, 0x5e, 0x2e,
	0x40, 0x81, 0xbd, 0x3d, 0xcb, 0x7b, 0x21, 0x08, 0xaf, 0xc9, 0xf7, 0xb6, 0x63, 0x88, 0xc1, 0x14,
	0xf8, 0xd5, 0xa9, 0xae, 0xed, 0x3d, 0x8c, 0x01, 0xe2, 0x94, 0xf9, 0x34, 0x4f, 0x7c, 0xca, 0x39,
	0x28, 0xaa, 0x12, 0xe0, 0xb2, 0xce, 0x92, 0x2f, 0x6d, 0xb4, 0x75, 0x22, 0xe3, 0x53, 0x01, 0x39,
	0x48, 0x76, 0x96, 0x53, 0x6e, 0xbf, 0x42, 0x37, 0x65, 0x4e, 0xf9, 0x30, 0x89, 0x1c, 0x6b, 0xc7,
	0xda, 0xeb, 0xf4, 0x1f, 0x15, 0x1a, 0xcf, 0x50, 0xa9, 0xf1, 0xd6, 0x25, 0xcd, 0xd2, 0x23, 0xd2,
	0x00, 0x12, 0x74, 0xab, 0xd3, 0x20, 0xb2, 0x5f, 0xa3, 0xf5, 0xbc, 0xd6, 0x08, 0x67, 0x6d, 0xc7,
	0xda, 0xdb, 0xe8, 0xe3, 0x42, 0xe3, 0x39, 0x2b, 0x35, 0xbe, 0x5d, 0xdf, 0x9c, 0x11, 0x12, 0xcc,
	0x93, 0xf6, 0x31, 0xda, 0x94, 0x8a, 0x0a, 0x35, 0x0c, 0x53, 0x18, 0x9d, 0x3b, 0x6d, 0xd3, 0xf8,
	0x69, 0xa1, 0xf1, 0x32, 0x2e, 0x35, 0xb6, 0x9b, 0xe6, 0x0b, 0x48, 0x02, 0x64, 0xa2, 0x7e, 0x15,
	0xd8, 0x6f, 0xd0, 0x06, 0xe3, 0x51, 0x63, 0xe9, 0x18, 0xcb, 0xe3, 0x42, 0xe3, 0x05, 0x2c, 0x35,
	0xbe, 0x53, 0x3b, 0xe6, 0x88, 0x04, 0xeb, 0x8c, 0x47, 0xf5, 0xfd, 0x01, 0xba, 0x15, 0x82, 0x18,
	0x8e, 0xc6, 0x34, 0x31, 0x3f, 0x70, 0xc3, 0x3c, 0x64, 0xb7, 0xd0, 0x78, 0x85, 0x97, 0x1a, 0xdf,
	0xad, 0x2d, 0xcb, 0x94, 0x04, 0x28, 0x04, 0xf1, 0xb6, 0x8a, 0x06, 0x91, 0xfd, 0x1c, 0x75, 0x24,
	0x63, 0x91, 0xd3, 0x35, 0x8a, 0x07, 0x85, 0xc6, 0x26, 0x2e, 0x35, 0xde, 0x6c, 0x1e, 0xc1, 0x58,
	0x44, 0x02, 0x03, 0x89, 0x83, 0xee, 0xaf, 0x8e, 0x21, 0x60, 0x32, 0x07, 0x2e, 0xd9, 0x8b, 0x1f,
	0x16, 0x6a, 0x9f, 0xc8, 0xd8, 0xfe, 0x6e, 0xa1, 0x7b, 0xa7, 0x20, 0xd5, 0x19, 0xe3, 0xd1, 0x52,
	0xdd, 0xbb, 0x0b, 0xfb, 0x89, 0x77, 0xdd, 0x3a, 0x78, 0xab, 0xbe, 0xde, 0xfe, 0xff, 0x54, 0xcd,
	0xba, 0x92, 0x83, 0xcf, 0xbf, 0xfe, 0x7e, 0x5b, 0xdb, 0x25, 0xc4, 0xbf, 0x76, 0x15, 0x9b, 0xb9,
	0x1d, 0x54, 0xa3, 0x3f, 0xb2, 0x9e, 0xf5, 0x8f, 0x7f, 0x4e, 0x5c, 0xeb, 0x6a, 0xe2, 0x5a, 0x7f,
	0x26, 0xae, 0xf5, 0x75, 0xea, 0xb6, 0xae, 0xa6, 0x6e, 0xeb, 0xf7, 0xd4, 0x6d, 0xbd, 0xdf, 0x8f,
	0x13, 0x35, 0xfe, 0x10, 0x7a, 0x23, 0xc8, 0xfc, 0x8c, 0xaa, 0x64, 0xc4, 0x99, 0xfa, 0x04, 0xe2,
	0x7c, 0xe1, 0xbd, 0x30, 0x66, 0x75, 0x99, 0x33, 0x19, 0x76, 0xcd, 0x56, 0xbe, 0xfc, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x20, 0xc9, 0xfd, 0xee, 0x01, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	PostSendProposeSpanTx(ctx context.Context, in *MsgProposeSpan, opts ...grpc.CallOption) (*MsgProposeSpanResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PostSendProposeSpanTx(ctx context.Context, in *MsgProposeSpan, opts ...grpc.CallOption) (*MsgProposeSpanResponse, error) {
	out := new(MsgProposeSpanResponse)
	err := c.cc.Invoke(ctx, "/heimdall.bor.v1beta1.Msg/PostSendProposeSpanTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	PostSendProposeSpanTx(context.Context, *MsgProposeSpan) (*MsgProposeSpanResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PostSendProposeSpanTx(ctx context.Context, req *MsgProposeSpan) (*MsgProposeSpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSendProposeSpanTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PostSendProposeSpanTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProposeSpan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PostSendProposeSpanTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.bor.v1beta1.Msg/PostSendProposeSpanTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PostSendProposeSpanTx(ctx, req.(*MsgProposeSpan))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall.bor.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostSendProposeSpanTx",
			Handler:    _Msg_PostSendProposeSpanTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heimdall/bor/v1beta1/msg.proto",
}

func (m *MsgProposeSpan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposeSpan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposeSpan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Seed) > 0 {
		i -= len(m.Seed)
		copy(dAtA[i:], m.Seed)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Seed)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.BorChainId) > 0 {
		i -= len(m.BorChainId)
		copy(dAtA[i:], m.BorChainId)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.BorChainId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.EndBlock != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.EndBlock))
		i--
		dAtA[i] = 0x20
	}
	if m.StartBlock != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if m.SpanId != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.SpanId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgProposeSpanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposeSpanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposeSpanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgProposeSpan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SpanId != 0 {
		n += 1 + sovMsg(uint64(m.SpanId))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.StartBlock != 0 {
		n += 1 + sovMsg(uint64(m.StartBlock))
	}
	if m.EndBlock != 0 {
		n += 1 + sovMsg(uint64(m.EndBlock))
	}
	l = len(m.BorChainId)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgProposeSpanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgProposeSpan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgProposeSpan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposeSpan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			m.SpanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpanId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			m.EndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BorChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgProposeSpanResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgProposeSpanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposeSpanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)