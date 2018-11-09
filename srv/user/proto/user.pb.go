// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/user/proto/user.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_81056cd21ad18ff4, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_81056cd21ad18ff4, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LoginResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserInfoRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRequest) Reset()         { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()    {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_81056cd21ad18ff4, []int{2}
}
func (m *UserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRequest.Unmarshal(m, b)
}
func (m *UserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRequest.Marshal(b, m, deterministic)
}
func (dst *UserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRequest.Merge(dst, src)
}
func (m *UserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoRequest.Size(m)
}
func (m *UserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRequest proto.InternalMessageInfo

func (m *UserInfoRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserInfo struct {
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_81056cd21ad18ff4, []int{3}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UserInfoResponse struct {
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_81056cd21ad18ff4, []int{4}
}
func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (dst *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(dst, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfoResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfoResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*UserInfoRequest)(nil), "user.UserInfoRequest")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*UserInfoResponse)(nil), "user.UserInfoResponse")
}

func init() { proto.RegisterFile("srv/user/proto/user.proto", fileDescriptor_user_81056cd21ad18ff4) }

var fileDescriptor_user_81056cd21ad18ff4 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x25, 0xa1, 0xa9, 0xca, 0x95, 0x2f, 0x1d, 0x1f, 0x0a, 0x99, 0x8a, 0x27, 0xa6, 0x56, 0x2a,
	0x23, 0x3b, 0x12, 0x12, 0x0b, 0x41, 0xfc, 0x80, 0xd0, 0x1c, 0xc8, 0x12, 0xb5, 0x83, 0x2f, 0x0d,
	0x1b, 0x3f, 0x80, 0x5f, 0x8d, 0x6c, 0x63, 0x88, 0xb3, 0x30, 0xc0, 0x76, 0xf7, 0xde, 0xf9, 0xdd,
	0xc7, 0x33, 0x9c, 0xb1, 0xe9, 0x16, 0x1b, 0x26, 0xb3, 0x68, 0x8c, 0x6e, 0xb5, 0x0b, 0xe7, 0x2e,
	0xc4, 0x91, 0x8d, 0xc5, 0x35, 0xec, 0xde, 0xea, 0x67, 0xa9, 0x4a, 0x7a, 0xdd, 0x10, 0xb7, 0x58,
	0xc0, 0xc4, 0xe2, 0xaa, 0x5a, 0x53, 0x9e, 0xcc, 0x92, 0x8b, 0x9d, 0xf2, 0x3b, 0xb7, 0x5c, 0x53,
	0x31, 0xbf, 0x69, 0x53, 0xe7, 0xa9, 0xe7, 0x42, 0x2e, 0xee, 0x60, 0xef, 0x4b, 0x87, 0x1b, 0xad,
	0x98, 0x70, 0x1f, 0x52, 0xe9, 0xcb, 0xb2, 0x32, 0x95, 0x75, 0x24, 0xbc, 0x3d, 0x10, 0x3e, 0x86,
	0x8c, 0xd6, 0x95, 0x7c, 0xc9, 0x47, 0x8e, 0xf0, 0x89, 0x38, 0x87, 0x83, 0x07, 0x26, 0x73, 0xa3,
	0x9e, 0x74, 0x98, 0xce, 0x8b, 0x26, 0x41, 0x54, 0x7c, 0x24, 0x30, 0x09, 0x35, 0x7f, 0xef, 0x18,
	0x2d, 0x98, 0xc5, 0x0b, 0xe2, 0x0c, 0xa6, 0x35, 0xf1, 0xca, 0xc8, 0xa6, 0x95, 0x5a, 0xe5, 0x63,
	0x47, 0xf7, 0x21, 0xd1, 0xc1, 0xe1, 0xcf, 0xbc, 0xff, 0x75, 0x85, 0xdf, 0xfb, 0x2e, 0xdf, 0x61,
	0x6a, 0xfb, 0xde, 0x93, 0xe9, 0xe4, 0x8a, 0x70, 0x09, 0x99, 0x73, 0x02, 0x71, 0xee, 0xdc, 0xee,
	0xdb, 0x5b, 0x1c, 0x45, 0x98, 0x1f, 0x52, 0x6c, 0xe1, 0x55, 0xef, 0x8c, 0x27, 0xbe, 0x64, 0x70,
	0xfa, 0xe2, 0x74, 0x08, 0x87, 0xc7, 0x8f, 0x63, 0xf7, 0x9f, 0x2e, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0x59, 0xda, 0x09, 0x6c, 0x02, 0x00, 0x00,
}