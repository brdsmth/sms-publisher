package services

import (
	"log"
	"sync"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var RabbitMQConnMutex sync.Mutex

func ConnectToRabbitMQ(rabbitMQURL string) {
	var err error
	RabbitMQConn, err = amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	log.Println("Connected to RabbitMQ successfully")
}
