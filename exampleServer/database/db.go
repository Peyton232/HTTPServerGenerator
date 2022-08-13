package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
	pet    *mongo.Collection
}

// make env vars
var mongoUri string = ""
var mongoDb string = ""

func Connect() (*DB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Print("\nDB connection failed in database package")
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	return &DB{
		client: client,
		pet:    client.Database(mongoDb).Collection("pet"),
	}, nil
}
