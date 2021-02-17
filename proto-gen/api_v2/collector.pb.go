// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: collector.proto

package api_v2

import (
	context "context"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	model "model"
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

type PostSpansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batch *model.Batch `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *PostSpansRequest) Reset() {
	*x = PostSpansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostSpansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSpansRequest) ProtoMessage() {}

func (x *PostSpansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSpansRequest.ProtoReflect.Descriptor instead.
func (*PostSpansRequest) Descriptor() ([]byte, []int) {
	return file_collector_proto_rawDescGZIP(), []int{0}
}

func (x *PostSpansRequest) GetBatch() *model.Batch {
	if x != nil {
		return x.Batch
	}
	return nil
}

type PostSpansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostSpansResponse) Reset() {
	*x = PostSpansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostSpansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSpansResponse) ProtoMessage() {}

func (x *PostSpansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSpansResponse.ProtoReflect.Descriptor instead.
func (*PostSpansResponse) Descriptor() ([]byte, []int) {
	return file_collector_proto_rawDescGZIP(), []int{1}
}

var File_collector_proto protoreflect.FileDescriptor

var file_collector_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32,
	0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x61,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x42, 0x04,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x13, 0x0a, 0x11, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x7c, 0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e,
	0x73, 0x12, 0x1f, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x32, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x32, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x76,
	0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x32, 0xc8, 0xe2, 0x1e, 0x01, 0xd0, 0xe2, 0x1e, 0x01, 0xe0, 0xe2, 0x1e, 0x01, 0xc0, 0xe3, 0x1e,
	0x01, 0x92, 0x41, 0x42, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x72,
	0x35, 0x0a, 0x0a, 0x4a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49, 0x12, 0x27, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f,
	0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collector_proto_rawDescOnce sync.Once
	file_collector_proto_rawDescData = file_collector_proto_rawDesc
)

func file_collector_proto_rawDescGZIP() []byte {
	file_collector_proto_rawDescOnce.Do(func() {
		file_collector_proto_rawDescData = protoimpl.X.CompressGZIP(file_collector_proto_rawDescData)
	})
	return file_collector_proto_rawDescData
}

var file_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_collector_proto_goTypes = []interface{}{
	(*PostSpansRequest)(nil),  // 0: jaeger.api_v2.PostSpansRequest
	(*PostSpansResponse)(nil), // 1: jaeger.api_v2.PostSpansResponse
	(*model.Batch)(nil),       // 2: jaeger.api_v2.Batch
}
var file_collector_proto_depIdxs = []int32{
	2, // 0: jaeger.api_v2.PostSpansRequest.batch:type_name -> jaeger.api_v2.Batch
	0, // 1: jaeger.api_v2.CollectorService.PostSpans:input_type -> jaeger.api_v2.PostSpansRequest
	1, // 2: jaeger.api_v2.CollectorService.PostSpans:output_type -> jaeger.api_v2.PostSpansResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_collector_proto_init() }
func file_collector_proto_init() {
	if File_collector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostSpansRequest); i {
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
		file_collector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostSpansResponse); i {
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
			RawDescriptor: file_collector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collector_proto_goTypes,
		DependencyIndexes: file_collector_proto_depIdxs,
		MessageInfos:      file_collector_proto_msgTypes,
	}.Build()
	File_collector_proto = out.File
	file_collector_proto_rawDesc = nil
	file_collector_proto_goTypes = nil
	file_collector_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CollectorServiceClient is the client API for CollectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectorServiceClient interface {
	PostSpans(ctx context.Context, in *PostSpansRequest, opts ...grpc.CallOption) (*PostSpansResponse, error)
}

type collectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorServiceClient(cc grpc.ClientConnInterface) CollectorServiceClient {
	return &collectorServiceClient{cc}
}

func (c *collectorServiceClient) PostSpans(ctx context.Context, in *PostSpansRequest, opts ...grpc.CallOption) (*PostSpansResponse, error) {
	out := new(PostSpansResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v2.CollectorService/PostSpans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorServiceServer is the server API for CollectorService service.
type CollectorServiceServer interface {
	PostSpans(context.Context, *PostSpansRequest) (*PostSpansResponse, error)
}

// UnimplementedCollectorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCollectorServiceServer struct {
}

func (*UnimplementedCollectorServiceServer) PostSpans(context.Context, *PostSpansRequest) (*PostSpansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSpans not implemented")
}

func RegisterCollectorServiceServer(s *grpc.Server, srv CollectorServiceServer) {
	s.RegisterService(&_CollectorService_serviceDesc, srv)
}

func _CollectorService_PostSpans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostSpansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServiceServer).PostSpans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v2.CollectorService/PostSpans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServiceServer).PostSpans(ctx, req.(*PostSpansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.api_v2.CollectorService",
	HandlerType: (*CollectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostSpans",
			Handler:    _CollectorService_PostSpans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector.proto",
}
