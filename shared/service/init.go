package service

import (
	// "micro_service_deployed_by_k8s//srv/user/model"
	"fmt"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"log"
	"os"
)

var (
	Svc micro.Service
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
		// initTestEnv()
	}

	service_app := os.Getenv("SERVICE_APP")
	// registryHost := os.Getenv("REGISTRY_SERVER_HOST")
	// registryPort := os.Getenv("REGISTRY_SERVER_PORT")

	Svc = micro.NewService(
		micro.Name(service_app),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": service_app,
		}),
		// micro.Registry(newRegistry(registryHost, registryPort)),
	)

	Svc.Init()
	// Svc.Init(micro.Registry(newRegistry(registryHost, registryPort)))
	// proto.RegisterServiceHandler(Svc.Server(), new(Service))
}

func Run() {
	if err := Svc.Run(); err != nil {
		log.Fatal(err)
	}
}
