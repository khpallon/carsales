package main

import (
	"database/sql"
	"log"
)

func createTable(db *sql.DB) {
	carTable := `CREATE TABLE cars (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"desc" TEXT,
		"price" TEXT		
	  );`

	log.Println("Create car table...")
	statement, err := db.Prepare(carTable) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("car table created")
}
