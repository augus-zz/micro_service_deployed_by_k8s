package main

import (
	_ "micro_service_deployed_by_k8s/shared/config"
	"micro_service_deployed_by_k8s/srv/user/service"
)

func main() {
	service.Run()
}
