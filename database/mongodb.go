package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client         *mongo.Client
	database       *mongo.Database
	contextMongoDB context.Context
)

func GetCollectionName(collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}

func MongoDisconect() {
	fmt.Println("Closing MongoDB connection")
	client.Disconnect(contextMongoDB)
}

func GetContextMongo() context.Context {
	return contextMongoDB
}

func Init() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	contextMongoDB = context.Background()
	err = client.Connect(contextMongoDB)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(contextMongoDB, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database = client.Database("crypto")

	fmt.Println("Connected to MongoDB!")

}
