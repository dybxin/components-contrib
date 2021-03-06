// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv_store.proto

package cloudstate

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DeleteStateEnvelope struct {
	Key                  string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Options              *StateOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	Etag                 string        `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeleteStateEnvelope) Reset()         { *m = DeleteStateEnvelope{} }
func (m *DeleteStateEnvelope) String() string { return proto.CompactTextString(m) }
func (*DeleteStateEnvelope) ProtoMessage()    {}
func (*DeleteStateEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{0}
}

func (m *DeleteStateEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStateEnvelope.Unmarshal(m, b)
}
func (m *DeleteStateEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStateEnvelope.Marshal(b, m, deterministic)
}
func (m *DeleteStateEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStateEnvelope.Merge(m, src)
}
func (m *DeleteStateEnvelope) XXX_Size() int {
	return xxx_messageInfo_DeleteStateEnvelope.Size(m)
}
func (m *DeleteStateEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStateEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStateEnvelope proto.InternalMessageInfo

func (m *DeleteStateEnvelope) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DeleteStateEnvelope) GetOptions() *StateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *DeleteStateEnvelope) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

type SaveStateEnvelope struct {
	Key                  string               `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *any.Any             `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Etag                 string               `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	Metadata             map[string]string    `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Options              *StateRequestOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SaveStateEnvelope) Reset()         { *m = SaveStateEnvelope{} }
func (m *SaveStateEnvelope) String() string { return proto.CompactTextString(m) }
func (*SaveStateEnvelope) ProtoMessage()    {}
func (*SaveStateEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{1}
}

func (m *SaveStateEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveStateEnvelope.Unmarshal(m, b)
}
func (m *SaveStateEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveStateEnvelope.Marshal(b, m, deterministic)
}
func (m *SaveStateEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveStateEnvelope.Merge(m, src)
}
func (m *SaveStateEnvelope) XXX_Size() int {
	return xxx_messageInfo_SaveStateEnvelope.Size(m)
}
func (m *SaveStateEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveStateEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_SaveStateEnvelope proto.InternalMessageInfo

func (m *SaveStateEnvelope) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SaveStateEnvelope) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SaveStateEnvelope) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func (m *SaveStateEnvelope) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SaveStateEnvelope) GetOptions() *StateRequestOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetStateEnvelope struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Etag                 string   `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateEnvelope) Reset()         { *m = GetStateEnvelope{} }
func (m *GetStateEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetStateEnvelope) ProtoMessage()    {}
func (*GetStateEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{2}
}

func (m *GetStateEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateEnvelope.Unmarshal(m, b)
}
func (m *GetStateEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateEnvelope.Marshal(b, m, deterministic)
}
func (m *GetStateEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateEnvelope.Merge(m, src)
}
func (m *GetStateEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetStateEnvelope.Size(m)
}
func (m *GetStateEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateEnvelope proto.InternalMessageInfo

func (m *GetStateEnvelope) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetStateEnvelope) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

type GetStateResponseEnvelope struct {
	Data                 *any.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Etag                 string   `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateResponseEnvelope) Reset()         { *m = GetStateResponseEnvelope{} }
func (m *GetStateResponseEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetStateResponseEnvelope) ProtoMessage()    {}
func (*GetStateResponseEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{3}
}

func (m *GetStateResponseEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateResponseEnvelope.Unmarshal(m, b)
}
func (m *GetStateResponseEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateResponseEnvelope.Marshal(b, m, deterministic)
}
func (m *GetStateResponseEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateResponseEnvelope.Merge(m, src)
}
func (m *GetStateResponseEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetStateResponseEnvelope.Size(m)
}
func (m *GetStateResponseEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateResponseEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateResponseEnvelope proto.InternalMessageInfo

func (m *GetStateResponseEnvelope) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetStateResponseEnvelope) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

type StateOptions struct {
	Concurrency          string   `protobuf:"bytes,1,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Consistency          string   `protobuf:"bytes,2,opt,name=consistency,proto3" json:"consistency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateOptions) Reset()         { *m = StateOptions{} }
func (m *StateOptions) String() string { return proto.CompactTextString(m) }
func (*StateOptions) ProtoMessage()    {}
func (*StateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{4}
}

func (m *StateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateOptions.Unmarshal(m, b)
}
func (m *StateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateOptions.Marshal(b, m, deterministic)
}
func (m *StateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateOptions.Merge(m, src)
}
func (m *StateOptions) XXX_Size() int {
	return xxx_messageInfo_StateOptions.Size(m)
}
func (m *StateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StateOptions proto.InternalMessageInfo

func (m *StateOptions) GetConcurrency() string {
	if m != nil {
		return m.Concurrency
	}
	return ""
}

func (m *StateOptions) GetConsistency() string {
	if m != nil {
		return m.Consistency
	}
	return ""
}

type StateRequestOptions struct {
	Concurrency          string   `protobuf:"bytes,1,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Consistency          string   `protobuf:"bytes,2,opt,name=consistency,proto3" json:"consistency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateRequestOptions) Reset()         { *m = StateRequestOptions{} }
func (m *StateRequestOptions) String() string { return proto.CompactTextString(m) }
func (*StateRequestOptions) ProtoMessage()    {}
func (*StateRequestOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{5}
}

func (m *StateRequestOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateRequestOptions.Unmarshal(m, b)
}
func (m *StateRequestOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateRequestOptions.Marshal(b, m, deterministic)
}
func (m *StateRequestOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateRequestOptions.Merge(m, src)
}
func (m *StateRequestOptions) XXX_Size() int {
	return xxx_messageInfo_StateRequestOptions.Size(m)
}
func (m *StateRequestOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StateRequestOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StateRequestOptions proto.InternalMessageInfo

func (m *StateRequestOptions) GetConcurrency() string {
	if m != nil {
		return m.Concurrency
	}
	return ""
}

func (m *StateRequestOptions) GetConsistency() string {
	if m != nil {
		return m.Consistency
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteStateEnvelope)(nil), "cloudstate.DeleteStateEnvelope")
	proto.RegisterType((*SaveStateEnvelope)(nil), "cloudstate.SaveStateEnvelope")
	proto.RegisterMapType((map[string]string)(nil), "cloudstate.SaveStateEnvelope.MetadataEntry")
	proto.RegisterType((*GetStateEnvelope)(nil), "cloudstate.GetStateEnvelope")
	proto.RegisterType((*GetStateResponseEnvelope)(nil), "cloudstate.GetStateResponseEnvelope")
	proto.RegisterType((*StateOptions)(nil), "cloudstate.StateOptions")
	proto.RegisterType((*StateRequestOptions)(nil), "cloudstate.StateRequestOptions")
}

func init() { proto.RegisterFile("kv_store.proto", fileDescriptor_cd3673e574d3b8b3) }

var fileDescriptor_cd3673e574d3b8b3 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x9b, 0xec, 0x56, 0xbb, 0x67, 0xad, 0xac, 0xd3, 0x52, 0x62, 0x54, 0x5c, 0x82, 0x17,
	0x8b, 0x42, 0x0a, 0xeb, 0x8d, 0x7f, 0x40, 0x50, 0xba, 0x54, 0x10, 0x51, 0x66, 0x41, 0xf4, 0xaa,
	0x4c, 0xd3, 0xe3, 0xb2, 0x6c, 0x3a, 0x13, 0x33, 0x27, 0x81, 0xbc, 0x85, 0x0f, 0xe2, 0xad, 0x0f,
	0xe6, 0x1b, 0x48, 0x26, 0x99, 0xdd, 0x6c, 0x1a, 0x4b, 0x2f, 0xbc, 0x0b, 0xe7, 0x7c, 0xf9, 0xe6,
	0x7c, 0xbf, 0x99, 0x03, 0x77, 0x57, 0xf9, 0x99, 0x26, 0x95, 0x62, 0x98, 0xa4, 0x8a, 0x14, 0x83,
	0x28, 0x56, 0xd9, 0x85, 0x26, 0x41, 0xe8, 0xdf, 0x5f, 0x28, 0xb5, 0x88, 0xf1, 0xd8, 0x74, 0xce,
	0xb3, 0xef, 0xc7, 0x42, 0x16, 0x95, 0xcc, 0x7f, 0xd0, 0x6e, 0xe1, 0x65, 0x42, 0xb6, 0x39, 0x42,
	0x49, 0x4b, 0x2a, 0xce, 0x56, 0x58, 0x57, 0x82, 0x0c, 0x0e, 0x4e, 0x30, 0x46, 0xc2, 0x79, 0x69,
	0x3c, 0x93, 0x39, 0xc6, 0x2a, 0x41, 0x76, 0x04, 0xbd, 0x15, 0x16, 0x9e, 0x33, 0x76, 0x26, 0x83,
	0x77, 0xfd, 0x9f, 0xbf, 0x3d, 0x87, 0x97, 0x05, 0x36, 0x85, 0xdb, 0x2a, 0xa1, 0xa5, 0x92, 0xda,
	0x73, 0xc7, 0xce, 0x64, 0x38, 0xf5, 0xc2, 0xcd, 0x58, 0xa1, 0xf1, 0xf8, 0x54, 0xf5, 0xb9, 0x15,
	0x32, 0x06, 0x7d, 0x24, 0xb1, 0xf0, 0x7a, 0xa5, 0x19, 0x37, 0xdf, 0xc1, 0x2f, 0x17, 0xee, 0xcd,
	0x45, 0x7e, 0xc3, 0x53, 0x9f, 0xc2, 0x6e, 0x2e, 0xe2, 0x0c, 0xeb, 0x33, 0x0f, 0xc3, 0x2a, 0x63,
	0x68, 0x33, 0x86, 0x6f, 0x65, 0xc1, 0x2b, 0x49, 0xd7, 0x69, 0xec, 0x14, 0xf6, 0x2e, 0x91, 0xc4,
	0x85, 0x20, 0xe1, 0xf5, 0xc7, 0xbd, 0xc9, 0x70, 0xfa, 0x6c, 0x6b, 0xec, 0xf6, 0x20, 0xe1, 0xc7,
	0x5a, 0x3d, 0x93, 0x94, 0x16, 0x7c, 0xfd, 0x33, 0x7b, 0xb9, 0x89, 0xbf, 0x6b, 0x46, 0x79, 0x7c,
	0x25, 0x3e, 0xc7, 0x1f, 0x19, 0x6a, 0x6a, 0x53, 0xf0, 0x5f, 0xc3, 0xfe, 0x96, 0x2b, 0x1b, 0x35,
	0xc2, 0x56, 0x31, 0x0f, 0x9b, 0x31, 0x07, 0x75, 0xa0, 0x57, 0xee, 0x0b, 0x27, 0x78, 0x03, 0xa3,
	0x53, 0xa4, 0x9b, 0xc1, 0xb2, 0x00, 0xdc, 0x06, 0xee, 0xaf, 0xe0, 0xd9, 0xff, 0x39, 0xea, 0x44,
	0x49, 0xbd, 0xf1, 0x99, 0x40, 0xdf, 0x80, 0x71, 0xae, 0x61, 0x6b, 0x14, 0x9d, 0xce, 0x1c, 0xee,
	0x34, 0x6f, 0x9d, 0x8d, 0x61, 0x18, 0x29, 0x19, 0x65, 0x69, 0x8a, 0x32, 0xb2, 0xe9, 0x9a, 0xa5,
	0x5a, 0xa1, 0x97, 0x9a, 0x8c, 0xc2, 0x5d, 0x2b, 0x6c, 0x29, 0xf8, 0x06, 0x07, 0x1d, 0x28, 0xff,
	0x87, 0xf5, 0xf4, 0x8f, 0x03, 0xfb, 0x1f, 0xb0, 0xf8, 0x52, 0x92, 0x9d, 0x97, 0xcb, 0xc5, 0x3e,
	0xc3, 0x9e, 0x45, 0xc3, 0x1e, 0x36, 0x6f, 0xb3, 0x0d, 0xdc, 0x7f, 0xd2, 0xd5, 0x6d, 0xe3, 0x0c,
	0x76, 0xd8, 0x09, 0x0c, 0xd6, 0x2f, 0x8a, 0x3d, 0xba, 0xf6, 0xa1, 0xf9, 0x47, 0x57, 0x70, 0xcf,
	0xca, 0x75, 0x0d, 0x76, 0xd8, 0x7b, 0x18, 0x36, 0x16, 0x93, 0x6d, 0x3d, 0xb4, 0x8e, 0x8d, 0xfd,
	0xb7, 0xd3, 0xf9, 0x2d, 0x53, 0x79, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x77, 0x8a, 0xdc, 0x45,
	0x51, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyValueStoreClient is the client API for KeyValueStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyValueStoreClient interface {
	GetState(ctx context.Context, in *GetStateEnvelope, opts ...grpc.CallOption) (*GetStateResponseEnvelope, error)
	SaveState(ctx context.Context, in *SaveStateEnvelope, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteState(ctx context.Context, in *DeleteStateEnvelope, opts ...grpc.CallOption) (*empty.Empty, error)
}

type keyValueStoreClient struct {
	cc *grpc.ClientConn
}

func NewKeyValueStoreClient(cc *grpc.ClientConn) KeyValueStoreClient {
	return &keyValueStoreClient{cc}
}

func (c *keyValueStoreClient) GetState(ctx context.Context, in *GetStateEnvelope, opts ...grpc.CallOption) (*GetStateResponseEnvelope, error) {
	out := new(GetStateResponseEnvelope)
	err := c.cc.Invoke(ctx, "/cloudstate.KeyValueStore/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreClient) SaveState(ctx context.Context, in *SaveStateEnvelope, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cloudstate.KeyValueStore/SaveState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreClient) DeleteState(ctx context.Context, in *DeleteStateEnvelope, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cloudstate.KeyValueStore/DeleteState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueStoreServer is the server API for KeyValueStore service.
type KeyValueStoreServer interface {
	GetState(context.Context, *GetStateEnvelope) (*GetStateResponseEnvelope, error)
	SaveState(context.Context, *SaveStateEnvelope) (*empty.Empty, error)
	DeleteState(context.Context, *DeleteStateEnvelope) (*empty.Empty, error)
}

// UnimplementedKeyValueStoreServer can be embedded to have forward compatible implementations.
type UnimplementedKeyValueStoreServer struct {
}

func (*UnimplementedKeyValueStoreServer) GetState(ctx context.Context, req *GetStateEnvelope) (*GetStateResponseEnvelope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (*UnimplementedKeyValueStoreServer) SaveState(ctx context.Context, req *SaveStateEnvelope) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveState not implemented")
}
func (*UnimplementedKeyValueStoreServer) DeleteState(ctx context.Context, req *DeleteStateEnvelope) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteState not implemented")
}

func RegisterKeyValueStoreServer(s *grpc.Server, srv KeyValueStoreServer) {
	s.RegisterService(&_KeyValueStore_serviceDesc, srv)
}

func _KeyValueStore_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudstate.KeyValueStore/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).GetState(ctx, req.(*GetStateEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStore_SaveState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveStateEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).SaveState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudstate.KeyValueStore/SaveState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).SaveState(ctx, req.(*SaveStateEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStore_DeleteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStateEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).DeleteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudstate.KeyValueStore/DeleteState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).DeleteState(ctx, req.(*DeleteStateEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyValueStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudstate.KeyValueStore",
	HandlerType: (*KeyValueStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _KeyValueStore_GetState_Handler,
		},
		{
			MethodName: "SaveState",
			Handler:    _KeyValueStore_SaveState_Handler,
		},
		{
			MethodName: "DeleteState",
			Handler:    _KeyValueStore_DeleteState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kv_store.proto",
}
