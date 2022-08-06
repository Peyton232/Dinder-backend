package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/alive", Alive)
	http.HandleFunc("/Join", JoinRoom)
	http.HandleFunc("/Create", CreateRoom)
	http.HandleFunc("/Get", GetRooms)
	http.HandleFunc("/Leave", LeaveRoom)
	http.HandleFunc("/SwipeRight", SwipeRight)
  	http.HandleFunc("/SwipeLeft", SwipeLeft)
	http.ListenAndServe(":42069", nil)
}

func Alive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//wat do: add user to a room in database
//wat need: user id, room id
//return: list of restaurants
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//wat do: creates a room
//wat need: location, user
//return: list of restaurants, room ID
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//wat do: list all current rooms
//wat need: nothing
//return: list of all rooms
func GetRooms(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//wat do: remove user from room
//wat need: user id, room id they're leaving from
//return: nothing
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//wat do: increment vote for selected restaurant
//wat need: restaurant id
//return: true or false for found
func SwipeRight(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//wat do: nothing
//wat need:
//return: true or false for found
func SwipeLeft(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
