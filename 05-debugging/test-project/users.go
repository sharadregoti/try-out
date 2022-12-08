package main

import "net/http"

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

func getAvenger(w http.ResponseWriter, req *http.Request) {

}
