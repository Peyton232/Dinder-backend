package database

import (
	"context"
	"log"
	"math/rand"
	"strconv"
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

func (DB *DB) CreateRoom(user string, location string) (string, error) {

	//generate random 4 digit code
	code := strconv.Itoa(rangeIn(1000, 9999))

	// create room in db

	return code, nil
}

// helper funcs
func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
