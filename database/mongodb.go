package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	uri := os.Getenv("URL_MONGO")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
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

	dbName := os.Getenv("DB_NAME")

	database = client.Database(dbName)

	fmt.Println("Connected to MongoDB!")

}
