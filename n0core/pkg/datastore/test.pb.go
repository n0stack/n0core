// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/datastore/test.proto

package datastore // import "github.com/n0stack/n0stack/n0core/pkg/datastore"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Test struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_770abde3a4767a4a, []int{0}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (dst *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(dst, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "n0stack.internal.n0core.datastore.Test")
}

func init() { proto.RegisterFile("pkg/datastore/test.proto", fileDescriptor_test_770abde3a4767a4a) }

var fileDescriptor_test_770abde3a4767a4a = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xc8, 0x4e, 0xd7,
	0x4f, 0x49, 0x2c, 0x49, 0x2c, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xcc, 0x33, 0x28, 0x2e, 0x49, 0x4c, 0xce, 0xd6, 0xcb, 0xcc,
	0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcb, 0x33, 0x48, 0xce, 0x2f, 0x4a, 0xd5, 0x83, 0xab,
	0x56, 0x92, 0xe2, 0x62, 0x09, 0x49, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0xac, 0xa3, 0x2c, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xa1, 0x66, 0x21, 0xd1, 0x20, 0xa3, 0xf4, 0x51,
	0x2c, 0xb7, 0x86, 0xb3, 0x92, 0xd8, 0xc0, 0x4e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3a,
	0x5c, 0xf6, 0xf1, 0x9e, 0x00, 0x00, 0x00,
}
