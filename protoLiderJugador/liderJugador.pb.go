// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: liderJugador.proto

package ProtoLiderJugadores

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

//PARTICIPAR DEL JUEGO
type Solicitud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Solicitud) Reset() {
	*x = Solicitud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liderJugador_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solicitud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solicitud) ProtoMessage() {}

func (x *Solicitud) ProtoReflect() protoreflect.Message {
	mi := &file_liderJugador_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solicitud.ProtoReflect.Descriptor instead.
func (*Solicitud) Descriptor() ([]byte, []int) {
	return file_liderJugador_proto_rawDescGZIP(), []int{0}
}

type RespuestaSolicitud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etapa int32 `protobuf:"varint,1,opt,name=etapa,proto3" json:"etapa,omitempty"`
}

func (x *RespuestaSolicitud) Reset() {
	*x = RespuestaSolicitud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liderJugador_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaSolicitud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaSolicitud) ProtoMessage() {}

func (x *RespuestaSolicitud) ProtoReflect() protoreflect.Message {
	mi := &file_liderJugador_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaSolicitud.ProtoReflect.Descriptor instead.
func (*RespuestaSolicitud) Descriptor() ([]byte, []int) {
	return file_liderJugador_proto_rawDescGZIP(), []int{1}
}

func (x *RespuestaSolicitud) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

//ENVIAR JUGADAS
type Jugada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jugada int32 `protobuf:"varint,1,opt,name=jugada,proto3" json:"jugada,omitempty"`
}

func (x *Jugada) Reset() {
	*x = Jugada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liderJugador_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jugada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jugada) ProtoMessage() {}

func (x *Jugada) ProtoReflect() protoreflect.Message {
	mi := &file_liderJugador_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jugada.ProtoReflect.Descriptor instead.
func (*Jugada) Descriptor() ([]byte, []int) {
	return file_liderJugador_proto_rawDescGZIP(), []int{2}
}

func (x *Jugada) GetJugada() int32 {
	if x != nil {
		return x.Jugada
	}
	return 0
}

type RespuestaJugada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eliminado bool  `protobuf:"varint,1,opt,name=eliminado,proto3" json:"eliminado,omitempty"`
	Etapa     int32 `protobuf:"varint,2,opt,name=etapa,proto3" json:"etapa,omitempty"`
}

func (x *RespuestaJugada) Reset() {
	*x = RespuestaJugada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liderJugador_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaJugada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaJugada) ProtoMessage() {}

func (x *RespuestaJugada) ProtoReflect() protoreflect.Message {
	mi := &file_liderJugador_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaJugada.ProtoReflect.Descriptor instead.
func (*RespuestaJugada) Descriptor() ([]byte, []int) {
	return file_liderJugador_proto_rawDescGZIP(), []int{3}
}

func (x *RespuestaJugada) GetEliminado() bool {
	if x != nil {
		return x.Eliminado
	}
	return false
}

func (x *RespuestaJugada) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

var File_liderJugador_proto protoreflect.FileDescriptor

var file_liderJugador_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x22, 0x2a, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x75,
	0x65, 0x73, 0x74, 0x61, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x74,
	0x61, 0x70, 0x61, 0x22, 0x20, 0x0a, 0x06, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6a,
	0x75, 0x67, 0x61, 0x64, 0x61, 0x22, 0x45, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x6e, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x6e, 0x61, 0x64, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x32, 0x7e, 0x0a, 0x07,
	0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x3e, 0x0a, 0x11, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x61, 0x72, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x12, 0x0f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x1a, 0x18, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x53, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x12, 0x33, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x69, 0x61,
	0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a,
	0x75, 0x67, 0x61, 0x64, 0x61, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x42, 0x3b, 0x5a, 0x39,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6e, 0x43, 0x72, 0x61, 0x73, 0x74, 0x69, 0x76, 0x2f, 0x49, 0x4e,
	0x46, 0x2d, 0x33, 0x34, 0x33, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x64, 0x65, 0x72,
	0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_liderJugador_proto_rawDescOnce sync.Once
	file_liderJugador_proto_rawDescData = file_liderJugador_proto_rawDesc
)

func file_liderJugador_proto_rawDescGZIP() []byte {
	file_liderJugador_proto_rawDescOnce.Do(func() {
		file_liderJugador_proto_rawDescData = protoimpl.X.CompressGZIP(file_liderJugador_proto_rawDescData)
	})
	return file_liderJugador_proto_rawDescData
}

var file_liderJugador_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_liderJugador_proto_goTypes = []interface{}{
	(*Solicitud)(nil),          // 0: grpc.Solicitud
	(*RespuestaSolicitud)(nil), // 1: grpc.RespuestaSolicitud
	(*Jugada)(nil),             // 2: grpc.Jugada
	(*RespuestaJugada)(nil),    // 3: grpc.RespuestaJugada
}
var file_liderJugador_proto_depIdxs = []int32{
	0, // 0: grpc.Jugador.IngresarSolicitud:input_type -> grpc.Solicitud
	2, // 1: grpc.Jugador.EnviarJugada:input_type -> grpc.Jugada
	1, // 2: grpc.Jugador.IngresarSolicitud:output_type -> grpc.RespuestaSolicitud
	3, // 3: grpc.Jugador.EnviarJugada:output_type -> grpc.RespuestaJugada
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_liderJugador_proto_init() }
func file_liderJugador_proto_init() {
	if File_liderJugador_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_liderJugador_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solicitud); i {
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
		file_liderJugador_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaSolicitud); i {
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
		file_liderJugador_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jugada); i {
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
		file_liderJugador_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaJugada); i {
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
			RawDescriptor: file_liderJugador_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_liderJugador_proto_goTypes,
		DependencyIndexes: file_liderJugador_proto_depIdxs,
		MessageInfos:      file_liderJugador_proto_msgTypes,
	}.Build()
	File_liderJugador_proto = out.File
	file_liderJugador_proto_rawDesc = nil
	file_liderJugador_proto_goTypes = nil
	file_liderJugador_proto_depIdxs = nil
}