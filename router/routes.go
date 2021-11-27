package router

import (
	"github.com/edwinnduti/gone-nats/middleware"
	"github.com/gorilla/mux"
)

func Route(m middleware.Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", m.WelcomeHandler).Methods("GET", "OPTIONS")
	// r.HandleFunc("/post", middleware.PostUserHandler).Methods("POST", "OPTIONS")

	return r

}
