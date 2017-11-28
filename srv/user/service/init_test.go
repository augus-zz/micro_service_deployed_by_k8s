package service

import (
	// "fmt"
	// micro "github.com/micro/go-micro"
	// "github.com/micro/go-micro/registry"
	// "log"
	"os"
	// "os/exec"
	"testing"
)

var (
// cmd *exec.Cmd
)

func initTestEnv() {
	// do something for test env
	// os.Setenv("REGISTRY_SERVER_HOST", registryHost)
	// os.Setenv("REGISTRY_SERVER_PORT", registryPort)

	startRegistry()
}

func startRegistry() {
	// 	cmd = exec.Command(fmt.Sprintf("consul agent -dev --advertise=%s -http-port=%s", registryHost, registryPort))
	// 	err := cmd.Start()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
}

func stopRegistry() {
	// 	// we need to stop Registry
	// 	if err := cmd.Process.Signal(os.Interrupt); err != nil {
	// 		if err := cmd.Process.Signal(os.Kill); err != nil {
	// 			log.Println("failed to stop register, register addr: %s:%d", registryHost, registryPort)
	// 		}
	// 	}
	// 	defer cmd.Wait()
}

func TestInit(t *testing.T) {
	// get service from registry
	registryHost := os.Getenv("REGISTRY_SERVER_HOST")
	registryPort := os.Getenv("REGISTRY_SERVER_PORT")

	r := newRegistry(registryHost, registryPort)

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
