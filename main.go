package main

import (
	"makeToon/database"
	"makeToon/handler"
	"makeToon/route"
)

func main() {
	handler.AwsConfig()
	database.MongoConn()
	router := route.Init()
	router.Logger.Fatal(router.Start(":3000"))
}