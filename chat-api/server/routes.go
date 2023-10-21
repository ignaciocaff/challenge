package server

import (
	"chatjobsity/handlers"
	"chatjobsity/services"
	"chatjobsity/services/ws"

	"github.com/gin-gonic/gin"
)

func (s *AppServer) routes() {
	healthRoutes(s)
	roomRoutes(s)
	websocketRoutes(s)
	authRoutes(s)
}

func healthRoutes(server *AppServer) {
	server.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func websocketRoutes(server *AppServer) {
	group := server.Engine.Group("/ws")
	service := ws.NewWsServer(server.ctx, server.db, server.config, server.rmq)
	group.GET("/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		service.ServeWs(c.Writer, c.Request, roomId)
	})
}

func authRoutes(server *AppServer) {
	group := server.Engine.Group("/api")
	service := services.New(server.db)
	handler := handlers.New(service, server.sm)
	group.Use(AuthRequired)
	group.POST("/auth", handler.Login)
	group.GET("/auth/me", handler.Me)
	group.GET("/auth/logout", handler.Logout)
}

func roomRoutes(server *AppServer) {
	group := server.Engine.Group("/api")
	service := services.NewRoomService(server.db, server.config)
	handler := handlers.NewRoomHandler(service)
	group.Use(AuthRequired)
	group.GET("/rooms", handler.Rooms)
	group.GET("/rooms/:roomId/messages", handler.Messages)
}
