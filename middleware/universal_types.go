// contains the universal vars and types in the entire middleware package/folder
package middleware

import (
	"github.com/nats-io/nats.go"
)

// nats server struct
type Server struct {
	Nc *nats.Conn
}
