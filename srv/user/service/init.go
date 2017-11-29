package service

import (
	srv "micro_service_deployed_by_k8s/shared/service"
	proto "micro_service_deployed_by_k8s/srv/user/proto"
)

type Service struct{}

func init() {
	srv.Svc.Init()
	proto.RegisterServiceHandler(srv.Svc.Server(), new(Service))
}
