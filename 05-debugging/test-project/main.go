package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/v1/avenger/{id}", getAvenger)
	
	// Math Endpoints
	r.HandleFunc("/v1/add/{num1}/{num2}", handleAdd)
	r.HandleFunc("/v1/sub/{num1}/{num2}", handleSub)
	
	http.ListenAndServe(":8090", r)
}
