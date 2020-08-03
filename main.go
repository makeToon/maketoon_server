package main

import (
	"maketoon/database"
	"maketoon/handler"
	"maketoon/route"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	handler.AwsConfig()
	database.MongoConn()
	router := route.Init()
	router.Logger.Fatal(router.Start(":" + port))
}