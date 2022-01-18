package reqres

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/models"
	"github.com/nats-io/nats.go"
)

func (s NatsServer) PostHouseReply(subject string) {
	// subscribe to nats subject add_house
	sub, err := s.Server.Nc.Subscribe("addHouse", func(msg *nats.Msg) {

		// new empty house
		house := new(models.House)
		// boolean alternative
		var parkingLot, hasWifi int = 0, 0

		// unmarshal incoming data to house
		if err := json.Unmarshal(msg.Data, &house); err != nil {
			lib.CheckErr(w, "Marshal to house error", http.StatusInternalServerError, "Json Marshal Error", err)
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
			lib.CheckErr(w, "DB connection Error", http.StatusInternalServerError, "DB Connection Error", err)
		}

		// close db connection
		defer db.Close()

		// add house to db
		result, err := db.Exec("INSERT INTO house (house_number, name, location, cost, year_of_construction, category, area, perimeter, number_of_floors, number_of_bedrooms, construction_material, roofing_type, fencing_type, parking_lot, source_of_water_supply, has_wifi) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", house.HouseNumber, house.Name, house.Location, house.Cost, house.YearOfConstruction, house.Category, house.Area, house.Perimeter, house.NumberOfFloors, house.NumberOfBedrooms, house.ConstructionMaterial, house.RoofingType, house.FenceType, parkingLot, house.SourceOfWaterSupply, hasWifi)
		if err != nil {
			lib.CheckErr(w, "House Insertion Error", http.StatusInternalServerError, "House Insertion Error", err)
		}

		// get last inserted id
		id, err := result.LastInsertId()
		if err != nil {
			lib.CheckErr(w, "Last Id Retrival Error", http.StatusInternalServerError, "Last Id Retrival Failed", err)
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
			lib.CheckErr(w, "Marshal error", http.StatusInternalServerError, "Json Marshal Error", err)
		}

		// publish message
		s.Server.Nc.Publish(msg.Reply, resInBytes)

	})

	// check error from subscribe
	if err != nil {
		lib.CheckErr(w, "Subscribe to postHouse Error", http.StatusInternalServerError, "Subscribe Error", err)
	}

	// unsubscribe from nats subject postHouse
	if err := sub.Unsubscribe(); err != nil {
		lib.CheckErr(w, "Cannot Unsubscribe to NATS error", http.StatusInternalServerError, "NATS Unsubscribe Error", err)
	}

}
