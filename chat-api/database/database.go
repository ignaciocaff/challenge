package database

import (
	"chatjobsity/env"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	ctx context.Context
	env env.EnvApp
}

func NewMongo(ctx context.Context, env env.EnvApp) *Mongo {
	return &Mongo{ctx: ctx, env: env}
}

func (m *Mongo) Connect() (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", m.env.MongoUser, m.env.MongoPass, m.env.MongoHost, m.env.MongoPort)
	client, err := mongo.Connect(m.ctx, options.Client().ApplyURI(url))
	if err != nil {
		fmt.Printf("Error connecting to MongoDB! %v\n", err)
		return nil, err
	}
	if err := client.Ping(m.ctx, readpref.Primary()); err != nil {
		fmt.Printf("Error pinging to MongoDB! %v\n", err)
		return nil, err
	}
	fmt.Printf("Connected to MongoDB!\n")
	return client, nil
}
