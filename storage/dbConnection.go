package storage

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	// MongoDb := "mongodb://localhost:27017"

	// client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("connected to mongodb")
	// return client

	// create connection options
	//mongoUrl := "mongodb://mongo:27017"
	//mongoUrl := "mongodb://admin:password@172.21.0.1:27017"

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	mongoUrl := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("error connecting", err)
		return nil
	}

	log.Println("Connected to mongo!")

	return c
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
