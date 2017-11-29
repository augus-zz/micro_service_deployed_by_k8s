package service

import (
	"context"
	"log"
	"micro_service_deployed_by_k8s/shared/db"
	"micro_service_deployed_by_k8s/srv/user/model"
	proto "micro_service_deployed_by_k8s/srv/user/proto"
)

func (s *Service) GetUsers(ctx context.Context, req *proto.UsersRequest, rsp *proto.UsersResponse) error {
	users := []model.User{}

	if err := db.DB.Where("id IN (?)", req.Ids).Find(&users).Error; err != nil || len(users) == 0 {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Users = []*proto.User{}
	for _, user := range users {
		u := &proto.User{
			Id:    user.ID,
			Name:  user.Name,
			Phone: user.Phone,
			Email: user.Email,
			Role:  user.Role,
		}

		rsp.Users = append(rsp.Users, u)
	}

	return nil
}

func (s *Service) SearchUsers(ctx context.Context, req *proto.UsersSearchRequest, rsp *proto.UsersResponse) error {
	return nil
}

func (s *Service) GetUser(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	user := model.User{}

	if err := db.DB.First(&user, req.User.Id).Error; err != nil || user.ID != req.User.Id {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.User = &proto.User{
		Id:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		Role:  user.Role,
	}
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

	if err := db.DB.Create(&user).Error; err != nil {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.User = &proto.User{
		Id:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		Role:  user.Role,
	}
	rsp.Success = true

	return nil
}

func (s *Service) UpdateUser(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	user := model.User{
		ID:    req.User.Id,
		Name:  req.User.Name,
		Phone: req.User.Phone,
		Email: req.User.Email,
		Role:  req.User.Role,
	}

	// should not update all attrs
	if err := db.DB.Save(&user).Error; err != nil {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.User = &proto.User{
		Id:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		Role:  user.Role,
	}
	rsp.Success = true
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, req *proto.UserRequest, rsp *proto.DeleteResponse) error {
	user := model.User{ID: req.User.Id}

	if err := db.DB.Delete(&user).Error; err != nil {
		rsp.Success = false
		rsp.ErrorMsg = "Failed to delete"
		return nil
	}

	log.Printf("req: %v", req)
	rsp.Success = true

	return nil
}
