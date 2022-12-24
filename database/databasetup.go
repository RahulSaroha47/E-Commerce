package database

import (
	"context"
	"fmt"
	"log"
	"time"	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client = DBSetup()

func DBSetup() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:3306"))

	if err!=nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
    if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(),nil)
	if err != nil {
		fmt.Println("failed to connect to mongodb :(")
		return nil
	}

	fmt.Println("succefully connected to database")
	return client
	
}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
  var userColllection = client.Database("ECommerce").Collection(collectionName)
  return userColllection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
  var productColllection = client.Database("ECommerce").Collection(collectionName)
  return productColllection
}