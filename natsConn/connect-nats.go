package natsConn

import (
	"log"

	"github.com/edwinnduti/gone-nats/secrets"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

// init function
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting ENV values: %v\n", err)
	}

	// loaded ENV variables
	log.Printf("Success getting the env values")
}

// ConnectNats function
func GetNatsConnection() (*nats.Conn, error) {
	// Connect NATS to nats.DefaultURL
	natsconn, err := nats.Connect(secrets.Configs.NatsUrl)
	if err != nil {
		return nil, err
	}

	// Check if connection is alive
	log.Println("NATS Connection Established on :", natsconn.ConnectedUrl())
	return natsconn, nil
}
