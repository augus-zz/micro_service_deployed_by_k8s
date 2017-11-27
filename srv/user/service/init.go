package service

import (
	// "micro_service_deployed_by_k8s//srv/user/model"
	micro "github.com/micro/go-micro"
	"log"
	proto "micro_service_deployed_by_k8s/srv/user/proto"
)

var (
	service micro.Service
)

type Service struct{}

func init() {
	service = micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "user",
		}),
	)

	service.Init()

	proto.RegisterServiceHandler(service.Server(), new(Service))
}

func Run() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
