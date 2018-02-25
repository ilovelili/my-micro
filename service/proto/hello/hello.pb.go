// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ilovelili/my-micro/service/proto/hello/hello.proto

/*
Package go_micro_srv_greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/ilovelili/my-micro/service/proto/hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Say service

type SayClient interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type sayClient struct {
	c           client.Client
	serviceName string
}

func NewSayClient(serviceName string, c client.Client) SayClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.greeter"
	}
	return &sayClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Say.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterSayHandler(s server.Server, hdlr SayHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Say{hdlr}, opts...))
}

type Say struct {
	SayHandler
}

func (h *Say) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.SayHandler.Hello(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/ilovelili/my-micro/service/proto/hello/hello.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xce, 0xbf, 0x0a, 0xc2, 0x30,
	0x10, 0x06, 0x70, 0x4b, 0xfd, 0x9b, 0x49, 0x82, 0x83, 0x88, 0x15, 0xe9, 0xe4, 0xe2, 0x05, 0x74,
	0x76, 0xef, 0x26, 0xd4, 0x27, 0x68, 0xcb, 0x91, 0x06, 0x92, 0x5e, 0x4d, 0xd2, 0x42, 0xdf, 0x5e,
	0x1a, 0x3b, 0xea, 0x72, 0x7c, 0xdc, 0x77, 0xf0, 0x3b, 0xf6, 0x90, 0xca, 0xd7, 0x5d, 0x09, 0x15,
	0x19, 0xa1, 0x34, 0xf5, 0xa8, 0x95, 0x56, 0xc2, 0x0c, 0x57, 0xa3, 0x2a, 0x4b, 0xc2, 0xa1, 0xed,
	0x55, 0x85, 0xa2, 0xb5, 0xe4, 0x49, 0xd4, 0xa8, 0xf5, 0x34, 0x21, 0x6c, 0xf8, 0x4e, 0x12, 0x84,
	0x4b, 0x70, 0xb6, 0x07, 0x69, 0x11, 0x3d, 0xda, 0x34, 0x61, 0xab, 0x1c, 0xdf, 0x1d, 0x3a, 0xcf,
	0x39, 0x9b, 0x37, 0x85, 0xc1, 0x7d, 0x74, 0x8e, 0x2e, 0x9b, 0x3c, 0xe4, 0xf4, 0xc8, 0xd6, 0x39,
	0xba, 0x96, 0x1a, 0x87, 0x7c, 0xcb, 0x62, 0xe3, 0xe4, 0x54, 0x8f, 0xf1, 0xf6, 0x64, 0xf1, 0xab,
	0x18, 0x78, 0xc6, 0x16, 0xd9, 0x08, 0xf1, 0x04, 0x7e, 0x19, 0x30, 0x01, 0x87, 0xd3, 0xbf, 0xfa,
	0x0b, 0xa4, 0xb3, 0x72, 0x19, 0x5e, 0xbd, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x07, 0x80,
	0x52, 0xeb, 0x00, 0x00, 0x00,
}
