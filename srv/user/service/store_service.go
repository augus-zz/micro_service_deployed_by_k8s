package service

import (
	// "micro-demo-deploy-by-k8s/srv/user/model"
	"context"
	proto "micro-demo-deploy-by-k8s/srv/user/proto"
)

func (s *Service) GetStores(ctx context.Context, req *proto.StoresRequest, rsp *proto.StoresResponse) error {
	return nil
}

func (s *Service) SearchStores(ctx context.Context, req *proto.StoresSearchRequest, rsp *proto.StoresResponse) error {
	return nil
}

func (s *Service) GetStore(ctx context.Context, req *proto.StoreRequest, rsp *proto.StoreResponse) error {
	return nil
}
