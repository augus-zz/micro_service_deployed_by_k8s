package service

import (
	// "micro_service_deployed_by_k8s//srv/user/model"
	"fmt"
	micro "github.com/micro/go-micro"
	registry "github.com/micro/go-micro/registry"
	"log"
	proto "micro_service_deployed_by_k8s/srv/user/proto"
	"os"
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

	registryServer := os.Getenv("REGISTRY_SERVER_HOST")
	registryPort := os.Getenv("REGISTRY_SERVER_PORT")
	registryAddr := fmt.Sprintf("%s:%s", server, registryPort)
	registry := registry.NewRegistry(
		registry.Addrs(registryAddr),
	)
	service.Init(micro.Registry(registry))

	proto.RegisterServiceHandler(service.Server(), new(Service))
}

func Run() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
