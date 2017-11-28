package config

import (
	"micro_service_deployed_by_k8s/shared/utils"
	"os"
)

var (
	ENV_PRODUCTION  string = "production"
	ENV_DEVELOPMENT string = "development"
	ENV_STAGING     string = "staging"
	ENV_TEST        string = "test"

	SERVICE_ENV string = ENV_TEST

	SERVICE_ENVS []string = []string{
		ENV_PRODUCTION,
		ENV_DEVELOPMENT,
		ENV_STAGING,
		ENV_TEST,
	}
)

func init() {
	initEnv()
}

func initEnv() {
	if !utils.Contains(SERVICE_ENVS, os.Getenv("SERVICE_ENV")) {
		os.Setenv("SERVICE_ENV", ENV_TEST)
		SERVICE_ENV = ENV_TEST
	}
}
