package service

import (
	"context"
	"log"
	"micro-demo-deploy-by-k8s/srv/user/model"
	proto "micro-demo-deploy-by-k8s/srv/user/proto"
)

func (s *Service) GetCustomers(ctx context.Context, req *proto.CustomersRequest, rsp *proto.CustomersResponse) error {
	customers := []model.Customer{}

	if err := model.DB.Where("id IN (?)", req.Ids).Find(&customers).Error; err != nil || len(customers) == 0 {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Customers = []*proto.Customer{}
	for _, customer := range customers {
		c := &proto.Customer{
			Id:      customer.ID,
			Name:    customer.Name,
			StoreId: customer.StoreID,
			UserId:  customer.UserID,
		}

		rsp.Customers = append(rsp.Customers, c)
	}
	return nil
}

func (s *Service) SearchCustomers(ctx context.Context, req *proto.CustomersSearchRequest, rsp *proto.CustomersResponse) error {
	return nil
}

func (s *Service) GetCustomer(ctx context.Context, req *proto.CustomerRequest, rsp *proto.CustomerResponse) error {
	customer := model.Customer{}

	if err := model.DB.First(&customer, req.Customer.Id).Error; err != nil || customer.ID != req.Customer.Id {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Customer = &proto.Customer{
		Id:      customer.ID,
		Name:    customer.Name,
		StoreId: customer.StoreID,
		UserId:  customer.UserID,
	}
	rsp.Success = true
	return nil
}

func (s *Service) CreateCustomer(ctx context.Context, req *proto.CustomerRequest, rsp *proto.CustomerResponse) error {
	customer := model.Customer{
		ID:      req.Customer.Id,
		Name:    req.Customer.Name,
		StoreID: req.Customer.Id,
		UserID:  req.Customer.UserId,
	}

	if err := model.DB.Create(&customer).Error; err != nil {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Customer = &proto.Customer{
		Id:      customer.ID,
		Name:    customer.Name,
		StoreId: customer.StoreID,
		UserId:  customer.UserID,
	}
	rsp.Success = true
	return nil
}

func (s *Service) UpdateCustomer(ctx context.Context, req *proto.CustomerRequest, rsp *proto.CustomerResponse) error {
	customer := model.Customer{
		ID:      req.Customer.Id,
		Name:    req.Customer.Name,
		StoreID: req.Customer.Id,
		UserID:  req.Customer.UserId,
	}

	// should not update all attrs
	if err := model.DB.Save(&customer).Error; err != nil {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Customer = &proto.Customer{
		Id:      customer.ID,
		Name:    customer.Name,
		StoreId: customer.StoreID,
		UserId:  customer.UserID,
	}
	rsp.Success = true

	return nil
}

func (s *Service) DeleteCustomer(ctx context.Context, req *proto.CustomerRequest, rsp *proto.DeleteResponse) error {
	customer := model.Customer{ID: req.Customer.Id}

	if err := model.DB.Delete(&customer).Error; err != nil {
		rsp.Success = false
		rsp.ErrorMsg = "Failed to delete"
		return nil
	}

	log.Printf("req: %v", req)
	rsp.Success = true

	return nil
}
