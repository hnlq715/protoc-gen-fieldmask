// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gogo.proto

package gogoproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
	Filename:      "gogo.proto",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65002,
	Name:          "gogoproto.embed",
	Tag:           "varint,65002,opt,name=embed",
	Filename:      "gogo.proto",
}

var E_Customname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65004,
	Name:          "gogoproto.customname",
	Tag:           "bytes,65004,opt,name=customname",
	Filename:      "gogo.proto",
}

func init() {
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Customname)
}

func init() { proto.RegisterFile("gogo.proto", fileDescriptor_592445b5231bc2b9) }

var fileDescriptor_592445b5231bc2b9 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcf, 0x4f, 0xcf,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xb1, 0xc1, 0x4c, 0x29, 0x85, 0xf4, 0xfc,
	0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x25, 0xb5, 0x38, 0xb9, 0x28,
	0xb3, 0xa0, 0x24, 0xbf, 0x08, 0xa2, 0xd8, 0xca, 0x9a, 0x8b, 0x23, 0xaf, 0x34, 0x27, 0x27, 0x31,
	0x29, 0x27, 0x55, 0x48, 0x56, 0x0f, 0xa2, 0x5c, 0x0f, 0xa6, 0x5c, 0xcf, 0x2d, 0x33, 0x35, 0x27,
	0xc5, 0xbf, 0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0xe2, 0xe5, 0x6f, 0x66, 0x05, 0x46, 0x0d, 0x8e,
	0x20, 0xb8, 0x06, 0x2b, 0x53, 0x2e, 0xd6, 0xd4, 0xdc, 0xa4, 0xd4, 0x14, 0x42, 0x3a, 0x5f, 0x41,
	0x75, 0x42, 0x54, 0x5b, 0xd9, 0x73, 0x71, 0x25, 0x97, 0x16, 0x97, 0xe4, 0xe7, 0xe6, 0x25, 0xe6,
	0x12, 0xb4, 0xf5, 0x0d, 0x58, 0x2f, 0x67, 0x10, 0x92, 0x16, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbc, 0xe8, 0x85, 0xa3, 0xee, 0x00, 0x00, 0x00,
}
