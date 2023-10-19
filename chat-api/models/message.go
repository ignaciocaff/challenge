package models

import "time"

type Message struct {
	Id     string    `json:"id" bson:"_id,omitempty"`
	Text   string    `json:"text" bson:"text"`
	Sender Sender    `json:"sender"  bson:"sender"`
	RoomId string    `json:"roomId" bson:"roomId"`
	Date   time.Time `json:"date" bson:"date"`
}

type Sender struct {
	Username string `json:"username" bson:"username"`
	Id       string `json:"id" bson:"id"`
}

type Post struct {
	Text   string `json:"text"`
	RoomId string `json:"roomId"`
}
