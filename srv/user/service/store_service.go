package service

import (
	"context"
	"log"
	"micro_service_deployed_by_k8s/srv/user/model"
	proto "micro_service_deployed_by_k8s/srv/user/proto"
)

func (s *Service) GetStores(ctx context.Context, req *proto.StoresRequest, rsp *proto.StoresResponse) error {
	stores := []model.Store{}

	if err := model.DB.Where("id IN (?)", req.Ids).Find(&stores).Error; err != nil || len(stores) == 0 {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Stores = []*proto.Store{}
	for _, store := range stores {
		s := &proto.Store{
			Id:      store.ID,
			Title:   store.Title,
			Address: store.Address,
		}

		rsp.Stores = append(rsp.Stores, s)
	}

	return nil
}

func (s *Service) SearchStores(ctx context.Context, req *proto.StoresSearchRequest, rsp *proto.StoresResponse) error {
	return nil
}

func (s *Service) GetStore(ctx context.Context, req *proto.StoreRequest, rsp *proto.StoreResponse) error {
	store := model.Store{}

	if err := model.DB.First(&store, req.Store.Id).Error; err != nil || store.ID != req.Store.Id {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Store = &proto.Store{
		Id:      store.ID,
		Title:   store.Title,
		Address: store.Address,
	}
	rsp.Success = true

	return nil
}

func (s *Service) CreateStore(ctx context.Context, req *proto.StoreRequest, rsp *proto.StoreResponse) error {
	store := model.Store{
		ID:      req.Store.Id,
		Title:   req.Store.Title,
		Address: req.Store.Address,
	}

	if err := model.DB.Create(&store).Error; err != nil {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Store = &proto.Store{
		Id:      store.ID,
		Title:   store.Title,
		Address: store.Address,
	}
	rsp.Success = true

	return nil
}

func (s *Service) UpdateStore(ctx context.Context, req *proto.StoreRequest, rsp *proto.StoreResponse) error {
	store := model.Store{
		ID:      req.Store.Id,
		Title:   req.Store.Title,
		Address: req.Store.Address,
	}

	// should not update all attrs
	if err := model.DB.Save(&store).Error; err != nil {
		return nil
	}

	log.Printf("req: %v", req)

	rsp.Store = &proto.Store{
		Id:      store.ID,
		Title:   store.Title,
		Address: store.Address,
	}
	rsp.Success = true

	return nil
}

func (s *Service) DeleteStore(ctx context.Context, req *proto.StoreRequest, rsp *proto.DeleteResponse) error {
	store := model.Store{ID: req.Store.Id}

	if err := model.DB.Delete(&store).Error; err != nil {
		rsp.Success = false
		rsp.ErrorMsg = "Failed to delete"
		return nil
	}

	log.Printf("req: %v", req)
	rsp.Success = true

	return nil
}
