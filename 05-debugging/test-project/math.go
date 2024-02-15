package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleAdd(w http.ResponseWriter, req *http.Request) {
	v := mux.Vars(req)

	num1, _ := strconv.Atoi(v["num1"])
	num2, _ := strconv.Atoi(v["num2"])

	SendResponse(req.Context(), w, http.StatusNotFound, map[string]interface{}{"message": add(num1, num2)})
}

func handleSub(w http.ResponseWriter, req *http.Request) {
	v := mux.Vars(req)

	num1, _ := strconv.Atoi(v["num1"])
	num2, _ := strconv.Atoi(v["num2"])

	SendResponse(req.Context(), w, http.StatusNotFound, map[string]interface{}{"message": sub(num1, num2)})
}

func add(num1, num2 int) int {
	result := num1 + num2
	return result
}

func sub(num1, num2 int) int {
	result := num1 - num2
	return result
}
