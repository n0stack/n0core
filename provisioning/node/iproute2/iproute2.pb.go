// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/iproute2/iproute2.proto

/*
Package iproute2 is a generated protocol buffer package.

It is generated from these files:
	proto/iproute2/iproute2.proto

It has these top-level messages:
	Tap
	ApplyTapRequest
	DeleteTapRequest
*/
package iproute2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type Tap_NetworkType int32

const (
	Tap_FLAT  Tap_NetworkType = 0
	Tap_VLAN  Tap_NetworkType = 1
	Tap_VXLAN Tap_NetworkType = 2
)

var Tap_NetworkType_name = map[int32]string{
	0: "FLAT",
	1: "VLAN",
	2: "VXLAN",
}
var Tap_NetworkType_value = map[string]int32{
	"FLAT":  0,
	"VLAN":  1,
	"VXLAN": 2,
}

func (x Tap_NetworkType) String() string {
	return proto.EnumName(Tap_NetworkType_name, int32(x))
}
func (Tap_NetworkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Tap struct {
	Name       string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	BridgeName string          `protobuf:"bytes,2,opt,name=bridge_name,json=bridgeName" json:"bridge_name,omitempty"`
	Type       Tap_NetworkType `protobuf:"varint,3,opt,name=type,enum=n0stack.n0core.iproute2.Tap_NetworkType" json:"type,omitempty"`
	// FLATの場合は未定義(0)
	NetworkId uint64 `protobuf:"varint,4,opt,name=network_id,json=networkId" json:"network_id,omitempty"`
}

func (m *Tap) Reset()                    { *m = Tap{} }
func (m *Tap) String() string            { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()               {}
func (*Tap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Tap) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tap) GetBridgeName() string {
	if m != nil {
		return m.BridgeName
	}
	return ""
}

func (m *Tap) GetType() Tap_NetworkType {
	if m != nil {
		return m.Type
	}
	return Tap_FLAT
}

func (m *Tap) GetNetworkId() uint64 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

type ApplyTapRequest struct {
	Tap *Tap `protobuf:"bytes,1,opt,name=tap" json:"tap,omitempty"`
}

func (m *ApplyTapRequest) Reset()                    { *m = ApplyTapRequest{} }
func (m *ApplyTapRequest) String() string            { return proto.CompactTextString(m) }
func (*ApplyTapRequest) ProtoMessage()               {}
func (*ApplyTapRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApplyTapRequest) GetTap() *Tap {
	if m != nil {
		return m.Tap
	}
	return nil
}

type DeleteTapRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteTapRequest) Reset()                    { *m = DeleteTapRequest{} }
func (m *DeleteTapRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTapRequest) ProtoMessage()               {}
func (*DeleteTapRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteTapRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Tap)(nil), "n0stack.n0core.iproute2.Tap")
	proto.RegisterType((*ApplyTapRequest)(nil), "n0stack.n0core.iproute2.ApplyTapRequest")
	proto.RegisterType((*DeleteTapRequest)(nil), "n0stack.n0core.iproute2.DeleteTapRequest")
	proto.RegisterEnum("n0stack.n0core.iproute2.Tap_NetworkType", Tap_NetworkType_name, Tap_NetworkType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Iproute2Service service

type Iproute2ServiceClient interface {
	ApplyTap(ctx context.Context, in *ApplyTapRequest, opts ...grpc.CallOption) (*Tap, error)
	DeleteTap(ctx context.Context, in *DeleteTapRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type iproute2ServiceClient struct {
	cc *grpc.ClientConn
}

func NewIproute2ServiceClient(cc *grpc.ClientConn) Iproute2ServiceClient {
	return &iproute2ServiceClient{cc}
}

func (c *iproute2ServiceClient) ApplyTap(ctx context.Context, in *ApplyTapRequest, opts ...grpc.CallOption) (*Tap, error) {
	out := new(Tap)
	err := grpc.Invoke(ctx, "/n0stack.n0core.iproute2.Iproute2Service/ApplyTap", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iproute2ServiceClient) DeleteTap(ctx context.Context, in *DeleteTapRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.n0core.iproute2.Iproute2Service/DeleteTap", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Iproute2Service service

type Iproute2ServiceServer interface {
	ApplyTap(context.Context, *ApplyTapRequest) (*Tap, error)
	DeleteTap(context.Context, *DeleteTapRequest) (*google_protobuf.Empty, error)
}

func RegisterIproute2ServiceServer(s *grpc.Server, srv Iproute2ServiceServer) {
	s.RegisterService(&_Iproute2Service_serviceDesc, srv)
}

func _Iproute2Service_ApplyTap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyTapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Iproute2ServiceServer).ApplyTap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.iproute2.Iproute2Service/ApplyTap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Iproute2ServiceServer).ApplyTap(ctx, req.(*ApplyTapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iproute2Service_DeleteTap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Iproute2ServiceServer).DeleteTap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.iproute2.Iproute2Service/DeleteTap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Iproute2ServiceServer).DeleteTap(ctx, req.(*DeleteTapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Iproute2Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.n0core.iproute2.Iproute2Service",
	HandlerType: (*Iproute2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyTap",
			Handler:    _Iproute2Service_ApplyTap_Handler,
		},
		{
			MethodName: "DeleteTap",
			Handler:    _Iproute2Service_DeleteTap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/iproute2/iproute2.proto",
}

func init() { proto.RegisterFile("proto/iproute2/iproute2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x6b, 0xe2, 0x40,
	0x14, 0x35, 0x9a, 0x5d, 0xcc, 0x15, 0x56, 0x99, 0x87, 0x5d, 0x71, 0x57, 0x56, 0xf2, 0xb0, 0x64,
	0xa1, 0x4c, 0x24, 0x2d, 0x7d, 0xea, 0x4b, 0x4a, 0x5b, 0x10, 0x44, 0x4a, 0x1a, 0xa4, 0xf4, 0x45,
	0xf2, 0x71, 0x9b, 0x0e, 0x6a, 0x66, 0x1a, 0x27, 0x96, 0xfc, 0xb4, 0xfe, 0x80, 0xfe, 0xaf, 0x92,
	0x89, 0x8a, 0x48, 0xf5, 0xed, 0x70, 0xee, 0x39, 0x27, 0x27, 0x73, 0x2f, 0xf4, 0x45, 0xc6, 0x25,
	0xb7, 0x99, 0xc8, 0x78, 0x2e, 0xd1, 0xd9, 0x01, 0xaa, 0x78, 0xf2, 0x2b, 0x1d, 0xae, 0x64, 0x10,
	0xcd, 0x69, 0x3a, 0x8c, 0x78, 0x86, 0x74, 0x3b, 0xee, 0xfd, 0x4e, 0x38, 0x4f, 0x16, 0x68, 0x2b,
	0x59, 0x98, 0x3f, 0xdb, 0xb8, 0x14, 0xb2, 0xa8, 0x5c, 0xe6, 0x87, 0x06, 0x0d, 0x3f, 0x10, 0x84,
	0x80, 0x9e, 0x06, 0x4b, 0xec, 0x6a, 0x03, 0xcd, 0x32, 0x3c, 0x85, 0xc9, 0x5f, 0x68, 0x85, 0x19,
	0x8b, 0x13, 0x9c, 0xa9, 0x51, 0x5d, 0x8d, 0xa0, 0xa2, 0x26, 0xa5, 0xe0, 0x0a, 0x74, 0x59, 0x08,
	0xec, 0x36, 0x06, 0x9a, 0xf5, 0xc3, 0xb1, 0xe8, 0x91, 0x06, 0xd4, 0x0f, 0x04, 0x9d, 0xa0, 0x7c,
	0xe3, 0xd9, 0xdc, 0x2f, 0x04, 0x7a, 0xca, 0x45, 0xfa, 0x00, 0x69, 0x45, 0xce, 0x58, 0xdc, 0xd5,
	0x07, 0x9a, 0xa5, 0x7b, 0xc6, 0x86, 0x19, 0xc5, 0xe6, 0x19, 0xb4, 0xf6, 0x3c, 0xa4, 0x09, 0xfa,
	0xdd, 0xd8, 0xf5, 0x3b, 0xb5, 0x12, 0x4d, 0xc7, 0xee, 0xa4, 0xa3, 0x11, 0x03, 0xbe, 0x4d, 0x1f,
	0x4b, 0x58, 0x37, 0x5d, 0x68, 0xbb, 0x42, 0x2c, 0x0a, 0x3f, 0x10, 0x1e, 0xbe, 0xe6, 0xb8, 0x92,
	0x84, 0x42, 0x43, 0x06, 0x42, 0xfd, 0x51, 0xcb, 0xf9, 0x73, 0xaa, 0x9c, 0x57, 0x0a, 0xcd, 0x7f,
	0xd0, 0xb9, 0xc1, 0x05, 0x4a, 0xdc, 0xcb, 0xf8, 0xe2, 0x59, 0x9c, 0x77, 0x0d, 0xda, 0xa3, 0x8d,
	0xfb, 0x01, 0xb3, 0x35, 0x8b, 0x90, 0xf8, 0xd0, 0xdc, 0x7e, 0x9e, 0x1c, 0x7f, 0x87, 0x83, 0x86,
	0xbd, 0x93, 0xa5, 0xcc, 0x1a, 0xb9, 0x07, 0x63, 0xd7, 0x88, 0xfc, 0x3f, 0x2a, 0x3e, 0x6c, 0xdd,
	0xfb, 0x49, 0xab, 0x95, 0xd3, 0xed, 0xca, 0xe9, 0x6d, 0xb9, 0x72, 0xb3, 0x76, 0x7d, 0xf9, 0x74,
	0x91, 0x30, 0xf9, 0x92, 0x87, 0x34, 0xe2, 0x4b, 0x7b, 0x13, 0x68, 0x57, 0x81, 0xe5, 0x81, 0xac,
	0xd9, 0x8a, 0xf1, 0x94, 0xa5, 0x89, 0x9d, 0xf2, 0x18, 0x77, 0x27, 0x16, 0x7e, 0x57, 0x49, 0xe7,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x70, 0x3e, 0x67, 0x84, 0x02, 0x00, 0x00,
}