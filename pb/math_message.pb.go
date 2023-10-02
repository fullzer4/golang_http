// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.3
// source: proto/math_message.proto

package pb

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

type Sum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float32 `protobuf:"fixed32,1,opt,name=a,proto3" json:"a,omitempty"`
	B float32 `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Sum) Reset() {
	*x = Sum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sum) ProtoMessage() {}

func (x *Sum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sum.ProtoReflect.Descriptor instead.
func (*Sum) Descriptor() ([]byte, []int) {
	return file_proto_math_message_proto_rawDescGZIP(), []int{0}
}

func (x *Sum) GetA() float32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Sum) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

type NewSumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum *Sum `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *NewSumRequest) Reset() {
	*x = NewSumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSumRequest) ProtoMessage() {}

func (x *NewSumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSumRequest.ProtoReflect.Descriptor instead.
func (*NewSumRequest) Descriptor() ([]byte, []int) {
	return file_proto_math_message_proto_rawDescGZIP(), []int{1}
}

func (x *NewSumRequest) GetSum() *Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

type NewSumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *NewSumResponse) Reset() {
	*x = NewSumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSumResponse) ProtoMessage() {}

func (x *NewSumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSumResponse.ProtoReflect.Descriptor instead.
func (*NewSumResponse) Descriptor() ([]byte, []int) {
	return file_proto_math_message_proto_rawDescGZIP(), []int{2}
}

func (x *NewSumResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_proto_math_message_proto protoreflect.FileDescriptor

var file_proto_math_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x22, 0x21, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x62, 0x22, 0x31, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x53,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x73, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x75, 0x6c, 0x6c, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x28, 0x0a, 0x0e, 0x4e,
	0x65, 0x77, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x4b, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x18, 0x2e, 0x66, 0x75,
	0x6c, 0x6c, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x75, 0x6c, 0x6c, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_math_message_proto_rawDescOnce sync.Once
	file_proto_math_message_proto_rawDescData = file_proto_math_message_proto_rawDesc
)

func file_proto_math_message_proto_rawDescGZIP() []byte {
	file_proto_math_message_proto_rawDescOnce.Do(func() {
		file_proto_math_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_math_message_proto_rawDescData)
	})
	return file_proto_math_message_proto_rawDescData
}

var file_proto_math_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_math_message_proto_goTypes = []interface{}{
	(*Sum)(nil),            // 0: fullcycle.Sum
	(*NewSumRequest)(nil),  // 1: fullcycle.NewSumRequest
	(*NewSumResponse)(nil), // 2: fullcycle.NewSumResponse
}
var file_proto_math_message_proto_depIdxs = []int32{
	0, // 0: fullcycle.NewSumRequest.sum:type_name -> fullcycle.Sum
	1, // 1: fullcycle.MathService.Sum:input_type -> fullcycle.NewSumRequest
	2, // 2: fullcycle.MathService.Sum:output_type -> fullcycle.NewSumResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_math_message_proto_init() }
func file_proto_math_message_proto_init() {
	if File_proto_math_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_math_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sum); i {
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
		file_proto_math_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSumRequest); i {
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
		file_proto_math_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSumResponse); i {
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
			RawDescriptor: file_proto_math_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_math_message_proto_goTypes,
		DependencyIndexes: file_proto_math_message_proto_depIdxs,
		MessageInfos:      file_proto_math_message_proto_msgTypes,
	}.Build()
	File_proto_math_message_proto = out.File
	file_proto_math_message_proto_rawDesc = nil
	file_proto_math_message_proto_goTypes = nil
	file_proto_math_message_proto_depIdxs = nil
}
