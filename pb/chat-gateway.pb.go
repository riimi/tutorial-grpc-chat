// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat-gateway.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b278c71b6605e99, []int{0}
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

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("chat-gateway.proto", fileDescriptor_4b278c71b6605e99) }

var fileDescriptor_4b278c71b6605e99 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xce, 0x3d, 0x4a, 0x04, 0x31,
	0x14, 0xc0, 0x71, 0x12, 0x16, 0x65, 0xb3, 0x60, 0xf1, 0x04, 0x19, 0xc6, 0x0f, 0x24, 0x36, 0xb2,
	0xb0, 0x89, 0x1f, 0xdd, 0xf6, 0x56, 0x62, 0xa3, 0x85, 0x75, 0x32, 0xf3, 0x1c, 0x03, 0x3a, 0x09,
	0xc9, 0xdb, 0x51, 0x5b, 0xaf, 0xe0, 0x51, 0x3c, 0x8a, 0x57, 0xf0, 0x20, 0x92, 0xec, 0x2c, 0x88,
	0x60, 0x97, 0xcf, 0xdf, 0xff, 0x09, 0x68, 0x1e, 0x0d, 0x2d, 0x3a, 0x43, 0xf8, 0x62, 0xde, 0x54,
	0x88, 0x9e, 0x3c, 0xf0, 0x60, 0xeb, 0x83, 0xce, 0xfb, 0xee, 0x09, 0xb5, 0x09, 0x4e, 0x9b, 0xbe,
	0xf7, 0x64, 0xc8, 0xf9, 0x3e, 0xad, 0x5f, 0xd4, 0xfb, 0xe3, 0x6d, 0xd9, 0xd9, 0xd5, 0x83, 0xc6,
	0xe7, 0x40, 0xe3, 0x77, 0xb9, 0x10, 0xdb, 0x37, 0x98, 0x92, 0xe9, 0x10, 0x76, 0x04, 0x77, 0x6d,
	0xc5, 0x8e, 0xd9, 0xe9, 0xf4, 0x96, 0xbb, 0x16, 0x40, 0x4c, 0x08, 0x5f, 0xa9, 0xe2, 0xe5, 0xa4,
	0xac, 0x2f, 0x3e, 0x99, 0x98, 0xe5, 0x21, 0xee, 0x30, 0x0e, 0xae, 0x41, 0xb8, 0x16, 0x93, 0x84,
	0x7d, 0x0b, 0x33, 0x15, 0xac, 0x1a, 0xa1, 0x7a, 0x4f, 0xad, 0x8b, 0x6a, 0x53, 0x54, 0x57, 0xb9,
	0x28, 0x8f, 0xde, 0xbf, 0xbe, 0x3f, 0x78, 0x25, 0x77, 0xf5, 0x70, 0xae, 0xb3, 0x92, 0x30, 0x0e,
	0x18, 0x75, 0x16, 0x96, 0x6c, 0x0e, 0xf7, 0x62, 0x9a, 0x56, 0x36, 0x35, 0xd1, 0x59, 0x84, 0x7f,
	0x90, 0xfa, 0x77, 0x49, 0x9e, 0x14, 0xf1, 0x50, 0x56, 0x7f, 0xc5, 0x0d, 0xb3, 0x64, 0xf3, 0x33,
	0x66, 0xb7, 0x8a, 0x71, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x4f, 0x79, 0xac, 0x40, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*empty.Empty, error)
	Subscribe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ChatService_SubscribeClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.chatService/send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Subscribe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ChatService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/pb.chatService/subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_SubscribeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *chatServiceSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	Send(context.Context, *Message) (*empty.Empty, error)
	Subscribe(*empty.Empty, ChatService_SubscribeServer) error
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.chatService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Send(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).Subscribe(m, &chatServiceSubscribeServer{stream})
}

type ChatService_SubscribeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *chatServiceSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.chatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send",
			Handler:    _ChatService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "subscribe",
			Handler:       _ChatService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat-gateway.proto",
}