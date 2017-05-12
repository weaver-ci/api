package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func encodeResponse(statusCode int, obj interface{}, w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(obj)

	w.Header().Set("Content-Type", "application/json")

	// If no errors, Just write the response
	if err == nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(res))
	}

	// If theres an error write the error out
	w.WriteHeader(http.StatusInternalServerError)

	// Don't show the error if the environment is not production
	if strings.ToLower(environment) != "development" {
		return
	}
	errMessage, errErr := json.Marshal(err)

	if errErr != nil {
		return
	}

	w.Write([]byte(errMessage))
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	encodeResponse(http.StatusInternalServerError, err, w, r)
}
