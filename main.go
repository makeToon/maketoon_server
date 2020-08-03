package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/makeToon/maketoon_server/database"
	"github.com/makeToon/maketoon_server/handler"
	"github.com/makeToon/maketoon_server/route"
	"os"
)

func main() {
	if err := env.Parse(&handler.Envs); err != nil {
		fmt.Printf("%+v\n", err)
	}
	port := os.Getenv("PORT")

	handler.AwsConfig()
	database.MongoConn()
	router := route.Init()
	router.Logger.Fatal(router.Start(port))
	//router.Logger.Fatal(router.Start(handler.Envs.Port))
}