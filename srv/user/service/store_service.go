package service

import (
	// "../model"
	"../proto"
)

type StoreService struct{}

func (us *StoreService) GetStores(ctx context.Context, req *proto.StoresRequest, rsp *proto.StoresResponse) error {
	return nil
}

func (us *StoreService) SearchStores(ctx context.Context, req *proto.StoresSearchRequest, rsp *proto.StoresResponse) error {
	return nil
}

func (us *StoreService) GetStore(ctx context.Context, req *proto.StoreRequest, rsp *proto.StoreResponse) error {
	return nil
}

func registerStoreService() {
	proto.RegisterStoreServiceHandler(service.Server(), new(StoreService))
}
