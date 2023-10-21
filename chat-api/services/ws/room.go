package ws

import (
	"chatjobsity/env"
	"chatjobsity/models"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Room represents a chat room

type Room struct {
	roomId string
	// forward is a channel that holds incoming messages that should be forwarded to the other clients
	forward chan []byte
	// join is a channel for clients wishing to join the room
	join chan *Client
	// leave is a channel for clients wishing to leave the room
	leave chan *Client
	// clients holds all current clients in this room
	clients map[*Client]bool
	// rabbitmq channel
	ch *amqp.Channel
	// env is the environment variables
	env env.EnvApp
}

// NewRoom makes a new room that is ready to go

func NewRoom(roomId string, ch *amqp.Channel, env env.EnvApp) *Room {
	return &Room{
		roomId:  roomId,
		forward: make(chan []byte),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
		ch:      ch,
		env:     env,
	}
}

// Start handles the connection of a client to a room

func (r *Room) Start() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			var message models.Message
			err := json.Unmarshal(msg, &message)
			if err != nil {
				fmt.Printf("Error unmarshalling message %v\n", err)
			}
			if strings.HasPrefix(message.Text, "/stock=") {
				// Go to bot stock
				stockSymbol := strings.TrimPrefix(message.Text, "/stock=")
				newPost := models.Post{
					Text:   stockSymbol,
					RoomId: r.roomId,
				}
				jsonBot, _ := json.Marshal(newPost)

				err := r.ch.PublishWithContext(
					context.Background(),
					"",
					r.env.BotQueue,
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        jsonBot,
					})
				if err != nil {
					fmt.Printf("Error on RabbitMQ publish: %v \n", err)
				}
				go r.startRabbitMQConsumer()
			} else {
				r.forwardMessageToClients(msg)
			}
		}
	}
}

func (r *Room) forwardMessageToClients(msg []byte) {
	for client := range r.clients {
		select {
		case client.send <- msg:
		default:
			delete(r.clients, client)
			close(client.send)
		}
	}
}

func (r *Room) startRabbitMQConsumer() {
	msgs, err := r.ch.Consume(
		r.roomId,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Error on RabbitMQ Consume: %v \n", err)
	}

	for d := range msgs {
		newMessage := models.Message{
			RoomId: r.roomId,
			Sender: models.Sender{Id: "", Username: "Botsity"},
			Date:   time.Now(),
		}
		if string(d.Body) == "N/D" {
			newMessage.Text = "The stock symbol does not exist"
		} else {
			newMessage.Text = string(d.Body)
		}

		jsonData, _ := json.Marshal(newMessage)
		r.forwardMessageToClients(jsonData)
	}
}
