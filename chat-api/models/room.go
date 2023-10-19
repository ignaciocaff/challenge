package models

type Room struct {
	Id   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
