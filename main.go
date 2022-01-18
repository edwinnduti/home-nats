/*
 ENTRY POINT
*/

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/natsConn"
	"github.com/edwinnduti/home-nats/reqres"
	"github.com/edwinnduti/home-nats/router"
)

func main() {
	// declare nats server struct
	var natsServer reqres.NatsServer

	// error handling
	var err error

	// writer for http response
	var w http.ResponseWriter

	// set nats server connection
	if natsServer.Server.Nc, err = natsConn.GetNatsConnection(); err != nil {
		lib.CheckErr(w, "NATS not Connecting error", http.StatusInternalServerError, "Nats Connection Error", err)
	}

	// subscribe to nats subject welcome
	natsServer.WelcomeReply("welcome")
	natsServer.PostHouseReply("addHouse")
	natsServer.GetHouseReply("getHouse")
	natsServer.GetAllHousesReply("getAllHouses")
	natsServer.DeleteHouseReply("deleteHouse")
	natsServer.PutHouseReply("putHouse")

	// get router and give it the service connection
	r := router.Route(natsServer.Server)

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
	consts.InfoLogger.Printf("Listening to logs on Port %v...", PORT)
	webServer.ListenAndServe()
}
