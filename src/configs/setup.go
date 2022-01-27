package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() *mongo.Client {

	DbUser := GetEnv("DB_USER")
	DbName := GetEnv("DB_NAME")
	DbHost := GetEnv("DB_HOST")
	DbPassword := GetEnv("DB_PASSWORD")

	MongoUri := "mongodb+srv://" + DbUser + ":" + DbPassword + "@" + DbHost + "/" + DbName + "?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	DbName := GetEnv("DB_NAME")

	collection := client.Database(DbName).Collection(collectionName)

	return collection
}
