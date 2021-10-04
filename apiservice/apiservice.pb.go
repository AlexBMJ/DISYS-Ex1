// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: apiservice/apiservice.proto

package go_apiservice_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiservice_apiservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_apiservice_apiservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_apiservice_apiservice_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NewCourse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *NewCourse) Reset() {
	*x = NewCourse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiservice_apiservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCourse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCourse) ProtoMessage() {}

func (x *NewCourse) ProtoReflect() protoreflect.Message {
	mi := &file_apiservice_apiservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCourse.ProtoReflect.Descriptor instead.
func (*NewCourse) Descriptor() ([]byte, []int) {
	return file_apiservice_apiservice_proto_rawDescGZIP(), []int{1}
}

func (x *NewCourse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCourseByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetCourseByIdRequest) Reset() {
	*x = GetCourseByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiservice_apiservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseByIdRequest) ProtoMessage() {}

func (x *GetCourseByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apiservice_apiservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCourseByIdRequest) Descriptor() ([]byte, []int) {
	return file_apiservice_apiservice_proto_rawDescGZIP(), []int{2}
}

func (x *GetCourseByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Course *Course `protobuf:"bytes,2,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *UpdateCourseRequest) Reset() {
	*x = UpdateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiservice_apiservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseRequest) ProtoMessage() {}

func (x *UpdateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apiservice_apiservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourseRequest) Descriptor() ([]byte, []int) {
	return file_apiservice_apiservice_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCourseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCourseRequest) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

var File_apiservice_apiservice_proto protoreflect.FileDescriptor

var file_apiservice_apiservice_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61,
	0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x51, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x32, 0xd3, 0x02, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x45, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12,
	0x1f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x6f, 0x5f, 0x61, 0x70, 0x69, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apiservice_apiservice_proto_rawDescOnce sync.Once
	file_apiservice_apiservice_proto_rawDescData = file_apiservice_apiservice_proto_rawDesc
)

func file_apiservice_apiservice_proto_rawDescGZIP() []byte {
	file_apiservice_apiservice_proto_rawDescOnce.Do(func() {
		file_apiservice_apiservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_apiservice_apiservice_proto_rawDescData)
	})
	return file_apiservice_apiservice_proto_rawDescData
}

var file_apiservice_apiservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apiservice_apiservice_proto_goTypes = []interface{}{
	(*Course)(nil),               // 0: apiservice.Course
	(*NewCourse)(nil),            // 1: apiservice.NewCourse
	(*GetCourseByIdRequest)(nil), // 2: apiservice.GetCourseByIdRequest
	(*UpdateCourseRequest)(nil),  // 3: apiservice.UpdateCourseRequest
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_apiservice_apiservice_proto_depIdxs = []int32{
	0, // 0: apiservice.UpdateCourseRequest.course:type_name -> apiservice.Course
	1, // 1: apiservice.RouteCourse.AddCourse:input_type -> apiservice.NewCourse
	1, // 2: apiservice.RouteCourse.DeleteCourse:input_type -> apiservice.NewCourse
	2, // 3: apiservice.RouteCourse.GetCourseById:input_type -> apiservice.GetCourseByIdRequest
	4, // 4: apiservice.RouteCourse.GetOverview:input_type -> google.protobuf.Empty
	3, // 5: apiservice.RouteCourse.UpdateCourse:input_type -> apiservice.UpdateCourseRequest
	0, // 6: apiservice.RouteCourse.AddCourse:output_type -> apiservice.Course
	0, // 7: apiservice.RouteCourse.DeleteCourse:output_type -> apiservice.Course
	0, // 8: apiservice.RouteCourse.GetCourseById:output_type -> apiservice.Course
	0, // 9: apiservice.RouteCourse.GetOverview:output_type -> apiservice.Course
	0, // 10: apiservice.RouteCourse.UpdateCourse:output_type -> apiservice.Course
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apiservice_apiservice_proto_init() }
func file_apiservice_apiservice_proto_init() {
	if File_apiservice_apiservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apiservice_apiservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_apiservice_apiservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCourse); i {
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
		file_apiservice_apiservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseByIdRequest); i {
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
		file_apiservice_apiservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseRequest); i {
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
			RawDescriptor: file_apiservice_apiservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apiservice_apiservice_proto_goTypes,
		DependencyIndexes: file_apiservice_apiservice_proto_depIdxs,
		MessageInfos:      file_apiservice_apiservice_proto_msgTypes,
	}.Build()
	File_apiservice_apiservice_proto = out.File
	file_apiservice_apiservice_proto_rawDesc = nil
	file_apiservice_apiservice_proto_goTypes = nil
	file_apiservice_apiservice_proto_depIdxs = nil
}