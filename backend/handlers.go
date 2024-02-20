package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Car struct {
	Title string `json:"title"`
	Price string `json:"price"`
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

	fmt.Println(newCar.Title)

	sqliteDatabase, _ := sql.Open("sqlite3", "./cars.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                          // Defer Closing the database
	createTable(sqliteDatabase)                           // Create Database Tables

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../frontend/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Println(err)
		return
	}
}
