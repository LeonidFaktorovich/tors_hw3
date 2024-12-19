// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: internal_service.proto

package internal_service

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

type VectorTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostTp []uint64 `protobuf:"varint,1,rep,packed,name=host_tp,json=hostTp,proto3" json:"host_tp,omitempty"`
}

func (x *VectorTime) Reset() {
	*x = VectorTime{}
	mi := &file_internal_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VectorTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorTime) ProtoMessage() {}

func (x *VectorTime) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorTime.ProtoReflect.Descriptor instead.
func (*VectorTime) Descriptor() ([]byte, []int) {
	return file_internal_service_proto_rawDescGZIP(), []int{0}
}

func (x *VectorTime) GetHostTp() []uint64 {
	if x != nil {
		return x.HostTp
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string      `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Tp   *VectorTime `protobuf:"bytes,2,opt,name=tp,proto3" json:"tp,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	mi := &file_internal_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_internal_service_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Value) GetTp() *VectorTime {
	if x != nil {
		return x.Tp
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	mi := &file_internal_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_internal_service_proto_rawDescGZIP(), []int{2}
}

func (x *Entry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Entry) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type EntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *EntriesRequest) Reset() {
	*x = EntriesRequest{}
	mi := &file_internal_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntriesRequest) ProtoMessage() {}

func (x *EntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntriesRequest.ProtoReflect.Descriptor instead.
func (*EntriesRequest) Descriptor() ([]byte, []int) {
	return file_internal_service_proto_rawDescGZIP(), []int{3}
}

func (x *EntriesRequest) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type EntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EntriesResponse) Reset() {
	*x = EntriesResponse{}
	mi := &file_internal_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntriesResponse) ProtoMessage() {}

func (x *EntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntriesResponse.ProtoReflect.Descriptor instead.
func (*EntriesResponse) Descriptor() ([]byte, []int) {
	return file_internal_service_proto_rawDescGZIP(), []int{4}
}

var File_internal_service_proto protoreflect.FileDescriptor

var file_internal_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0a, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x74,
	0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x54, 0x70, 0x22,
	0x38, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x02,
	0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x02, 0x74, 0x70, 0x22, 0x37, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x32, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x45, 0x0a, 0x0f, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0b,
	0x53, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x13, 0x5a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_service_proto_rawDescOnce sync.Once
	file_internal_service_proto_rawDescData = file_internal_service_proto_rawDesc
)

func file_internal_service_proto_rawDescGZIP() []byte {
	file_internal_service_proto_rawDescOnce.Do(func() {
		file_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_service_proto_rawDescData)
	})
	return file_internal_service_proto_rawDescData
}

var file_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_service_proto_goTypes = []any{
	(*VectorTime)(nil),      // 0: VectorTime
	(*Value)(nil),           // 1: Value
	(*Entry)(nil),           // 2: Entry
	(*EntriesRequest)(nil),  // 3: EntriesRequest
	(*EntriesResponse)(nil), // 4: EntriesResponse
}
var file_internal_service_proto_depIdxs = []int32{
	0, // 0: Value.tp:type_name -> VectorTime
	1, // 1: Entry.value:type_name -> Value
	2, // 2: EntriesRequest.entries:type_name -> Entry
	3, // 3: InternalService.SendEntries:input_type -> EntriesRequest
	4, // 4: InternalService.SendEntries:output_type -> EntriesResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_service_proto_init() }
func file_internal_service_proto_init() {
	if File_internal_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_service_proto_goTypes,
		DependencyIndexes: file_internal_service_proto_depIdxs,
		MessageInfos:      file_internal_service_proto_msgTypes,
	}.Build()
	File_internal_service_proto = out.File
	file_internal_service_proto_rawDesc = nil
	file_internal_service_proto_goTypes = nil
	file_internal_service_proto_depIdxs = nil
}
