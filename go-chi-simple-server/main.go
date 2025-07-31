// Create a simple server using chi that responds with "Hello, World!" on the root path.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var port = ":3000"

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Respond with {Message: "Hello, World!"}
		w.Header().Set("Content-Type", "application/json")

		// Create an object with a Message property
		response := map[string]string{
			"Message": "Hello, World!",
		}

		// Convert the object to JSON
		json.NewEncoder(w).Encode(response)
	})

	fmt.Printf("Server is running on http://localhost%s", port)
	http.ListenAndServe(port, r)
}
