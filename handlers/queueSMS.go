package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sms-publisher/services"

	gonanoid "github.com/matoous/go-nanoid"
	"github.com/streadway/amqp"
)

var (
	SCHEDULE_SMS_QUEUE = "SCHEDULED_SMS_QUEUE"
	SMS_QUEUE          = "SMS_QUEUE"
)

func QueueSMSHandler(w http.ResponseWriter, r *http.Request) {
	id, err := gonanoid.Nanoid(6)
	if err != nil {
		log.Printf("error adding nanoid: %s", err)
		return
	}

	log.Printf("Queueing SMS:\t%s", id)

	// Process the incoming request
	sms, err := ProcessSMS(r, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.RabbitMQConnMutex.Lock()
	defer services.RabbitMQConnMutex.Unlock()

	if services.RabbitMQConn == nil {
		log.Println("RabbitMQ connection is not available")
		http.Error(w, "RabbitMQ connection is not available", http.StatusInternalServerError)
		return
	}

	// Create a channel
	ch, err := services.RabbitMQConn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return
	}
	defer ch.Close()

	// Declare the channel
	_, err = ch.QueueDeclare(SCHEDULE_SMS_QUEUE, false, false, false, false, nil)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return
	}

	// Send SMS requests to the queue
	smsJSON, err := json.Marshal(sms)
	log.Println(string(smsJSON))
	if err != nil {
		http.Error(w, "Failed to marshal sms", http.StatusInternalServerError)
		log.Printf("Failed to marshal sms: %v", err)
		return
	}

	err = ch.Publish("", SMS_QUEUE, false, false, amqp.Publishing{
		ContentType: "text/plain",
		// Body:        []byte(body),
		Body: smsJSON,
	})
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return
	}

	// Respond to the HTTP request indicating success
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "OK")
	log.Printf("Queued SMS:\t\t%s", id)
}
