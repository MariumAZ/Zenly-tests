// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: gopher.proto

package Go_store

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

//Define query
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Area string `protobuf:"bytes,3,opt,name=area,proto3" json:"area,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gopher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_gopher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_gopher_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Query) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Query) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

// The request message containing the user's name.
type GopherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *GopherRequest) Reset() {
	*x = GopherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gopher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GopherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GopherRequest) ProtoMessage() {}

func (x *GopherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gopher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GopherRequest.ProtoReflect.Descriptor instead.
func (*GopherRequest) Descriptor() ([]byte, []int) {
	return file_gopher_proto_rawDescGZIP(), []int{1}
}

func (x *GopherRequest) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

// The response message containing the greetings
type GopherReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   *Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Created bool   `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated bool   `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *GopherReply) Reset() {
	*x = GopherReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gopher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GopherReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GopherReply) ProtoMessage() {}

func (x *GopherReply) ProtoReflect() protoreflect.Message {
	mi := &file_gopher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GopherReply.ProtoReflect.Descriptor instead.
func (*GopherReply) Descriptor() ([]byte, []int) {
	return file_gopher_proto_rawDescGZIP(), []int{2}
}

func (x *GopherReply) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *GopherReply) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

func (x *GopherReply) GetUpdated() bool {
	if x != nil {
		return x.Updated
	}
	return false
}

var File_gopher_proto protoreflect.FileDescriptor

var file_gopher_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x22, 0x34, 0x0a, 0x0d, 0x47, 0x6f, 0x70, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x66, 0x0a,
	0x0b, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x6f,
	0x70, 0x68, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x32, 0xbc, 0x01, 0x0a, 0x06, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x15, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e,
	0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x67,
	0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x70,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e,
	0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x41, 0x5a, 0x2f, 0x5a, 0x65, 0x6e, 0x6c,
	0x79, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x47, 0x6f, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gopher_proto_rawDescOnce sync.Once
	file_gopher_proto_rawDescData = file_gopher_proto_rawDesc
)

func file_gopher_proto_rawDescGZIP() []byte {
	file_gopher_proto_rawDescOnce.Do(func() {
		file_gopher_proto_rawDescData = protoimpl.X.CompressGZIP(file_gopher_proto_rawDescData)
	})
	return file_gopher_proto_rawDescData
}

var file_gopher_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gopher_proto_goTypes = []interface{}{
	(*Query)(nil),         // 0: gopher.Query
	(*GopherRequest)(nil), // 1: gopher.GopherRequest
	(*GopherReply)(nil),   // 2: gopher.GopherReply
}
var file_gopher_proto_depIdxs = []int32{
	0, // 0: gopher.GopherRequest.query:type_name -> gopher.Query
	0, // 1: gopher.GopherReply.query:type_name -> gopher.Query
	1, // 2: gopher.Gopher.CreateQuery:input_type -> gopher.GopherRequest
	1, // 3: gopher.Gopher.UpdateQuery:input_type -> gopher.GopherRequest
	1, // 4: gopher.Gopher.GetQuery:input_type -> gopher.GopherRequest
	2, // 5: gopher.Gopher.CreateQuery:output_type -> gopher.GopherReply
	2, // 6: gopher.Gopher.UpdateQuery:output_type -> gopher.GopherReply
	2, // 7: gopher.Gopher.GetQuery:output_type -> gopher.GopherReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gopher_proto_init() }
func file_gopher_proto_init() {
	if File_gopher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gopher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_gopher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GopherRequest); i {
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
		file_gopher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GopherReply); i {
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
			RawDescriptor: file_gopher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gopher_proto_goTypes,
		DependencyIndexes: file_gopher_proto_depIdxs,
		MessageInfos:      file_gopher_proto_msgTypes,
	}.Build()
	File_gopher_proto = out.File
	file_gopher_proto_rawDesc = nil
	file_gopher_proto_goTypes = nil
	file_gopher_proto_depIdxs = nil
}