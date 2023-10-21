package services

import (
	"chatjobsity/env"
	"chatjobsity/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomService interface {
	Messages(roomId string) ([]*models.Message, error)
	Rooms() ([]*models.Room, error)
}

type roomService struct {
	db  *mongo.Client
	env env.EnvApp
}

func NewRoomService(db *mongo.Client, env env.EnvApp) *roomService {
	return &roomService{db: db, env: env}
}

func (s *roomService) Rooms() ([]*models.Room, error) {
	var rooms []*models.Room
	collection := s.db.Database(s.env.MongoDbName).Collection("rooms")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var room models.Room
		if err := cursor.Decode(&room); err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
func (s *roomService) Messages(roomId string) ([]*models.Message, error) {
	var messages []*models.Message

	collection := s.db.Database("jobsity").Collection("messages")
	filter := bson.M{"roomId": roomId}
	options := options.Find()
	options.SetSort(bson.D{{"date", -1}})

	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append([]*models.Message{&message}, messages...)
		if len(messages) >= 50 {
			break
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
