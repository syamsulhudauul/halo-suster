// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: backend.proto

package pb

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

type NipState int32

const (
	NipState_PROCESS       NipState = 0
	NipState_FINISH_UNUSED NipState = 1
	NipState_FINISH_USED   NipState = 2
)

// Enum value maps for NipState.
var (
	NipState_name = map[int32]string{
		0: "PROCESS",
		1: "FINISH_UNUSED",
		2: "FINISH_USED",
	}
	NipState_value = map[string]int32{
		"PROCESS":       0,
		"FINISH_UNUSED": 1,
		"FINISH_USED":   2,
	}
)

func (x NipState) Enum() *NipState {
	p := new(NipState)
	*p = x
	return p
}

func (x NipState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NipState) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_proto_enumTypes[0].Descriptor()
}

func (NipState) Type() protoreflect.EnumType {
	return &file_backend_proto_enumTypes[0]
}

func (x NipState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NipState.Descriptor instead.
func (NipState) EnumDescriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{0}
}

type InitDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitDataRequest) Reset() {
	*x = InitDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDataRequest) ProtoMessage() {}

func (x *InitDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitDataRequest.ProtoReflect.Descriptor instead.
func (*InitDataRequest) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{0}
}

type InitDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitDataResponse) Reset() {
	*x = InitDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDataResponse) ProtoMessage() {}

func (x *InitDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitDataResponse.ProtoReflect.Descriptor instead.
func (*InitDataResponse) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{1}
}

type GetUnusedNipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsIt   bool     `protobuf:"varint,1,opt,name=is_it,json=isIt,proto3" json:"is_it,omitempty"`
	IsUsed bool     `protobuf:"varint,2,opt,name=is_used,json=isUsed,proto3" json:"is_used,omitempty"`
	State  NipState `protobuf:"varint,3,opt,name=state,proto3,enum=halo.suster.testcases.NipState" json:"state,omitempty"`
}

func (x *GetUnusedNipRequest) Reset() {
	*x = GetUnusedNipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnusedNipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnusedNipRequest) ProtoMessage() {}

func (x *GetUnusedNipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnusedNipRequest.ProtoReflect.Descriptor instead.
func (*GetUnusedNipRequest) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{2}
}

func (x *GetUnusedNipRequest) GetIsIt() bool {
	if x != nil {
		return x.IsIt
	}
	return false
}

func (x *GetUnusedNipRequest) GetIsUsed() bool {
	if x != nil {
		return x.IsUsed
	}
	return false
}

func (x *GetUnusedNipRequest) GetState() NipState {
	if x != nil {
		return x.State
	}
	return NipState_PROCESS
}

type GetUnusedNipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nip int64 `protobuf:"varint,1,opt,name=nip,proto3" json:"nip,omitempty"`
}

func (x *GetUnusedNipResponse) Reset() {
	*x = GetUnusedNipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnusedNipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnusedNipResponse) ProtoMessage() {}

func (x *GetUnusedNipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnusedNipResponse.ProtoReflect.Descriptor instead.
func (*GetUnusedNipResponse) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{3}
}

func (x *GetUnusedNipResponse) GetNip() int64 {
	if x != nil {
		return x.Nip
	}
	return 0
}

var File_backend_proto protoreflect.FileDescriptor

var file_backend_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x73, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x4e, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x69, 0x73, 0x49, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x75, 0x73, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12,
	0x35, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x73, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4e, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x75,
	0x73, 0x65, 0x64, 0x4e, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x69, 0x70,
	0x2a, 0x3b, 0x0a, 0x08, 0x4e, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x49, 0x4e,
	0x49, 0x53, 0x48, 0x5f, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x32, 0xc2, 0x01,
	0x0a, 0x0a, 0x4e, 0x49, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x08,
	0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x2e, 0x68, 0x61, 0x6c, 0x6f, 0x2e,
	0x73, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2e, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x68, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4e, 0x49, 0x50, 0x12, 0x2a, 0x2e, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x73, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x4e, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x73, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x75,
	0x73, 0x65, 0x64, 0x4e, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_backend_proto_rawDescOnce sync.Once
	file_backend_proto_rawDescData = file_backend_proto_rawDesc
)

func file_backend_proto_rawDescGZIP() []byte {
	file_backend_proto_rawDescOnce.Do(func() {
		file_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_proto_rawDescData)
	})
	return file_backend_proto_rawDescData
}

var file_backend_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_backend_proto_goTypes = []interface{}{
	(NipState)(0),                // 0: halo.suster.testcases.NipState
	(*InitDataRequest)(nil),      // 1: halo.suster.testcases.InitDataRequest
	(*InitDataResponse)(nil),     // 2: halo.suster.testcases.InitDataResponse
	(*GetUnusedNipRequest)(nil),  // 3: halo.suster.testcases.GetUnusedNipRequest
	(*GetUnusedNipResponse)(nil), // 4: halo.suster.testcases.GetUnusedNipResponse
	(*emptypb.Empty)(nil),        // 5: google.protobuf.Empty
}
var file_backend_proto_depIdxs = []int32{
	0, // 0: halo.suster.testcases.GetUnusedNipRequest.state:type_name -> halo.suster.testcases.NipState
	1, // 1: halo.suster.testcases.NIPService.InitData:input_type -> halo.suster.testcases.InitDataRequest
	3, // 2: halo.suster.testcases.NIPService.StreamNIP:input_type -> halo.suster.testcases.GetUnusedNipRequest
	5, // 3: halo.suster.testcases.NIPService.InitData:output_type -> google.protobuf.Empty
	4, // 4: halo.suster.testcases.NIPService.StreamNIP:output_type -> halo.suster.testcases.GetUnusedNipResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_backend_proto_init() }
func file_backend_proto_init() {
	if File_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitDataRequest); i {
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
		file_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitDataResponse); i {
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
		file_backend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnusedNipRequest); i {
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
		file_backend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnusedNipResponse); i {
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
			RawDescriptor: file_backend_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_proto_goTypes,
		DependencyIndexes: file_backend_proto_depIdxs,
		EnumInfos:         file_backend_proto_enumTypes,
		MessageInfos:      file_backend_proto_msgTypes,
	}.Build()
	File_backend_proto = out.File
	file_backend_proto_rawDesc = nil
	file_backend_proto_goTypes = nil
	file_backend_proto_depIdxs = nil
}
