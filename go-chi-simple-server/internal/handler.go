package internal

import (
	"encoding/json"
	"net/http"
)

// HelloHandler handles the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with {Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")

	// Create an object with a Message property
	response := map[string]string{
		"Message": "Hello, World!",
	}

	// Convert the object to JSON
	json.NewEncoder(w).Encode(response)
}

// UserHandler handles the /user endpoint
func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}
	json.NewEncoder(w).Encode(user)
}
