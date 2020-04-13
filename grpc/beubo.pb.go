// Code generated by protoc-gen-go. DO NOT EDIT.
// source: beubo.proto

package beubo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request is a basic representation of a http request
type Request struct {
	Url                  string    `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method               string    `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers              []*Header `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_45059721eb3bc8d4, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Header represents all the headers of a request
type Header struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_45059721eb3bc8d4, []int{1}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Header) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// The request message containing the user's name.
type PluginMessage struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Identifier           string                    `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Host                 string                    `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Distributed          bool                      `protobuf:"varint,4,opt,name=distributed,proto3" json:"distributed,omitempty"`
	Caching              bool                      `protobuf:"varint,5,opt,name=caching,proto3" json:"caching,omitempty"`
	Endpoints            []*PluginMessage_Endpoint `protobuf:"bytes,6,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PluginMessage) Reset()         { *m = PluginMessage{} }
func (m *PluginMessage) String() string { return proto.CompactTextString(m) }
func (*PluginMessage) ProtoMessage()    {}
func (*PluginMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_45059721eb3bc8d4, []int{2}
}

func (m *PluginMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginMessage.Unmarshal(m, b)
}
func (m *PluginMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginMessage.Marshal(b, m, deterministic)
}
func (m *PluginMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginMessage.Merge(m, src)
}
func (m *PluginMessage) XXX_Size() int {
	return xxx_messageInfo_PluginMessage.Size(m)
}
func (m *PluginMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PluginMessage proto.InternalMessageInfo

func (m *PluginMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginMessage) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *PluginMessage) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PluginMessage) GetDistributed() bool {
	if m != nil {
		return m.Distributed
	}
	return false
}

func (m *PluginMessage) GetCaching() bool {
	if m != nil {
		return m.Caching
	}
	return false
}

func (m *PluginMessage) GetEndpoints() []*PluginMessage_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type PluginMessage_Endpoint struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Priority             int32    `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginMessage_Endpoint) Reset()         { *m = PluginMessage_Endpoint{} }
func (m *PluginMessage_Endpoint) String() string { return proto.CompactTextString(m) }
func (*PluginMessage_Endpoint) ProtoMessage()    {}
func (*PluginMessage_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_45059721eb3bc8d4, []int{2, 0}
}

func (m *PluginMessage_Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginMessage_Endpoint.Unmarshal(m, b)
}
func (m *PluginMessage_Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginMessage_Endpoint.Marshal(b, m, deterministic)
}
func (m *PluginMessage_Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginMessage_Endpoint.Merge(m, src)
}
func (m *PluginMessage_Endpoint) XXX_Size() int {
	return xxx_messageInfo_PluginMessage_Endpoint.Size(m)
}
func (m *PluginMessage_Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginMessage_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_PluginMessage_Endpoint proto.InternalMessageInfo

func (m *PluginMessage_Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginMessage_Endpoint) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

// Event is the struct used for plugins to communicate with Beubo.
type Event struct {
	Key                  string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data                 string     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Values               []*any.Any `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_45059721eb3bc8d4, []int{3}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Event) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Event) GetValues() []*any.Any {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "beubo.Request")
	proto.RegisterType((*Header)(nil), "beubo.Header")
	proto.RegisterType((*PluginMessage)(nil), "beubo.PluginMessage")
	proto.RegisterType((*PluginMessage_Endpoint)(nil), "beubo.PluginMessage.Endpoint")
	proto.RegisterType((*Event)(nil), "beubo.Event")
}

func init() { proto.RegisterFile("beubo.proto", fileDescriptor_45059721eb3bc8d4) }

var fileDescriptor_45059721eb3bc8d4 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0x7c, 0x4e, 0x28, 0x42, 0xa3, 0x82, 0x16, 0x4b, 0x20, 0xcb, 0x17, 0x22, 0x81,
	0xdc, 0x2a, 0x70, 0x82, 0x13, 0xa9, 0x2a, 0xb8, 0x20, 0x45, 0xbe, 0xc2, 0x65, 0x1d, 0x4f, 0x9d,
	0x55, 0xdd, 0xdd, 0xe0, 0x5d, 0x57, 0xf2, 0x7f, 0xe0, 0x47, 0xa3, 0xfd, 0x70, 0xdb, 0x48, 0xb9,
	0xcd, 0x7b, 0x3b, 0x5f, 0xef, 0xcd, 0xc2, 0xb2, 0xa4, 0xae, 0x54, 0xf9, 0xa1, 0x55, 0x46, 0xe1,
	0xc4, 0x81, 0xe4, 0x6d, 0xad, 0x54, 0xdd, 0xd0, 0xa5, 0x23, 0xcb, 0xee, 0xf6, 0x92, 0xcb, 0xde,
	0x67, 0x64, 0x7f, 0x60, 0x56, 0xd0, 0xdf, 0x8e, 0xb4, 0xc1, 0x57, 0x10, 0x77, 0x6d, 0xc3, 0xa2,
	0x34, 0x5a, 0x2d, 0x0a, 0x1b, 0xe2, 0x1b, 0x98, 0xde, 0x93, 0xd9, 0xab, 0x8a, 0x8d, 0x1c, 0x19,
	0x10, 0x7e, 0x80, 0xd9, 0x9e, 0x78, 0x45, 0xad, 0x66, 0x71, 0x1a, 0xaf, 0x96, 0xeb, 0xf3, 0xdc,
	0x4f, 0xfd, 0xe9, 0xd8, 0x62, 0x78, 0xcd, 0xd6, 0x30, 0xf5, 0x94, 0x6d, 0x7e, 0x47, 0xfd, 0xd0,
	0xfc, 0x8e, 0x7a, 0xdb, 0xfc, 0x81, 0x37, 0x1d, 0x69, 0x36, 0x4a, 0x63, 0xdb, 0xdc, 0xa3, 0xec,
	0xdf, 0x08, 0xce, 0xb7, 0x4d, 0x57, 0x0b, 0xf9, 0x8b, 0xb4, 0xe6, 0x35, 0x21, 0xc2, 0x58, 0xf2,
	0x7b, 0x0a, 0xc5, 0x2e, 0xc6, 0xf7, 0x00, 0xa2, 0x22, 0x69, 0xc4, 0xad, 0xa0, 0x36, 0xac, 0xf7,
	0x8c, 0xb1, 0x35, 0x7b, 0xa5, 0x0d, 0x8b, 0x7d, 0x8d, 0x8d, 0x31, 0x85, 0x65, 0x25, 0xb4, 0x69,
	0x45, 0xd9, 0x19, 0xaa, 0xd8, 0x38, 0x8d, 0x56, 0xf3, 0xe2, 0x39, 0x85, 0x0c, 0x66, 0x3b, 0xbe,
	0xdb, 0x0b, 0x59, 0xb3, 0x89, 0x7b, 0x1d, 0x20, 0x7e, 0x83, 0x05, 0xc9, 0xea, 0xa0, 0x84, 0x34,
	0x9a, 0x4d, 0x9d, 0xe8, 0x77, 0x41, 0xf4, 0xd1, 0xb2, 0xf9, 0x4d, 0xc8, 0x2a, 0x9e, 0xf2, 0x93,
	0xaf, 0x30, 0x1f, 0xe8, 0x93, 0x62, 0x12, 0x98, 0x1f, 0x5a, 0xa1, 0x5a, 0x61, 0x7a, 0x27, 0x65,
	0x52, 0x3c, 0xe2, 0xec, 0x37, 0x4c, 0x6e, 0x1e, 0x48, 0x9a, 0x13, 0x0e, 0x22, 0x8c, 0x2b, 0x6e,
	0x78, 0x50, 0xef, 0x62, 0xfc, 0xf4, 0xe8, 0xaa, 0xbf, 0xcc, 0x45, 0xee, 0x6f, 0x9f, 0x0f, 0xb7,
	0xcf, 0xbf, 0xcb, 0x7e, 0xf0, 0x7a, 0x2d, 0x61, 0xb1, 0xb1, 0x1a, 0x7e, 0x14, 0xdb, 0x6b, 0xfc,
	0x08, 0xb3, 0x6b, 0x25, 0x25, 0xed, 0x0c, 0xbe, 0x08, 0xd2, 0xdc, 0xe4, 0xe4, 0x08, 0x65, 0x67,
	0xab, 0xe8, 0x2a, 0xc2, 0x2f, 0x30, 0x0f, 0xff, 0x46, 0xe3, 0xc5, 0x29, 0x23, 0x92, 0x97, 0x81,
	0x0d, 0x69, 0xd9, 0xd9, 0x55, 0xb4, 0x79, 0x0d, 0xfe, 0x47, 0x6e, 0xc0, 0x8d, 0xdd, 0xda, 0x9d,
	0xb6, 0x51, 0x39, 0x75, 0xcb, 0x7d, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x06, 0x5c, 0xea, 0x05,
	0xbc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BeuboGRPCClient is the client API for BeuboGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeuboGRPCClient interface {
	// Bidirectional plugin connection
	Connect(ctx context.Context, opts ...grpc.CallOption) (BeuboGRPC_ConnectClient, error)
	// For plugins that want to listen to incoming requests and perform action on the data
	Requests(ctx context.Context, in *PluginMessage, opts ...grpc.CallOption) (BeuboGRPC_RequestsClient, error)
}

type beuboGRPCClient struct {
	cc *grpc.ClientConn
}

func NewBeuboGRPCClient(cc *grpc.ClientConn) BeuboGRPCClient {
	return &beuboGRPCClient{cc}
}

func (c *beuboGRPCClient) Connect(ctx context.Context, opts ...grpc.CallOption) (BeuboGRPC_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BeuboGRPC_serviceDesc.Streams[0], "/beubo.BeuboGRPC/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &beuboGRPCConnectClient{stream}
	return x, nil
}

type BeuboGRPC_ConnectClient interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type beuboGRPCConnectClient struct {
	grpc.ClientStream
}

func (x *beuboGRPCConnectClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *beuboGRPCConnectClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *beuboGRPCClient) Requests(ctx context.Context, in *PluginMessage, opts ...grpc.CallOption) (BeuboGRPC_RequestsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BeuboGRPC_serviceDesc.Streams[1], "/beubo.BeuboGRPC/Requests", opts...)
	if err != nil {
		return nil, err
	}
	x := &beuboGRPCRequestsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BeuboGRPC_RequestsClient interface {
	Recv() (*Request, error)
	grpc.ClientStream
}

type beuboGRPCRequestsClient struct {
	grpc.ClientStream
}

func (x *beuboGRPCRequestsClient) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BeuboGRPCServer is the server API for BeuboGRPC service.
type BeuboGRPCServer interface {
	// Bidirectional plugin connection
	Connect(BeuboGRPC_ConnectServer) error
	// For plugins that want to listen to incoming requests and perform action on the data
	Requests(*PluginMessage, BeuboGRPC_RequestsServer) error
}

// UnimplementedBeuboGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedBeuboGRPCServer struct {
}

func (*UnimplementedBeuboGRPCServer) Connect(srv BeuboGRPC_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedBeuboGRPCServer) Requests(req *PluginMessage, srv BeuboGRPC_RequestsServer) error {
	return status.Errorf(codes.Unimplemented, "method Requests not implemented")
}

func RegisterBeuboGRPCServer(s *grpc.Server, srv BeuboGRPCServer) {
	s.RegisterService(&_BeuboGRPC_serviceDesc, srv)
}

func _BeuboGRPC_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BeuboGRPCServer).Connect(&beuboGRPCConnectServer{stream})
}

type BeuboGRPC_ConnectServer interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type beuboGRPCConnectServer struct {
	grpc.ServerStream
}

func (x *beuboGRPCConnectServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *beuboGRPCConnectServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BeuboGRPC_Requests_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PluginMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BeuboGRPCServer).Requests(m, &beuboGRPCRequestsServer{stream})
}

type BeuboGRPC_RequestsServer interface {
	Send(*Request) error
	grpc.ServerStream
}

type beuboGRPCRequestsServer struct {
	grpc.ServerStream
}

func (x *beuboGRPCRequestsServer) Send(m *Request) error {
	return x.ServerStream.SendMsg(m)
}

var _BeuboGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "beubo.BeuboGRPC",
	HandlerType: (*BeuboGRPCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _BeuboGRPC_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Requests",
			Handler:       _BeuboGRPC_Requests_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "beubo.proto",
}
