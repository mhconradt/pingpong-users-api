package main

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/mhconradt/grpc-statuses"
	proto "github.com/mhconradt/pingpong-users-api/proto"
	"github.com/mhconradt/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server/grpc"
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"
	"log"
)

type Users struct{}

var pool *gorm.DB
// How to share DB connection between requests?

func (u *Users) CreateUser(ctx context.Context, req *proto.CreateUserRequest, res *proto.CreateUserResponse) error {
	if usr, err := user.DefaultCreateUser(ctx, req.User, pool); err != nil {
		res.Status = status.BadRequest("request failed to complete")
	} else {
		res.User = usr
		res.Status = status.Success()
	}
	return nil
}

func (u *Users) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, res *proto.UpdateUserResponse) error {
	if usr, err := user.DefaultPatchUser(ctx, req.User, req.UpdateMask, pool); err != nil {
		res.Status = status.BadRequest("invalid request")
	} else {
		res.User = usr
		res.Status = status.Success()
	}
	return nil
}

func (u *Users) GetUser(ctx context.Context, req *proto.GetUserRequest, res *proto.GetUserResponse) error {
	q := &user.User{Id: req.Id}
	usr, err := user.DefaultReadUser(ctx, q, pool)
	if err != nil {
		res.Status = status.NotFound("user with id: %v not found", req.Id)
		return nil
	}
	if usr, err = user.DefaultApplyFieldMaskUser(ctx, new(user.User), usr, req.GetMask, "", pool); err != nil {
		res.Status = status.InternalServerError("failed to apply field mask to result")
	} else {
		res.User = usr
		res.Status = status.Success()
	}
	return nil
}

func (u *Users) DeactivateUser(ctx context.Context, req *proto.DeactivateUserRequest, res *proto.DeactivateUserResponse) error {
	orm, err := (&user.User{Id: req.Id}).ToORM(ctx)
	if err != nil {
		res.Status = status.BadRequest("invalid request")
	}
	if err := pool.Model(orm).Update("status", 0).Error; err != nil {
		res.Status = status.InternalServerError("failed to deactivate user")
	} else {
		res.User = &user.User{}
		res.Status = status.Success()
	}
	return nil
}

func main() {
	// Use the protobuf codec
	c := encoding.GetCodec("proto")
	codecOpt := grpc.Codec("application/protobuf", c)
	// Use one port for consistency's sake
	p := LookupWithDefault("PORT", ":3500")
	portOpt := micro.Address(p)
	srv := grpc.NewServer(codecOpt)
	service := micro.NewService(
		micro.Server(srv),
		micro.Name("users"),
		micro.Version("latest"),
		portOpt,
	)
	// Define err before so we don't re-declare pool.
	var err error
	pool, err = InitDB()
	if err != nil {
		log.Fatal(err)
	}
	service.Init()
	if err = proto.RegisterUsersHandler(service.Server(), new(Users)); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
