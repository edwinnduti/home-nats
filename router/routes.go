package router

import (
	"github.com/edwinnduti/home-nats/middleware"
	"github.com/gorilla/mux"
)

func Route(m middleware.Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", m.WelcomeHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/house/{house_id}", m.GetHouseHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/add-house", m.PostHouseHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/houses", m.GetAllHousesHandler).Methods("GET", "OPTIONS")

	return r

}
