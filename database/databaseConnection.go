package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	connectionString := "mongodb+srv://Soumya:test123@cluster0.3or3tpi.mongodb.net/?retryWrites=true&w=majority"
	fmt.Println(connectionString)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clienterr := client.Connect(ctx)

	if clienterr != nil {
		log.Fatal(clienterr)
	}
	fmt.Println("connected to MongoDB")
	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("Pizzy").Collection(collectionName)
	return collection

}
