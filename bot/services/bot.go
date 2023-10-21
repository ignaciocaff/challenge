package services

import (
	"botjobsity/env"
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Bot struct {
	Env     env.EnvApp
	Forever chan struct{}
}

type Post struct {
	Text   string `json:"text"`
	RoomId string `json:"roomId"`
}

func NewBot(env env.EnvApp) *Bot {
	return &Bot{Env: env}
}

func (b *Bot) Connect() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", b.Env.RmqUser, b.Env.RmqPass, b.Env.RmqHost, b.Env.RmqPort))
	if err != nil {
		fmt.Printf("Failed to connect to RabbitMQ.\n")
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Failed to open a channel....\n")
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		b.Env.BotQueue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		fmt.Printf("Failed to declare queue %s.\n", b.Env.BotQueue)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Printf("Failed to register a consumer. Retrying ...\n")
	}
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		fmt.Printf("Received a message: %s", d.Body)
		var post Post
		err := json.Unmarshal(d.Body, &post)
		if err != nil {
			fmt.Printf("Received bad response : %s %v ", string(d.Body), err)
			return
		}
		go b.createRoomQueue(post.RoomId, ch)
		go b.handleMessage(post, ch)
	}

}

func (b *Bot) handleMessage(post Post, ch *amqp.Channel) {
	data, err := b.FetchFile(post.Text)
	if err != nil {
		fmt.Printf("Error fetching file: %v", err)
		b.sendMessageToRoom(post.RoomId, "Error fetching file", ch)
		return
	}
	msg, err := b.Process(data, post.Text)
	if err != nil {
		fmt.Printf("Error processing file: %v", err)
		b.sendMessageToRoom(post.RoomId, "Error processing file", ch)
		return
	}
	fmt.Printf("Message to send: %s", msg)
	b.sendMessageToRoom(post.RoomId, msg, ch)
}

func (b *Bot) createRoomQueue(roomID string, ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		roomID, // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	b.failOnError(err, "Failed to declare a queue")
	fmt.Printf("Queue %s created", q.Name)
}

func (b *Bot) sendMessageToRoom(roomID, message string, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := ch.PublishWithContext(ctx,
		"",     // exchange
		roomID, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	b.failOnError(err, "Failed to publish a message")
}

func (b *Bot) failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}
