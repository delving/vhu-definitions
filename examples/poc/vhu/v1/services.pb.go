// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: vhu/v1/services.proto

package vhu

import (
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

// [start] object service
type GetObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetObjectRequest) Reset() {
	*x = GetObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vhu_v1_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectRequest) ProtoMessage() {}

func (x *GetObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vhu_v1_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectRequest.ProtoReflect.Descriptor instead.
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return file_vhu_v1_services_proto_rawDescGZIP(), []int{0}
}

func (x *GetObjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *ConceptualObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *GetObjectResponse) Reset() {
	*x = GetObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vhu_v1_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectResponse) ProtoMessage() {}

func (x *GetObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vhu_v1_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectResponse.ProtoReflect.Descriptor instead.
func (*GetObjectResponse) Descriptor() ([]byte, []int) {
	return file_vhu_v1_services_proto_rawDescGZIP(), []int{1}
}

func (x *GetObjectResponse) GetObject() *ConceptualObject {
	if x != nil {
		return x.Object
	}
	return nil
}

type ListObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListObjectsRequest) Reset() {
	*x = ListObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vhu_v1_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectsRequest) ProtoMessage() {}

func (x *ListObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vhu_v1_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectsRequest.ProtoReflect.Descriptor instead.
func (*ListObjectsRequest) Descriptor() ([]byte, []int) {
	return file_vhu_v1_services_proto_rawDescGZIP(), []int{2}
}

func (x *ListObjectsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListObjectsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListObjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*ConceptualObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *ListObjectsResponse) Reset() {
	*x = ListObjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vhu_v1_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectsResponse) ProtoMessage() {}

func (x *ListObjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vhu_v1_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectsResponse.ProtoReflect.Descriptor instead.
func (*ListObjectsResponse) Descriptor() ([]byte, []int) {
	return file_vhu_v1_services_proto_rawDescGZIP(), []int{3}
}

func (x *ListObjectsResponse) GetObjects() []*ConceptualObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

// [start] timeline service
type GetTimelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTimelineRequest) Reset() {
	*x = GetTimelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vhu_v1_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimelineRequest) ProtoMessage() {}

func (x *GetTimelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vhu_v1_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimelineRequest.ProtoReflect.Descriptor instead.
func (*GetTimelineRequest) Descriptor() ([]byte, []int) {
	return file_vhu_v1_services_proto_rawDescGZIP(), []int{4}
}

func (x *GetTimelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTimelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *Timeline `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *GetTimelineResponse) Reset() {
	*x = GetTimelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vhu_v1_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimelineResponse) ProtoMessage() {}

func (x *GetTimelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vhu_v1_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimelineResponse.ProtoReflect.Descriptor instead.
func (*GetTimelineResponse) Descriptor() ([]byte, []int) {
	return file_vhu_v1_services_proto_rawDescGZIP(), []int{5}
}

func (x *GetTimelineResponse) GetObject() *Timeline {
	if x != nil {
		return x.Object
	}
	return nil
}

var File_vhu_v1_services_proto protoreflect.FileDescriptor

var file_vhu_v1_services_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x68, 0x75, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x68, 0x75, 0x2e, 0x76, 0x31, 0x1a,
	0x18, 0x76, 0x68, 0x75, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x68, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63,
	0x65, 0x70, 0x74, 0x75, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76,
	0x68, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x75, 0x61, 0x6c,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76,
	0x68, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x32, 0xa3, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x63, 0x65,
	0x70, 0x74, 0x75, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x18, 0x2e, 0x76, 0x68, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x68, 0x75, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x76, 0x68, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x76, 0x68, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x59, 0x0a, 0x0f,
	0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a,
	0x2e, 0x76, 0x68, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x68, 0x75,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x07, 0x70, 0x6f, 0x63, 0x2f, 0x76,
	0x68, 0x75, 0xca, 0x02, 0x10, 0x41, 0x70, 0x70, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5c, 0x56, 0x48, 0x55, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vhu_v1_services_proto_rawDescOnce sync.Once
	file_vhu_v1_services_proto_rawDescData = file_vhu_v1_services_proto_rawDesc
)

func file_vhu_v1_services_proto_rawDescGZIP() []byte {
	file_vhu_v1_services_proto_rawDescOnce.Do(func() {
		file_vhu_v1_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_vhu_v1_services_proto_rawDescData)
	})
	return file_vhu_v1_services_proto_rawDescData
}

var file_vhu_v1_services_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vhu_v1_services_proto_goTypes = []interface{}{
	(*GetObjectRequest)(nil),    // 0: vhu.v1.GetObjectRequest
	(*GetObjectResponse)(nil),   // 1: vhu.v1.GetObjectResponse
	(*ListObjectsRequest)(nil),  // 2: vhu.v1.ListObjectsRequest
	(*ListObjectsResponse)(nil), // 3: vhu.v1.ListObjectsResponse
	(*GetTimelineRequest)(nil),  // 4: vhu.v1.GetTimelineRequest
	(*GetTimelineResponse)(nil), // 5: vhu.v1.GetTimelineResponse
	(*ConceptualObject)(nil),    // 6: vhu.v1.ConceptualObject
	(*Timeline)(nil),            // 7: vhu.v1.Timeline
}
var file_vhu_v1_services_proto_depIdxs = []int32{
	6, // 0: vhu.v1.GetObjectResponse.object:type_name -> vhu.v1.ConceptualObject
	6, // 1: vhu.v1.ListObjectsResponse.objects:type_name -> vhu.v1.ConceptualObject
	7, // 2: vhu.v1.GetTimelineResponse.object:type_name -> vhu.v1.Timeline
	0, // 3: vhu.v1.ConceptualObjectService.GetObject:input_type -> vhu.v1.GetObjectRequest
	2, // 4: vhu.v1.ConceptualObjectService.ListObjects:input_type -> vhu.v1.ListObjectsRequest
	4, // 5: vhu.v1.TimelineService.GetTimeline:input_type -> vhu.v1.GetTimelineRequest
	1, // 6: vhu.v1.ConceptualObjectService.GetObject:output_type -> vhu.v1.GetObjectResponse
	3, // 7: vhu.v1.ConceptualObjectService.ListObjects:output_type -> vhu.v1.ListObjectsResponse
	5, // 8: vhu.v1.TimelineService.GetTimeline:output_type -> vhu.v1.GetTimelineResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_vhu_v1_services_proto_init() }
func file_vhu_v1_services_proto_init() {
	if File_vhu_v1_services_proto != nil {
		return
	}
	file_vhu_v1_definitions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vhu_v1_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectRequest); i {
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
		file_vhu_v1_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectResponse); i {
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
		file_vhu_v1_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectsRequest); i {
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
		file_vhu_v1_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectsResponse); i {
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
		file_vhu_v1_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimelineRequest); i {
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
		file_vhu_v1_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimelineResponse); i {
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
			RawDescriptor: file_vhu_v1_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_vhu_v1_services_proto_goTypes,
		DependencyIndexes: file_vhu_v1_services_proto_depIdxs,
		MessageInfos:      file_vhu_v1_services_proto_msgTypes,
	}.Build()
	File_vhu_v1_services_proto = out.File
	file_vhu_v1_services_proto_rawDesc = nil
	file_vhu_v1_services_proto_goTypes = nil
	file_vhu_v1_services_proto_depIdxs = nil
}
