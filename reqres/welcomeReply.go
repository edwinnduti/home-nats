package reqres

import (
	"encoding/json"
	"net/http"

	"github.com/edwinnduti/home-nats/lib"
	"github.com/edwinnduti/home-nats/models"
	"github.com/nats-io/nats.go"
)

func (s NatsServer) WelcomeReply(subject string) {
	// subscribe to nats subject welcome
	sub, err := s.Server.Nc.Subscribe("welcome", func(msg *nats.Msg) {
		message := models.Info{
			Message: "Welcome to my NATS API!",
		}

		byteMsg, err := json.Marshal(message)
		if err != nil {
			lib.CheckErr(w, "Marshal error", http.StatusInternalServerError, "Json Marshal Error", err)
		}

		// publish message
		s.Server.Nc.Publish(msg.Reply, byteMsg)
	})

	// check error from subscribe
	if err != nil {
		lib.CheckErr(w, "Subscribe to welcome Error", http.StatusInternalServerError, "Subscribe Error", err)
	}

	// unsubscribe from nats subject welcome
	if err := sub.Unsubscribe(); err != nil {
		lib.CheckErr(w, "Cannot Unsubscribe to NATS error", http.StatusInternalServerError, "NATS Unsubscribe Error", err)
	}

}
