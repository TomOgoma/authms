// Code generated by protoc-gen-go.
// source: authms.proto
// DO NOT EDIT!

/*
Package authms is a generated protocol buffer package.

It is generated from these files:
	authms.proto

It has these top-level messages:
	History
	OAuth
	Value
	User
	BasicAuthRequest
	OAuthRequest
	TokenRequest
	Response
*/
package authms

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type History struct {
	Id            int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserID        int64  `protobuf:"varint,2,opt,name=userID" json:"userID,omitempty"`
	IpAddress     string `protobuf:"bytes,3,opt,name=ipAddress" json:"ipAddress,omitempty"`
	Date          string `protobuf:"bytes,4,opt,name=date" json:"date,omitempty"`
	AccessType    string `protobuf:"bytes,5,opt,name=accessType" json:"accessType,omitempty"`
	SuccessStatus bool   `protobuf:"varint,6,opt,name=successStatus" json:"successStatus,omitempty"`
}

func (m *History) Reset()                    { *m = History{} }
func (m *History) String() string            { return proto.CompactTextString(m) }
func (*History) ProtoMessage()               {}
func (*History) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type OAuth struct {
	AppName   string `protobuf:"bytes,1,opt,name=appName" json:"appName,omitempty"`
	AppUserID string `protobuf:"bytes,2,opt,name=appUserID" json:"appUserID,omitempty"`
	AppToken  string `protobuf:"bytes,3,opt,name=appToken" json:"appToken,omitempty"`
}

func (m *OAuth) Reset()                    { *m = OAuth{} }
func (m *OAuth) String() string            { return proto.CompactTextString(m) }
func (*OAuth) ProtoMessage()               {}
func (*OAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Value struct {
	Value    string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Verified bool   `protobuf:"varint,2,opt,name=verified" json:"verified,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type User struct {
	Id           int64      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Token        string     `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Password     string     `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Username     string     `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	LoginHistory []*History `protobuf:"bytes,5,rep,name=loginHistory" json:"loginHistory,omitempty"`
	Phone        *Value     `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
	Mail         *Value     `protobuf:"bytes,7,opt,name=mail" json:"mail,omitempty"`
	OAuth        *OAuth     `protobuf:"bytes,8,opt,name=oAuth" json:"oAuth,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *User) GetLoginHistory() []*History {
	if m != nil {
		return m.LoginHistory
	}
	return nil
}

func (m *User) GetPhone() *Value {
	if m != nil {
		return m.Phone
	}
	return nil
}

func (m *User) GetMail() *Value {
	if m != nil {
		return m.Mail
	}
	return nil
}

func (m *User) GetOAuth() *OAuth {
	if m != nil {
		return m.OAuth
	}
	return nil
}

type BasicAuthRequest struct {
	DeviceID         string `protobuf:"bytes,1,opt,name=deviceID" json:"deviceID,omitempty"`
	ForServiceID     string `protobuf:"bytes,2,opt,name=forServiceID" json:"forServiceID,omitempty"`
	BasicID          string `protobuf:"bytes,3,opt,name=basicID" json:"basicID,omitempty"`
	Password         string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	RefererServiceID string `protobuf:"bytes,5,opt,name=refererServiceID" json:"refererServiceID,omitempty"`
}

func (m *BasicAuthRequest) Reset()                    { *m = BasicAuthRequest{} }
func (m *BasicAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*BasicAuthRequest) ProtoMessage()               {}
func (*BasicAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type OAuthRequest struct {
	DeviceID         string `protobuf:"bytes,1,opt,name=deviceID" json:"deviceID,omitempty"`
	ForServiceID     string `protobuf:"bytes,2,opt,name=forServiceID" json:"forServiceID,omitempty"`
	AppName          string `protobuf:"bytes,3,opt,name=appName" json:"appName,omitempty"`
	AppUserID        string `protobuf:"bytes,4,opt,name=appUserID" json:"appUserID,omitempty"`
	AppToken         string `protobuf:"bytes,5,opt,name=appToken" json:"appToken,omitempty"`
	RefererServiceID string `protobuf:"bytes,6,opt,name=refererServiceID" json:"refererServiceID,omitempty"`
}

func (m *OAuthRequest) Reset()                    { *m = OAuthRequest{} }
func (m *OAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*OAuthRequest) ProtoMessage()               {}
func (*OAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type TokenRequest struct {
	DeviceID         string `protobuf:"bytes,1,opt,name=deviceID" json:"deviceID,omitempty"`
	ForServiceID     string `protobuf:"bytes,2,opt,name=forServiceID" json:"forServiceID,omitempty"`
	Token            string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	RefererServiceID string `protobuf:"bytes,4,opt,name=refererServiceID" json:"refererServiceID,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Response struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	User   *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	Detail string `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*History)(nil), "History")
	proto.RegisterType((*OAuth)(nil), "OAuth")
	proto.RegisterType((*Value)(nil), "Value")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*BasicAuthRequest)(nil), "BasicAuthRequest")
	proto.RegisterType((*OAuthRequest)(nil), "OAuthRequest")
	proto.RegisterType((*TokenRequest)(nil), "TokenRequest")
	proto.RegisterType((*Response)(nil), "Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthMS service

type AuthMSClient interface {
	RegisterUserName(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error)
	RegisterEmail(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error)
	RegisterPhone(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error)
	RegisterOAuth(ctx context.Context, in *OAuthRequest, opts ...client.CallOption) (*Response, error)
	LoginUserName(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error)
	LoginEmail(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error)
	LoginPhone(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error)
	LoginOAuth(ctx context.Context, in *OAuthRequest, opts ...client.CallOption) (*Response, error)
	ValidateToken(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*Response, error)
}

type authMSClient struct {
	c           client.Client
	serviceName string
}

func NewAuthMSClient(serviceName string, c client.Client) AuthMSClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "authms"
	}
	return &authMSClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authMSClient) RegisterUserName(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.RegisterUserName", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMSClient) RegisterEmail(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.RegisterEmail", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMSClient) RegisterPhone(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.RegisterPhone", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMSClient) RegisterOAuth(ctx context.Context, in *OAuthRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.RegisterOAuth", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMSClient) LoginUserName(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.LoginUserName", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMSClient) LoginEmail(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.LoginEmail", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMSClient) LoginPhone(ctx context.Context, in *BasicAuthRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.LoginPhone", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMSClient) LoginOAuth(ctx context.Context, in *OAuthRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.LoginOAuth", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMSClient) ValidateToken(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.ValidateToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthMS service

type AuthMSHandler interface {
	RegisterUserName(context.Context, *BasicAuthRequest, *Response) error
	RegisterEmail(context.Context, *BasicAuthRequest, *Response) error
	RegisterPhone(context.Context, *BasicAuthRequest, *Response) error
	RegisterOAuth(context.Context, *OAuthRequest, *Response) error
	LoginUserName(context.Context, *BasicAuthRequest, *Response) error
	LoginEmail(context.Context, *BasicAuthRequest, *Response) error
	LoginPhone(context.Context, *BasicAuthRequest, *Response) error
	LoginOAuth(context.Context, *OAuthRequest, *Response) error
	ValidateToken(context.Context, *TokenRequest, *Response) error
}

func RegisterAuthMSHandler(s server.Server, hdlr AuthMSHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthMS{hdlr}, opts...))
}

type AuthMS struct {
	AuthMSHandler
}

func (h *AuthMS) RegisterUserName(ctx context.Context, in *BasicAuthRequest, out *Response) error {
	return h.AuthMSHandler.RegisterUserName(ctx, in, out)
}

func (h *AuthMS) RegisterEmail(ctx context.Context, in *BasicAuthRequest, out *Response) error {
	return h.AuthMSHandler.RegisterEmail(ctx, in, out)
}

func (h *AuthMS) RegisterPhone(ctx context.Context, in *BasicAuthRequest, out *Response) error {
	return h.AuthMSHandler.RegisterPhone(ctx, in, out)
}

func (h *AuthMS) RegisterOAuth(ctx context.Context, in *OAuthRequest, out *Response) error {
	return h.AuthMSHandler.RegisterOAuth(ctx, in, out)
}

func (h *AuthMS) LoginUserName(ctx context.Context, in *BasicAuthRequest, out *Response) error {
	return h.AuthMSHandler.LoginUserName(ctx, in, out)
}

func (h *AuthMS) LoginEmail(ctx context.Context, in *BasicAuthRequest, out *Response) error {
	return h.AuthMSHandler.LoginEmail(ctx, in, out)
}

func (h *AuthMS) LoginPhone(ctx context.Context, in *BasicAuthRequest, out *Response) error {
	return h.AuthMSHandler.LoginPhone(ctx, in, out)
}

func (h *AuthMS) LoginOAuth(ctx context.Context, in *OAuthRequest, out *Response) error {
	return h.AuthMSHandler.LoginOAuth(ctx, in, out)
}

func (h *AuthMS) ValidateToken(ctx context.Context, in *TokenRequest, out *Response) error {
	return h.AuthMSHandler.ValidateToken(ctx, in, out)
}

func init() { proto.RegisterFile("authms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x5c, 0xff, 0xd4, 0xbe, 0x4d, 0x3e, 0x95, 0x51, 0x85, 0x4c, 0x55, 0xa1, 0xc8, 0x62,
	0x11, 0xd1, 0xca, 0x48, 0x61, 0xc5, 0xb2, 0xa8, 0x48, 0x54, 0x02, 0x82, 0x26, 0xa5, 0x1b, 0x56,
	0xd3, 0xf8, 0xa6, 0x19, 0x91, 0xc4, 0x83, 0xc7, 0x0e, 0xea, 0x5b, 0xf0, 0x1a, 0xac, 0xd8, 0xf0,
	0x1c, 0x3c, 0x08, 0x4f, 0x81, 0xe6, 0xda, 0x4e, 0xe2, 0xfe, 0xa4, 0x41, 0xea, 0xce, 0xe7, 0x5c,
	0xdf, 0x3b, 0x67, 0x8e, 0xcf, 0x8c, 0xa1, 0x25, 0x8a, 0x7c, 0x3c, 0xd5, 0xb1, 0xca, 0xd2, 0x3c,
	0x8d, 0x7e, 0x58, 0xb0, 0xfd, 0x56, 0xea, 0x3c, 0xcd, 0xae, 0xd8, 0xff, 0xb0, 0x25, 0x93, 0xd0,
	0xea, 0x58, 0x5d, 0x9b, 0x6f, 0xc9, 0x84, 0x3d, 0x06, 0xaf, 0xd0, 0x98, 0x9d, 0x9e, 0x84, 0x5b,
	0xc4, 0x55, 0x88, 0x1d, 0x40, 0x20, 0xd5, 0x71, 0x92, 0x64, 0xa8, 0x75, 0x68, 0x77, 0xac, 0x6e,
	0xc0, 0x97, 0x04, 0x63, 0xe0, 0x24, 0x22, 0xc7, 0xd0, 0xa1, 0x02, 0x3d, 0xb3, 0xa7, 0x00, 0x62,
	0x38, 0x44, 0xad, 0xcf, 0xae, 0x14, 0x86, 0x2e, 0x55, 0x56, 0x18, 0xf6, 0x0c, 0xda, 0xba, 0x20,
	0x38, 0xc8, 0x45, 0x5e, 0xe8, 0xd0, 0xeb, 0x58, 0x5d, 0x9f, 0x37, 0xc9, 0xe8, 0x33, 0xb8, 0xfd,
	0xe3, 0x22, 0x1f, 0xb3, 0x10, 0xb6, 0x85, 0x52, 0x1f, 0xc4, 0x14, 0x49, 0x6d, 0xc0, 0x6b, 0x68,
	0xa4, 0x09, 0xa5, 0x3e, 0x2d, 0x55, 0x07, 0x7c, 0x49, 0xb0, 0x7d, 0xf0, 0x85, 0x52, 0x67, 0xe9,
	0x17, 0x9c, 0x55, 0xba, 0x17, 0x38, 0x7a, 0x05, 0xee, 0xb9, 0x98, 0x14, 0xc8, 0xf6, 0xc0, 0x9d,
	0x9b, 0x87, 0x6a, 0x74, 0x09, 0x4c, 0xeb, 0x1c, 0x33, 0x39, 0x92, 0x98, 0xd0, 0x5c, 0x9f, 0x2f,
	0x70, 0xf4, 0xc7, 0x02, 0xc7, 0xac, 0x70, 0xc3, 0xc0, 0x3d, 0x70, 0x73, 0x5a, 0xac, 0x54, 0x52,
	0x02, 0x33, 0x4a, 0x09, 0xad, 0xbf, 0xa5, 0x59, 0x52, 0xab, 0xa8, 0xb1, 0xa9, 0x19, 0x93, 0x67,
	0x66, 0x6b, 0xa5, 0x81, 0x0b, 0xcc, 0x8e, 0xa0, 0x35, 0x49, 0x2f, 0xe5, 0xac, 0xfa, 0x5c, 0xa1,
	0xdb, 0xb1, 0xbb, 0x3b, 0x3d, 0x3f, 0xae, 0x30, 0x6f, 0x54, 0xd9, 0x01, 0xb8, 0x6a, 0x9c, 0xce,
	0x90, 0xac, 0xdc, 0xe9, 0x79, 0x31, 0xed, 0x8e, 0x97, 0x24, 0xdb, 0x07, 0x67, 0x2a, 0xe4, 0x24,
	0xdc, 0x6e, 0x14, 0x89, 0x33, 0x9d, 0xa9, 0xb1, 0x39, 0xf4, 0xab, 0x22, 0x99, 0xce, 0x4b, 0x32,
	0xfa, 0x65, 0xc1, 0xee, 0x6b, 0xa1, 0xe5, 0x90, 0x48, 0xfc, 0x5a, 0xa0, 0xce, 0x8d, 0xec, 0x04,
	0xe7, 0x72, 0x88, 0xa7, 0x27, 0x95, 0x6d, 0x0b, 0xcc, 0x22, 0x68, 0x8d, 0xd2, 0x6c, 0x80, 0x59,
	0x55, 0x2f, 0xbd, 0x68, 0x70, 0xe6, 0x83, 0x5e, 0x98, 0x99, 0xa7, 0x27, 0x95, 0x23, 0x35, 0x6c,
	0x98, 0xe5, 0x5c, 0x33, 0xeb, 0x39, 0xec, 0x66, 0x38, 0xc2, 0x0c, 0x57, 0xa6, 0x97, 0xd9, 0xba,
	0xc1, 0x47, 0xbf, 0x2d, 0x68, 0xf5, 0x1f, 0x58, 0x72, 0x9d, 0x41, 0x7b, 0x4d, 0x06, 0x9d, 0x75,
	0x19, 0x74, 0x9b, 0x19, 0xbc, 0x75, 0x43, 0xde, 0x1d, 0x1b, 0xfa, 0x6e, 0x41, 0x8b, 0xba, 0x1e,
	0x6a, 0x43, 0x8b, 0xb0, 0xda, 0xab, 0x61, 0xbd, 0x4d, 0x92, 0x73, 0x87, 0x24, 0x01, 0x3e, 0x47,
	0xad, 0xd2, 0x99, 0xc6, 0x95, 0xa3, 0x10, 0xd0, 0x51, 0x60, 0xe0, 0x0c, 0xd3, 0x04, 0x69, 0x65,
	0x97, 0xd3, 0x33, 0x7b, 0x02, 0x8e, 0x09, 0x37, 0x2d, 0xb8, 0xd3, 0x73, 0x63, 0xe3, 0x10, 0x27,
	0xca, 0x5c, 0x3d, 0x09, 0xe6, 0x26, 0xa1, 0xe5, 0x62, 0x15, 0xea, 0xfd, 0xb4, 0xc1, 0x33, 0x5f,
	0xf1, 0xfd, 0x80, 0xf5, 0x60, 0x97, 0xe3, 0xa5, 0xd4, 0x39, 0x66, 0xa6, 0x91, 0xac, 0x7f, 0x14,
	0x5f, 0x8f, 0xe6, 0x7e, 0x10, 0xd7, 0x9a, 0xa2, 0xff, 0xd8, 0x0b, 0x68, 0xd7, 0x3d, 0x6f, 0x28,
	0xeb, 0xff, 0xd0, 0xf0, 0x91, 0x0e, 0xce, 0x7d, 0x0d, 0x87, 0xcb, 0x86, 0xf2, 0xae, 0x6a, 0xc7,
	0xfd, 0x75, 0xd3, 0xdf, 0x99, 0x33, 0xbb, 0xb1, 0xfe, 0x23, 0x00, 0x6a, 0xd8, 0x4c, 0x7c, 0xfd,
	0xf6, 0x66, 0xca, 0xbb, 0xd5, 0xdb, 0xf7, 0xcb, 0x3e, 0x84, 0xf6, 0xb9, 0x98, 0x48, 0x73, 0xb3,
	0x97, 0xb9, 0x6d, 0xc7, 0xab, 0x49, 0x6c, 0xbc, 0x7c, 0xe1, 0xd1, 0x7f, 0xe6, 0xe5, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x03, 0x4b, 0x94, 0x10, 0x77, 0x06, 0x00, 0x00,
}
