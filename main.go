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

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/middleware"
	"github.com/edwinnduti/home-nats/models"
	"github.com/edwinnduti/home-nats/natsConn"
	"github.com/edwinnduti/home-nats/router"
	"github.com/nats-io/nats.go"
)

func main() {
	// declare server
	var server middleware.Server

	// error handling
	var err error

	// writer for http response
	var w http.ResponseWriter

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

	// subscribe to nats subject welcome
	server.Nc.Subscribe("welcome", func(msg *nats.Msg) {
		message := models.Info{
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
			w.WriteHeader(http.StatusOK)
			// send response
			json.NewEncoder(w).Encode(response)
		}

		// publish message
		server.Nc.Publish(msg.Reply, byteMsg)

	})

	// subscribe to nats subject add_house
	server.Nc.Subscribe("addHouse", func(msg *nats.Msg) {

		// new empty house
		house := new(models.House)
		// boolean alternative
		var parkingLot, hasWifi int = 0, 0

		// unmarshal incoming data to house
		err := json.Unmarshal(msg.Data, &house)
		if err != nil {
			// log error
			log.Printf("Marshal error: %v", err)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "Json Marshal Error",
			}
			//response code
			w.WriteHeader(http.StatusOK)
			// send response
			json.NewEncoder(w).Encode(response)
		}

		// SHIFT BOOLEAN TO 1,0
		if house.HasWifi {
			hasWifi = 1
		}
		if house.ParkingLot {
			parkingLot = 1
		}

		//connect to database
		db, err := lib.ConnectDB()
		if err != nil {
			// log error
			log.Printf("DB connection Error: %v", err)

			// response code
			w.WriteHeader(http.StatusOK)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "DB Connection Error",
			}

			// write response
			json.NewEncoder(w).Encode(response)
		}

		// close db connection
		defer db.Close()

		// add house to db
		result, err := db.Exec("INSERT INTO house (house_number, name, location, cost, year_of_construction, category, area, perimeter, number_of_floors, number_of_bedrooms, construction_material, roofing_type, fencing_type, parking_lot, source_of_water_supply, has_wifi) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", house.HouseNumber, house.Name, house.Location, house.Cost, house.YearOfConstruction, house.Category, house.Area, house.Perimeter, house.NumberOfFloors, house.NumberOfBedrooms, house.ConstructionMaterial, house.RoofingType, house.FenceType, parkingLot, house.SourceOfWaterSupply, hasWifi)
		if err != nil {
			// log error
			log.Printf("House Insertion Error: %v", err)

			// response code
			w.WriteHeader(http.StatusOK)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "House Insertion Error",
			}

			// write response
			json.NewEncoder(w).Encode(response)
		}

		// get last inserted id
		id, err := result.LastInsertId()
		if err != nil {
			// log error
			log.Printf("Last Id Retrival Error: %v", err)

			// response code
			w.WriteHeader(http.StatusOK)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "Last Id Retrival Failed",
			}

			// write response
			json.NewEncoder(w).Encode(response)
		}

		// return response to user
		idMsg := fmt.Sprintf("House ID %v inserted successfully", id)
		res := models.NewResponse{
			Code:    http.StatusOK,
			Message: idMsg,
		}

		// marshal res data
		resInBytes, err := json.Marshal(res)
		if err != nil {
			// log error
			log.Printf("Marshal error: %v", err)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "Json Marshal Error",
			}
			//response code
			w.WriteHeader(http.StatusOK)
			// send response
			json.NewEncoder(w).Encode(response)
		}

		// publish message
		server.Nc.Publish(msg.Reply, resInBytes)

	})

	// subscribe to nats subject getHouse
	server.Nc.Subscribe("getHouse", func(msg *nats.Msg) {

		idHolder := new(models.Info)
		err := json.Unmarshal(msg.Data, &idHolder)
		if err != nil {
			// log error
			log.Printf("Marshal error: %v", err)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "Json Marshal Error",
			}
			//response code
			w.WriteHeader(http.StatusOK)
			// send response
			json.NewEncoder(w).Encode(response)
		}

		//connect to database
		db, err := lib.ConnectDB()
		if err != nil {
			// log error
			log.Printf("DB connection Error: %v", err)

			// response code
			w.WriteHeader(http.StatusOK)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "DB Connection Error",
			}

			// write response
			json.NewEncoder(w).Encode(response)
		}

		// close db connection
		defer db.Close()

		// empty house struct
		house := new(models.House)

		// int holders
		var hasWifi, hasParkingLot int

		// query db using db
		row := db.QueryRow("SELECT * FROM house WHERE id = ?", idHolder.Message)

		// unmarshal row data to house struct
		err = row.Scan(
			&house.Id,
			&house.HouseNumber,
			&house.Name,
			&house.Location,
			&house.Cost,
			&house.YearOfConstruction,
			&house.Category,
			&house.Area,
			&house.Perimeter,
			&house.NumberOfFloors,
			&house.NumberOfBedrooms,
			&house.ConstructionMaterial,
			&house.RoofingType,
			&house.FenceType,
			&hasParkingLot,
			&house.SourceOfWaterSupply,
			&hasWifi,
		)
		if err != nil {
			// log error
			log.Printf("Row Scan Error: %v", err)

			// response code
			w.WriteHeader(http.StatusOK)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "Row Scan Error",
			}

			// write response
			json.NewEncoder(w).Encode(response)
		}

		// convert ints to boolean for user in haswifi
		if hasWifi == 0 {
			house.HasWifi = false
		} else {
			house.HasWifi = true
		}

		// convert ints to boolean for user in parking lot
		if hasParkingLot == 0 {
			house.ParkingLot = false
		} else {
			house.ParkingLot = true
		}

		// convert house struct type to byte data
		houseInBytes, err := json.Marshal(house)
		if err != nil {
			// log error
			log.Printf("Marshal House error: %v", err)

			// return error
			response := models.NewResponse{
				Code:    http.StatusInternalServerError,
				Message: "Json Marshal Error",
			}
			//response code
			w.WriteHeader(http.StatusOK)
			// send response
			json.NewEncoder(w).Encode(response)
		}

		// publish message
		server.Nc.Publish(msg.Reply, houseInBytes)

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
	consts.InfoLogger.Printf("Listening to logs on Port %v...", PORT)
	webServer.ListenAndServe()
}
