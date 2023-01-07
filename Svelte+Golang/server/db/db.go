package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URL = "mongodb+srv://aaa:aaa@cluster0.retrres.mongodb.net/?retryWrites=true&w=majority"
const dbName = "todo"
const collectionName = "todos"

func Connect() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(URL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance created")
	return collection
}
