package rabbitmq

import (
	"chatjobsity/env"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(env env.EnvApp) (*RabbitMQ, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", env.RmqUser,
		env.RmqPass, env.RmqHost, env.RmqPort))
	if err != nil {
		fmt.Printf("Error connecting to RabbitMQ! %v\n", err)
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Error opening a channel! %v\n", err)
		return nil, err
	}
	_, err = ch.QueueDeclare(
		env.BotQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Error declaring bot queue channel! %v\n", err)
		return nil, err
	}
	return &RabbitMQ{conn: conn, ch: ch}, nil
}

func (r *RabbitMQ) GetChannel() *amqp.Channel {
	return r.ch
}

func (r *RabbitMQ) GetConnection() *amqp.Connection {
	return r.conn
}
