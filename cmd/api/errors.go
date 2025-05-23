package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %v path: %s", err, r.URL.Path)

	w.WriteHeader(http.StatusInternalServerError)
	data := map[string]string{
		"error": "The server encountered a problem and could not process your request",
	}
	json.NewEncoder(w).Encode(data)
}
