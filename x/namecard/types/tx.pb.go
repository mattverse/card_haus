// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: namecard/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type MsgCreateNamecard struct {
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl  string `protobuf:"bytes,3,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	SelfIntro string `protobuf:"bytes,4,opt,name=selfIntro,proto3" json:"selfIntro,omitempty"`
}

func (m *MsgCreateNamecard) Reset()         { *m = MsgCreateNamecard{} }
func (m *MsgCreateNamecard) String() string { return proto.CompactTextString(m) }
func (*MsgCreateNamecard) ProtoMessage()    {}
func (*MsgCreateNamecard) Descriptor() ([]byte, []int) {
	return fileDescriptor_f393939cc93b52bd, []int{0}
}
func (m *MsgCreateNamecard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateNamecard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateNamecard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateNamecard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateNamecard.Merge(m, src)
}
func (m *MsgCreateNamecard) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateNamecard) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateNamecard.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateNamecard proto.InternalMessageInfo

func (m *MsgCreateNamecard) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgCreateNamecard) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreateNamecard) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *MsgCreateNamecard) GetSelfIntro() string {
	if m != nil {
		return m.SelfIntro
	}
	return ""
}

type MsgCreateNamecardResponse struct {
}

func (m *MsgCreateNamecardResponse) Reset()         { *m = MsgCreateNamecardResponse{} }
func (m *MsgCreateNamecardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateNamecardResponse) ProtoMessage()    {}
func (*MsgCreateNamecardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f393939cc93b52bd, []int{1}
}
func (m *MsgCreateNamecardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateNamecardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateNamecardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateNamecardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateNamecardResponse.Merge(m, src)
}
func (m *MsgCreateNamecardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateNamecardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateNamecardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateNamecardResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateNamecard)(nil), "mattverse.cardhaus.namecard.MsgCreateNamecard")
	proto.RegisterType((*MsgCreateNamecardResponse)(nil), "mattverse.cardhaus.namecard.MsgCreateNamecardResponse")
}

func init() { proto.RegisterFile("namecard/tx.proto", fileDescriptor_f393939cc93b52bd) }

var fileDescriptor_f393939cc93b52bd = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x1b, 0x5a, 0x01, 0xf5, 0x80, 0x54, 0x8b, 0x21, 0xa4, 0xc8, 0xa0, 0x4e, 0x4c, 0xb6,
	0x04, 0x12, 0x0f, 0x00, 0x0b, 0x0c, 0x65, 0xa8, 0xc4, 0xc2, 0xe6, 0x34, 0x57, 0x37, 0x52, 0x1c,
	0x47, 0x3e, 0x07, 0x05, 0x31, 0xf0, 0x0a, 0x3c, 0x16, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08, 0x8a,
	0x4b, 0x5a, 0x89, 0x4a, 0x48, 0x6c, 0xf7, 0xff, 0xf7, 0xfb, 0xbb, 0xd3, 0x99, 0x8c, 0x72, 0xa9,
	0x61, 0x2e, 0x6d, 0x22, 0x5c, 0xc5, 0x0b, 0x6b, 0x9c, 0xa1, 0x63, 0x2d, 0x9d, 0x7b, 0x06, 0x8b,
	0xc0, 0xdb, 0xc6, 0x52, 0x96, 0xc8, 0xbb, 0x54, 0x74, 0xac, 0x8c, 0x32, 0x3e, 0x27, 0xda, 0x6a,
	0xfd, 0x24, 0x3a, 0x53, 0xc6, 0xa8, 0x0c, 0x84, 0x57, 0x71, 0xb9, 0x10, 0x2e, 0xd5, 0x80, 0x4e,
	0xea, 0x62, 0x1d, 0x98, 0xbc, 0x92, 0xd1, 0x14, 0xd5, 0xad, 0x05, 0xe9, 0xe0, 0xe1, 0x87, 0x45,
	0x43, 0x72, 0x20, 0x93, 0xc4, 0x02, 0x62, 0x18, 0x9c, 0x07, 0x17, 0xc3, 0x59, 0x27, 0x29, 0x25,
	0x83, 0x76, 0x62, 0xb8, 0xe7, 0x6d, 0x5f, 0xd3, 0x88, 0x1c, 0xa6, 0x5a, 0x2a, 0x78, 0xb4, 0x59,
	0xd8, 0xf7, 0xfe, 0x46, 0xd3, 0x53, 0x32, 0x44, 0xc8, 0x16, 0xf7, 0xb9, 0xb3, 0x26, 0x1c, 0xf8,
	0xe6, 0xd6, 0x98, 0x8c, 0xc9, 0xc9, 0xce, 0xf0, 0x19, 0x60, 0x61, 0x72, 0x84, 0xcb, 0x37, 0xd2,
	0x9f, 0xa2, 0xa2, 0x15, 0x39, 0xfa, 0xb5, 0x1d, 0xe7, 0x7f, 0xdc, 0x81, 0xef, 0x00, 0xa3, 0xeb,
	0xff, 0xe5, 0xbb, 0x05, 0x6e, 0xee, 0x3e, 0x6a, 0x16, 0xac, 0x6a, 0x16, 0x7c, 0xd5, 0x2c, 0x78,
	0x6f, 0x58, 0x6f, 0xd5, 0xb0, 0xde, 0x67, 0xc3, 0x7a, 0x4f, 0x5c, 0xa5, 0x6e, 0x59, 0xc6, 0x7c,
	0x6e, 0xb4, 0xd8, 0xb0, 0x45, 0xc7, 0x16, 0x95, 0xd8, 0xfe, 0xdd, 0x4b, 0x01, 0x18, 0xef, 0xfb,
	0x5b, 0x5f, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xd8, 0x61, 0xa0, 0xd4, 0x01, 0x00, 0x00,
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
	CreateNamecard(ctx context.Context, in *MsgCreateNamecard, opts ...grpc.CallOption) (*MsgCreateNamecardResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateNamecard(ctx context.Context, in *MsgCreateNamecard, opts ...grpc.CallOption) (*MsgCreateNamecardResponse, error) {
	out := new(MsgCreateNamecardResponse)
	err := c.cc.Invoke(ctx, "/mattverse.cardhaus.namecard.Msg/CreateNamecard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateNamecard(context.Context, *MsgCreateNamecard) (*MsgCreateNamecardResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateNamecard(ctx context.Context, req *MsgCreateNamecard) (*MsgCreateNamecardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamecard not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateNamecard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateNamecard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateNamecard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mattverse.cardhaus.namecard.Msg/CreateNamecard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateNamecard(ctx, req.(*MsgCreateNamecard))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mattverse.cardhaus.namecard.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamecard",
			Handler:    _Msg_CreateNamecard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namecard/tx.proto",
}

func (m *MsgCreateNamecard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateNamecard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateNamecard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SelfIntro) > 0 {
		i -= len(m.SelfIntro)
		copy(dAtA[i:], m.SelfIntro)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SelfIntro)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ImageUrl) > 0 {
		i -= len(m.ImageUrl)
		copy(dAtA[i:], m.ImageUrl)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ImageUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateNamecardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateNamecardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateNamecardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgCreateNamecard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ImageUrl)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SelfIntro)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateNamecardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateNamecard) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateNamecard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateNamecard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUrl", wireType)
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
			m.ImageUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfIntro", wireType)
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
			m.SelfIntro = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateNamecardResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateNamecardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateNamecardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
