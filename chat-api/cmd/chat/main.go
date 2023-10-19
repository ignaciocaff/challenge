package main

import (
	"chatjobsity/database"
	"chatjobsity/env"
	"chatjobsity/server"
	"chatjobsity/services/rabbitmq"

	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	_env := env.GetEnv(".env")

	client, err := database.NewMongo(ctx, _env).Connect()
	if err != nil {
		fmt.Printf("Error connecting to MongoDB %v", err)
		panic(err)
	}
	rabbitmq, err := rabbitmq.NewRabbitMQ(_env)
	defer func() {
		rabbitmq.GetConnection().Close()
		rabbitmq.GetChannel().Close()
	}()
	if err != nil {
		fmt.Printf("Error connecting to RabbitMQ %v", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Printf("Error disconnecting to MongoDB %v", err)
		}
	}()

	server.New(ctx, client, _env, rabbitmq).Run()
}
