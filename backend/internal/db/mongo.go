package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect(uri string) (*mongo.Client, error) {
	var err error
	clientOptions := options.Client().ApplyURI(uri)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return &mongo.Client{}, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return &mongo.Client{}, err
	}

	log.Printf("Connected to MongoDB on %s", uri)
	return client, nil
}
