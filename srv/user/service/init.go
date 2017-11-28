package service

import (
	// "micro_service_deployed_by_k8s//srv/user/model"
	"fmt"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"log"
	proto "micro_service_deployed_by_k8s/srv/user/proto"
	"os"
)

var (
	service micro.Service
)

type Service struct{}

func newRegistry(registryHost, registryPort string) registry.Registry {
	registryAddr := fmt.Sprintf("%s:%s", registryHost, registryPort)
	return registry.NewRegistry(
		registry.Addrs(registryAddr),
	)
}

func init() {
	env := os.Getenv("SERVICE_ENV")
	if env == "test" {
		initTestEnv()
	}

	registryHost := os.Getenv("REGISTRY_SERVER_HOST")
	registryPort := os.Getenv("REGISTRY_SERVER_PORT")

	service = micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "user",
		}),
		micro.Registry(newRegistry(registryHost, registryPort)),
	)

	service.Init()
	// service.Init(micro.Registry(newRegistry(registryHost, registryPort)))
	proto.RegisterServiceHandler(service.Server(), new(Service))
}

func Run() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
