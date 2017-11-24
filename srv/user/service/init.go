package service

import (
	"../model"
	"../proto"
	micro "github.com/micro/go-micro"
	"log"
)

var (
	service micro.Service
)

func init() {
	service = micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "user",
		}),
	)

	service.Init()

	registerUserService()
	registerStoreService()
	registerCustomerService()
}

func Run() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
