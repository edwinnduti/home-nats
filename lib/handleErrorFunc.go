package lib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edwinnduti/home-nats/consts"
	"github.com/edwinnduti/home-nats/models"
)

// HandleErrorFunc is a middleware function that handles errors by returning a response and logging them
func CheckErr(w http.ResponseWriter, printedLog string, responseCode int, responseMessage string, errorMessage error) {

	// create log for request performed
	logErr := fmt.Sprintf("%v: %v", printedLog, errorMessage)
	consts.ErrorLogger.Printf(logErr)

	// response code
	w.WriteHeader(responseCode)

	// return error
	response := models.NewResponse{
		Code:    responseCode,
		Message: responseMessage,
	}

	// write response
	json.NewEncoder(w).Encode(response)
}
