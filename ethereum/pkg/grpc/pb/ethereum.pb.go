// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ethereum.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateAccountRequest struct {
	Passphrase           string   `protobuf:"bytes,1,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9885406d38046041, []int{0}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

type CreateAccountReply struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountReply) Reset()         { *m = CreateAccountReply{} }
func (m *CreateAccountReply) String() string { return proto.CompactTextString(m) }
func (*CreateAccountReply) ProtoMessage()    {}
func (*CreateAccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9885406d38046041, []int{1}
}

func (m *CreateAccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountReply.Unmarshal(m, b)
}
func (m *CreateAccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountReply.Marshal(b, m, deterministic)
}
func (m *CreateAccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountReply.Merge(m, src)
}
func (m *CreateAccountReply) XXX_Size() int {
	return xxx_messageInfo_CreateAccountReply.Size(m)
}
func (m *CreateAccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountReply proto.InternalMessageInfo

func (m *CreateAccountReply) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CreateAccountReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAccountRequest)(nil), "pb.CreateAccountRequest")
	proto.RegisterType((*CreateAccountReply)(nil), "pb.CreateAccountReply")
}

func init() { proto.RegisterFile("ethereum.proto", fileDescriptor_9885406d38046041) }

var fileDescriptor_9885406d38046041 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2d, 0xc9, 0x48,
	0x2d, 0x4a, 0x2d, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32,
	0xe3, 0x12, 0x71, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x09,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe3, 0xe2, 0x2a, 0x48, 0x2c, 0x2e, 0x2e, 0xc8,
	0x28, 0x4a, 0x2c, 0x4e, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x12, 0x51, 0x72, 0xe0,
	0x12, 0x42, 0xd3, 0x57, 0x90, 0x53, 0x29, 0x24, 0xc1, 0xc5, 0x9e, 0x98, 0x92, 0x52, 0x94, 0x5a,
	0x5c, 0x0c, 0xd5, 0x02, 0xe3, 0x0a, 0x09, 0x70, 0x31, 0xa7, 0x16, 0x15, 0x49, 0x30, 0x81, 0x45,
	0x41, 0x4c, 0x23, 0x5f, 0x2e, 0x0e, 0x57, 0xa8, 0x7b, 0x84, 0x1c, 0xb9, 0x78, 0x51, 0x4c, 0x13,
	0x92, 0xd0, 0x2b, 0x48, 0xd2, 0xc3, 0xe6, 0x30, 0x29, 0x31, 0x2c, 0x32, 0x05, 0x39, 0x95, 0x49,
	0x6c, 0x60, 0x3f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0x86, 0x23, 0xbf, 0xe5, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EthereumClient is the client API for Ethereum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EthereumClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error)
}

type ethereumClient struct {
	cc *grpc.ClientConn
}

func NewEthereumClient(cc *grpc.ClientConn) EthereumClient {
	return &ethereumClient{cc}
}

func (c *ethereumClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error) {
	out := new(CreateAccountReply)
	err := c.cc.Invoke(ctx, "/pb.Ethereum/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EthereumServer is the server API for Ethereum service.
type EthereumServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error)
}

func RegisterEthereumServer(s *grpc.Server, srv EthereumServer) {
	s.RegisterService(&_Ethereum_serviceDesc, srv)
}

func _Ethereum_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Ethereum/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ethereum_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Ethereum",
	HandlerType: (*EthereumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Ethereum_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethereum.proto",
}
