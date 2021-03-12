// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: sports/sports.proto

package sports

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListSportsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSportsRequest) Reset() {
	*x = ListSportsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSportsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSportsRequest) ProtoMessage() {}

func (x *ListSportsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSportsRequest.ProtoReflect.Descriptor instead.
func (*ListSportsRequest) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{0}
}

// Response to ListSports call.
type ListSportsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sports []*Sport `protobuf:"bytes,1,rep,name=sports,proto3" json:"sports,omitempty"`
}

func (x *ListSportsResponse) Reset() {
	*x = ListSportsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSportsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSportsResponse) ProtoMessage() {}

func (x *ListSportsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSportsResponse.ProtoReflect.Descriptor instead.
func (*ListSportsResponse) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{1}
}

func (x *ListSportsResponse) GetSports() []*Sport {
	if x != nil {
		return x.Sports
	}
	return nil
}

// A sports resource.
type Sport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the sport.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type is the type of the sport
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Name is the official name given to the sport.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// AdvertisedStartTime is the time the sport is advertised to start.
	AdvertisedStartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
}

func (x *Sport) Reset() {
	*x = Sport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sport) ProtoMessage() {}

func (x *Sport) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sport.ProtoReflect.Descriptor instead.
func (*Sport) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{2}
}

func (x *Sport) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sport) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Sport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sport) GetAdvertisedStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

var File_sports_sports_proto protoreflect.FileDescriptor

var file_sports_sports_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x8f,
	0x01, 0x0a, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x4e, 0x0a, 0x15, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x32, 0x69, 0x0a, 0x06, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sports_sports_proto_rawDescOnce sync.Once
	file_sports_sports_proto_rawDescData = file_sports_sports_proto_rawDesc
)

func file_sports_sports_proto_rawDescGZIP() []byte {
	file_sports_sports_proto_rawDescOnce.Do(func() {
		file_sports_sports_proto_rawDescData = protoimpl.X.CompressGZIP(file_sports_sports_proto_rawDescData)
	})
	return file_sports_sports_proto_rawDescData
}

var file_sports_sports_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sports_sports_proto_goTypes = []interface{}{
	(*ListSportsRequest)(nil),   // 0: sports.ListSportsRequest
	(*ListSportsResponse)(nil),  // 1: sports.ListSportsResponse
	(*Sport)(nil),               // 2: sports.Sport
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_sports_sports_proto_depIdxs = []int32{
	2, // 0: sports.ListSportsResponse.sports:type_name -> sports.Sport
	3, // 1: sports.Sport.advertised_start_time:type_name -> google.protobuf.Timestamp
	0, // 2: sports.Sports.ListSports:input_type -> sports.ListSportsRequest
	1, // 3: sports.Sports.ListSports:output_type -> sports.ListSportsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sports_sports_proto_init() }
func file_sports_sports_proto_init() {
	if File_sports_sports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sports_sports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSportsRequest); i {
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
		file_sports_sports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSportsResponse); i {
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
		file_sports_sports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sport); i {
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
			RawDescriptor: file_sports_sports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sports_sports_proto_goTypes,
		DependencyIndexes: file_sports_sports_proto_depIdxs,
		MessageInfos:      file_sports_sports_proto_msgTypes,
	}.Build()
	File_sports_sports_proto = out.File
	file_sports_sports_proto_rawDesc = nil
	file_sports_sports_proto_goTypes = nil
	file_sports_sports_proto_depIdxs = nil
}
