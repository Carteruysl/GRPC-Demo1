// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: file.proto

package calculator

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

//傳入參數的spec
type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *CalculatorRequest) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *CalculatorRequest) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

//回傳值的spec
type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type GetFibonacciRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GetFibonacciRequest) Reset() {
	*x = GetFibonacciRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFibonacciRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFibonacciRequest) ProtoMessage() {}

func (x *GetFibonacciRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFibonacciRequest.ProtoReflect.Descriptor instead.
func (*GetFibonacciRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *GetFibonacciRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type GetFibonacciResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GetFibonacciResponse) Reset() {
	*x = GetFibonacciResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFibonacciResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFibonacciResponse) ProtoMessage() {}

func (x *GetFibonacciResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFibonacciResponse.ProtoReflect.Descriptor instead.
func (*GetFibonacciResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *GetFibonacciResponse) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0xb2, 0x01, 0x0a, 0x11, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63,
	0x63, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61,
	0x63, 0x63, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_proto_goTypes = []interface{}{
	(*CalculatorRequest)(nil),    // 0: calculator.CalculatorRequest
	(*CalculatorResponse)(nil),   // 1: calculator.CalculatorResponse
	(*GetFibonacciRequest)(nil),  // 2: calculator.GetFibonacciRequest
	(*GetFibonacciResponse)(nil), // 3: calculator.GetFibonacciResponse
}
var file_file_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.Sum:input_type -> calculator.CalculatorRequest
	2, // 1: calculator.CalculatorService.GetFibonacci:input_type -> calculator.GetFibonacciRequest
	1, // 2: calculator.CalculatorService.Sum:output_type -> calculator.CalculatorResponse
	3, // 3: calculator.CalculatorService.GetFibonacci:output_type -> calculator.GetFibonacciResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFibonacciRequest); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFibonacciResponse); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
