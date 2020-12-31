// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: example/example.proto

package proto

import (
	context "context"
	_ "github.com/cgentron/api/proto"
	_ "github.com/cgentron/pluginamzn/proto"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Song ...
type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Artist ...
	Artist string `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
	// SongTitle ...
	SongTitle string `protobuf:"bytes,2,opt,name=song_title,json=songTitle,proto3" json:"song_title,omitempty"`
	// AlbumTitle ...
	AlbumTitle string `protobuf:"bytes,3,opt,name=album_title,json=albumTitle,proto3" json:"album_title,omitempty"`
	// Year ...
	Year string `protobuf:"bytes,4,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *Song) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *Song) GetSongTitle() string {
	if x != nil {
		return x.SongTitle
	}
	return ""
}

func (x *Song) GetAlbumTitle() string {
	if x != nil {
		return x.AlbumTitle
	}
	return ""
}

func (x *Song) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

// Insert ...
type Insert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Insert) Reset() {
	*x = Insert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insert) ProtoMessage() {}

func (x *Insert) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insert.ProtoReflect.Descriptor instead.
func (*Insert) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1}
}

// Request ...
type Insert_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
}

func (x *Insert_Request) Reset() {
	*x = Insert_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insert_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insert_Request) ProtoMessage() {}

func (x *Insert_Request) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insert_Request.ProtoReflect.Descriptor instead.
func (*Insert_Request) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Insert_Request) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

// Response ...
type Insert_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Insert_Response) Reset() {
	*x = Insert_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insert_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insert_Response) ProtoMessage() {}

func (x *Insert_Response) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insert_Response.ProtoReflect.Descriptor instead.
func (*Insert_Response) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Insert_Response) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_example_example_proto protoreflect.FileDescriptor

var file_example_example_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x72, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x6e, 0x67, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0xa6, 0x02, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x1a, 0xfb, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x04, 0x73, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x3a, 0xce,
	0x01, 0xf2, 0x94, 0x8c, 0x2f, 0x80, 0x01, 0x52, 0x7e, 0x0a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x6d, 0x61, 0x7a, 0x6e, 0x12, 0x5b, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x6d,
	0x7a, 0x6e, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x63, 0x35, 0x36, 0x37, 0x62,
	0x35, 0x66, 0x37, 0x66, 0x33, 0x34, 0x37, 0x62, 0x34, 0x62, 0x36, 0x31, 0x35, 0x33, 0x66, 0x63,
	0x63, 0x31, 0x39, 0x37, 0x30, 0x30, 0x61, 0x35, 0x39, 0x36, 0x31, 0x31, 0x62, 0x39, 0x62, 0x64,
	0x34, 0x63, 0x39, 0x2e, 0x7a, 0x69, 0x70, 0xfa, 0x94, 0x8c, 0x2f, 0x43, 0x52, 0x41, 0x0a, 0x36,
	0x61, 0x72, 0x6e, 0x3a, 0x61, 0x77, 0x73, 0x3a, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x3a, 0x65,
	0x75, 0x2d, 0x77, 0x65, 0x73, 0x74, 0x2d, 0x31, 0x3a, 0x32, 0x39, 0x31, 0x33, 0x33, 0x39, 0x30,
	0x38, 0x38, 0x39, 0x33, 0x35, 0x3a, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x6d,
	0x79, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x12, 0x07, 0x24, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x54, 0x1a,
	0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32,
	0x44, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_example_proto_rawDescOnce sync.Once
	file_example_example_proto_rawDescData = file_example_example_proto_rawDesc
)

func file_example_example_proto_rawDescGZIP() []byte {
	file_example_example_proto_rawDescOnce.Do(func() {
		file_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_example_proto_rawDescData)
	})
	return file_example_example_proto_rawDescData
}

var file_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_example_proto_goTypes = []interface{}{
	(*Song)(nil),            // 0: proto.Song
	(*Insert)(nil),          // 1: proto.Insert
	(*Insert_Request)(nil),  // 2: proto.Insert.Request
	(*Insert_Response)(nil), // 3: proto.Insert.Response
}
var file_example_example_proto_depIdxs = []int32{
	0, // 0: proto.Insert.Request.song:type_name -> proto.Song
	2, // 1: proto.Example.Insert:input_type -> proto.Insert.Request
	3, // 2: proto.Example.Insert:output_type -> proto.Insert.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_example_proto_init() }
func file_example_example_proto_init() {
	if File_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Song); i {
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
		file_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insert); i {
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
		file_example_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insert_Request); i {
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
		file_example_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insert_Response); i {
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
			RawDescriptor: file_example_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_example_proto_goTypes,
		DependencyIndexes: file_example_example_proto_depIdxs,
		MessageInfos:      file_example_example_proto_msgTypes,
	}.Build()
	File_example_example_proto = out.File
	file_example_example_proto_rawDesc = nil
	file_example_example_proto_goTypes = nil
	file_example_example_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	// Insert ...
	Insert(ctx context.Context, in *Insert_Request, opts ...grpc.CallOption) (*Insert_Response, error)
}

type exampleClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleClient(cc grpc.ClientConnInterface) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Insert(ctx context.Context, in *Insert_Request, opts ...grpc.CallOption) (*Insert_Response, error) {
	out := new(Insert_Response)
	err := c.cc.Invoke(ctx, "/proto.Example/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	// Insert ...
	Insert(context.Context, *Insert_Request) (*Insert_Response, error)
}

// UnimplementedExampleServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (*UnimplementedExampleServer) Insert(context.Context, *Insert_Request) (*Insert_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Insert_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Example/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Insert(ctx, req.(*Insert_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Example_Insert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/example.proto",
}
