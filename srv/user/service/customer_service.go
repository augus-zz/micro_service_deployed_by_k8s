package service

import (
	// "micro-demo-deploy-by-k8s/srv/user/model"
	"context"
	proto "micro-demo-deploy-by-k8s/srv/user/proto"
)

func (s *Service) GetCustomers(ctx context.Context, req *proto.CustomersRequest, rsp *proto.CustomersResponse) error {
	return nil
}

func (s *Service) SearchCustomers(ctx context.Context, req *proto.CustomersSearchRequest, rsp *proto.CustomersResponse) error {
	return nil
}

func (s *Service) GetCustomer(ctx context.Context, req *proto.CustomerRequest, rsp *proto.CustomerResponse) error {
	return nil
}
