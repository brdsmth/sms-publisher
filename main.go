package main

import (
	"log"
	"net/http"

	"sms-publisher/config"
	"sms-publisher/handlers"
	"sms-publisher/services"
)

func main() {
	var err error

	rabbitMQURL := config.ReadEnv("RABBITMQ_URL")
	// rabbitMQURL := "amqp://guest:guest@localhost:5672/" // local rabbitmq url
	if rabbitMQURL == "" {
		log.Fatal("RABBITMQ_URL environment variable not set")
	}

	services.ConnectToRabbitMQ(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer services.RabbitMQConn.Close()

	// Define routes and handlers for sending SMS messages
	http.HandleFunc("/queue-sms", handlers.QueueSMSHandler)
	http.HandleFunc("/schedule-sms", handlers.ScheduleSMSHandler)

	// Start the HTTP server for the publisher microservice
	log.Println("Server listening on 8080")
	http.ListenAndServe(":8080", nil)
}
