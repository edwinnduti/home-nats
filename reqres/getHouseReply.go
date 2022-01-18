package reqres

import (
	"encoding/json"
	"net/http"

	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/models"
	"github.com/nats-io/nats.go"
)

func (s NatsServer) GetHouseReply(subject string) {
	// subscribe to nats subject getHouse
	s.Server.Nc.Subscribe("getHouse", func(msg *nats.Msg) {

		idHolder := new(models.Info)
		if err := json.Unmarshal(msg.Data, &idHolder); err != nil {
			lib.CheckErr(w, "Marshal to info struct error", http.StatusInternalServerError, "Json Marshal Error", err)
		}

		//connect to database
		db, err := lib.ConnectDB()
		if err != nil {
			lib.CheckErr(w, "DB connection Error", http.StatusInternalServerError, "DB Connection Error", err)
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
		if err = row.Scan(
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
			lib.CheckErr(w, "Row Scan in GET house Error", http.StatusInternalServerError, "Row Scan Error", err)
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
			lib.CheckErr(w, "Marshal House error", http.StatusInternalServerError, "Json Marshal Error", err)
		}

		// publish message
		s.Server.Nc.Publish(msg.Reply, houseInBytes)

	})

}
