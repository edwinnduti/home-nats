package lib

import (
	"database/sql"
	"os"

	"github.com/edwinnduti/gone-nats/consts"
	"github.com/go-sql-driver/mysql"
)

// connection to database
func ConnectDB() (*sql.DB, error) {

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("ADDR"),
		DBName: os.Getenv("DBNAME"),
	}

	// open database by driver and source name
	db, err := sql.Open("mysql", cfg.FormatDSN())

	// log db result
	if dbPingErr := db.Ping(); dbPingErr == nil {
		consts.InfoLogger.Println("Database connection successful!")
	}

	// return database connection and errors if present
	return db, err
}
