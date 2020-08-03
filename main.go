package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/makeToon/maketoon_server/database"
	"github.com/makeToon/maketoon_server/handler"
	"github.com/makeToon/maketoon_server/route"
)

func main() {
	if err := env.Parse(&handler.Envs); err != nil {
		fmt.Printf("%+v\n", err)
	}

	handler.AwsConfig()
	database.MongoConn()
	router := route.Init()
	router.Logger.Fatal(router.Start(":8080"))
	//router.Logger.Fatal(router.Start(handler.Envs.Port))
}