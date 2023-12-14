package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/streadway/amqp"
	"sms-publisher/config"
)

// Define a global variable to hold the RabbitMQ connection
var rabbitMQConn *amqp.Connection
var rabbitMQConnMutex sync.Mutex

func connectToRabbitMQ() (*amqp.Connection, error) {
	// Retrieve RabbitMQ connection URL from environment variables for security
	// rabbitMQURL := readEnv("RABBITMQ_URL")
	rabbitMQURL := config.ReadEnv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}
	// rabbitMQURL := "amqp://guest:guest@localhost:5672/" // local rabbitmq url

	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to RabbitMQ successfully")

	return conn, nil
}

func sendSMSHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SMS HANDLER")
	// Parse request data (e.g., recipient phone number, message content)
	// Validate input data

	// Acquire a lock to ensure safe access to the RabbitMQ connection
	rabbitMQConnMutex.Lock()
	defer rabbitMQConnMutex.Unlock()

	// Ensure that the RabbitMQ connection is open
	if rabbitMQConn == nil {
		log.Println("RabbitMQ connection is not available")
		http.Error(w, "RabbitMQ connection is not available", http.StatusInternalServerError)
		return
	}

	// Create a channel
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		// Handle the error and respond accordingly
		return
	}
	defer ch.Close()

	// Declare a queue
	queueName := "SMS_QUEUE"
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		// Handle the error and respond accordingly
		return
	}

	// Send SMS requests to the queue
	message := "THIS IS A TEST SMS MESSAGE"
	err = ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		// Handle the error and respond accordingly
		return
	}

	fmt.Println("SMS request sent to the queue")
	// Respond to the HTTP request indicating success
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "SMS request sent to the queue")
}

func main() {
	// ... Initialize any other necessary components ...
	// Initialize the RabbitMQ connection in the main function
	var err error
	rabbitMQConn, err = connectToRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()

	// Define routes and handlers for sending SMS messages
	http.HandleFunc("/publish-sms", sendSMSHandler)

	// Start the HTTP server for the publisher microservice
	fmt.Println("server listening on 8080")
	http.ListenAndServe(":8080", nil)
}
