// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

/*
Package serbench is a generated protocol buffer package.

It is generated from these files:
	proto.proto

It has these top-level messages:
	UserProto
*/
package serbench

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

type UserProto struct {
	Id   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Item []string `protobuf:"bytes,3,rep,name=item" json:"item,omitempty"`
}

func (m *UserProto) Reset()                    { *m = UserProto{} }
func (m *UserProto) String() string            { return proto.CompactTextString(m) }
func (*UserProto) ProtoMessage()               {}
func (*UserProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserProto) GetItem() []string {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*UserProto)(nil), "serbench.UserProto")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0x1c, 0xc5, 0xa9, 0x45, 0x49, 0xa9, 0x79, 0xc9, 0x19, 0x4a, 0xce,
	0x5c, 0x9c, 0xa1, 0xc5, 0xa9, 0x45, 0x01, 0x60, 0x61, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x1b, 0x24, 0x96, 0x59, 0x92, 0x9a, 0x2b, 0xc1,
	0xac, 0xc0, 0x0c, 0x12, 0x03, 0xb1, 0x93, 0xd8, 0xc0, 0xa6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x86, 0xa2, 0xcd, 0x72, 0x64, 0x00, 0x00, 0x00,
}
