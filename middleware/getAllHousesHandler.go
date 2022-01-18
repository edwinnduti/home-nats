package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/models"
)

// GetAllHousesHandler is a middleware function that handles the request
func (srv Server) GetAllHousesHandler(w http.ResponseWriter, r *http.Request) {

	// make nats request-reply
	msg, err := srv.Nc.Request("getAllHouses", nil, time.Second)
	if err != nil {
		// log error
		consts.ErrorLogger.Printf("NATS getAllHouses request Error: %v", err)

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

	// create log for request performed
	consts.InfoLogger.Printf("%v : %v OK / ALL HOUSES REQUESTED", http.MethodGet, http.StatusOK)

	// write to rw
	w.Write(msg.Data)
}
