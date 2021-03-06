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

type DeviceStatus struct {
	PowerOn              bool     `protobuf:"varint,1,opt,name=PowerOn,proto3" json:"PowerOn,omitempty"`
	HeatingMode          bool     `protobuf:"varint,2,opt,name=HeatingMode,proto3" json:"HeatingMode,omitempty"`
	Power                int32    `protobuf:"varint,3,opt,name=Power,proto3" json:"Power,omitempty"`
	AutoMode             bool     `protobuf:"varint,4,opt,name=AutoMode,proto3" json:"AutoMode,omitempty"`
	DesiredTemp          int64    `protobuf:"varint,5,opt,name=DesiredTemp,proto3" json:"DesiredTemp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceStatus) Reset()         { *m = DeviceStatus{} }
func (m *DeviceStatus) String() string { return proto.CompactTextString(m) }
func (*DeviceStatus) ProtoMessage()    {}
func (*DeviceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05856c5e66c1def, []int{0}
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

func (m *DeviceStatus) GetAutoMode() bool {
	if m != nil {
		return m.AutoMode
	}
	return false
}

func (m *DeviceStatus) GetDesiredTemp() int64 {
	if m != nil {
		return m.DesiredTemp
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
	return fileDescriptor_f05856c5e66c1def, []int{1}
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
	proto.RegisterType((*DeviceStatus)(nil), "iots.temperature.srv.AirConditioner.DeviceStatus")
	proto.RegisterType((*Empty)(nil), "iots.temperature.srv.AirConditioner.Empty")
}

func init() {
	proto.RegisterFile("AirConditioner.proto", fileDescriptor_f05856c5e66c1def)
}

var fileDescriptor_f05856c5e66c1def = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0xd6, 0xd8, 0x32, 0x8a, 0xc2, 0xd2, 0xc3, 0x92, 0x53, 0x88, 0x97, 0xe0, 0x61,
	0x41, 0x7d, 0x82, 0x62, 0x45, 0x2f, 0xa2, 0xa4, 0xbe, 0x40, 0x34, 0x83, 0xec, 0x21, 0x3b, 0xcb,
	0xec, 0xa4, 0xe2, 0xcb, 0xf8, 0x9c, 0x1e, 0xa5, 0x2b, 0x4a, 0xe2, 0xa9, 0x3d, 0xfe, 0xb3, 0xf3,
	0xcd, 0x7c, 0xcb, 0xc0, 0x62, 0xe9, 0xf8, 0x86, 0x7c, 0xeb, 0xc4, 0x91, 0x47, 0xb6, 0x81, 0x49,
	0x48, 0x9f, 0x3b, 0x92, 0x68, 0x05, 0xbb, 0x80, 0xdc, 0x48, 0xcf, 0x68, 0x23, 0x6f, 0xec, 0xb8,
	0xb5, 0xfc, 0x54, 0x70, 0xb2, 0xc2, 0x8d, 0x7b, 0xc5, 0xb5, 0x34, 0xd2, 0x47, 0x6d, 0x60, 0xf6,
	0x44, 0xef, 0xc8, 0x8f, 0xde, 0xa8, 0x42, 0x55, 0xf3, 0xfa, 0x37, 0xea, 0x02, 0x8e, 0xef, 0xb1,
	0x11, 0xe7, 0xdf, 0x1e, 0xa8, 0x45, 0x73, 0x90, 0x5e, 0x87, 0x25, 0xbd, 0x80, 0x2c, 0x35, 0x9b,
	0x69, 0xa1, 0xaa, 0xac, 0xfe, 0x09, 0x3a, 0x87, 0xf9, 0xb2, 0x17, 0x4a, 0xd0, 0x61, 0x82, 0xfe,
	0xf2, 0x76, 0xe6, 0x0a, 0xa3, 0x63, 0x6c, 0x9f, 0xb1, 0x0b, 0x26, 0x2b, 0x54, 0x35, 0xad, 0x87,
	0xa5, 0x72, 0x06, 0xd9, 0x6d, 0x17, 0xe4, 0xe3, 0xea, 0x4b, 0xc1, 0xe9, 0x58, 0x5e, 0x33, 0x9c,
	0xad, 0x51, 0x46, 0xfa, 0x97, 0x76, 0x87, 0x5f, 0xdb, 0x21, 0x92, 0x5f, 0xec, 0x84, 0x24, 0x89,
	0x72, 0xb2, 0xdd, 0x79, 0xf7, 0x6f, 0xe7, 0x1e, 0x03, 0xf2, 0xfd, 0xfd, 0xca, 0xc9, 0xcb, 0x51,
	0x3a, 0xe8, 0xf5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x66, 0xe4, 0x63, 0xe8, 0x01, 0x00,
	0x00,
}
