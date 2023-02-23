// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: proto/vision.proto

package proto

import (
	visionpb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GoogleCloudVision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoundingBox *visionpb.Vertex     `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	Duration    *durationpb.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *GoogleCloudVision) Reset() {
	*x = GoogleCloudVision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vision_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleCloudVision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleCloudVision) ProtoMessage() {}

func (x *GoogleCloudVision) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vision_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleCloudVision.ProtoReflect.Descriptor instead.
func (*GoogleCloudVision) Descriptor() ([]byte, []int) {
	return file_proto_vision_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleCloudVision) GetBoundingBox() *visionpb.Vertex {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

func (x *GoogleCloudVision) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

var File_proto_vision_proto protoreflect.FileDescriptor

var file_proto_vision_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0c, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52, 0x0b,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x12, 0x35, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x79, 0x61, 0x74, 0x74, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2d, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_vision_proto_rawDescOnce sync.Once
	file_proto_vision_proto_rawDescData = file_proto_vision_proto_rawDesc
)

func file_proto_vision_proto_rawDescGZIP() []byte {
	file_proto_vision_proto_rawDescOnce.Do(func() {
		file_proto_vision_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vision_proto_rawDescData)
	})
	return file_proto_vision_proto_rawDescData
}

var file_proto_vision_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_vision_proto_goTypes = []interface{}{
	(*GoogleCloudVision)(nil),   // 0: proto.GoogleCloudVision
	(*visionpb.Vertex)(nil),     // 1: google.cloud.vision.v1.Vertex
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_proto_vision_proto_depIdxs = []int32{
	1, // 0: proto.GoogleCloudVision.bounding_box:type_name -> google.cloud.vision.v1.Vertex
	2, // 1: proto.GoogleCloudVision.duration:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_vision_proto_init() }
func file_proto_vision_proto_init() {
	if File_proto_vision_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_vision_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleCloudVision); i {
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
			RawDescriptor: file_proto_vision_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_vision_proto_goTypes,
		DependencyIndexes: file_proto_vision_proto_depIdxs,
		MessageInfos:      file_proto_vision_proto_msgTypes,
	}.Build()
	File_proto_vision_proto = out.File
	file_proto_vision_proto_rawDesc = nil
	file_proto_vision_proto_goTypes = nil
	file_proto_vision_proto_depIdxs = nil
}
