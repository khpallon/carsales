package main

import (
	"fmt"
	"html/template"
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
