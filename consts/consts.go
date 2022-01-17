package consts

import (
	"log"
	"os"
)

// for logging mechanism
var (
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
)
