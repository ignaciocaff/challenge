package services

import (
	"chatjobsity/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService interface {
	Login(username, password string) (*models.User, error)
	SignUp() (string, error)
}

type authService struct {
	db *mongo.Client
}

func New(db *mongo.Client) *authService {
	return &authService{db: db}
}

func (s *authService) Login(username, password string) (*models.User, error) {
	filter := bson.M{
		"username": username,
		"password": password,
	}
	var user models.User
	collection := s.db.Database("jobsity").Collection("users")
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (s *authService) SignUp() (string, error) {
	return "", nil
}
