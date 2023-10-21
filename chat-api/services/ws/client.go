package ws

import (
	"chatjobsity/env"
	"chatjobsity/models"
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

// Client is a single chatting user

type Client struct {
	name string
	// socket is the web socket for this client
	socket *websocket.Conn
	// send is a channel on which messages are sent
	send chan []byte
	// room is the room this client is chatting in
	room *Room
	// db is the database connection
	db *mongo.Client
	// env is the environment variables
	env env.EnvApp
}

// NewClient makes a new client that is ready to chat

func NewClient(name string, socket *websocket.Conn, room *Room, db *mongo.Client, env env.EnvApp) *Client {
	return &Client{
		name:   name,
		socket: socket,
		send:   make(chan []byte),
		room:   room,
		db:     db,
		env:    env,
	}
}

// Read allows a client to read messages from the socket

func (c *Client) Read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		go c.InsertMessage(msg)
		c.room.forward <- msg
	}
}

// Write allows a client to write messages to the socket

func (c *Client) Write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

func (c *Client) InsertMessage(msg []byte) error {
	var message models.Message
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(message.Text, "/stock=") {
		collection := c.db.Database(c.env.MongoDbName).Collection("messages")
		newMessage := models.Message{
			RoomId: c.room.roomId,
			Sender: models.Sender{Id: message.Sender.Id, Username: message.Sender.Username},
			Text:   message.Text,
			Date:   time.Now(),
		}
		_, err = collection.InsertOne(context.Background(), newMessage)
	}
	return err
}
