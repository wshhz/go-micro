// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

/*
Package consignment is a generated protocol buffer package.

It is generated from these files:
	consignment.proto

It has these top-level messages:
	Consignment
	Container
	GetRequest
	Response
*/
package consignment

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

// 委托信息
type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

// 容器
type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// 创建一个空白的获取请求
type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 响应对象
type Response struct {
	Created      bool           `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment  *Consignment   `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	// 创建委托
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	// 获取委托
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	// 创建委托
	CreateConsignment(context.Context, *Consignment, *Response) error
	// 获取委托
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}

func init() { proto.RegisterFile("consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x6e, 0xe2, 0x30,
	0x14, 0x85, 0x27, 0xfc, 0xe7, 0x06, 0x0d, 0xc2, 0xd2, 0x80, 0x35, 0xb3, 0x98, 0x28, 0x2b, 0x56,
	0x2c, 0xa8, 0xd4, 0x45, 0xd5, 0x5d, 0xa4, 0x22, 0xb6, 0xe6, 0x01, 0x2a, 0x1a, 0x5f, 0x85, 0x2b,
	0x15, 0x3b, 0xb5, 0x0d, 0x7d, 0x1b, 0x9e, 0xa0, 0x0f, 0x59, 0xe1, 0x90, 0x62, 0x5a, 0xb1, 0xcb,
	0xb9, 0xe7, 0x1c, 0xe7, 0xf3, 0x4d, 0x60, 0x5c, 0x68, 0x65, 0xa9, 0x54, 0x3b, 0x54, 0x6e, 0x5e,
	0x19, 0xed, 0x34, 0x4b, 0x82, 0x51, 0xf6, 0x11, 0x41, 0x92, 0x5f, 0x34, 0xfb, 0x0d, 0x2d, 0x92,
	0x3c, 0x4a, 0xa3, 0x59, 0x2c, 0x5a, 0x24, 0x59, 0x0a, 0x89, 0x44, 0x5b, 0x18, 0xaa, 0x1c, 0x69,
	0xc5, 0x5b, 0xde, 0x08, 0x47, 0x6c, 0x02, 0xbd, 0x77, 0xa4, 0x72, 0xeb, 0x78, 0x3b, 0x8d, 0x66,
	0x5d, 0x71, 0x56, 0xec, 0x1e, 0xa0, 0xd0, 0xca, 0x6d, 0x48, 0xa1, 0xb1, 0xbc, 0x93, 0xb6, 0x67,
	0xc9, 0x62, 0x32, 0x0f, 0x71, 0xf2, 0xc6, 0x16, 0x41, 0x92, 0xfd, 0x83, 0xf8, 0x80, 0xd6, 0xe2,
	0xeb, 0x33, 0x49, 0xde, 0xf5, 0xef, 0x1b, 0xd4, 0x83, 0x95, 0xcc, 0x76, 0x10, 0x7f, 0xb5, 0x7e,
	0xb0, 0xfe, 0x87, 0xa4, 0xd8, 0x5b, 0xa7, 0x77, 0x68, 0x4e, 0xdd, 0x9a, 0x15, 0x9a, 0xd1, 0x4a,
	0x9e, 0x50, 0xb5, 0xa1, 0x92, 0x94, 0x47, 0x8d, 0xc5, 0x59, 0xb1, 0x29, 0xf4, 0xf7, 0xb6, 0x2e,
	0x75, 0x6a, 0xe3, 0x24, 0x57, 0x32, 0x1b, 0x02, 0x2c, 0xd1, 0x09, 0x7c, 0xdb, 0xa3, 0x75, 0xd9,
	0x31, 0x82, 0x81, 0x40, 0x5b, 0x69, 0x65, 0x91, 0x71, 0xe8, 0x17, 0x06, 0x37, 0x0e, 0x6b, 0x82,
	0x81, 0x68, 0x24, 0x7b, 0x80, 0x70, 0xc3, 0x1e, 0x23, 0x59, 0xf0, 0xef, 0x37, 0x6f, 0x9e, 0x45,
	0x18, 0x66, 0x8f, 0x30, 0x0c, 0xa4, 0xe5, 0x6d, 0xbf, 0xb6, 0xdb, 0xe5, 0xab, 0xf4, 0xe2, 0x18,
	0xc1, 0x68, 0xbd, 0xa5, 0xaa, 0x22, 0x55, 0xae, 0xd1, 0x1c, 0xa8, 0x40, 0xf6, 0x04, 0xe3, 0xdc,
	0x83, 0x85, 0x5f, 0xf9, 0xe6, 0x81, 0x7f, 0xff, 0x5c, 0x39, 0xcd, 0x6d, 0xb3, 0x5f, 0x2c, 0x87,
	0xd1, 0x12, 0x5d, 0x10, 0xb5, 0x6c, 0x7a, 0x95, 0xbd, 0x2c, 0xea, 0xe6, 0x21, 0x2f, 0x3d, 0xff,
	0x07, 0xde, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x05, 0x18, 0x52, 0x10, 0x96, 0x02, 0x00, 0x00,
}
