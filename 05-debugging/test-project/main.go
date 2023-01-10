package main

import (
	"net/http"

	"github.com/gorilla/mux"
)



func (a *Avenger) isAlive() {
	a.Alive = true
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/avenger/{id}", getAvenger)
	http.ListenAndServe(":8090", r)
}
