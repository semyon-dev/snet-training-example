// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: service.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	training "snet-training-example/training"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SttResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SttResp) Reset() {
	*x = SttResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SttResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SttResp) ProtoMessage() {}

func (x *SttResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SttResp.ProtoReflect.Descriptor instead.
func (*SttResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *SttResp) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type SttInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelId *training.ModelID `protobuf:"bytes,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	Speech  []byte            `protobuf:"bytes,2,opt,name=speech,proto3" json:"speech,omitempty"`
}

func (x *SttInput) Reset() {
	*x = SttInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SttInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SttInput) ProtoMessage() {}

func (x *SttInput) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SttInput.ProtoReflect.Descriptor instead.
func (*SttInput) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *SttInput) GetModelId() *training.ModelID {
	if x != nil {
		return x.ModelId
	}
	return nil
}

func (x *SttInput) GetSpeech() []byte {
	if x != nil {
		return x.Speech
	}
	return nil
}

type TtsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TtsInput) Reset() {
	*x = TtsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtsInput) ProtoMessage() {}

func (x *TtsInput) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtsInput.ProtoReflect.Descriptor instead.
func (*TtsInput) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *TtsInput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TtsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TtsResponse) Reset() {
	*x = TtsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtsResponse) ProtoMessage() {}

func (x *TtsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtsResponse.ProtoReflect.Descriptor instead.
func (*TtsResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *TtsResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x07, 0x73,
	0x74, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x52,
	0x0a, 0x08, 0x73, 0x74, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x32, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x44, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x22, 0x1e, 0x0a, 0x08, 0x74, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x21, 0x0a, 0x0b, 0x74, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x96, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x12, 0x87, 0x01, 0x0a, 0x03, 0x73, 0x74, 0x74, 0x12, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x5b, 0x8a, 0xb5, 0x18, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x90, 0xb5,
	0x18, 0x05, 0x98, 0xb5, 0x18, 0x64, 0xa0, 0xb5, 0x18, 0x64, 0xa8, 0xb5, 0x18, 0x0a, 0xb2, 0xb5,
	0x18, 0x12, 0x70, 0x6e, 0x67, 0x2c, 0x20, 0x6d, 0x70, 0x34, 0x2c, 0x20, 0x74, 0x78, 0x74, 0x2c,
	0x20, 0x6d, 0x70, 0x33, 0xba, 0xb5, 0x18, 0x0b, 0x7a, 0x69, 0x70, 0x2c, 0x20, 0x74, 0x61, 0x72,
	0x2e, 0x67, 0x7a, 0xc2, 0xb5, 0x18, 0x17, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x40,
	0x0a, 0x0c, 0x42, 0x61, 0x73, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x30,
	0x0a, 0x03, 0x74, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x74, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x74, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_proto_goTypes = []any{
	(*SttResp)(nil),          // 0: service.sttResp
	(*SttInput)(nil),         // 1: service.sttInput
	(*TtsInput)(nil),         // 2: service.ttsInput
	(*TtsResponse)(nil),      // 3: service.ttsResponse
	(*training.ModelID)(nil), // 4: trainingV2.ModelID
}
var file_service_proto_depIdxs = []int32{
	4, // 0: service.sttInput.model_id:type_name -> trainingV2.ModelID
	1, // 1: service.ProMethods.stt:input_type -> service.sttInput
	2, // 2: service.BasicMethods.tts:input_type -> service.ttsInput
	0, // 3: service.ProMethods.stt:output_type -> service.sttResp
	3, // 4: service.BasicMethods.tts:output_type -> service.ttsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SttResp); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SttInput); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TtsInput); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TtsResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
