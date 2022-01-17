package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/edwinnduti/gone-nats/consts"
	"github.com/edwinnduti/gone-nats/models"
)

func (srv Server) PostHouseHandler(w http.ResponseWriter, r *http.Request) {

	// empty house
	house := new(models.House)

	// get data from user
	err := json.NewDecoder(r.Body).Decode(&house)
	if err != nil {
		// log error
		log.Printf("Posthouse Decode Error: %v", err)

		// response code
		w.WriteHeader(http.StatusOK)

		// return error
		response := models.NewResponse{
			Code:    http.StatusInternalServerError,
			Message: "House Creation Error",
		}

		// write response
		json.NewEncoder(w).Encode(response)
	}

	// marshal house into bytes
	houseInBytes, err := json.Marshal(&house)
	if err != nil {
		// log error
		log.Printf("Marshal house Error: %v", err)

		// response code
		w.WriteHeader(http.StatusOK)

		// return error
		response := models.NewResponse{
			Code:    http.StatusInternalServerError,
			Message: "Marshal house Error",
		}

		// write response
		json.NewEncoder(w).Encode(response)
	}

	// NATS request method
	msg, err := srv.Nc.Request("addHouse", houseInBytes, time.Second)
	if err != nil {
		// log error
		log.Printf("NATS addHouse request Error: %v", err)

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
	info := new(models.Info)
	err = json.Unmarshal(msg.Data, &info)
	if err != nil {
		// log error
		log.Printf("NATS Unmarshal for Info.Message Error: %v", err)

		// response code
		w.WriteHeader(http.StatusOK)

		// return error
		response := models.NewResponse{
			Code:    http.StatusInternalServerError,
			Message: "NATS Unmarshal Error",
		}

		// write response
		json.NewEncoder(w).Encode(response)
	}

	// create log for request performed
	consts.InfoLogger.Printf("%v : %v OK / HOUSE ID %v CREATED", http.MethodPost, http.StatusCreated, info.Message)

	// write to rw
	w.Write(msg.Data)

}
