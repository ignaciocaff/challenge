package ws

import (
	"chatjobsity/env"
	"chatjobsity/services/rabbitmq"
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsServer struct {
	ctx    context.Context
	db     *mongo.Client
	config env.EnvApp
	rooms  map[string]*Room
	rmq    *rabbitmq.RabbitMQ
}

func NewWsServer(ctx context.Context, db *mongo.Client, config env.EnvApp, rmq *rabbitmq.RabbitMQ) *WsServer {
	return &WsServer{
		ctx:    ctx,
		db:     db,
		config: config,
		rooms:  make(map[string]*Room),
		rmq:    rmq,
	}
}

func (ws *WsServer) ServeWs(w http.ResponseWriter, r *http.Request, roomID string) {
	// Upgrade the HTTP connection to a WebSocket connection
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading connection to WebSocket: %v", err)
		return
	}

	// Check if the room exists on the map, if it doesn't exist, create a new one
	room, exists := ws.rooms[roomID]
	if !exists {
		room = NewRoom(roomID, ws.rmq.GetChannel(), ws.config)
		_, err := room.ch.QueueDeclare(
			roomID, // name
			false,  // durable
			false,  // delete when unused
			false,  // exclusive
			false,  // no-wait
			nil,    // arguments
		)
		if err != nil {
			fmt.Printf("Error declaring queue %v\n", err)
		}
		go room.Start()
		ws.rooms[roomID] = room
	}

	// Continue handling WebSocket communication here
	client := NewClient("clientID", conn, room, ws.db, ws.config)
	room.join <- client
	defer func() { room.leave <- client }()
	go client.Write()
	client.Read()
}
