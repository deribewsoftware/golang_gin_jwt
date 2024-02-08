package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnection() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading in .env file: ")
	}

	mongoDB := os.Getenv("DB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDB))
	if err != nil {
		log.Fatal(err)

	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	// database connection
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connect to mongodb database")

	return client

}

var Client *mongo.Client = DBConnection()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection

}
