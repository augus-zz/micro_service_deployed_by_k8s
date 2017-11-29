package service

import (
	srv "micro_service_deployed_by_k8s/shared/service"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	// get service from registry
	registryHost := os.Getenv("REGISTRY_SERVER_HOST")
	registryPort := os.Getenv("REGISTRY_SERVER_PORT")

	r := srv.NewRegistry(registryHost, registryPort)

	if services, err := r.GetService("user"); err != nil {
		t.Fatal("failed to query service: user")

		if len(services) != 1 || services[0].Name != "user" {
			t.Fatal("failed to register service: user")
		}

		r.Deregister(services[0])
	}
}

func TestRun(t *testing.T) {
	t.Skip("skip this, because we can not call service.Run now")
}
