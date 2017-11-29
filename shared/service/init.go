package service

import (
	"fmt"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"log"
	"os"
)

var (
	Svc micro.Service
)

func NewRegistry(registryHost, registryPort string) registry.Registry {
	registryAddr := fmt.Sprintf("%s:%s", registryHost, registryPort)
	return registry.NewRegistry(
		registry.Addrs(registryAddr),
	)
}

func initTestEnv() {
}

func init() {
	env := os.Getenv("SERVICE_ENV")
	if env == "test" {
		initTestEnv()
	}

	service_app := os.Getenv("SERVICE_APP")
	registryHost := os.Getenv("REGISTRY_SERVER_HOST")
	registryPort := os.Getenv("REGISTRY_SERVER_PORT")

	Svc = micro.NewService(
		micro.Name(service_app),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": service_app,
		}),
		micro.Registry(
			NewRegistry(registryHost, registryPort),
		),
	)

	Svc.Init()
}

func Run() {
	if err := Svc.Run(); err != nil {
		log.Fatal(err)
	}
}
