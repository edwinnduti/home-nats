package middleware

import (
	"net/http"
	"time"

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/lib"
)

// GetAllHousesHandler is a middleware function that handles the request
func (srv Server) GetAllHousesHandler(w http.ResponseWriter, r *http.Request) {

	// make nats request-reply
	msg, err := srv.Nc.Request("getAllHouses", nil, time.Second)
	if err != nil {
		lib.CheckErr(w, "NATS getAllHouses request Error", http.StatusInternalServerError, "NATS request Error", err)
	}

	// create log for request performed
	consts.InfoLogger.Printf("%v : %v OK / ALL HOUSES REQUESTED", http.MethodGet, http.StatusOK)

	// write to rw
	w.Write(msg.Data)
}
