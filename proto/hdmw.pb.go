// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hdmw.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Wallet struct {
	Mnemonic             string   `protobuf:"bytes,1,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	Seed                 string   `protobuf:"bytes,2,opt,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Wallet) Reset()         { *m = Wallet{} }
func (m *Wallet) String() string { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()    {}
func (*Wallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_hdmw_83c3afef72312a65, []int{0}
}
func (m *Wallet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wallet.Unmarshal(m, b)
}
func (m *Wallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wallet.Marshal(b, m, deterministic)
}
func (dst *Wallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wallet.Merge(dst, src)
}
func (m *Wallet) XXX_Size() int {
	return xxx_messageInfo_Wallet.Size(m)
}
func (m *Wallet) XXX_DiscardUnknown() {
	xxx_messageInfo_Wallet.DiscardUnknown(m)
}

var xxx_messageInfo_Wallet proto.InternalMessageInfo

func (m *Wallet) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

func (m *Wallet) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

type Entropy struct {
	Entropy              string   `protobuf:"bytes,1,opt,name=entropy,proto3" json:"entropy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entropy) Reset()         { *m = Entropy{} }
func (m *Entropy) String() string { return proto.CompactTextString(m) }
func (*Entropy) ProtoMessage()    {}
func (*Entropy) Descriptor() ([]byte, []int) {
	return fileDescriptor_hdmw_83c3afef72312a65, []int{1}
}
func (m *Entropy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entropy.Unmarshal(m, b)
}
func (m *Entropy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entropy.Marshal(b, m, deterministic)
}
func (dst *Entropy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entropy.Merge(dst, src)
}
func (m *Entropy) XXX_Size() int {
	return xxx_messageInfo_Entropy.Size(m)
}
func (m *Entropy) XXX_DiscardUnknown() {
	xxx_messageInfo_Entropy.DiscardUnknown(m)
}

var xxx_messageInfo_Entropy proto.InternalMessageInfo

func (m *Entropy) GetEntropy() string {
	if m != nil {
		return m.Entropy
	}
	return ""
}

func init() {
	proto.RegisterType((*Wallet)(nil), "proto.Wallet")
	proto.RegisterType((*Entropy)(nil), "proto.Entropy")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Bip44WalletClient is the client API for Bip44Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Bip44WalletClient interface {
	InitializeWallet(ctx context.Context, in *Entropy, opts ...grpc.CallOption) (*Wallet, error)
}

type bip44WalletClient struct {
	cc *grpc.ClientConn
}

func NewBip44WalletClient(cc *grpc.ClientConn) Bip44WalletClient {
	return &bip44WalletClient{cc}
}

func (c *bip44WalletClient) InitializeWallet(ctx context.Context, in *Entropy, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/proto.Bip44Wallet/InitializeWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Bip44WalletServer is the server API for Bip44Wallet service.
type Bip44WalletServer interface {
	InitializeWallet(context.Context, *Entropy) (*Wallet, error)
}

func RegisterBip44WalletServer(s *grpc.Server, srv Bip44WalletServer) {
	s.RegisterService(&_Bip44Wallet_serviceDesc, srv)
}

func _Bip44Wallet_InitializeWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entropy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Bip44WalletServer).InitializeWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bip44Wallet/InitializeWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Bip44WalletServer).InitializeWallet(ctx, req.(*Entropy))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bip44Wallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Bip44Wallet",
	HandlerType: (*Bip44WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeWallet",
			Handler:    _Bip44Wallet_InitializeWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hdmw.proto",
}

func init() { proto.RegisterFile("proto/hdmw.proto", fileDescriptor_hdmw_83c3afef72312a65) }

var fileDescriptor_hdmw_83c3afef72312a65 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xc9, 0x2d, 0xd7, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x05, 0x17,
	0x5b, 0x78, 0x62, 0x4e, 0x4e, 0x6a, 0x89, 0x90, 0x14, 0x17, 0x47, 0x6e, 0x5e, 0x6a, 0x6e, 0x7e,
	0x5e, 0x66, 0xb2, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f, 0x24, 0xc4, 0xc5, 0x52,
	0x9c, 0x9a, 0x9a, 0x22, 0xc1, 0x04, 0x16, 0x07, 0xb3, 0x95, 0x94, 0xb9, 0xd8, 0x5d, 0xf3, 0x4a,
	0x8a, 0xf2, 0x0b, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x53, 0x21, 0x4c, 0xa8, 0x4e, 0x18, 0xd7, 0xc8,
	0x89, 0x8b, 0xdb, 0x29, 0xb3, 0xc0, 0xc4, 0x04, 0x6a, 0x87, 0x31, 0x97, 0x80, 0x67, 0x5e, 0x66,
	0x49, 0x66, 0x62, 0x4e, 0x66, 0x55, 0x2a, 0x54, 0x8c, 0x0f, 0xe2, 0x20, 0x3d, 0xa8, 0x61, 0x52,
	0xbc, 0x50, 0x3e, 0x44, 0x5a, 0x89, 0x21, 0x89, 0x0d, 0xcc, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0xa3, 0x69, 0x26, 0xc4, 0x00, 0x00, 0x00,
}