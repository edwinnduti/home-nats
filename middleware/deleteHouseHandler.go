package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/models"
	"github.com/gorilla/mux"
)

func (srv Server) DeleteHouseHandler(w http.ResponseWriter, r *http.Request) {
	mapForId := mux.Vars(r)
	id := mapForId["house_id"]

	idMsg := models.Info{
		Message: id,
	}

	// marshal idMsg to bytes type
	idMsgInBytes, err := json.Marshal(idMsg)
	if err != nil {
		lib.CheckErr(w, "Marshal IdMsg Error", http.StatusInternalServerError, "Marshal Id Error", err)
	}

	// make nats request-reply
	msg, err := srv.Nc.Request("deleteHouse", idMsgInBytes, time.Second)
	if err != nil {
		lib.CheckErr(w, "NATS deleteHouse request Error", http.StatusInternalServerError, "NATS request Error", err)
	}

	// create log for request performed
	consts.InfoLogger.Printf("%v : %v OK / HOUSE ID %v DELETED", http.MethodGet, http.StatusOK, id)

	// write to rw
	w.Write(msg.Data)
}
