package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:alpha123@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")

	rows, err := db.Query("SELECT * FROM users ORDER BY id ASC")
	if err != nil {
		log.Fatal("Error querying users:", err)
	}
	defer rows.Close()

	fmt.Print("\n⮕ Users in Database:\n\n")
	for rows.Next() {
		var id int
		var name string
		var age int
		var address string
		var salary int
		err := rows.Scan(&id, &name, &age, &address, &salary)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Printf("ID: %-6d Name: %-20s Age: %-5d Address: %-25s Salary: %d\n", id, name, age, address, salary)

	}

	// Check for errors after iterating
	if err := rows.Err(); err != nil {
		log.Fatal("Error in rows:", err)
	}
}
