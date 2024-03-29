// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: users.proto

package users

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/mhconradt/proto/status"
	_ "github.com/mhconradt/proto/user"
	_ "google.golang.org/genproto/protobuf/field_mask"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Users service

type UsersService interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UpdateUserResponse, error)
	DeactivateUser(ctx context.Context, in *DeactivateUserRequest, opts ...client.CallOption) (*DeactivateUserResponse, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "users"
	}
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "Users.GetUser", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error) {
	req := c.c.NewRequest(c.name, "Users.CreateUser", in)
	out := new(CreateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UpdateUserResponse, error) {
	req := c.c.NewRequest(c.name, "Users.UpdateUser", in)
	out := new(UpdateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) DeactivateUser(ctx context.Context, in *DeactivateUserRequest, opts ...client.CallOption) (*DeactivateUserResponse, error) {
	req := c.c.NewRequest(c.name, "Users.DeactivateUser", in)
	out := new(DeactivateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	GetUser(context.Context, *GetUserRequest, *GetUserResponse) error
	CreateUser(context.Context, *CreateUserRequest, *CreateUserResponse) error
	UpdateUser(context.Context, *UpdateUserRequest, *UpdateUserResponse) error
	DeactivateUser(context.Context, *DeactivateUserRequest, *DeactivateUserResponse) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) error {
	type users interface {
		GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error
		CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error
		UpdateUser(ctx context.Context, in *UpdateUserRequest, out *UpdateUserResponse) error
		DeactivateUser(ctx context.Context, in *DeactivateUserRequest, out *DeactivateUserResponse) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	return s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.UsersHandler.GetUser(ctx, in, out)
}

func (h *usersHandler) CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error {
	return h.UsersHandler.CreateUser(ctx, in, out)
}

func (h *usersHandler) UpdateUser(ctx context.Context, in *UpdateUserRequest, out *UpdateUserResponse) error {
	return h.UsersHandler.UpdateUser(ctx, in, out)
}

func (h *usersHandler) DeactivateUser(ctx context.Context, in *DeactivateUserRequest, out *DeactivateUserResponse) error {
	return h.UsersHandler.DeactivateUser(ctx, in, out)
}
