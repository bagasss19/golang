// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The request message containing the user's name.
type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Create
type CreateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{1}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Response             string   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{2}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

// Read
type ReadUserId struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserId) Reset()         { *m = ReadUserId{} }
func (m *ReadUserId) String() string { return proto.CompactTextString(m) }
func (*ReadUserId) ProtoMessage()    {}
func (*ReadUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{3}
}
func (m *ReadUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserId.Unmarshal(m, b)
}
func (m *ReadUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserId.Marshal(b, m, deterministic)
}
func (dst *ReadUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserId.Merge(dst, src)
}
func (m *ReadUserId) XXX_Size() int {
	return xxx_messageInfo_ReadUserId.Size(m)
}
func (m *ReadUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserId.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserId proto.InternalMessageInfo

func (m *ReadUserId) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadUserId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponseId struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponseId) Reset()         { *m = ReadResponseId{} }
func (m *ReadResponseId) String() string { return proto.CompactTextString(m) }
func (*ReadResponseId) ProtoMessage()    {}
func (*ReadResponseId) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{4}
}
func (m *ReadResponseId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponseId.Unmarshal(m, b)
}
func (m *ReadResponseId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponseId.Marshal(b, m, deterministic)
}
func (dst *ReadResponseId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponseId.Merge(dst, src)
}
func (m *ReadResponseId) XXX_Size() int {
	return xxx_messageInfo_ReadResponseId.Size(m)
}
func (m *ReadResponseId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponseId.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponseId proto.InternalMessageInfo

func (m *ReadResponseId) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponseId) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Update
type UpdateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{5}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed updated
	// Equals 1 in case of succesfull update
	Updated              int64    `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{6}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

// Delete
type DeleteRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{7}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResponse struct {
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed deleted
	// Equals 1 in case of succesfull delete
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{8}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

// Request data to read all todo task
type ReadAllRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{9}
}
func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (dst *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(dst, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

// Contains list of all todo tasks
type ReadAllResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 []*User  `protobuf:"bytes,2,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ee8be50d839374b7, []int{10}
}
func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (dst *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(dst, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*CreateRequest)(nil), "user.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "user.CreateResponse")
	proto.RegisterType((*ReadUserId)(nil), "user.ReadUserId")
	proto.RegisterType((*ReadResponseId)(nil), "user.ReadResponseId")
	proto.RegisterType((*UpdateRequest)(nil), "user.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "user.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "user.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "user.DeleteResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "user.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "user.ReadAllResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// Create new
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read
	Read(ctx context.Context, in *ReadUserId, opts ...grpc.CallOption) (*ReadResponseId, error)
	// Update
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Read all
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *ReadUserId, opts ...grpc.CallOption) (*ReadResponseId, error) {
	out := new(ReadResponseId)
	err := c.cc.Invoke(ctx, "/user.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// Create new
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read
	Read(context.Context, *ReadUserId) (*ReadResponseId, error)
	// Update
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Read all
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*ReadUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _UserService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_ee8be50d839374b7) }

var fileDescriptor_user_ee8be50d839374b7 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x4e, 0xab, 0x40,
	0x14, 0x4d, 0xa1, 0x8f, 0xf6, 0x9d, 0xa6, 0xbc, 0x66, 0x5e, 0x4d, 0x08, 0x0b, 0x63, 0x66, 0xe5,
	0x8a, 0xc4, 0x92, 0xb8, 0x6a, 0x4c, 0x6a, 0xdd, 0xb8, 0x45, 0xfb, 0x01, 0x28, 0x77, 0x41, 0x42,
	0x0b, 0xce, 0xb4, 0xfa, 0x23, 0x7e, 0xb0, 0x99, 0x19, 0x86, 0x0a, 0xa1, 0x75, 0xe1, 0x8e, 0x33,
	0x73, 0xcf, 0x3d, 0x67, 0xce, 0xbd, 0x00, 0x07, 0x49, 0x22, 0xaa, 0x44, 0xb9, 0x2f, 0xd9, 0x50,
	0x7d, 0xf3, 0x67, 0x0c, 0x37, 0x92, 0x04, 0x0b, 0x31, 0x56, 0x78, 0x97, 0x6e, 0x29, 0x18, 0x5c,
	0x0d, 0xae, 0xff, 0x26, 0x0d, 0x66, 0x73, 0xfc, 0xa1, 0x6d, 0x9a, 0x17, 0x81, 0xa3, 0x2f, 0x0c,
	0x50, 0x8c, 0x2a, 0x95, 0xf2, 0xa3, 0x14, 0x59, 0xe0, 0x1a, 0x86, 0xc5, 0x7c, 0x85, 0xe9, 0x5a,
	0x50, 0xba, 0xa7, 0x84, 0xde, 0x0e, 0x24, 0xf7, 0x6c, 0x06, 0x37, 0xad, 0xf2, 0xba, 0xb3, 0xfa,
	0x64, 0x97, 0xd0, 0x06, 0x74, 0xcf, 0xc9, 0x02, 0x91, 0x76, 0xa6, 0xac, 0x24, 0xc6, 0xd8, 0x1d,
	0x7c, 0xdb, 0x42, 0x56, 0xe5, 0x4e, 0x52, 0x4f, 0x8f, 0x10, 0x63, 0x51, 0xdf, 0xd6, 0xde, 0x1a,
	0xcc, 0x23, 0x20, 0xa1, 0x34, 0x53, 0x1d, 0x1f, 0xb3, 0x1e, 0xae, 0x0f, 0x27, 0xcf, 0x34, 0xcb,
	0x4d, 0x9c, 0x3c, 0xe3, 0xf7, 0xf0, 0x55, 0xbd, 0x55, 0xeb, 0xe5, 0xfc, 0xe4, 0x79, 0x85, 0xe9,
	0xa6, 0xca, 0x7e, 0xf5, 0xec, 0x25, 0x7c, 0xdb, 0xe2, 0xe4, 0xb3, 0x03, 0x8c, 0x0e, 0xba, 0xc6,
	0xfa, 0xb7, 0x90, 0xdf, 0x60, 0xfa, 0x40, 0x05, 0x9d, 0x33, 0xd0, 0x7d, 0xf7, 0x12, 0xbe, 0xa5,
	0x9c, 0x13, 0xcc, 0x74, 0x4d, 0x23, 0x58, 0x43, 0xce, 0x4d, 0x6a, 0xab, 0xa2, 0x38, 0xa9, 0xc8,
	0xd7, 0xf8, 0xd7, 0xd4, 0x9c, 0x94, 0x38, 0xe6, 0xe2, 0xf6, 0xe5, 0xb2, 0xf8, 0x74, 0x30, 0x51,
	0xf0, 0x89, 0xc4, 0x7b, 0xfe, 0x4a, 0x2c, 0x86, 0x67, 0xd6, 0x83, 0xfd, 0x37, 0xb5, 0xad, 0x7d,
	0x0b, 0xe7, 0xed, 0xc3, 0x5a, 0x36, 0xc2, 0x50, 0x39, 0x61, 0x33, 0x73, 0x7b, 0xdc, 0x0f, 0x5b,
	0xdf, 0xd9, 0x80, 0x18, 0x9e, 0x19, 0x86, 0x15, 0x69, 0x4d, 0xd7, 0x92, 0x3a, 0xf3, 0x8a, 0xe1,
	0x99, 0x40, 0x2d, 0xa9, 0x35, 0x11, 0x4b, 0xea, 0x64, 0x7e, 0x8b, 0x51, 0x9d, 0x11, 0xfb, 0x66,
	0xe5, 0x18, 0x6b, 0x78, 0xd1, 0x39, 0x35, 0xbc, 0x17, 0x4f, 0xff, 0xcb, 0xf1, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x67, 0x7a, 0xe6, 0x0d, 0xd9, 0x03, 0x00, 0x00,
}