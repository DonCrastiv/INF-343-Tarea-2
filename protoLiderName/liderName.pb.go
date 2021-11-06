// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: liderName.proto

package protoLiderName

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

//ENVIAR JUGADAS
type JugadaToName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdJugador int32 `protobuf:"varint,1,opt,name=idJugador,proto3" json:"idJugador,omitempty"`
	Jugada    int32 `protobuf:"varint,2,opt,name=jugada,proto3" json:"jugada,omitempty"`
	Etapa     int32 `protobuf:"varint,3,opt,name=etapa,proto3" json:"etapa,omitempty"`
}

func (x *JugadaToName) Reset() {
	*x = JugadaToName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liderName_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JugadaToName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JugadaToName) ProtoMessage() {}

func (x *JugadaToName) ProtoReflect() protoreflect.Message {
	mi := &file_liderName_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JugadaToName.ProtoReflect.Descriptor instead.
func (*JugadaToName) Descriptor() ([]byte, []int) {
	return file_liderName_proto_rawDescGZIP(), []int{0}
}

func (x *JugadaToName) GetIdJugador() int32 {
	if x != nil {
		return x.IdJugador
	}
	return 0
}

func (x *JugadaToName) GetJugada() int32 {
	if x != nil {
		return x.Jugada
	}
	return 0
}

func (x *JugadaToName) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

type RespuestaJugadas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jugadas  []int32 `protobuf:"varint,1,rep,packed,name=jugadas,proto3" json:"jugadas,omitempty"`
	Cantidad int32   `protobuf:"varint,2,opt,name=cantidad,proto3" json:"cantidad,omitempty"`
}

func (x *RespuestaJugadas) Reset() {
	*x = RespuestaJugadas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liderName_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaJugadas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaJugadas) ProtoMessage() {}

func (x *RespuestaJugadas) ProtoReflect() protoreflect.Message {
	mi := &file_liderName_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaJugadas.ProtoReflect.Descriptor instead.
func (*RespuestaJugadas) Descriptor() ([]byte, []int) {
	return file_liderName_proto_rawDescGZIP(), []int{1}
}

func (x *RespuestaJugadas) GetJugadas() []int32 {
	if x != nil {
		return x.Jugadas
	}
	return nil
}

func (x *RespuestaJugadas) GetCantidad() int32 {
	if x != nil {
		return x.Cantidad
	}
	return 0
}

var File_liderName_proto protoreflect.FileDescriptor

var file_liderName_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x5a, 0x0a, 0x0c, 0x4a, 0x75, 0x67, 0x61, 0x64,
	0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x4a, 0x75, 0x67,
	0x61, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x64, 0x4a, 0x75,
	0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x74,
	0x61, 0x70, 0x61, 0x22, 0x48, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x75, 0x67, 0x61, 0x64,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x32, 0x4f, 0x0a,
	0x10, 0x4c, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64,
	0x61, 0x73, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61,
	0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x73, 0x42, 0x36,
	0x5a, 0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6e, 0x43, 0x72, 0x61, 0x73, 0x74, 0x69, 0x76, 0x2f,
	0x49, 0x4e, 0x46, 0x2d, 0x33, 0x34, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_liderName_proto_rawDescOnce sync.Once
	file_liderName_proto_rawDescData = file_liderName_proto_rawDesc
)

func file_liderName_proto_rawDescGZIP() []byte {
	file_liderName_proto_rawDescOnce.Do(func() {
		file_liderName_proto_rawDescData = protoimpl.X.CompressGZIP(file_liderName_proto_rawDescData)
	})
	return file_liderName_proto_rawDescData
}

var file_liderName_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_liderName_proto_goTypes = []interface{}{
	(*JugadaToName)(nil),     // 0: grpc.JugadaToName
	(*RespuestaJugadas)(nil), // 1: grpc.RespuestaJugadas
}
var file_liderName_proto_depIdxs = []int32{
	0, // 0: grpc.LiderNameService.EnviarJugadas:input_type -> grpc.JugadaToName
	1, // 1: grpc.LiderNameService.EnviarJugadas:output_type -> grpc.RespuestaJugadas
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_liderName_proto_init() }
func file_liderName_proto_init() {
	if File_liderName_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_liderName_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JugadaToName); i {
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
		file_liderName_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaJugadas); i {
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
			RawDescriptor: file_liderName_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_liderName_proto_goTypes,
		DependencyIndexes: file_liderName_proto_depIdxs,
		MessageInfos:      file_liderName_proto_msgTypes,
	}.Build()
	File_liderName_proto = out.File
	file_liderName_proto_rawDesc = nil
	file_liderName_proto_goTypes = nil
	file_liderName_proto_depIdxs = nil
}
