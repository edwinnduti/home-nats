package reqres

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/models"
	"github.com/nats-io/nats.go"
)

func (s NatsServer) PutHouseReply(subject string) {
	// subscribe to nats subject getHouse
	s.Server.Nc.Subscribe("putHouse", func(msg *nats.Msg) {

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

		// new empty house
		house := new(models.House)
		// boolean alternative
		var parkingLot, hasWifi int = 0, 0

		// SHIFT BOOLEAN TO 1,0
		if house.HasWifi {
			hasWifi = 1
		}
		if house.ParkingLot {
			parkingLot = 1
		}

		// update db using db.exec
		result, err := db.Exec("UPDATE house set house_number = ?, name = ?, location = ?, cost = ?, year_of_construction = ?, category = ?, area = ?, perimeter = ?, number_of_floors = ?, number_of_bedrooms = ?, construction_material = ?, roofing_type = ?, fencing_type = ?, parking_lot = ?, source_of_water_supply = ?, has_wifi = ? WHERE id = ?", house.HouseNumber, house.Name, house.Location, house.Cost, house.YearOfConstruction, house.Category, house.Area, house.Perimeter, house.NumberOfFloors, house.NumberOfBedrooms, house.ConstructionMaterial, house.RoofingType, house.FenceType, parkingLot, house.SourceOfWaterSupply, hasWifi, idHolder.Message)

		// unmarshal row data to house struct
		if err != nil {
			lib.CheckErr(w, "Update Row in PUT house Error", http.StatusInternalServerError, "Row Exec Error", err)
		}

		// convert response struct type to byte data
		var res models.NewResponse
		numberOfRowsAffected, err := result.RowsAffected()
		if err != nil {
			lib.CheckErr(w, "Rows Updated in house Error", http.StatusInternalServerError, "Rows Updated Error", err)
		}
		res.Code = http.StatusOK
		res.Message = fmt.Sprintf("%v House of ID %v Updated", numberOfRowsAffected, idHolder.Message)

		resInBytes, err := json.Marshal(res)
		if err != nil {
			lib.CheckErr(w, "Marshal House error", http.StatusInternalServerError, "Json Marshal Error", err)
		}

		// publish message
		s.Server.Nc.Publish(msg.Reply, resInBytes)

	})

}
