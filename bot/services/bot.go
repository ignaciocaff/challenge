package services

import (
	"botjobsity/env"
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	b.failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	b.failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		b.Env.BotQueue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	b.failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	b.failOnError(err, "Failed to register a consumer")
	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		go b.handleMessage(d, ch)
	}
}

func (b *Bot) handleMessage(d amqp.Delivery, ch *amqp.Channel) {
	log.Printf("Received a message: %s", d.Body)
	var post Post
	err := json.Unmarshal(d.Body, &post)
	if err != nil {
		log.Printf("Received bad response : %s %v ", string(d.Body), err)
		return
	}
	data, err := b.FetchFile(post.Text)
	if err != nil {
		log.Printf("Error fetching file: %v", err)
		return
	}
	msg, err := b.Process(data, post.Text)
	if err != nil {
		log.Printf("Error processing file: %v", err)
		return
	}
	log.Printf("Message to send: %s", msg)

	go func() {
		b.sendMessageToRoom(post.RoomId, msg, ch)
	}()
}

func (b *Bot) sendMessageToRoom(roomID, message string, ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		roomID, // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	b.failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
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
		log.Panicf("%s: %s", msg, err)
	}
}
