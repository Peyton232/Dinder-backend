package main

import (
	"fmt"
	"net/http"

	"github.com/henlegay/diner-api/database"
)

var DB *database.DB

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
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:]) // remove this line

	// get 'user' query param
	// get 'room' query param

}

// wat do: creates a room
// wat need: location, user
// return: list of restaurants, room ID
func CreateRoom(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])
}

// wat do: list all current rooms
// wat need: nothing
// return: list of all rooms
func GetRooms(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])
}

// wat do: remove user from room
// wat need: user id, room id they're leaving from
// return: nothing
func LeaveRoom(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])
}

// wat do: increment vote for selected restaurant
// wat need: restaurant id
// return: true or false for found
func SwipeRight(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])
}

// wat do: nothing
// wat need:
// return: true or false for found
func SwipeLeft(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])
}
