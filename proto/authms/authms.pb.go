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
	Register(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
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

func (c *authMSClient) Register(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthMS.Register", in)
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
	Register(context.Context, *User, *Response) error
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

func (h *AuthMS) Register(ctx context.Context, in *User, out *Response) error {
	return h.AuthMSHandler.Register(ctx, in, out)
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
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x89, 0xed, 0xda, 0xd3, 0x04, 0x95, 0x55, 0x85, 0x4c, 0x55, 0x55, 0x91, 0xc5, 0x21,
	0x82, 0xca, 0x48, 0xe1, 0xc4, 0xb1, 0xa8, 0x48, 0x54, 0x02, 0x8a, 0x36, 0xa5, 0x17, 0x4e, 0xdb,
	0x78, 0xd2, 0xac, 0x48, 0xe2, 0xc5, 0x6b, 0x07, 0xf5, 0x1f, 0x38, 0xf0, 0x1b, 0xdc, 0xf9, 0x0e,
	0x3e, 0x84, 0xaf, 0x40, 0x3b, 0x6b, 0x27, 0x71, 0xdb, 0x14, 0x0e, 0xbd, 0xed, 0x7b, 0xb3, 0xb3,
	0x33, 0xfb, 0xf6, 0xcd, 0x42, 0x47, 0x94, 0xc5, 0x64, 0xa6, 0x13, 0x95, 0x67, 0x45, 0x16, 0xff,
	0x74, 0x60, 0xeb, 0xad, 0xd4, 0x45, 0x96, 0x5f, 0xb1, 0x87, 0xd0, 0x92, 0x69, 0xe4, 0xf4, 0x9c,
	0x7e, 0x9b, 0xb7, 0x64, 0xca, 0x1e, 0x83, 0x5f, 0x6a, 0xcc, 0x4f, 0x8e, 0xa3, 0x16, 0x71, 0x15,
	0x62, 0xfb, 0x10, 0x4a, 0x75, 0x94, 0xa6, 0x39, 0x6a, 0x1d, 0xb5, 0x7b, 0x4e, 0x3f, 0xe4, 0x2b,
	0x82, 0x31, 0x70, 0x53, 0x51, 0x60, 0xe4, 0x52, 0x80, 0xd6, 0xec, 0x00, 0x40, 0x8c, 0x46, 0xa8,
	0xf5, 0xd9, 0x95, 0xc2, 0xc8, 0xa3, 0xc8, 0x1a, 0xc3, 0x9e, 0x42, 0x57, 0x97, 0x04, 0x87, 0x85,
	0x28, 0x4a, 0x1d, 0xf9, 0x3d, 0xa7, 0x1f, 0xf0, 0x26, 0x19, 0x7f, 0x06, 0xef, 0xf4, 0xa8, 0x2c,
	0x26, 0x2c, 0x82, 0x2d, 0xa1, 0xd4, 0x07, 0x31, 0x43, 0xea, 0x36, 0xe4, 0x35, 0x34, 0xad, 0x09,
	0xa5, 0x3e, 0xad, 0xba, 0x0e, 0xf9, 0x8a, 0x60, 0x7b, 0x10, 0x08, 0xa5, 0xce, 0xb2, 0x2f, 0x38,
	0xaf, 0xfa, 0x5e, 0xe2, 0xf8, 0x15, 0x78, 0xe7, 0x62, 0x5a, 0x22, 0xdb, 0x05, 0x6f, 0x61, 0x16,
	0xd5, 0xd1, 0x16, 0x98, 0xd4, 0x05, 0xe6, 0x72, 0x2c, 0x31, 0xa5, 0x73, 0x03, 0xbe, 0xc4, 0xf1,
	0x1f, 0x07, 0x5c, 0x53, 0xe1, 0x86, 0x80, 0xbb, 0xe0, 0x15, 0x54, 0xcc, 0x76, 0x62, 0x81, 0x39,
	0x4a, 0x09, 0xad, 0xbf, 0x65, 0x79, 0x5a, 0x77, 0x51, 0x63, 0x13, 0x33, 0x22, 0xcf, 0xcd, 0xd5,
	0xac, 0x80, 0x4b, 0xcc, 0x0e, 0xa1, 0x33, 0xcd, 0x2e, 0xe5, 0xbc, 0x7a, 0xae, 0xc8, 0xeb, 0xb5,
	0xfb, 0xdb, 0x83, 0x20, 0xa9, 0x30, 0x6f, 0x44, 0xd9, 0x3e, 0x78, 0x6a, 0x92, 0xcd, 0x91, 0xa4,
	0xdc, 0x1e, 0xf8, 0x09, 0xdd, 0x8e, 0x5b, 0x92, 0xed, 0x81, 0x3b, 0x13, 0x72, 0x1a, 0x6d, 0x35,
	0x82, 0xc4, 0x99, 0xcc, 0xcc, 0xc8, 0x1c, 0x05, 0x55, 0x90, 0x44, 0xe7, 0x96, 0x8c, 0x7f, 0x39,
	0xb0, 0xf3, 0x5a, 0x68, 0x39, 0x22, 0x12, 0xbf, 0x96, 0xa8, 0x0b, 0xd3, 0x76, 0x8a, 0x0b, 0x39,
	0xc2, 0x93, 0xe3, 0x4a, 0xb6, 0x25, 0x66, 0x31, 0x74, 0xc6, 0x59, 0x3e, 0xc4, 0xbc, 0x8a, 0x5b,
	0x2d, 0x1a, 0x9c, 0x79, 0xd0, 0x0b, 0x73, 0xe6, 0xc9, 0x71, 0xa5, 0x48, 0x0d, 0x1b, 0x62, 0xb9,
	0xd7, 0xc4, 0x7a, 0x06, 0x3b, 0x39, 0x8e, 0x31, 0xc7, 0xb5, 0xd3, 0xad, 0xb7, 0x6e, 0xf0, 0xf1,
	0x6f, 0x07, 0x3a, 0xa7, 0xf7, 0xdc, 0x72, 0xed, 0xc1, 0xf6, 0x1d, 0x1e, 0x74, 0xef, 0xf2, 0xa0,
	0xd7, 0xf4, 0xe0, 0xad, 0x17, 0xf2, 0x37, 0x5c, 0xe8, 0x87, 0x03, 0x1d, 0xca, 0xba, 0xaf, 0x0b,
	0x2d, 0xcd, 0xda, 0x5e, 0x37, 0xeb, 0x6d, 0x2d, 0xb9, 0x1b, 0x5a, 0x12, 0x10, 0x70, 0xd4, 0x2a,
	0x9b, 0x6b, 0x5c, 0x1b, 0x85, 0x90, 0x46, 0x81, 0x81, 0x3b, 0xca, 0x52, 0xa4, 0xca, 0x1e, 0xa7,
	0x35, 0x7b, 0x02, 0xae, 0x31, 0x37, 0x15, 0xdc, 0x1e, 0x78, 0x89, 0x51, 0x88, 0x13, 0x65, 0xbe,
	0x9e, 0x14, 0x0b, 0xe3, 0x50, 0x5b, 0xac, 0x42, 0x83, 0xef, 0x2d, 0xf0, 0xcd, 0x2b, 0xbe, 0x1f,
	0xb2, 0x03, 0x53, 0xed, 0x52, 0xea, 0x02, 0x73, 0x66, 0x73, 0xf7, 0xc2, 0xa4, 0xae, 0x1f, 0x3f,
	0x60, 0x2f, 0xa0, 0xfb, 0xce, 0x0c, 0x84, 0x89, 0xd0, 0xbb, 0x3c, 0x4a, 0xae, 0xfb, 0xb6, 0x99,
	0x70, 0x08, 0x40, 0x09, 0x6f, 0x68, 0x0a, 0xfe, 0x77, 0xf7, 0x47, 0x9a, 0xa7, 0x7f, 0xed, 0xee,
	0x57, 0xbb, 0xed, 0xff, 0xd5, 0x4d, 0x4e, 0x37, 0xee, 0x7c, 0x0e, 0xdd, 0x73, 0x31, 0x95, 0xe6,
	0xdb, 0xb4, 0xa6, 0xe8, 0x26, 0xeb, 0xcf, 0xdc, 0xd8, 0x7c, 0xe1, 0xd3, 0x27, 0xfe, 0xf2, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0xde, 0xfb, 0xa0, 0xd4, 0x05, 0x00, 0x00,
}