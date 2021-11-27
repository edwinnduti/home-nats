package secrets

import (
	"os"

	"github.com/edwinnduti/gone-nats/models"
)

// secret keys
var (
	Configs = models.Configs{
		NatsUrl: os.Getenv("NATSURL"),
		Port:    os.Getenv("PORT"),
	}
)
