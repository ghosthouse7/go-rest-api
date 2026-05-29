package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json: "email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is running"))
}

func getsUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/users", getsUsersHandler)
	log.Println("Server starting on port 8080")
	log.Println(http.ListenAndServe(":8080", nil))
}
