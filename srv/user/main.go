package main

import (
	_ "micro_service_deployed_by_k8s/shared/config"
	_ "micro_service_deployed_by_k8s/shared/db"
	srv "micro_service_deployed_by_k8s/shared/service"
	_ "micro_service_deployed_by_k8s/srv/user/service"
)

func main() {
	srv.Svc.Run()
}
