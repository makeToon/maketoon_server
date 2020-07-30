package database

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"makeToon/handler"
	"makeToon/model"
	"time"
)

var Client *mongo.Client

func MongoConn() (client *mongo.Client){
	if err := env.Parse(&handler.Envs); err != nil {
		fmt.Printf("%+v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, connectErr := mongo.Connect(ctx, options.Client().ApplyURI(
		handler.Envs.DbLocation,
	))
	if connectErr != nil { log.Fatal(connectErr) }
	defer cancel()

	// Check the connection
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	Client = client
	return client
}

func SetPhoto(userId, area, photo, width, height string) {
	findUser := model.User{}
	areaDocument := make(map[string]string)
	areaDocument["area"] = area
	areaDocument["imgUrl"] = photo
	areaDocument["width"] = width
	areaDocument["height"] = height

	user := model.User{UserID: userId, Area: []map[string]string{areaDocument}}

	// collection connect
	client := Client
	c := client.Database("maketoon").Collection("users")

	filter := bson.D{{"userId", userId}}
	checkErr := c.FindOne(context.TODO(), filter).Decode(&findUser)
	if checkErr != nil {
		if checkErr != mongo.ErrNoDocuments {
			panic(checkErr)
		}
	}

	if findUser.UserID == "" {
		_, err := c.InsertOne(context.TODO(), &user)
		if err != nil {
			panic(err)
		}
	} else {
		var copyArea = findUser.Area
		var isExist = false

		for i, _ := range copyArea {
			if copyArea[i]["area"] == area {
				copyArea[i] = areaDocument
				isExist = true
				break
			}
		}

		if isExist == false {
			copyArea = append(copyArea, areaDocument)
		}

		update := bson.D{
			{"$set", bson.D{
				{"area", copyArea},
			}},
		}
		_, err := c.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
	}
}

func GetFunc(userId string) []map[string]string {
	findUser := model.User{}

	client := Client
	c := client.Database("maketoon").Collection("users")

	filter := bson.D{{"userId", userId}}
	checkErr := c.FindOne(context.TODO(), filter).Decode(&findUser)
	if checkErr != nil {
		panic(checkErr)
	}

	if findUser.UserID == "" {
		var dummyUser []map[string]string = nil
		return dummyUser
	} else {
		return findUser.Area
	}
}