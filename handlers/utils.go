package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type SMS struct {
	ID         string    `json:"id"`
	Recipient  string    `json:"recipient"`
	Content    string    `json:"content"`
	TimeToSend time.Time `json:"timeToSend"`
}

func ProcessSMS(r *http.Request, id string) (*SMS, error) {
	log.Printf("Processing SMS:\t%s", id)

	// Parse request data (e.g., recipient phone number, message content)
	var sms *SMS
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&sms)
	if err != nil {
		if err == io.EOF {
			// Empty body error handling if necessary
			return nil, fmt.Errorf("empty request body")
		}
		return nil, err
	}

	sms.ID = id

	return sms, nil
}
