package reqres

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/models"
	"github.com/nats-io/nats.go"
)

func (s NatsServer) DeleteHouseReply(subject string) {
	// subscribe to nats subject getHouse
	s.Server.Nc.Subscribe("deleteHouse", func(msg *nats.Msg) {

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

		// query db using db
		result, err := db.Exec("DELETE FROM house WHERE id = ?", idHolder.Message)

		// unmarshal row data to house struct
		if err != nil {
			lib.CheckErr(w, "Delete Row in GET house Error", http.StatusInternalServerError, "Row Scan Error", err)
		}

		// convert response struct type to byte data
		var res models.NewResponse
		numberOfRowsAffected, err := result.RowsAffected()
		if err != nil {
			lib.CheckErr(w, "Rows Affected in house Error", http.StatusInternalServerError, "Rows Affected Error", err)
		}
		res.Code = http.StatusOK
		res.Message = fmt.Sprintf("%v House of ID %v Deleted", numberOfRowsAffected, idHolder.Message)

		resInBytes, err := json.Marshal(res)
		if err != nil {
			lib.CheckErr(w, "Marshal House error", http.StatusInternalServerError, "Json Marshal Error", err)
		}

		// publish message
		s.Server.Nc.Publish(msg.Reply, resInBytes)

	})

}
