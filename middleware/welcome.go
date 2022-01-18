package middleware

import (
	"net/http"
	"time"

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/lib"
)

// welcome message
func (srv Server) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	// status code
	w.Header().Set("Content-Type", "application/json")

	// NATS request method
	msg, err := srv.Nc.Request("welcome", nil, time.Second)
	if err != nil {
		lib.CheckErr(w, "NATS request Error", http.StatusInternalServerError, "NATS request Error", err)
	}

	// NEW USER
	consts.InfoLogger.Printf("%v : %v OK / WELCOME HANDLER REQUESTED", http.MethodGet, http.StatusOK)

	// write to rw
	w.Write(msg.Data)

}
