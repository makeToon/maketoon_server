package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"makeToon/database"
	"makeToon/handler"
	"makeToon/route"
)

func main() {
	if err := env.Parse(&handler.Envs); err != nil {
		fmt.Printf("%+v\n", err)
	}

	handler.AwsConfig()
	database.MongoConn()
	router := route.Init()
	router.Logger.Fatal(router.Start(handler.Envs.Port))
}