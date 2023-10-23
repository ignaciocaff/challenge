package services

import (
	"chatjobsity/env"
	"chatjobsity/models"
	"chatjobsity/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService interface {
	Login(username, password string) (*models.User, error)
	SignUp(user models.User) error
}

type authService struct {
	db  *mongo.Client
	env env.EnvApp
}

func New(db *mongo.Client, env env.EnvApp) *authService {
	return &authService{db: db, env: env}
}

func (s *authService) Login(username, password string) (*models.User, error) {

	filter := bson.M{
		"username": username,
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
	decryptedPwd, err := utils.Decrypt(user.Password, []byte(s.env.EncryptKey))
	if err != nil {
		return nil, err
	}
	if password != decryptedPwd {
		return nil, err
	}

	return &user, nil
}

func (s *authService) SignUp(user models.User) error {
	collection := s.db.Database(s.env.MongoDbName).Collection("users")
	filter := bson.M{
		"username": user.Username,
	}
	var u models.User
	err := collection.FindOne(context.Background(), filter).Decode(&u)

	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if u != (models.User{}) {
		return errors.New("user already exists")
	}
	encrypted, err := utils.Encrypt(user.Password, []byte(s.env.EncryptKey))
	if err != nil {
		return err
	}
	user.Password = encrypted
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
