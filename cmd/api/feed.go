package main

import "net/http"

func (app *application) feedView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ALl the posts are here"))
}
