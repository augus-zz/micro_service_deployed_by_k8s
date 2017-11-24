package service

import (
	// "micro-demo-deploy-by-k8s/srv/user/model"
	"context"
	proto "micro-demo-deploy-by-k8s/srv/user/proto"
)

func (s *Service) GetUsers(ctx context.Context, req *proto.UsersRequest, rsp *proto.UsersResponse) error {
	return nil
}

func (s *Service) SearchUsers(ctx context.Context, req *proto.UsersSearchRequest, rsp *proto.UsersResponse) error {
	return nil
}

func (s *Service) GetUser(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	return nil
}
