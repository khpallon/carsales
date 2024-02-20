package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/cars", CreateCar)

	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Starting server at - http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
