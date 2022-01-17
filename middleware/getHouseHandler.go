package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/models"
	"github.com/gorilla/mux"
)

func (srv Server) GetHouseHandler(w http.ResponseWriter, r *http.Request) {
	mapForId := mux.Vars(r)
	id := mapForId["house_id"]

	idMsg := models.Info{
		Message: id,
	}

	// marshal idMsg to bytes type
	idMsgInBytes, err := json.Marshal(idMsg)
	if err != nil {
		// log error
		log.Printf("Marshal IdMsg Error: %v", err)

		// response code
		w.WriteHeader(http.StatusOK)

		// return error
		response := models.NewResponse{
			Code:    http.StatusInternalServerError,
			Message: "Marshal Id Error",
		}

		// write response
		json.NewEncoder(w).Encode(response)
	}

	// make nats request-reply
	msg, err := srv.Nc.Request("getHouse", idMsgInBytes, time.Second)
	if err != nil {
		// log error
		log.Printf("NATS getHouse request Error: %v", err)

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
	consts.InfoLogger.Printf("%v : %v OK / HOUSE ID %v REQUESTED", http.MethodGet, http.StatusOK, id)

	// write to rw
	w.Write(msg.Data)
}
