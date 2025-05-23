package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) feedView(w http.ResponseWriter, r *http.Request) {
	posts, err := app.store.PostStoreAllMethods.GetAll()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
