package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", add)
	http.ListenAndServe(":2112", nil)
}

func add(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req.URL.Query())

	num1, _ := strconv.Atoi(req.URL.Query().Get("num1"))
	num2, _ := strconv.Atoi(req.URL.Query().Get("num2"))

	json.NewEncoder(w).Encode(map[string]interface{}{"result": num1 + num2})
}
