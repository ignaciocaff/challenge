package server

import (
	"context"
	"fmt"

	"chatjobsity/env"
	"chatjobsity/services/rabbitmq"
	"chatjobsity/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Starter interface {
	Run()
}

type AppServer struct {
	*gin.Engine
	ctx    context.Context
	db     *mongo.Client
	config env.EnvApp
	sm     *utils.SessionManager
	rmq    *rabbitmq.RabbitMQ
}

func (s *AppServer) configure() {
	s.Engine.Use(gin.Recovery())
	s.Engine.SetTrustedProxies([]string{"*"})
	s.Engine.Use(func(c *gin.Context) {
		c.Set("session_manager", s.sm)
		c.Next()
	})
}

func New(ctx context.Context, db *mongo.Client, config env.EnvApp, rmq *rabbitmq.RabbitMQ) Starter {
	gin.SetMode(config.GinMode)
	server := &AppServer{
		gin.Default(),
		ctx,
		db,
		config,
		utils.NewSessionManager(),
		rmq,
	}

	server.configure()
	server.routes()
	return server
}

func (s *AppServer) Run() {
	fmt.Printf("Server running on port %s\n", s.config.Port)
	s.Engine.Run(":" + s.config.Port)
}
