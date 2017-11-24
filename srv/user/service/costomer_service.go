package service

import (
	"../model"
	"../proto"
)

type CustomerService struct{}

func (us *CustomerService) GetCustomers(ctx context.Context, req *proto.CustomersRequest, rsp *proto.CustomersResponse) error {
	return nil
}

func (us *CustomerService) SearchCustomers(ctx context.Context, req *proto.CustomersSearchRequest, rsp *proto.CustomersResponse) error {
	return nil
}

func (us *CustomerService) GetCustomer(ctx context.Context, req *proto.CustomerRequest, rsp *proto.CustomerResponse) error {
	return nil
}

func registerCustomerService() {
	proto.RegisterCustomerServiceHandler(service.Server(), new(CustomerService))
}
