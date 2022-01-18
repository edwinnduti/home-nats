package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/models"
)

func (srv Server) PostHouseHandler(w http.ResponseWriter, r *http.Request) {

	// empty house
	house := new(models.House)

	// get data from user
	err := json.NewDecoder(r.Body).Decode(&house)
	if err != nil {
		lib.CheckErr(w, "Posthouse Decode Error", http.StatusInternalServerError, "House Creation Error", err)
	}

	// marshal house into bytes
	houseInBytes, err := json.Marshal(&house)
	if err != nil {
		lib.CheckErr(w, "Marshal house Error", http.StatusInternalServerError, "Marshal house Error", err)
	}

	// NATS request method
	msg, err := srv.Nc.Request("addHouse", houseInBytes, time.Second)
	if err != nil {
		lib.CheckErr(w, "NATS addHouse request Error", http.StatusInternalServerError, "NATS request Error", err)
	}

	// empty info struct
	info := new(models.Info)
	err = json.Unmarshal(msg.Data, &info)
	if err != nil {
		lib.CheckErr(w, "NATS Unmarshal for Info.Message Error", http.StatusInternalServerError, "NATS Unmarshal Error", err)
	}

	// create log for request performed
	consts.InfoLogger.Printf("%v : %v OK / HOUSE ID %v CREATED", http.MethodPost, http.StatusCreated, info.Message)

	// write to rw
	w.Write(msg.Data)

}
