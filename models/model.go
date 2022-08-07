package model

type Room struct {
	RoomID       string   `json:"room_id"`
	Users        []string `json:"users"`
	Location     string   `json:"location"`
	Restauraunts []string `json:"restauraunts"`
	Votes        []int    `json:"votes"`
}
