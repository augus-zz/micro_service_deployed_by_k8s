package service

import (
	// "../model"
	"../proto"
)

type UserService struct{}

func (us *UserService) GetUsers(ctx context.Context, req *proto.UsersRequest, rsp *proto.UsersResponse) error {
	return nil
}

func (us *UserService) SearchUsers(ctx context.Context, req *proto.UsersSearchRequest, rsp *proto.UsersResponse) error {
	return nil
}

func (us *UserService) GetUser(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	return nil
}

func registerUserService() {
	proto.RegisterUserServiceHandler(service.Server(), new(UserService))
}
