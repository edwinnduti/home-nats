/*
 ENTRY POINT
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edwinnduti/gone-nats/middleware"
	"github.com/edwinnduti/gone-nats/models"
	"github.com/edwinnduti/gone-nats/natsConn"
	"github.com/edwinnduti/gone-nats/router"
	"github.com/nats-io/nats.go"
)

func main() {
	// declare server
	var server middleware.Server

	// error handling
	var err error

	// writer for http response
	var rw http.ResponseWriter

	// set nats server connection
	server.Nc, err = natsConn.GetNatsConnection()
	if err != nil {
		// log error
		log.Printf("NATS not Connecting error: %v", err)

		// return error
		response := models.NewResponse{
			Code:    http.StatusInternalServerError,
			Message: "Nats Connection Error",
		}

		// send response
		json.NewEncoder(os.Stdout).Encode(response)
		return
	}

	// subscribe to nats
	server.Nc.Subscribe("welcome", func(msg *nats.Msg) {
		message := models.NewUser{
			Message: "Welcome to my NATS API!",
		}

		byteMsg, err := json.Marshal(message)
		if err != nil {
			// log error
			log.Printf("Marshal error: %v", err)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "Json Marshal Error",
			}
			//response code
			rw.WriteHeader(http.StatusOK)
			// send response
			json.NewEncoder(rw).Encode(response)
		}

		// publish message
		server.Nc.Publish(msg.Reply, byteMsg)

	})

	// get router and give it the service connection
	r := router.Route(server)

	// port number
	var PORT string = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// simple web server
	webServer := &http.Server{
		Addr:    fmt.Sprint(":", PORT),
		Handler: r,
	}

	// log info and listen to connection
	log.Printf("Listening to logs on Port %v\n", PORT)
	webServer.ListenAndServe()
}
