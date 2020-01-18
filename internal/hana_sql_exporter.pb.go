// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hana_sql_exporter.proto

package internal

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Secret struct {
	Name                 map[string][]byte `protobuf:"bytes,1,rep,name=Name,proto3" json:"Name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_7faba8c3cc2ecea9, []int{0}
}

func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetName() map[string][]byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*Secret)(nil), "internal.Secret")
	proto.RegisterMapType((map[string][]byte)(nil), "internal.Secret.NameEntry")
}

func init() { proto.RegisterFile("hana_sql_exporter.proto", fileDescriptor_7faba8c3cc2ecea9) }

var fileDescriptor_7faba8c3cc2ecea9 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x48, 0xcc, 0x4b,
	0x8c, 0x2f, 0x2e, 0xcc, 0x89, 0x4f, 0xad, 0x28, 0xc8, 0x2f, 0x2a, 0x49, 0x2d, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x51, 0x2a, 0xe4,
	0x62, 0x0b, 0x4e, 0x4d, 0x2e, 0x4a, 0x2d, 0x11, 0xd2, 0xe3, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd2, 0x83, 0x29, 0xd1, 0x83, 0xc8, 0xeb, 0x81, 0x24,
	0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xc0, 0xea, 0xa4, 0xcc, 0xb9, 0x38, 0xe1, 0x42, 0x42, 0x02,
	0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x08,
	0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x84, 0x63,
	0xc5, 0x64, 0xc1, 0x98, 0xc4, 0x06, 0x76, 0x83, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x45,
	0x58, 0xad, 0x9e, 0x00, 0x00, 0x00,
}
