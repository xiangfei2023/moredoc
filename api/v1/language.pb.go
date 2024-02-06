// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/language.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Language struct {
	Id        int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Language  string     `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Enable    bool       `protobuf:"varint,3,opt,name=enable,proto3" json:"enable,omitempty"`
	Code      string     `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Total     int32      `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	Sort      int32      `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`
	CreatedAt *time.Time `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at,omitempty"`
}

func (m *Language) Reset()         { *m = Language{} }
func (m *Language) String() string { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()    {}
func (*Language) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7f0bb6aa63037c, []int{0}
}
func (m *Language) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Language) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Language.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Language) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Language.Merge(m, src)
}
func (m *Language) XXX_Size() int {
	return m.Size()
}
func (m *Language) XXX_DiscardUnknown() {
	xxx_messageInfo_Language.DiscardUnknown(m)
}

var xxx_messageInfo_Language proto.InternalMessageInfo

func (m *Language) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Language) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Language) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Language) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Language) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Language) GetSort() int32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *Language) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Language) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateLanguageStatusRequest struct {
	Id     []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Enable bool    `protobuf:"varint,2,opt,name=enable,proto3" json:"enable,omitempty"`
}

func (m *UpdateLanguageStatusRequest) Reset()         { *m = UpdateLanguageStatusRequest{} }
func (m *UpdateLanguageStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateLanguageStatusRequest) ProtoMessage()    {}
func (*UpdateLanguageStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7f0bb6aa63037c, []int{1}
}
func (m *UpdateLanguageStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateLanguageStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateLanguageStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateLanguageStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateLanguageStatusRequest.Merge(m, src)
}
func (m *UpdateLanguageStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdateLanguageStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateLanguageStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateLanguageStatusRequest proto.InternalMessageInfo

func (m *UpdateLanguageStatusRequest) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UpdateLanguageStatusRequest) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

type ListLanguageRequest struct {
	Wd     string `protobuf:"bytes,1,opt,name=wd,proto3" json:"wd,omitempty"`
	Enable []bool `protobuf:"varint,2,rep,packed,name=enable,proto3" json:"enable,omitempty"`
}

func (m *ListLanguageRequest) Reset()         { *m = ListLanguageRequest{} }
func (m *ListLanguageRequest) String() string { return proto.CompactTextString(m) }
func (*ListLanguageRequest) ProtoMessage()    {}
func (*ListLanguageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7f0bb6aa63037c, []int{2}
}
func (m *ListLanguageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListLanguageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListLanguageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListLanguageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLanguageRequest.Merge(m, src)
}
func (m *ListLanguageRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListLanguageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLanguageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListLanguageRequest proto.InternalMessageInfo

func (m *ListLanguageRequest) GetWd() string {
	if m != nil {
		return m.Wd
	}
	return ""
}

func (m *ListLanguageRequest) GetEnable() []bool {
	if m != nil {
		return m.Enable
	}
	return nil
}

type ListLanguageReply struct {
	Total    int64       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Language []*Language `protobuf:"bytes,2,rep,name=language,proto3" json:"language,omitempty"`
}

func (m *ListLanguageReply) Reset()         { *m = ListLanguageReply{} }
func (m *ListLanguageReply) String() string { return proto.CompactTextString(m) }
func (*ListLanguageReply) ProtoMessage()    {}
func (*ListLanguageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7f0bb6aa63037c, []int{3}
}
func (m *ListLanguageReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListLanguageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListLanguageReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListLanguageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLanguageReply.Merge(m, src)
}
func (m *ListLanguageReply) XXX_Size() int {
	return m.Size()
}
func (m *ListLanguageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLanguageReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListLanguageReply proto.InternalMessageInfo

func (m *ListLanguageReply) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListLanguageReply) GetLanguage() []*Language {
	if m != nil {
		return m.Language
	}
	return nil
}

func init() {
	proto.RegisterType((*Language)(nil), "api.v1.Language")
	proto.RegisterType((*UpdateLanguageStatusRequest)(nil), "api.v1.UpdateLanguageStatusRequest")
	proto.RegisterType((*ListLanguageRequest)(nil), "api.v1.ListLanguageRequest")
	proto.RegisterType((*ListLanguageReply)(nil), "api.v1.ListLanguageReply")
}

func init() { proto.RegisterFile("api/v1/language.proto", fileDescriptor_4b7f0bb6aa63037c) }

var fileDescriptor_4b7f0bb6aa63037c = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0x49, 0x5b, 0xdb, 0xa9, 0x88, 0x3b, 0x76, 0x77, 0x63, 0xba, 0x66, 0x43, 0xbc,
	0x04, 0x91, 0x84, 0x56, 0xbc, 0x28, 0x22, 0x5d, 0xd8, 0x83, 0xb0, 0x87, 0x25, 0x2a, 0x82, 0x17,
	0x99, 0x36, 0x63, 0x08, 0xa4, 0x99, 0xd8, 0xbc, 0xb4, 0xf4, 0xea, 0x07, 0x90, 0x05, 0xbf, 0x94,
	0xc7, 0x05, 0x2f, 0xde, 0x94, 0xd6, 0xab, 0xdf, 0x41, 0x3a, 0x99, 0xc9, 0xb6, 0xdb, 0x45, 0xbc,
	0xbd, 0xf7, 0xe6, 0xfd, 0xff, 0x99, 0xf7, 0x9b, 0x17, 0xbc, 0x4f, 0xb3, 0xd8, 0x9f, 0xf5, 0xfd,
	0x84, 0xa6, 0x51, 0x41, 0x23, 0xe6, 0x65, 0x53, 0x0e, 0x9c, 0x34, 0x69, 0x16, 0x7b, 0xb3, 0xbe,
	0x79, 0x1c, 0x71, 0x1e, 0x25, 0xcc, 0x17, 0xd5, 0x51, 0xf1, 0xd1, 0x87, 0x78, 0xc2, 0x72, 0xa0,
	0x93, 0xac, 0x6c, 0x34, 0xbb, 0x11, 0x8f, 0xb8, 0x08, 0xfd, 0x75, 0x24, 0xab, 0x47, 0x52, 0xb6,
	0x36, 0xa7, 0x69, 0xca, 0x81, 0x42, 0xcc, 0xd3, 0x5c, 0x9e, 0xf6, 0xae, 0x9b, 0xb2, 0x49, 0x06,
	0x8b, 0xf2, 0xd0, 0xf9, 0xa2, 0xe1, 0xd6, 0x99, 0xbc, 0x0c, 0xb9, 0x83, 0xb5, 0x38, 0x34, 0x90,
	0x8d, 0x5c, 0x3d, 0xd0, 0xe2, 0x90, 0x98, 0xb8, 0xa5, 0x2e, 0x6a, 0x68, 0x36, 0x72, 0xdb, 0x41,
	0x95, 0x93, 0x03, 0xdc, 0x64, 0x29, 0x1d, 0x25, 0xcc, 0xd0, 0x6d, 0xe4, 0xb6, 0x02, 0x99, 0x11,
	0x82, 0xeb, 0x63, 0x1e, 0x32, 0xa3, 0x2e, 0xfa, 0x45, 0x4c, 0xba, 0xb8, 0x01, 0x1c, 0x68, 0x62,
	0x34, 0x6c, 0xe4, 0x36, 0x82, 0x32, 0x59, 0x77, 0xe6, 0x7c, 0x0a, 0x46, 0x53, 0x14, 0x45, 0x4c,
	0x5e, 0x62, 0x3c, 0x9e, 0x32, 0x0a, 0x2c, 0xfc, 0x40, 0xc1, 0xb8, 0x65, 0x23, 0xb7, 0x33, 0x30,
	0xbd, 0x72, 0x00, 0x4f, 0x0d, 0xe0, 0xbd, 0x51, 0x54, 0x4e, 0xea, 0x17, 0x3f, 0x8f, 0x51, 0xd0,
	0x96, 0x9a, 0xa1, 0x30, 0x28, 0xb2, 0x50, 0x19, 0xb4, 0xfe, 0xd7, 0x40, 0x6a, 0x86, 0xe0, 0x9c,
	0xe2, 0xde, 0x5b, 0x91, 0x28, 0x2a, 0xaf, 0x81, 0x42, 0x91, 0x07, 0xec, 0x53, 0xc1, 0x72, 0xa8,
	0x10, 0xe9, 0x12, 0xd1, 0x15, 0x06, 0x6d, 0x13, 0x83, 0xf3, 0x02, 0xdf, 0x3b, 0x8b, 0x73, 0x50,
	0x26, 0x1b, 0xf2, 0x79, 0x49, 0xb8, 0x1d, 0x68, 0xf3, 0x6d, 0xb9, 0xbe, 0x21, 0x7f, 0x87, 0xf7,
	0xb6, 0xe5, 0x59, 0xb2, 0xb8, 0xc2, 0x58, 0xbe, 0x90, 0xc4, 0xf8, 0x78, 0xeb, 0x91, 0x74, 0xb7,
	0x33, 0xb8, 0xeb, 0x95, 0xeb, 0xe4, 0x55, 0xf2, 0xaa, 0x63, 0xf0, 0x07, 0xe1, 0x8e, 0x2a, 0x0f,
	0xcf, 0x5f, 0x91, 0x39, 0xee, 0xde, 0x34, 0x2e, 0x79, 0xa8, 0x3c, 0xfe, 0x01, 0xc3, 0x3c, 0xd8,
	0x01, 0x7b, 0xba, 0x5e, 0x2d, 0xc7, 0xf9, 0xfc, 0xfd, 0xf7, 0x57, 0xed, 0xc8, 0x3c, 0xf4, 0xaf,
	0xad, 0xbb, 0x9f, 0x0b, 0xfd, 0x33, 0xf4, 0x88, 0x30, 0x7c, 0x7b, 0x73, 0x42, 0xd2, 0xab, 0x2e,
	0xbd, 0x8b, 0xcd, 0xbc, 0x7f, 0xf3, 0x61, 0x96, 0x2c, 0x9c, 0x07, 0xe2, 0x5b, 0x87, 0x64, 0x7f,
	0xe7, 0x5b, 0x49, 0x9c, 0xc3, 0xc9, 0xd3, 0x6f, 0x4b, 0x0b, 0x5d, 0x2e, 0x2d, 0xf4, 0x6b, 0x69,
	0xa1, 0x8b, 0x95, 0x55, 0xbb, 0x5c, 0x59, 0xb5, 0x1f, 0x2b, 0xab, 0x86, 0xe5, 0x3f, 0x77, 0x8e,
	0xde, 0xef, 0x4d, 0xf8, 0x94, 0x85, 0x7c, 0x2c, 0x1d, 0x9e, 0xcf, 0xfa, 0xa3, 0xa6, 0x98, 0xe8,
	0xc9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x7e, 0xc6, 0x97, 0xb0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LanguageAPIClient is the client API for LanguageAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LanguageAPIClient interface {
	UpdateLanguageStatus(ctx context.Context, in *UpdateLanguageStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListLanguage(ctx context.Context, in *ListLanguageRequest, opts ...grpc.CallOption) (*ListLanguageReply, error)
}

type languageAPIClient struct {
	cc *grpc.ClientConn
}

func NewLanguageAPIClient(cc *grpc.ClientConn) LanguageAPIClient {
	return &languageAPIClient{cc}
}

func (c *languageAPIClient) UpdateLanguageStatus(ctx context.Context, in *UpdateLanguageStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.LanguageAPI/UpdateLanguageStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageAPIClient) ListLanguage(ctx context.Context, in *ListLanguageRequest, opts ...grpc.CallOption) (*ListLanguageReply, error) {
	out := new(ListLanguageReply)
	err := c.cc.Invoke(ctx, "/api.v1.LanguageAPI/ListLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanguageAPIServer is the server API for LanguageAPI service.
type LanguageAPIServer interface {
	UpdateLanguageStatus(context.Context, *UpdateLanguageStatusRequest) (*emptypb.Empty, error)
	ListLanguage(context.Context, *ListLanguageRequest) (*ListLanguageReply, error)
}

// UnimplementedLanguageAPIServer can be embedded to have forward compatible implementations.
type UnimplementedLanguageAPIServer struct {
}

func (*UnimplementedLanguageAPIServer) UpdateLanguageStatus(ctx context.Context, req *UpdateLanguageStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLanguageStatus not implemented")
}
func (*UnimplementedLanguageAPIServer) ListLanguage(ctx context.Context, req *ListLanguageRequest) (*ListLanguageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLanguage not implemented")
}

func RegisterLanguageAPIServer(s *grpc.Server, srv LanguageAPIServer) {
	s.RegisterService(&_LanguageAPI_serviceDesc, srv)
}

func _LanguageAPI_UpdateLanguageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLanguageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageAPIServer).UpdateLanguageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.LanguageAPI/UpdateLanguageStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageAPIServer).UpdateLanguageStatus(ctx, req.(*UpdateLanguageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageAPI_ListLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageAPIServer).ListLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.LanguageAPI/ListLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageAPIServer).ListLanguage(ctx, req.(*ListLanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LanguageAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.LanguageAPI",
	HandlerType: (*LanguageAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateLanguageStatus",
			Handler:    _LanguageAPI_UpdateLanguageStatus_Handler,
		},
		{
			MethodName: "ListLanguage",
			Handler:    _LanguageAPI_ListLanguage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/language.proto",
}

func (m *Language) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Language) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Language) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.UpdatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintLanguage(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x42
	}
	if m.CreatedAt != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintLanguage(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x3a
	}
	if m.Sort != 0 {
		i = encodeVarintLanguage(dAtA, i, uint64(m.Sort))
		i--
		dAtA[i] = 0x30
	}
	if m.Total != 0 {
		i = encodeVarintLanguage(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintLanguage(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x22
	}
	if m.Enable {
		i--
		if m.Enable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Language) > 0 {
		i -= len(m.Language)
		copy(dAtA[i:], m.Language)
		i = encodeVarintLanguage(dAtA, i, uint64(len(m.Language)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLanguage(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UpdateLanguageStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateLanguageStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateLanguageStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enable {
		i--
		if m.Enable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		dAtA4 := make([]byte, len(m.Id)*10)
		var j3 int
		for _, num1 := range m.Id {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintLanguage(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListLanguageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListLanguageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListLanguageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Enable) > 0 {
		for iNdEx := len(m.Enable) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.Enable[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintLanguage(dAtA, i, uint64(len(m.Enable)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Wd) > 0 {
		i -= len(m.Wd)
		copy(dAtA[i:], m.Wd)
		i = encodeVarintLanguage(dAtA, i, uint64(len(m.Wd)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListLanguageReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListLanguageReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListLanguageReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Language) > 0 {
		for iNdEx := len(m.Language) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Language[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLanguage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Total != 0 {
		i = encodeVarintLanguage(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLanguage(dAtA []byte, offset int, v uint64) int {
	offset -= sovLanguage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Language) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLanguage(uint64(m.Id))
	}
	l = len(m.Language)
	if l > 0 {
		n += 1 + l + sovLanguage(uint64(l))
	}
	if m.Enable {
		n += 2
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovLanguage(uint64(l))
	}
	if m.Total != 0 {
		n += 1 + sovLanguage(uint64(m.Total))
	}
	if m.Sort != 0 {
		n += 1 + sovLanguage(uint64(m.Sort))
	}
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovLanguage(uint64(l))
	}
	if m.UpdatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt)
		n += 1 + l + sovLanguage(uint64(l))
	}
	return n
}

func (m *UpdateLanguageStatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Id) > 0 {
		l = 0
		for _, e := range m.Id {
			l += sovLanguage(uint64(e))
		}
		n += 1 + sovLanguage(uint64(l)) + l
	}
	if m.Enable {
		n += 2
	}
	return n
}

func (m *ListLanguageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Wd)
	if l > 0 {
		n += 1 + l + sovLanguage(uint64(l))
	}
	if len(m.Enable) > 0 {
		n += 1 + sovLanguage(uint64(len(m.Enable))) + len(m.Enable)*1
	}
	return n
}

func (m *ListLanguageReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovLanguage(uint64(m.Total))
	}
	if len(m.Language) > 0 {
		for _, e := range m.Language {
			l = e.Size()
			n += 1 + l + sovLanguage(uint64(l))
		}
	}
	return n
}

func sovLanguage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLanguage(x uint64) (n int) {
	return sovLanguage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Language) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLanguage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Language: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Language: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Language", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLanguage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLanguage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Language = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enable = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLanguage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLanguage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sort", wireType)
			}
			m.Sort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sort |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLanguage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLanguage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLanguage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLanguage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatedAt == nil {
				m.UpdatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLanguage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLanguage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdateLanguageStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLanguage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateLanguageStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateLanguageStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLanguage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Id = append(m.Id, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLanguage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLanguage
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLanguage
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Id) == 0 {
					m.Id = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLanguage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Id = append(m.Id, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLanguage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLanguage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListLanguageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLanguage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListLanguageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListLanguageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLanguage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLanguage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLanguage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Enable = append(m.Enable, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLanguage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLanguage
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLanguage
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.Enable) == 0 {
					m.Enable = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLanguage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Enable = append(m.Enable, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Enable", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLanguage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLanguage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListLanguageReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLanguage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListLanguageReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListLanguageReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Language", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLanguage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLanguage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Language = append(m.Language, &Language{})
			if err := m.Language[len(m.Language)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLanguage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLanguage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLanguage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLanguage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLanguage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLanguage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLanguage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLanguage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLanguage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLanguage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLanguage = fmt.Errorf("proto: unexpected end of group")
)
