package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/edwinnduti/home-nats/models"
)

// welcome message
func (srv Server) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	// status code
	w.Header().Set("Content-Type", "application/json")

	// NATS request method
	msg, err := srv.Nc.Request("welcome", nil, time.Second)
	if err != nil {
		// log error
		log.Printf("NATS request Error: %v", err)

		// response code
		w.WriteHeader(http.StatusOK)

		// return error
		response := models.NewResponse{
			Code:    http.StatusInternalServerError,
			Message: "NATS request Error",
		}

		// write response
		json.NewEncoder(w).Encode(response)
	}
	log.Println(string(msg.Data))

	// write to rw
	w.Write(msg.Data)

}
