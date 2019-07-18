// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

package shippy_service_consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

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

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

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

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

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
	proto.RegisterType((*Consignment)(nil), "shippy.service.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "shippy.service.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "shippy.service.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "shippy.service.consignment.Response")
}

func init() { proto.RegisterFile("consignment.proto", fileDescriptor_3804bf87090b51a9) }

var fileDescriptor_3804bf87090b51a9 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4e, 0xeb, 0x30,
	0x10, 0xc7, 0x5f, 0xfa, 0x9d, 0x49, 0xf5, 0xaa, 0x7a, 0xf1, 0x5e, 0x54, 0x16, 0x44, 0x11, 0x1f,
	0x5d, 0x65, 0x51, 0x8e, 0x50, 0xa1, 0x2a, 0x62, 0xe7, 0x1e, 0x00, 0x95, 0x78, 0x94, 0x5a, 0xa2,
	0x76, 0xc8, 0xb8, 0x45, 0x9c, 0x0d, 0x4e, 0xc2, 0x69, 0x50, 0x9c, 0x86, 0x1a, 0x21, 0x4a, 0x77,
	0xf9, 0x8f, 0xe7, 0x67, 0xff, 0x32, 0x1a, 0x18, 0x67, 0x5a, 0x91, 0xcc, 0xd5, 0x06, 0x95, 0x49,
	0x8a, 0x52, 0x1b, 0xcd, 0x26, 0xb4, 0x96, 0x45, 0xf1, 0x92, 0x10, 0x96, 0x3b, 0x99, 0x61, 0xe2,
	0x74, 0xc4, 0xaf, 0x1e, 0x04, 0xf3, 0x43, 0x66, 0x7f, 0xa1, 0x25, 0x45, 0xe8, 0x45, 0xde, 0xd4,
	0xe7, 0x2d, 0x29, 0x58, 0x04, 0x81, 0x40, 0xca, 0x4a, 0x59, 0x18, 0xa9, 0x55, 0xd8, 0xb2, 0x07,
	0x6e, 0x89, 0xfd, 0x83, 0xde, 0x33, 0xca, 0x7c, 0x6d, 0xc2, 0x76, 0xe4, 0x4d, 0xbb, 0x7c, 0x9f,
	0xd8, 0x2d, 0x40, 0xa6, 0x95, 0x59, 0x49, 0x85, 0x25, 0x85, 0x9d, 0xa8, 0x3d, 0x0d, 0x66, 0x97,
	0xc9, 0xcf, 0x2a, 0xc9, 0xbc, 0xe9, 0xe6, 0x0e, 0xc8, 0xce, 0xc0, 0xdf, 0x21, 0x11, 0x3e, 0xde,
	0x4b, 0x11, 0x76, 0xed, 0xf3, 0x83, 0xba, 0x90, 0x8a, 0x78, 0x03, 0xfe, 0x27, 0xf5, 0x4d, 0xfd,
	0x1c, 0x82, 0x6c, 0x4b, 0x46, 0x6f, 0xb0, 0xac, 0xd8, 0x5a, 0x1d, 0x9a, 0x52, 0x2a, 0x2a, 0x73,
	0x5d, 0xca, 0x5c, 0x2a, 0x6b, 0xee, 0xf3, 0x7d, 0x62, 0xff, 0xa1, 0xbf, 0xa5, 0x1a, 0xea, 0xd4,
	0x07, 0x55, 0x4c, 0x45, 0x3c, 0x04, 0x58, 0xa0, 0xe1, 0xf8, 0xb4, 0x45, 0x32, 0xf1, 0x9b, 0x07,
	0x03, 0x8e, 0x54, 0x68, 0x45, 0xc8, 0x42, 0xe8, 0x67, 0x25, 0xae, 0x0c, 0xd6, 0x06, 0x03, 0xde,
	0x44, 0x96, 0x42, 0xe0, 0xfc, 0xa5, 0xd5, 0x08, 0x66, 0xd7, 0xbf, 0x0c, 0xa2, 0xf9, 0xe6, 0x2e,
	0xcb, 0xee, 0x60, 0xe8, 0x44, 0x0a, 0xdb, 0x76, 0xa8, 0x27, 0xdf, 0xf5, 0x05, 0x9e, 0xbd, 0x7b,
	0x30, 0x5a, 0x56, 0xa0, 0x54, 0xf9, 0xb2, 0x26, 0x99, 0x80, 0xf1, 0xdc, 0x6a, 0xbb, 0x2b, 0x71,
	0xea, 0xfd, 0x93, 0x8b, 0x63, 0x8d, 0xcd, 0xa4, 0xe2, 0x3f, 0x6c, 0x05, 0xa3, 0x05, 0x1a, 0x87,
	0x24, 0x76, 0x75, 0x0c, 0x3d, 0xcc, 0xfc, 0xd4, 0x27, 0x1e, 0x7a, 0x76, 0xf3, 0x6f, 0x3e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xea, 0x87, 0x17, 0xe9, 0x0e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	// get Consignments
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
		serviceName = "shippy.service.consignment"
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
	CreateConsignment(context.Context, *Consignment, *Response) error
	// get Consignments
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
