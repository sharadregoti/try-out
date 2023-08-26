package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Avenger represents a single hero
type Avenger struct {
	RealName string `json:"real_name"`
	HeroName string `json:"hero_name"`
	Planet   string `json:"planet"`
	Alive    bool   `json:"alive"`
}

var avengers = []Avenger{
	{
		RealName: "Dr. Bruce Banner",
		HeroName: "Hulk",
		Planet:   "Midgard",
	},
	{
		RealName: "Tony Stark",
		HeroName: "Iron Man",
		Planet:   "Midgard",
	},
	{
		RealName: "Thor Odinson",
		HeroName: "Thor",
		Planet:   "Midgard",
	},
}

func (a *Avenger) isAlive() {
	a.Alive = true
}

func getAvenger(w http.ResponseWriter, req *http.Request) {

	v := mux.Vars(req)

	// id := v["id"]
	id, _ := strconv.Atoi(v["id"])
	for i, avg := range avengers {
		if id == i {
			SendResponse(req.Context(), w, http.StatusOK, avg)
			return
		}
	}

	SendResponse(req.Context(), w, http.StatusNotFound, map[string]string{"message": "avenger not found"})

}

// SendResponse sends an http response
func SendResponse(ctx context.Context, w http.ResponseWriter, statusCode int, body interface{}) error {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(body)
}