package service

import (
	"context"
	"log"
	"micro-demo-deploy-by-k8s/srv/user/model"
	proto "micro-demo-deploy-by-k8s/srv/user/proto"
)

func (s *Service) GetUsers(ctx context.Context, req *proto.UsersRequest, rsp *proto.UsersResponse) error {
	users := []model.User{}

	if err := model.DB.Where("id IN (?)", req.Ids).Find(&users).Error; err != nil || len(users) == 0 {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Users = []*proto.User{}
	for _, user := range users {
		u := &proto.User{}
		u.Id = user.ID
		u.Name = user.Name
		u.Phone = user.Phone
		u.Email = user.Email
		u.Role = user.Role

		rsp.Users = append(rsp.Users, u)
	}

	return nil
}

func (s *Service) SearchUsers(ctx context.Context, req *proto.UsersSearchRequest, rsp *proto.UsersResponse) error {
	return nil
}

func (s *Service) GetUser(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	user := model.User{}

	if err := model.DB.First(&user, req.User.Id).Error; err != nil || user.ID != req.User.Id {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.User = &proto.User{}
	rsp.User.Id = user.ID
	rsp.User.Name = user.Name
	rsp.User.Phone = user.Phone
	rsp.User.Email = user.Email
	rsp.User.Role = user.Role
	rsp.Success = true

	return nil
}

func (s *Service) CreateUser(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	user := model.User{
		ID:    req.User.Id,
		Name:  req.User.Name,
		Phone: req.User.Phone,
		Email: req.User.Email,
		Role:  req.User.Role,
	}

	if err := model.DB.Create(&user).Error; err != nil {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.User = &proto.User{}
	rsp.User.Id = user.ID
	rsp.User.Name = user.Name
	rsp.User.Phone = user.Phone
	rsp.User.Email = user.Email
	rsp.User.Role = user.Role
	rsp.Success = true

	return nil
}

func (s *Service) UpdateUser(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, req *proto.UserRequest, rsp *proto.DeleteResponse) error {
	user := model.User{ID: req.User.Id}

	if err := model.DB.Delete(&user).Error; err != nil {
		rsp.Success = false
		rsp.ErrorMsg = "Failed to delete"
		return nil
	}

	log.Printf("req: %v", req)
	rsp.Success = true

	return nil
}
