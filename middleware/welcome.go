package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/edwinnduti/home-nats/consts"
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
		consts.ErrorLogger.Printf("NATS request Error: %v", err)

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
	// NEW USER
	consts.InfoLogger.Printf("%v : %v OK / WELCOME HANDLER REQUESTED", http.MethodGet, http.StatusOK)

	// write to rw
	w.Write(msg.Data)

}
