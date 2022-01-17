package router

import (
	"github.com/edwinnduti/gone-nats/middleware"
	"github.com/gorilla/mux"
)

func Route(m middleware.Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", m.WelcomeHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/house/{house_id}", m.GetHouseHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/add-house", m.PostHouseHandler).Methods("POST", "OPTIONS")

	return r

}
