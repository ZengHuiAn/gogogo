// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

package greeter

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

type MessageType int32

const (
	MessageType_Request  MessageType = 0
	MessageType_Response MessageType = 1
	MessageType_Notify   MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "Request",
	1: "Response",
	2: "Notify",
}

var MessageType_value = map[string]int32{
	"Request":  0,
	"Response": 1,
	"Notify":   2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{0}
}

type SendRequest struct {
	MessageId            int32       `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Type                 MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Body                 []byte      `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *SendRequest) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_Request
}

func (m *SendRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type SendResponse struct {
	MessageId            int32       `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Type                 MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Body                 []byte      `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{1}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *SendResponse) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_Request
}

func (m *SendResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterEnum("MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*SendRequest)(nil), "SendRequest")
	proto.RegisterType((*SendResponse)(nil), "SendResponse")
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor_e585294ab3f34af5) }

var fileDescriptor_e585294ab3f34af5 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x9b, 0x5a, 0xab, 0x4e, 0x52, 0x59, 0xe6, 0xb4, 0x88, 0x87, 0xb0, 0xa7, 0x45, 0x30,
	0x87, 0xea, 0xd1, 0x17, 0xf0, 0xa0, 0x87, 0xd8, 0x17, 0xd8, 0x75, 0xc7, 0x92, 0x43, 0x93, 0x98,
	0x89, 0x42, 0xde, 0x5e, 0x48, 0x15, 0xf7, 0x05, 0x7a, 0x9b, 0x99, 0x7f, 0xe0, 0xe3, 0xff, 0x60,
	0xb3, 0x4f, 0x44, 0x99, 0x92, 0x89, 0x29, 0xe4, 0xd0, 0x0d, 0x20, 0xdf, 0xc8, 0x4f, 0x96, 0x3e,
	0xbf, 0x88, 0x33, 0xde, 0xc2, 0xd5, 0x81, 0x98, 0x87, 0x3d, 0x3d, 0x4f, 0xad, 0xd0, 0xa2, 0x3f,
	0xb7, 0xff, 0x07, 0xd4, 0xb0, 0xca, 0x25, 0x52, 0xbb, 0xd4, 0xa2, 0xbf, 0xde, 0x2a, 0xf3, 0x72,
	0x4c, 0x76, 0x25, 0x92, 0xad, 0x09, 0x22, 0xac, 0xc6, 0x30, 0x95, 0xf6, 0x4c, 0x8b, 0x5e, 0xd9,
	0x3a, 0x77, 0x23, 0xa8, 0x23, 0x82, 0x63, 0xf0, 0x4c, 0xa7, 0x60, 0xdc, 0x3d, 0x82, 0x9c, 0x3d,
	0xa2, 0x84, 0x8b, 0xdf, 0x46, 0xcd, 0x02, 0x15, 0x5c, 0xfe, 0xb1, 0x1b, 0x81, 0x00, 0xeb, 0xd7,
	0x90, 0xdd, 0x47, 0x69, 0x96, 0xdb, 0x27, 0x90, 0xbb, 0x34, 0x78, 0x7e, 0x27, 0xf7, 0x4d, 0x09,
	0xef, 0x41, 0xd5, 0xf5, 0xe0, 0x98, 0x5d, 0xf0, 0xa8, 0xcc, 0x4c, 0xcd, 0xcd, 0xc6, 0xcc, 0x5b,
	0x74, 0x8b, 0x71, 0x5d, 0x0d, 0x3e, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x79, 0x17, 0xc5, 0x14,
	0x52, 0x01, 0x00, 0x00,
}
