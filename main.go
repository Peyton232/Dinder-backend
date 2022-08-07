package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/henlegay/diner-api/database"
)

var DB *database.DB
var restaurant [12]string = [12]string{"Mcdonalds", "Chick-fil-a", "Taco Bell", "Olive Garden", "Chilis", "Braums", "Mooyah", "Canes", "5 guys", "in-n-out", "chiptole", "Crooked Crust"}

func main() {
	// connect to database
	DB = database.Connect()

	// setup endpoints
	http.HandleFunc("/alive", Alive)
	http.HandleFunc("/create", CreateRoom)
	http.HandleFunc("/join", JoinRoom)
	http.HandleFunc("/leave", LeaveRoom)
	http.HandleFunc("/get", GetRooms)
	http.HandleFunc("/swipeRight", SwipeRight)
	http.HandleFunc("/swipeLeft", SwipeLeft)

	// start server
	http.ListenAndServe(":42069", nil)
}

// response  is the response back to user
// r is the request we recieve
func Alive(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, I'm still alive!") // prints string to response
}

// wat do: add user to a room in database
// wat need: user id, room id
// return: list of restaurants
func JoinRoom(response http.ResponseWriter, request *http.Request) {
	//set response type to json (only do if you need to return stuff)
	response.Header().Set("Content-Type", "application/json")

	// get 'user' query param
	user := request.URL.Query().Get("user")
	// get 'room' query param
	room := request.URL.Query().Get("room")

	// call join room db function, params are user and room
	err := DB.JoinRoom(user, room)

	//if there was error report it
	if err != nil {
		// write error code
		response.WriteHeader(http.StatusInternalServerError)
		// write error in response
		json.NewEncoder(response).Encode(struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	// send back a successful OK code
	response.WriteHeader(http.StatusOK)

	// write list of restauraunts into response
	json.NewEncoder(response).Encode(struct {
		Restauraunts []string `json:"restauraunts"`
	}{
		Restauraunts: restaurant[:],
	})
	//TODO: refactor so we have a model to return so this is cleaned up a bit, instead of being anonamyous

}

// wat do: creates a room
// wat need: location, user
// return: list of restaurants, room ID
func CreateRoom(response http.ResponseWriter, request *http.Request) {
	//set response type to json (only do if you need to return stuff)
	response.Header().Set("Content-Type", "application/json")

	// get 'user' query param and make sure it's provided
	user := request.URL.Query().Get("user")

	if user == "" {
		// write error code
		response.WriteHeader(http.StatusBadRequest)
		// write error in response
		json.NewEncoder(response).Encode(struct {
			Error string `json:"error"`
		}{
			Error: "no user provided",
		})
		return
	}

	// get 'location' query param
	location := request.URL.Query().Get("location")

	// call create room db function, params are user and location
	roomID, err := DB.CreateRoom(user, location)
	//if there was error report it
	if err != nil {
		// write error code
		response.WriteHeader(http.StatusInternalServerError)
		// write error in response
		json.NewEncoder(response).Encode(struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	// send back a successful created code
	response.WriteHeader(http.StatusCreated)

	// write list of restauraunts and room ID into response
	json.NewEncoder(response).Encode(struct {
		Restauraunts []string `json:"restauraunts"`
		RoomID       string   `json:"roomID"`
	}{
		Restauraunts: restaurant[:],
		RoomID:       roomID,
	})
	//TODO: refactor so we have a model to return so this is cleaned up a bit, instead of being anonamyous
}

// wat do: list all current rooms
// wat need: nothing
// return: list of all rooms
func GetRooms(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])

	// call get all rooms db function
	// reurn result
}

// wat do: remove user from room
// wat need: user id, room id they're leaving from
// return: nothing
func LeaveRoom(response http.ResponseWriter, request *http.Request) {
	// get 'user' query param
	user := request.URL.Query().Get("user")
	// get 'room' query param
	room := request.URL.Query().Get("room")

	// call remove user db function, params are user and room
	DB.LeaveRoom(user, room)
}

// wat do: increment vote for selected restaurant
// wat need: restaurant id
// return: true or false for found
func SwipeRight(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])

	// get 'restaurant' query param

	// call db add vote function

	// call helper function to determine if result is found yet
}

// wat do: nothing
// wat need:
// return: true or false for found
func SwipeLeft(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])

	// call helper function to determine if result is found yet
}

// async wait function (Peyton will do this )
func FinalCountdown(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])
	// hold person here until ready to return with result

	// get 'room' query param
	room := request.URL.Query().Get("room")

	winner, err := DB.FinalCountdown(room)
	//if there was error report it
	if err != nil {
		// write error code
		response.WriteHeader(http.StatusInternalServerError)
		// write error in response
		json.NewEncoder(response).Encode(struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	// send back a successful OK code
	response.WriteHeader(http.StatusOK)

	// write list of restauraunts and room ID into response
	json.NewEncoder(response).Encode(struct {
		Winner string `json:"winner"`
	}{
		Winner: winner,
	})
	//TODO: refactor so we have a model to return so this is cleaned up a bit, instead of being anonamyous
}
