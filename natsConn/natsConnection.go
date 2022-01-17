package natsConn

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"

	"github.com/edwinnduti/gone-nats/consts"
)

// init function
func init() {
	// START OF APPLICATION
	consts.InfoLogger.Println("Starting the application...")

	// load env file (.env by default)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting ENV values: %v\n", err)
	}

	// log for loaded ENV variables
	consts.InfoLogger.Println("Env Values Acquired Successfully")
}

// ConnectNats function
func GetNatsConnection() (*nats.Conn, error) {
	// Connect NATS to nats.DefaultURL
	natsconn, err := nats.Connect(os.Getenv("NATSURL"))
	if err != nil {
		return nil, err
	}

	// log for if connection is alive
	consts.InfoLogger.Printf("NATS Connection Established on: %v", natsconn.ConnectedUrl())
	return natsconn, nil
}
