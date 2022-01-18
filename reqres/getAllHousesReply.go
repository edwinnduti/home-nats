package reqres

import (
	"encoding/json"
	"net/http"

	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/models"
	"github.com/nats-io/nats.go"
)

func (s NatsServer) GetAllHousesReply(subject string) {
	// subscribe to nats subject getAllHouses
	sub, err := s.Server.Nc.Subscribe("getAllHouses", func(msg *nats.Msg) {

		//connect to database
		db, err := lib.ConnectDB()
		if err != nil {
			lib.CheckErr(w, "DB connection Error", http.StatusInternalServerError, "DB Connection Error", err)
		}

		// close db connection
		defer db.Close()

		// // An houses slice to hold data from returned rows.
		var houses []models.House

		// int holders
		var hasWifi, hasParkingLot int

		// query db using db
		rows, err := db.Query("SELECT * FROM house")
		if err != nil {
			lib.CheckErr(w, "Query All houses error", http.StatusInternalServerError, "Query Error", err)
		}

		defer rows.Close()

		// unmarshal row data to house struct
		for rows.Next() {
			var house models.House

			if err = rows.Scan(
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
			); err != nil {
				// If there is an error while iterating, stop the loop and print error
				lib.CheckErr(w, "Row Scan in GET All houses Error", http.StatusInternalServerError, "Row Scan Error", err)

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

			// append house to houses slice
			houses = append(houses, house)
		}

		// convert houses slice type to byte data
		housesInBytes, err := json.Marshal(houses)
		if err != nil {
			lib.CheckErr(w, "Marshal All houses error", http.StatusInternalServerError, "Json Marshal Error", err)
		}

		// publish message
		s.Server.Nc.Publish(msg.Reply, housesInBytes)

	})

	// check error from subscribe
	if err != nil {
		lib.CheckErr(w, "Subscribe to getAllHouses Error", http.StatusInternalServerError, "Subscribe Error", err)
	}

	// unsubscribe from nats subject getAllHouses
	if err := sub.Unsubscribe(); err != nil {
		lib.CheckErr(w, "Cannot Unsubscribe to NATS error", http.StatusInternalServerError, "NATS Unsubscribe Error", err)
	}
}
