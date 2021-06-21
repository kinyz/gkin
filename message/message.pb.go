// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: message.proto

package message

import (
	context "context"
	"gkin/connect"
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

type MessageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List map[uint64]*Message `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MessageList) Reset() {
	*x = MessageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageList) ProtoMessage() {}

func (x *MessageList) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageList.ProtoReflect.Descriptor instead.
func (*MessageList) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageList) GetList() map[uint64]*Message {
	if x != nil {
		return x.List
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string            `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Headers   map[string]string `protobuf:"bytes,2,rep,name=Headers,proto3" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Key       string            `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Value     []byte            `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
	TimesTamp int64             `protobuf:"varint,5,opt,name=TimesTamp,proto3" json:"TimesTamp,omitempty"` //时间戳
	IsConsume bool              `protobuf:"varint,6,opt,name=IsConsume,proto3" json:"IsConsume,omitempty"` //是否消费
	IsWrite   bool              `protobuf:"varint,7,opt,name=IsWrite,proto3" json:"IsWrite,omitempty"`     //是否写入硬盘
	Producer  string            `protobuf:"bytes,8,opt,name=Producer,proto3" json:"Producer,omitempty"`    //生产者
	Sequence  int64             `protobuf:"varint,9,opt,name=Sequence,proto3" json:"Sequence,omitempty"`   //序列号
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Message) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Message) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Message) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Message) GetTimesTamp() int64 {
	if x != nil {
		return x.TimesTamp
	}
	return 0
}

func (x *Message) GetIsConsume() bool {
	if x != nil {
		return x.IsConsume
	}
	return false
}

func (x *Message) GetIsWrite() bool {
	if x != nil {
		return x.IsWrite
	}
	return false
}

func (x *Message) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *Message) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type RespSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence  int64 `protobuf:"varint,1,opt,name=Sequence,proto3" json:"Sequence,omitempty"`   //序列号
	IsConsume bool  `protobuf:"varint,2,opt,name=IsConsume,proto3" json:"IsConsume,omitempty"` //是否消费
}

func (x *RespSend) Reset() {
	*x = RespSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespSend) ProtoMessage() {}

func (x *RespSend) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespSend.ProtoReflect.Descriptor instead.
func (*RespSend) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *RespSend) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *RespSend) GetIsConsume() bool {
	if x != nil {
		return x.IsConsume
	}
	return false
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x44, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc5, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x32, 0x0a, 0x07, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x54, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x54,
	0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x44, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x32, 0x92, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x28, 0x0a, 0x09, 0x41, 0x73, 0x79, 0x6e, 0x63,
	0x53, 0x65, 0x6e, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x28,
	0x01, 0x12, 0x25, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(*MessageList)(nil),        // 0: pb.MessageList
	(*Message)(nil),            // 1: pb.Message
	(*RespSend)(nil),           // 2: pb.RespSend
	nil,                        // 3: pb.MessageList.ListEntry
	nil,                        // 4: pb.Message.HeadersEntry
	(*connect.Connection)(nil), // 5: pb.Connection
}
var file_message_proto_depIdxs = []int32{
	3, // 0: pb.MessageList.List:type_name -> pb.MessageList.ListEntry
	4, // 1: pb.Message.Headers:type_name -> pb.Message.HeadersEntry
	1, // 2: pb.MessageList.ListEntry.value:type_name -> pb.Message
	1, // 3: pb.MessageStream.AsyncSend:input_type -> pb.Message
	1, // 4: pb.MessageStream.SyncSend:input_type -> pb.Message
	5, // 5: pb.MessageStream.RequestConnect:input_type -> pb.Connection
	2, // 6: pb.MessageStream.AsyncSend:output_type -> pb.RespSend
	2, // 7: pb.MessageStream.SyncSend:output_type -> pb.RespSend
	5, // 8: pb.MessageStream.RequestConnect:output_type -> pb.Connection
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	connect.File_connect_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageList); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespSend); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageStreamClient is the client API for MessageStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageStreamClient interface {
	//  rpc StreamChannel(stream Message)returns(stream Message);
	//  rpc SendChannel(stream Message)returns(Resp);
	AsyncSend(ctx context.Context, opts ...grpc.CallOption) (MessageStream_AsyncSendClient, error)
	SyncSend(ctx context.Context, in *Message, opts ...grpc.CallOption) (*RespSend, error)
	RequestConnect(ctx context.Context, in *connect.Connection, opts ...grpc.CallOption) (*connect.Connection, error)
}

type messageStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageStreamClient(cc grpc.ClientConnInterface) MessageStreamClient {
	return &messageStreamClient{cc}
}

func (c *messageStreamClient) AsyncSend(ctx context.Context, opts ...grpc.CallOption) (MessageStream_AsyncSendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageStream_serviceDesc.Streams[0], "/pb.MessageStream/AsyncSend", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageStreamAsyncSendClient{stream}
	return x, nil
}

type MessageStream_AsyncSendClient interface {
	Send(*Message) error
	CloseAndRecv() (*RespSend, error)
	grpc.ClientStream
}

type messageStreamAsyncSendClient struct {
	grpc.ClientStream
}

func (x *messageStreamAsyncSendClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageStreamAsyncSendClient) CloseAndRecv() (*RespSend, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RespSend)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageStreamClient) SyncSend(ctx context.Context, in *Message, opts ...grpc.CallOption) (*RespSend, error) {
	out := new(RespSend)
	err := c.cc.Invoke(ctx, "/pb.MessageStream/SyncSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStreamClient) RequestConnect(ctx context.Context, in *connect.Connection, opts ...grpc.CallOption) (*connect.Connection, error) {
	out := new(connect.Connection)
	err := c.cc.Invoke(ctx, "/pb.MessageStream/RequestConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageStreamServer is the server API for MessageStream service.
type MessageStreamServer interface {
	//  rpc StreamChannel(stream Message)returns(stream Message);
	//  rpc SendChannel(stream Message)returns(Resp);
	AsyncSend(MessageStream_AsyncSendServer) error
	SyncSend(context.Context, *Message) (*RespSend, error)
	RequestConnect(context.Context, *connect.Connection) (*connect.Connection, error)
}

// UnimplementedMessageStreamServer can be embedded to have forward compatible implementations.
type UnimplementedMessageStreamServer struct {
}

func (*UnimplementedMessageStreamServer) AsyncSend(MessageStream_AsyncSendServer) error {
	return status.Errorf(codes.Unimplemented, "method AsyncSend not implemented")
}
func (*UnimplementedMessageStreamServer) SyncSend(context.Context, *Message) (*RespSend, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncSend not implemented")
}
func (*UnimplementedMessageStreamServer) RequestConnect(context.Context, *connect.Connection) (*connect.Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestConnect not implemented")
}

func RegisterMessageStreamServer(s *grpc.Server, srv MessageStreamServer) {
	s.RegisterService(&_MessageStream_serviceDesc, srv)
}

func _MessageStream_AsyncSend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageStreamServer).AsyncSend(&messageStreamAsyncSendServer{stream})
}

type MessageStream_AsyncSendServer interface {
	SendAndClose(*RespSend) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messageStreamAsyncSendServer struct {
	grpc.ServerStream
}

func (x *messageStreamAsyncSendServer) SendAndClose(m *RespSend) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageStreamAsyncSendServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageStream_SyncSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStreamServer).SyncSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageStream/SyncSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStreamServer).SyncSend(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStream_RequestConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(connect.Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStreamServer).RequestConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageStream/RequestConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStreamServer).RequestConnect(ctx, req.(*connect.Connection))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageStream",
	HandlerType: (*MessageStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncSend",
			Handler:    _MessageStream_SyncSend_Handler,
		},
		{
			MethodName: "RequestConnect",
			Handler:    _MessageStream_RequestConnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AsyncSend",
			Handler:       _MessageStream_AsyncSend_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
