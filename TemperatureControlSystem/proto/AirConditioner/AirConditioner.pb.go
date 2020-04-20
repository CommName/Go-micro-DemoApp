// Code generated by protoc-gen-go. DO NOT EDIT.
// source: AirConditioner.proto

package iots_temperature_srv_AirConditioner

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
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05856c5e66c1def, []int{0}
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

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type DeviceStatus struct {
	PowerOn              bool     `protobuf:"varint,1,opt,name=PowerOn,proto3" json:"PowerOn,omitempty"`
	HeatingMode          bool     `protobuf:"varint,2,opt,name=HeatingMode,proto3" json:"HeatingMode,omitempty"`
	Power                int32    `protobuf:"varint,3,opt,name=Power,proto3" json:"Power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceStatus) Reset()         { *m = DeviceStatus{} }
func (m *DeviceStatus) String() string { return proto.CompactTextString(m) }
func (*DeviceStatus) ProtoMessage()    {}
func (*DeviceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05856c5e66c1def, []int{1}
}

func (m *DeviceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceStatus.Unmarshal(m, b)
}
func (m *DeviceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceStatus.Marshal(b, m, deterministic)
}
func (m *DeviceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceStatus.Merge(m, src)
}
func (m *DeviceStatus) XXX_Size() int {
	return xxx_messageInfo_DeviceStatus.Size(m)
}
func (m *DeviceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceStatus proto.InternalMessageInfo

func (m *DeviceStatus) GetPowerOn() bool {
	if m != nil {
		return m.PowerOn
	}
	return false
}

func (m *DeviceStatus) GetHeatingMode() bool {
	if m != nil {
		return m.HeatingMode
	}
	return false
}

func (m *DeviceStatus) GetPower() int32 {
	if m != nil {
		return m.Power
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05856c5e66c1def, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "iots.temperature.srv.AirConditioner.Message")
	proto.RegisterType((*DeviceStatus)(nil), "iots.temperature.srv.AirConditioner.DeviceStatus")
	proto.RegisterType((*Empty)(nil), "iots.temperature.srv.AirConditioner.Empty")
}

func init() {
	proto.RegisterFile("AirConditioner.proto", fileDescriptor_f05856c5e66c1def)
}

var fileDescriptor_f05856c5e66c1def = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x71, 0xcc, 0x2c, 0x72,
	0xce, 0xcf, 0x4b, 0xc9, 0x2c, 0xc9, 0xcc, 0xcf, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0xce, 0xcc, 0x2f, 0x29, 0xd6, 0x2b, 0x49, 0xcd, 0x2d, 0x48, 0x2d, 0x4a, 0x2c, 0x29,
	0x2d, 0x4a, 0xd5, 0x2b, 0x2e, 0x2a, 0xd3, 0x43, 0x55, 0xaa, 0x24, 0xcd, 0xc5, 0xee, 0x9b, 0x5a,
	0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0x24, 0xc0, 0xc5, 0x5c, 0x9c, 0x58, 0x29, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0x62, 0x2a, 0x25, 0x70, 0xf1, 0xb8, 0xa4, 0x96, 0x65, 0x26, 0xa7, 0x06, 0x97,
	0x24, 0x96, 0x94, 0x16, 0x0b, 0x49, 0x70, 0xb1, 0x07, 0xe4, 0x97, 0xa7, 0x16, 0xf9, 0xe7, 0x81,
	0x55, 0x71, 0x04, 0xc1, 0xb8, 0x42, 0x0a, 0x5c, 0xdc, 0x1e, 0xa9, 0x89, 0x25, 0x99, 0x79, 0xe9,
	0xbe, 0xf9, 0x29, 0xa9, 0x12, 0x4c, 0x60, 0x59, 0x64, 0x21, 0x21, 0x11, 0x2e, 0x56, 0xb0, 0x62,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x89, 0x9d, 0x8b, 0xd5, 0x35, 0xb7, 0xa0,
	0xa4, 0xd2, 0xe8, 0x07, 0x23, 0x17, 0x1f, 0xaa, 0xd3, 0x84, 0x8a, 0xb8, 0xf8, 0x83, 0x53, 0x4b,
	0x50, 0x1c, 0x60, 0xa8, 0x47, 0x84, 0x9f, 0xf4, 0x90, 0xb5, 0x48, 0x69, 0x11, 0xa5, 0x05, 0xec,
	0x08, 0x25, 0x06, 0x90, 0x9d, 0xee, 0x68, 0x76, 0x92, 0x60, 0x80, 0x14, 0xe9, 0xee, 0x53, 0x62,
	0x48, 0x62, 0x03, 0x47, 0x97, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x7e, 0xb4, 0xbb, 0xc6,
	0x01, 0x00, 0x00,
}