// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: stream.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestListenTopic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conn       *Connection       `protobuf:"bytes,1,opt,name=Conn,proto3" json:"Conn,omitempty"`
	TopicGroup map[string]string `protobuf:"bytes,2,rep,name=TopicGroup,proto3" json:"TopicGroup,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequestListenTopic) Reset() {
	*x = RequestListenTopic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestListenTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestListenTopic) ProtoMessage() {}

func (x *RequestListenTopic) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestListenTopic.ProtoReflect.Descriptor instead.
func (*RequestListenTopic) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *RequestListenTopic) GetConn() *Connection {
	if x != nil {
		return x.Conn
	}
	return nil
}

func (x *RequestListenTopic) GetTopicGroup() map[string]string {
	if x != nil {
		return x.TopicGroup
	}
	return nil
}

type RequestSendStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conn *Connection `protobuf:"bytes,1,opt,name=Conn,proto3" json:"Conn,omitempty"` //Message Message=2;
}

func (x *RequestSendStream) Reset() {
	*x = RequestSendStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSendStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSendStream) ProtoMessage() {}

func (x *RequestSendStream) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSendStream.ProtoReflect.Descriptor instead.
func (*RequestSendStream) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *RequestSendStream) GetConn() *Connection {
	if x != nil {
		return x.Conn
	}
	return nil
}

type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`          //消息编号
	Sequence int64  `protobuf:"varint,2,opt,name=Sequence,proto3" json:"Sequence,omitempty"` //序列号
	Result   bool   `protobuf:"varint,3,opt,name=Result,proto3" json:"Result,omitempty"`     //返回结果
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseMessage) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ResponseMessage) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *ResponseMessage) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbf, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x46, 0x0a, 0x0a, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x1a, 0x3d, 0x0a, 0x0f, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x37, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x22, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x22, 0x59, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa4, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x30, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0b, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x32, 0x0a, 0x0a, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData = file_stream_proto_rawDesc
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proto_rawDescData)
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stream_proto_goTypes = []interface{}{
	(*RequestListenTopic)(nil), // 0: pb.RequestListenTopic
	(*RequestSendStream)(nil),  // 1: pb.RequestSendStream
	(*ResponseMessage)(nil),    // 2: pb.ResponseMessage
	nil,                        // 3: pb.RequestListenTopic.TopicGroupEntry
	(*Connection)(nil),         // 4: pb.Connection
	(*Message)(nil),            // 5: pb.Message
}
var file_stream_proto_depIdxs = []int32{
	4, // 0: pb.RequestListenTopic.Conn:type_name -> pb.Connection
	3, // 1: pb.RequestListenTopic.TopicGroup:type_name -> pb.RequestListenTopic.TopicGroupEntry
	4, // 2: pb.RequestSendStream.Conn:type_name -> pb.Connection
	4, // 3: pb.Stream.RequestConnect:input_type -> pb.Connection
	0, // 4: pb.Stream.WatchStream:input_type -> pb.RequestListenTopic
	5, // 5: pb.Stream.SendStream:input_type -> pb.Message
	4, // 6: pb.Stream.RequestConnect:output_type -> pb.Connection
	5, // 7: pb.Stream.WatchStream:output_type -> pb.Message
	2, // 8: pb.Stream.SendStream:output_type -> pb.ResponseMessage
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	file_message_proto_init()
	file_connect_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestListenTopic); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSendStream); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_rawDesc = nil
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
	RequestConnect(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Connection, error)
	WatchStream(ctx context.Context, in *RequestListenTopic, opts ...grpc.CallOption) (Stream_WatchStreamClient, error)
	SendStream(ctx context.Context, opts ...grpc.CallOption) (Stream_SendStreamClient, error)
}

type streamClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) RequestConnect(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := c.cc.Invoke(ctx, "/pb.Stream/RequestConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClient) WatchStream(ctx context.Context, in *RequestListenTopic, opts ...grpc.CallOption) (Stream_WatchStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/pb.Stream/WatchStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamWatchStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stream_WatchStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type streamWatchStreamClient struct {
	grpc.ClientStream
}

func (x *streamWatchStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamClient) SendStream(ctx context.Context, opts ...grpc.CallOption) (Stream_SendStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[1], "/pb.Stream/SendStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamSendStreamClient{stream}
	return x, nil
}

type Stream_SendStreamClient interface {
	Send(*Message) error
	Recv() (*ResponseMessage, error)
	grpc.ClientStream
}

type streamSendStreamClient struct {
	grpc.ClientStream
}

func (x *streamSendStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamSendStreamClient) Recv() (*ResponseMessage, error) {
	m := new(ResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
	RequestConnect(context.Context, *Connection) (*Connection, error)
	WatchStream(*RequestListenTopic, Stream_WatchStreamServer) error
	SendStream(Stream_SendStreamServer) error
}

// UnimplementedStreamServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (*UnimplementedStreamServer) RequestConnect(context.Context, *Connection) (*Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestConnect not implemented")
}
func (*UnimplementedStreamServer) WatchStream(*RequestListenTopic, Stream_WatchStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchStream not implemented")
}
func (*UnimplementedStreamServer) SendStream(Stream_SendStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SendStream not implemented")
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_RequestConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServer).RequestConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Stream/RequestConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServer).RequestConnect(ctx, req.(*Connection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stream_WatchStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestListenTopic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServer).WatchStream(m, &streamWatchStreamServer{stream})
}

type Stream_WatchStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type streamWatchStreamServer struct {
	grpc.ServerStream
}

func (x *streamWatchStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Stream_SendStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServer).SendStream(&streamSendStreamServer{stream})
}

type Stream_SendStreamServer interface {
	Send(*ResponseMessage) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type streamSendStreamServer struct {
	grpc.ServerStream
}

func (x *streamSendStreamServer) Send(m *ResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamSendStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestConnect",
			Handler:    _Stream_RequestConnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchStream",
			Handler:       _Stream_WatchStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendStream",
			Handler:       _Stream_SendStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
