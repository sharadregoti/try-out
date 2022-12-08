package main

import (
	"net/http"
)

func (a *Avenger) isAlive() {
	a.Alive = true
}

func main() {
	http.HandleFunc("/v1/avenger/{id}", getUsers)
	http.ListenAndServe(":8090", nil)
}
