// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deployment/v0/flavor.proto

package pdeployment // import "github.com/n0stack/n0stack/n0proto.go/deployment/v0"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import v0 "github.com/n0stack/n0stack/n0proto.go/provisioning/v0"

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

type Flavor struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	LimitCpuMilliCore    uint32            `protobuf:"varint,10,opt,name=limit_cpu_milli_core,json=limitCpuMilliCore" json:"limit_cpu_milli_core,omitempty"`
	LimitMemoryBytes     uint64            `protobuf:"varint,11,opt,name=limit_memory_bytes,json=limitMemoryBytes" json:"limit_memory_bytes,omitempty"`
	LimitStorageBytes    uint64            `protobuf:"varint,12,opt,name=limit_storage_bytes,json=limitStorageBytes" json:"limit_storage_bytes,omitempty"`
	NetworkName          string            `protobuf:"bytes,13,opt,name=network_name,json=networkName" json:"network_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Flavor) Reset()         { *m = Flavor{} }
func (m *Flavor) String() string { return proto.CompactTextString(m) }
func (*Flavor) ProtoMessage()    {}
func (*Flavor) Descriptor() ([]byte, []int) {
	return fileDescriptor_flavor_a7b5f18efbb567f0, []int{0}
}
func (m *Flavor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flavor.Unmarshal(m, b)
}
func (m *Flavor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flavor.Marshal(b, m, deterministic)
}
func (dst *Flavor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flavor.Merge(dst, src)
}
func (m *Flavor) XXX_Size() int {
	return xxx_messageInfo_Flavor.Size(m)
}
func (m *Flavor) XXX_DiscardUnknown() {
	xxx_messageInfo_Flavor.DiscardUnknown(m)
}

var xxx_messageInfo_Flavor proto.InternalMessageInfo

func (m *Flavor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Flavor) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Flavor) GetLimitCpuMilliCore() uint32 {
	if m != nil {
		return m.LimitCpuMilliCore
	}
	return 0
}

func (m *Flavor) GetLimitMemoryBytes() uint64 {
	if m != nil {
		return m.LimitMemoryBytes
	}
	return 0
}

func (m *Flavor) GetLimitStorageBytes() uint64 {
	if m != nil {
		return m.LimitStorageBytes
	}
	return 0
}

func (m *Flavor) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

type ListFlavorsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFlavorsRequest) Reset()         { *m = ListFlavorsRequest{} }
func (m *ListFlavorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListFlavorsRequest) ProtoMessage()    {}
func (*ListFlavorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flavor_a7b5f18efbb567f0, []int{1}
}
func (m *ListFlavorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFlavorsRequest.Unmarshal(m, b)
}
func (m *ListFlavorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFlavorsRequest.Marshal(b, m, deterministic)
}
func (dst *ListFlavorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFlavorsRequest.Merge(dst, src)
}
func (m *ListFlavorsRequest) XXX_Size() int {
	return xxx_messageInfo_ListFlavorsRequest.Size(m)
}
func (m *ListFlavorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFlavorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFlavorsRequest proto.InternalMessageInfo

type ListFlavorsResponse struct {
	Flavors              []*Flavor `protobuf:"bytes,1,rep,name=flavors" json:"flavors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListFlavorsResponse) Reset()         { *m = ListFlavorsResponse{} }
func (m *ListFlavorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListFlavorsResponse) ProtoMessage()    {}
func (*ListFlavorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_flavor_a7b5f18efbb567f0, []int{2}
}
func (m *ListFlavorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFlavorsResponse.Unmarshal(m, b)
}
func (m *ListFlavorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFlavorsResponse.Marshal(b, m, deterministic)
}
func (dst *ListFlavorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFlavorsResponse.Merge(dst, src)
}
func (m *ListFlavorsResponse) XXX_Size() int {
	return xxx_messageInfo_ListFlavorsResponse.Size(m)
}
func (m *ListFlavorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFlavorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFlavorsResponse proto.InternalMessageInfo

func (m *ListFlavorsResponse) GetFlavors() []*Flavor {
	if m != nil {
		return m.Flavors
	}
	return nil
}

type GetFlavorRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlavorRequest) Reset()         { *m = GetFlavorRequest{} }
func (m *GetFlavorRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlavorRequest) ProtoMessage()    {}
func (*GetFlavorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flavor_a7b5f18efbb567f0, []int{3}
}
func (m *GetFlavorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlavorRequest.Unmarshal(m, b)
}
func (m *GetFlavorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlavorRequest.Marshal(b, m, deterministic)
}
func (dst *GetFlavorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlavorRequest.Merge(dst, src)
}
func (m *GetFlavorRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlavorRequest.Size(m)
}
func (m *GetFlavorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlavorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlavorRequest proto.InternalMessageInfo

func (m *GetFlavorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyFlavorRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	LimitCpuMilliCore    uint32            `protobuf:"varint,10,opt,name=limit_cpu_milli_core,json=limitCpuMilliCore" json:"limit_cpu_milli_core,omitempty"`
	LimitMemoryBytes     uint64            `protobuf:"varint,11,opt,name=limit_memory_bytes,json=limitMemoryBytes" json:"limit_memory_bytes,omitempty"`
	LimitStorageBytes    uint64            `protobuf:"varint,12,opt,name=limit_storage_bytes,json=limitStorageBytes" json:"limit_storage_bytes,omitempty"`
	NetworkName          string            `protobuf:"bytes,13,opt,name=network_name,json=networkName" json:"network_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApplyFlavorRequest) Reset()         { *m = ApplyFlavorRequest{} }
func (m *ApplyFlavorRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyFlavorRequest) ProtoMessage()    {}
func (*ApplyFlavorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flavor_a7b5f18efbb567f0, []int{4}
}
func (m *ApplyFlavorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyFlavorRequest.Unmarshal(m, b)
}
func (m *ApplyFlavorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyFlavorRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyFlavorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyFlavorRequest.Merge(dst, src)
}
func (m *ApplyFlavorRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyFlavorRequest.Size(m)
}
func (m *ApplyFlavorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyFlavorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyFlavorRequest proto.InternalMessageInfo

func (m *ApplyFlavorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplyFlavorRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ApplyFlavorRequest) GetLimitCpuMilliCore() uint32 {
	if m != nil {
		return m.LimitCpuMilliCore
	}
	return 0
}

func (m *ApplyFlavorRequest) GetLimitMemoryBytes() uint64 {
	if m != nil {
		return m.LimitMemoryBytes
	}
	return 0
}

func (m *ApplyFlavorRequest) GetLimitStorageBytes() uint64 {
	if m != nil {
		return m.LimitStorageBytes
	}
	return 0
}

func (m *ApplyFlavorRequest) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

type DeleteFlavorRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFlavorRequest) Reset()         { *m = DeleteFlavorRequest{} }
func (m *DeleteFlavorRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFlavorRequest) ProtoMessage()    {}
func (*DeleteFlavorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flavor_a7b5f18efbb567f0, []int{5}
}
func (m *DeleteFlavorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFlavorRequest.Unmarshal(m, b)
}
func (m *DeleteFlavorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFlavorRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteFlavorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFlavorRequest.Merge(dst, src)
}
func (m *DeleteFlavorRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFlavorRequest.Size(m)
}
func (m *DeleteFlavorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFlavorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFlavorRequest proto.InternalMessageInfo

func (m *DeleteFlavorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GenerateVirtualMachineRequest struct {
	FlavorName           string            `protobuf:"bytes,1,opt,name=flavor_name,json=flavorName" json:"flavor_name,omitempty"`
	VirtualMachineName   string            `protobuf:"bytes,2,opt,name=virtual_machine_name,json=virtualMachineName" json:"virtual_machine_name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	RequestCpuMilliCore  uint32            `protobuf:"varint,4,opt,name=request_cpu_milli_core,json=requestCpuMilliCore" json:"request_cpu_milli_core,omitempty"`
	RequestMemoryBytes   uint64            `protobuf:"varint,5,opt,name=request_memory_bytes,json=requestMemoryBytes" json:"request_memory_bytes,omitempty"`
	RequestStorageBytes  uint64            `protobuf:"varint,6,opt,name=request_storage_bytes,json=requestStorageBytes" json:"request_storage_bytes,omitempty"`
	ImageName            string            `protobuf:"bytes,7,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	ImageTag             string            `protobuf:"bytes,8,opt,name=image_tag,json=imageTag" json:"image_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GenerateVirtualMachineRequest) Reset()         { *m = GenerateVirtualMachineRequest{} }
func (m *GenerateVirtualMachineRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateVirtualMachineRequest) ProtoMessage()    {}
func (*GenerateVirtualMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flavor_a7b5f18efbb567f0, []int{6}
}
func (m *GenerateVirtualMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateVirtualMachineRequest.Unmarshal(m, b)
}
func (m *GenerateVirtualMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateVirtualMachineRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateVirtualMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateVirtualMachineRequest.Merge(dst, src)
}
func (m *GenerateVirtualMachineRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateVirtualMachineRequest.Size(m)
}
func (m *GenerateVirtualMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateVirtualMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateVirtualMachineRequest proto.InternalMessageInfo

func (m *GenerateVirtualMachineRequest) GetFlavorName() string {
	if m != nil {
		return m.FlavorName
	}
	return ""
}

func (m *GenerateVirtualMachineRequest) GetVirtualMachineName() string {
	if m != nil {
		return m.VirtualMachineName
	}
	return ""
}

func (m *GenerateVirtualMachineRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *GenerateVirtualMachineRequest) GetRequestCpuMilliCore() uint32 {
	if m != nil {
		return m.RequestCpuMilliCore
	}
	return 0
}

func (m *GenerateVirtualMachineRequest) GetRequestMemoryBytes() uint64 {
	if m != nil {
		return m.RequestMemoryBytes
	}
	return 0
}

func (m *GenerateVirtualMachineRequest) GetRequestStorageBytes() uint64 {
	if m != nil {
		return m.RequestStorageBytes
	}
	return 0
}

func (m *GenerateVirtualMachineRequest) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *GenerateVirtualMachineRequest) GetImageTag() string {
	if m != nil {
		return m.ImageTag
	}
	return ""
}

func init() {
	proto.RegisterType((*Flavor)(nil), "n0stack.deployment.Flavor")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.deployment.Flavor.AnnotationsEntry")
	proto.RegisterType((*ListFlavorsRequest)(nil), "n0stack.deployment.ListFlavorsRequest")
	proto.RegisterType((*ListFlavorsResponse)(nil), "n0stack.deployment.ListFlavorsResponse")
	proto.RegisterType((*GetFlavorRequest)(nil), "n0stack.deployment.GetFlavorRequest")
	proto.RegisterType((*ApplyFlavorRequest)(nil), "n0stack.deployment.ApplyFlavorRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.deployment.ApplyFlavorRequest.AnnotationsEntry")
	proto.RegisterType((*DeleteFlavorRequest)(nil), "n0stack.deployment.DeleteFlavorRequest")
	proto.RegisterType((*GenerateVirtualMachineRequest)(nil), "n0stack.deployment.GenerateVirtualMachineRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.deployment.GenerateVirtualMachineRequest.AnnotationsEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FlavorService service

type FlavorServiceClient interface {
	ListFlavors(ctx context.Context, in *ListFlavorsRequest, opts ...grpc.CallOption) (*ListFlavorsResponse, error)
	GetFlavor(ctx context.Context, in *GetFlavorRequest, opts ...grpc.CallOption) (*Flavor, error)
	ApplyFlavor(ctx context.Context, in *ApplyFlavorRequest, opts ...grpc.CallOption) (*Flavor, error)
	DeleteFlavor(ctx context.Context, in *DeleteFlavorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GenerateVirtualMachine(ctx context.Context, in *GenerateVirtualMachineRequest, opts ...grpc.CallOption) (*v0.VirtualMachine, error)
}

type flavorServiceClient struct {
	cc *grpc.ClientConn
}

func NewFlavorServiceClient(cc *grpc.ClientConn) FlavorServiceClient {
	return &flavorServiceClient{cc}
}

func (c *flavorServiceClient) ListFlavors(ctx context.Context, in *ListFlavorsRequest, opts ...grpc.CallOption) (*ListFlavorsResponse, error) {
	out := new(ListFlavorsResponse)
	err := grpc.Invoke(ctx, "/n0stack.deployment.FlavorService/ListFlavors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flavorServiceClient) GetFlavor(ctx context.Context, in *GetFlavorRequest, opts ...grpc.CallOption) (*Flavor, error) {
	out := new(Flavor)
	err := grpc.Invoke(ctx, "/n0stack.deployment.FlavorService/GetFlavor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flavorServiceClient) ApplyFlavor(ctx context.Context, in *ApplyFlavorRequest, opts ...grpc.CallOption) (*Flavor, error) {
	out := new(Flavor)
	err := grpc.Invoke(ctx, "/n0stack.deployment.FlavorService/ApplyFlavor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flavorServiceClient) DeleteFlavor(ctx context.Context, in *DeleteFlavorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.deployment.FlavorService/DeleteFlavor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flavorServiceClient) GenerateVirtualMachine(ctx context.Context, in *GenerateVirtualMachineRequest, opts ...grpc.CallOption) (*v0.VirtualMachine, error) {
	out := new(v0.VirtualMachine)
	err := grpc.Invoke(ctx, "/n0stack.deployment.FlavorService/GenerateVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FlavorService service

type FlavorServiceServer interface {
	ListFlavors(context.Context, *ListFlavorsRequest) (*ListFlavorsResponse, error)
	GetFlavor(context.Context, *GetFlavorRequest) (*Flavor, error)
	ApplyFlavor(context.Context, *ApplyFlavorRequest) (*Flavor, error)
	DeleteFlavor(context.Context, *DeleteFlavorRequest) (*empty.Empty, error)
	GenerateVirtualMachine(context.Context, *GenerateVirtualMachineRequest) (*v0.VirtualMachine, error)
}

func RegisterFlavorServiceServer(s *grpc.Server, srv FlavorServiceServer) {
	s.RegisterService(&_FlavorService_serviceDesc, srv)
}

func _FlavorService_ListFlavors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFlavorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlavorServiceServer).ListFlavors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.FlavorService/ListFlavors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlavorServiceServer).ListFlavors(ctx, req.(*ListFlavorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlavorService_GetFlavor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlavorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlavorServiceServer).GetFlavor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.FlavorService/GetFlavor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlavorServiceServer).GetFlavor(ctx, req.(*GetFlavorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlavorService_ApplyFlavor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyFlavorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlavorServiceServer).ApplyFlavor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.FlavorService/ApplyFlavor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlavorServiceServer).ApplyFlavor(ctx, req.(*ApplyFlavorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlavorService_DeleteFlavor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFlavorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlavorServiceServer).DeleteFlavor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.FlavorService/DeleteFlavor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlavorServiceServer).DeleteFlavor(ctx, req.(*DeleteFlavorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlavorService_GenerateVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlavorServiceServer).GenerateVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.FlavorService/GenerateVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlavorServiceServer).GenerateVirtualMachine(ctx, req.(*GenerateVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlavorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.deployment.FlavorService",
	HandlerType: (*FlavorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFlavors",
			Handler:    _FlavorService_ListFlavors_Handler,
		},
		{
			MethodName: "GetFlavor",
			Handler:    _FlavorService_GetFlavor_Handler,
		},
		{
			MethodName: "ApplyFlavor",
			Handler:    _FlavorService_ApplyFlavor_Handler,
		},
		{
			MethodName: "DeleteFlavor",
			Handler:    _FlavorService_DeleteFlavor_Handler,
		},
		{
			MethodName: "GenerateVirtualMachine",
			Handler:    _FlavorService_GenerateVirtualMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deployment/v0/flavor.proto",
}

func init() { proto.RegisterFile("deployment/v0/flavor.proto", fileDescriptor_flavor_a7b5f18efbb567f0) }

var fileDescriptor_flavor_a7b5f18efbb567f0 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0x47, 0x3f, 0xc7, 0xad, 0x14, 0x36, 0xa1, 0xb2, 0x5c, 0x55, 0x04, 0x0b, 0x4a, 0x10,
	0xc8, 0x0e, 0x2d, 0x12, 0x08, 0x24, 0x50, 0x5b, 0x4a, 0x0f, 0x10, 0x10, 0x29, 0x42, 0x82, 0x8b,
	0x71, 0xd2, 0xa9, 0xbb, 0xaa, 0xed, 0x35, 0xf6, 0xda, 0xc8, 0x3f, 0x16, 0xf1, 0x13, 0x38, 0x71,
	0x47, 0xde, 0x75, 0x88, 0x9d, 0x98, 0x46, 0x88, 0x2b, 0x27, 0x7b, 0x67, 0xde, 0x9b, 0x9d, 0x9d,
	0x37, 0x3b, 0x0b, 0xda, 0x19, 0x06, 0x2e, 0x4b, 0x3d, 0xf4, 0xb9, 0x99, 0xf4, 0xcd, 0x73, 0xd7,
	0x4e, 0x58, 0x68, 0x04, 0x21, 0xe3, 0x8c, 0x10, 0xbf, 0x1f, 0x71, 0x7b, 0x7c, 0x69, 0x4c, 0x31,
	0xda, 0xb6, 0xc3, 0x98, 0xe3, 0xa2, 0x29, 0x10, 0xa3, 0xf8, 0xdc, 0x44, 0x2f, 0xe0, 0xa9, 0x24,
	0x68, 0xb7, 0x83, 0x90, 0x25, 0x34, 0xa2, 0xcc, 0xa7, 0xbe, 0x93, 0x85, 0x4b, 0x68, 0xc8, 0x63,
	0xdb, 0xb5, 0x3c, 0x7b, 0x7c, 0x41, 0x7d, 0x94, 0x30, 0xfd, 0x5b, 0x03, 0x56, 0x5e, 0x8a, 0x8d,
	0x08, 0x81, 0x25, 0xdf, 0xf6, 0x50, 0xad, 0x77, 0xeb, 0xbd, 0xf5, 0xa1, 0xf8, 0x27, 0x03, 0x50,
	0x6c, 0xdf, 0x67, 0xdc, 0xe6, 0x94, 0xf9, 0x91, 0xda, 0xec, 0x36, 0x7b, 0xca, 0xde, 0x3d, 0x63,
	0x3e, 0x19, 0x43, 0x06, 0x31, 0x0e, 0xa6, 0xe8, 0x63, 0x9f, 0x87, 0xe9, 0xb0, 0xc8, 0x27, 0x26,
	0x74, 0x5c, 0xea, 0x51, 0x6e, 0x8d, 0x83, 0xd8, 0xf2, 0xa8, 0xeb, 0x52, 0x6b, 0xcc, 0x42, 0x54,
	0xa1, 0x5b, 0xef, 0x6d, 0x0e, 0xaf, 0x09, 0xdf, 0x51, 0x10, 0x0f, 0x32, 0xcf, 0x11, 0x0b, 0x91,
	0xdc, 0x07, 0x22, 0x09, 0x1e, 0x7a, 0x2c, 0x4c, 0xad, 0x51, 0xca, 0x31, 0x52, 0x95, 0x6e, 0xbd,
	0xb7, 0x34, 0x6c, 0x09, 0xcf, 0x40, 0x38, 0x0e, 0x33, 0x3b, 0x31, 0xa0, 0x2d, 0xd1, 0x11, 0x67,
	0xa1, 0xed, 0x60, 0x0e, 0xdf, 0x10, 0x70, 0x19, 0xfd, 0x54, 0x7a, 0x24, 0xfe, 0x26, 0x6c, 0xf8,
	0xc8, 0xbf, 0xb2, 0xf0, 0xd2, 0x12, 0x27, 0xdf, 0x14, 0x27, 0x57, 0x72, 0xdb, 0x1b, 0xdb, 0x43,
	0xed, 0x19, 0xb4, 0x66, 0x8f, 0x44, 0x5a, 0xd0, 0xbc, 0xc4, 0x34, 0xaf, 0x53, 0xf6, 0x4b, 0x3a,
	0xb0, 0x9c, 0xd8, 0x6e, 0x8c, 0x6a, 0x43, 0xd8, 0xe4, 0xe2, 0x49, 0xe3, 0x71, 0x5d, 0xef, 0x00,
	0x79, 0x4d, 0x23, 0x2e, 0xab, 0x13, 0x0d, 0xf1, 0x4b, 0x8c, 0x11, 0xd7, 0x5f, 0x41, 0xbb, 0x64,
	0x8d, 0x02, 0xe6, 0x47, 0x48, 0x1e, 0xc2, 0xaa, 0x14, 0x3d, 0x52, 0xeb, 0xa2, 0xd2, 0xda, 0x9f,
	0x2b, 0x3d, 0x9c, 0x40, 0xf5, 0x5d, 0x68, 0x9d, 0x60, 0x1e, 0x2b, 0xdf, 0xa0, 0x4a, 0x4b, 0xfd,
	0x67, 0x03, 0xc8, 0x41, 0x10, 0xb8, 0xe9, 0x42, 0x28, 0xf9, 0x58, 0x25, 0xfb, 0xa3, 0xaa, 0x64,
	0xe6, 0x03, 0xfe, 0x6f, 0x81, 0x99, 0x16, 0xb8, 0x0b, 0xed, 0x17, 0xe8, 0x22, 0xc7, 0xc5, 0x12,
	0xfd, 0x68, 0xc2, 0xce, 0x09, 0xfa, 0x18, 0xda, 0x1c, 0x3f, 0xc8, 0xfb, 0x3a, 0x90, 0xd7, 0x75,
	0xc2, 0xba, 0x01, 0x8a, 0xd4, 0xdd, 0x2a, 0x90, 0x41, 0x9a, 0xb2, 0x6c, 0x49, 0x1f, 0x3a, 0x33,
	0x37, 0x5d, 0x22, 0x65, 0x5a, 0x24, 0x29, 0x45, 0x15, 0x8c, 0xb3, 0x2a, 0xb1, 0x0f, 0xab, 0xc4,
	0xbe, 0x32, 0xb5, 0x05, 0xba, 0xef, 0xc3, 0x56, 0x28, 0x81, 0xb3, 0xca, 0x2f, 0x09, 0xe5, 0xdb,
	0xb9, 0xb7, 0xa4, 0x7d, 0x1f, 0x3a, 0x13, 0x52, 0x49, 0xfd, 0x65, 0x21, 0x27, 0xc9, 0x7d, 0x45,
	0xfd, 0xf7, 0xe0, 0xfa, 0x84, 0x51, 0xee, 0x80, 0x15, 0x41, 0x99, 0xec, 0x52, 0xea, 0x81, 0x1d,
	0x00, 0xea, 0x65, 0x48, 0x51, 0xa8, 0x55, 0x51, 0xa8, 0x75, 0x61, 0x11, 0xf5, 0xd9, 0x06, 0xb9,
	0xb0, 0xb8, 0xed, 0xa8, 0x6b, 0xc2, 0xbb, 0x26, 0x0c, 0xef, 0x6d, 0xe7, 0x5f, 0x9b, 0x63, 0xef,
	0x7b, 0x13, 0x36, 0x65, 0x5f, 0x9c, 0x62, 0x98, 0xd0, 0x31, 0x92, 0xcf, 0xa0, 0x14, 0x66, 0x03,
	0xd9, 0xad, 0x12, 0x62, 0x7e, 0xa4, 0x68, 0x77, 0x16, 0xe2, 0xe4, 0x90, 0xd1, 0x6b, 0xe4, 0x2d,
	0xac, 0xff, 0x1e, 0x18, 0xe4, 0x56, 0xb5, 0xd0, 0xe5, 0x79, 0xa2, 0x5d, 0x31, 0x88, 0xf4, 0x1a,
	0x39, 0x05, 0xa5, 0x30, 0x07, 0xaa, 0x53, 0x9e, 0x1f, 0x14, 0x0b, 0x82, 0xbe, 0x83, 0x8d, 0xe2,
	0xb5, 0x21, 0x95, 0x07, 0xac, 0xb8, 0x58, 0xda, 0x96, 0x21, 0xdf, 0x45, 0x63, 0xf2, 0x2e, 0x1a,
	0xc7, 0xd9, 0xbb, 0xa8, 0xd7, 0x48, 0x04, 0x5b, 0xd5, 0x2d, 0x4c, 0x1e, 0xfc, 0x75, 0xbb, 0x6b,
	0xd3, 0xc2, 0x15, 0x5f, 0x5a, 0xa3, 0x0c, 0xd6, 0x6b, 0x87, 0x07, 0x9f, 0x9e, 0x3b, 0x94, 0x5f,
	0xc4, 0x23, 0x63, 0xcc, 0x3c, 0x33, 0xe7, 0x14, 0xbe, 0x22, 0x4b, 0xc3, 0x61, 0x66, 0xe9, 0xf5,
	0x7f, 0x1a, 0x4c, 0x97, 0xa3, 0x15, 0x81, 0xd9, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x38, 0x6e,
	0x1e, 0xe3, 0x21, 0x08, 0x00, 0x00,
}
