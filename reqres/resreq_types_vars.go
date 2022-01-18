// contains the universal vars and types in the entire middleware package/folder
package reqres

import (
	"net/http"

	"github.com/edwinnduti/home-nats/middleware"
)

// nats server struct
type NatsServer struct {
	Server middleware.Server
}

// response writer struct
var w http.ResponseWriter
