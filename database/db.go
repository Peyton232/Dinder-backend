package database

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	model "github.com/henlegay/diner-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
	rooms  *mongo.Collection
}

func Connect() *DB {

	// if time permits actually pull this info from env file
	// setup login info for db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:zV4eLIazskx5dbQm@cluster0.5zlwicm.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Print(err)
		log.Print("\nDB connection failed in database package")
		return nil
	}

	// setup context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to DB
	client.Connect(ctx)

	// return connection info for db
	return &DB{
		rooms:  client.Database("killabytez").Collection("rooms"),
		client: client,
	}
}

func (db DB) CreateRoom(user string, location string) (string, error) {
	//select collection
	collection := db.rooms

	// setup context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//generate random 4 digit code
	code := strconv.Itoa(rangeIn(1000, 9999))

	// put data into model
	roomData := &model.Room{
		Users:    []string{user},
		RoomID:   code,
		Location: location,
	}

	// create room in db
	res, err := collection.InsertOne(ctx, roomData)
	if err != nil || res == nil {
		log.Print(err)
		log.Print("\nunable to insert room into DB in database package\n")
		return "errors", err
	}

	// return code and that ther is no error
	return code, nil
}

func (db DB) JoinRoom(user string, room string) error {
	// select collection
	collection := db.rooms

	// setup context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check if room exists
	roomModel := model.Room{}
	filter := bson.M{"roomid": room}
	res := collection.FindOne(ctx, filter).Decode(&roomModel)

	// if no result
	if res == mongo.ErrNoDocuments {
		return res
	}

	// add user to room
	update := bson.M{"$addToSet": bson.M{"users": user}}
	collection.FindOneAndUpdate(ctx, filter, update)

	return nil
}

func (db DB) LeaveRoom(user string, room string) error {
	// select collection
	collection := db.rooms

	// setup context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// remove user from room
	filter := bson.M{"roomid": room}
	update := bson.M{"$pull": bson.M{"users": user}}
	collection.FindOneAndUpdate(ctx, filter, update)

	return nil
}

func (db DB) FinalCountdown(room string) (string, error) {
	// select collection
	collection := db.rooms

	// setup context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// check if room exists
	roomModel := model.Room{}
	filter := bson.M{"roomid": room}
	res := collection.FindOne(ctx, filter).Decode(&roomModel)

	// if no result
	if res == mongo.ErrNoDocuments {
		return "", res
	}

	// loop until room has completed
	for {
		collection.FindOne(ctx, filter).Decode(&roomModel)
		if roomModel.Found {
			return roomModel.Winner, nil
		}
		time.Sleep(10 * time.Second)
	}
}

func (db DB) Vote(room string, restaurant string) (string, error) {

	// select collection
	collection := db.rooms

	// setup context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check if room exists
	roomModel := model.Room{}
	filter := bson.M{"roomid": room}
	res := collection.FindOne(ctx, filter).Decode(&roomModel)
	// if no result
	if res == mongo.ErrNoDocuments {
		return "", res
	}

	for i := 0; i < len(roomModel.Restauraunts); i++ {
		if roomModel.Restauraunts[i] == restaurant {
			roomModel.Votes[i]++
		}
		//final winning condition
		if roomModel.Votes[i] == len(roomModel.Users) {
			//decision reached, return restaurant name
			//set winner to restaurant name
			roomModel.Winner = roomModel.Restauraunts[i]
			//set found to true
			roomModel.Found = true
			collection.FindOneAndUpdate(ctx, filter, roomModel)
			return restaurant, nil
		}
	}
	return "", nil
}

// ---------------------------------------------------------- helper funcs ----------------------------------------------------------
func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
