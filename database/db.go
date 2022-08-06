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
	rooms  *mongo.Collection
}

func Connect() *DB {

	// if time permits actually pull this info from env file
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:zV4eLIazskx5dbQm@cluster0.ubxcq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Print(err)
		log.Print("\nDB connection failed in database package")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	return &DB{
		rooms:  client.Database("killabytez").Collection("rooms"),
		client: client,
	}
}
