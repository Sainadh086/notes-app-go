package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insert(ctx context.Context, client *mongo.Client, data map[string]string) (interface{}, error) {
	collection := client.Database("baz").Collection("qux")
	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	return res.InsertedID, err
}

func main() {
	log.Println("Starting Notes App")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"))
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println("Connected to MongoDB")
	} else {
		log.Fatal("Failed to connect to MongoDB")
	}
	var data = map[string]string{"notes": "true", "data": "app"}
	_, err = insert(ctx, mongoClient, data)
	if err != nil {
		log.Fatal("Failed to insert data")
	}
}
