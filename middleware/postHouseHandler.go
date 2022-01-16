package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/edwinnduti/gone-nats/lib"
	"github.com/edwinnduti/gone-nats/models"
)

func PostHouseHandler(w http.ResponseWriter, r *http.Request) {

	// empty house
	house := new(models.House)

	// boolean alternative
	var parkingLot, hasWifi int = 0, 0

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
	idMsg := fmt.Sprintf("%v inserted successfully", id)
	res := models.NewResponse{
		Code:    http.StatusOK,
		Message: idMsg,
	}
	json.NewEncoder(w).Encode(res)

}
