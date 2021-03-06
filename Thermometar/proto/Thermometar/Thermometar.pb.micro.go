// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: Thermometar.proto

package iots_temperature_srv_Thermometar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Thermometar service

type ThermometarService interface {
	GetStatus(ctx context.Context, in *Empty, opts ...client.CallOption) (*RoomTemperatrue, error)
	CoolTheRoom(ctx context.Context, in *Degrees, opts ...client.CallOption) (*RoomTemperatrue, error)
	HeatTheRoom(ctx context.Context, in *Degrees, opts ...client.CallOption) (*RoomTemperatrue, error)
}

type thermometarService struct {
	c    client.Client
	name string
}

func NewThermometarService(name string, c client.Client) ThermometarService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "iots.temperature.srv.Thermometar"
	}
	return &thermometarService{
		c:    c,
		name: name,
	}
}

func (c *thermometarService) GetStatus(ctx context.Context, in *Empty, opts ...client.CallOption) (*RoomTemperatrue, error) {
	req := c.c.NewRequest(c.name, "Thermometar.GetStatus", in)
	out := new(RoomTemperatrue)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thermometarService) CoolTheRoom(ctx context.Context, in *Degrees, opts ...client.CallOption) (*RoomTemperatrue, error) {
	req := c.c.NewRequest(c.name, "Thermometar.CoolTheRoom", in)
	out := new(RoomTemperatrue)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thermometarService) HeatTheRoom(ctx context.Context, in *Degrees, opts ...client.CallOption) (*RoomTemperatrue, error) {
	req := c.c.NewRequest(c.name, "Thermometar.HeatTheRoom", in)
	out := new(RoomTemperatrue)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Thermometar service

type ThermometarHandler interface {
	GetStatus(context.Context, *Empty, *RoomTemperatrue) error
	CoolTheRoom(context.Context, *Degrees, *RoomTemperatrue) error
	HeatTheRoom(context.Context, *Degrees, *RoomTemperatrue) error
}

func RegisterThermometarHandler(s server.Server, hdlr ThermometarHandler, opts ...server.HandlerOption) error {
	type thermometar interface {
		GetStatus(ctx context.Context, in *Empty, out *RoomTemperatrue) error
		CoolTheRoom(ctx context.Context, in *Degrees, out *RoomTemperatrue) error
		HeatTheRoom(ctx context.Context, in *Degrees, out *RoomTemperatrue) error
	}
	type Thermometar struct {
		thermometar
	}
	h := &thermometarHandler{hdlr}
	return s.Handle(s.NewHandler(&Thermometar{h}, opts...))
}

type thermometarHandler struct {
	ThermometarHandler
}

func (h *thermometarHandler) GetStatus(ctx context.Context, in *Empty, out *RoomTemperatrue) error {
	return h.ThermometarHandler.GetStatus(ctx, in, out)
}

func (h *thermometarHandler) CoolTheRoom(ctx context.Context, in *Degrees, out *RoomTemperatrue) error {
	return h.ThermometarHandler.CoolTheRoom(ctx, in, out)
}

func (h *thermometarHandler) HeatTheRoom(ctx context.Context, in *Degrees, out *RoomTemperatrue) error {
	return h.ThermometarHandler.HeatTheRoom(ctx, in, out)
}
