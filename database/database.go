package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Replace with your MongoDB connection string
	client, _ = mongo.Connect(context.Background(), clientOptions)
}

type Collection struct {
	*mongo.Collection
}

func NewCollection(collectionName string) *Collection {
	return &Collection{
		Collection: client.Database("algoTrading").Collection(collectionName), // Replace with your database name
	}
}
