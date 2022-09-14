// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Proto/message.proto

package proto

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

type Message struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_9442aed4be7f2922, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "grpc.Message")
}

func init() { proto.RegisterFile("Proto/message.proto", fileDescriptor_9442aed4be7f2922) }

var fileDescriptor_9442aed4be7f2922 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x00, 0xf1, 0x84, 0x58, 0xd2,
	0x8b, 0x0a, 0x92, 0x95, 0x64, 0xb9, 0xd8, 0x7d, 0x21, 0xc2, 0x42, 0x42, 0x5c, 0x2c, 0x49, 0xf9,
	0x29, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x51, 0x36, 0x17, 0x1f, 0x54,
	0x3a, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x9b, 0x8b, 0xdb, 0x33, 0xaf, 0x24, 0xb5,
	0x28, 0x39, 0x31, 0x37, 0x29, 0x33, 0x5f, 0x88, 0x57, 0x0f, 0x64, 0x8c, 0x1e, 0x54, 0x91, 0x14,
	0x2a, 0x57, 0x48, 0x87, 0x8b, 0x07, 0x49, 0xb1, 0x11, 0x7e, 0xd5, 0x4e, 0xfc, 0x51, 0xbc, 0xc1,
	0x2e, 0xfa, 0x3e, 0x8e, 0x4e, 0x86, 0xfa, 0x60, 0x27, 0x26, 0xb1, 0x81, 0x29, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0xb5, 0x0e, 0x7c, 0xc0, 0x00, 0x00, 0x00,
}
