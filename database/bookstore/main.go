package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

/*

CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  username VARCHAR(128) NOT NULL,
  email VARCHAR(128) NOT NULL,
  password VARCHAR(128) NOT NULL,
  shipping_address VARCHAR(255) NOT NULL
);

*/
const usersCanCreateAccount = `
INSERT INTO customers (username, email, password, shipping_address) VALUES (?, ?, ?, ?);
`

func main() {
	// Define the PostgreSQL connection string.
	connStr := "host=localhost port=5432 user=peter.duke dbname=bookstore sslmode=disable"

	// Connect to the database.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to ensure it's connected.
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Now you can use `db` to interact with your PostgreSQL database.
	// Example query
	rows, err := db.Query("SELECT id, title FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop through and scan the results.
	for rows.Next() {
		var id int
		var title string
		err = rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, title)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
