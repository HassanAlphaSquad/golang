package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// User model
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Salary  int    `json:"salary"`
}

// Database connection
var db *sql.DB

func main() {
	var err error
	conn := "postgres://postgres:alpha123@localhost:5432/postgres?sslmode=disable"
	db, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Create router
	r := mux.NewRouter()
	r.HandleFunc("/users", get_all_users).Methods("GET")
	r.HandleFunc("/user/{id}", get_user).Methods("GET")
	r.HandleFunc("/user", create_user).Methods("POST")
	r.HandleFunc("/user/{id}", update_user).Methods("PUT")
	r.HandleFunc("/user/{id}", delete_user).Methods("DELETE")

	// Start server
	fmt.Println("🚀 Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Get all users
func get_all_users(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, age, address, salary FROM users ORDER BY id")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Address, &u.Salary); err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	json.NewEncoder(w).Encode(users)
}

// Get a single user by ID
func get_user(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	row := db.QueryRow("SELECT id, name, age, address, salary FROM users WHERE id = $1", id)

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Age, &u.Address, &u.Salary)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}

// Create a new user
func create_user(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	err := db.QueryRow(
		"INSERT INTO users (name, age, address, salary) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Name, u.Age, u.Address, u.Salary).Scan(&u.ID)

	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(u)
}

// Update an existing user
func update_user(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	_, err := db.Exec("UPDATE users SET name=$1, age=$2, address=$3, salary=$4 WHERE id=$5",
		u.Name, u.Age, u.Address, u.Salary, id)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

// Delete a user
func delete_user(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
}
