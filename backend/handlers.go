package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Title string `json:"title"`
	Price int    `json:"price"`
	Desc  string `json:"description"`
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload from the request body into a Car struct
	var newCar Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// You can perform further validation or processing here

	// Respond with a success message and status code 201 (Created)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
	fmt.Println(newCar)

}
