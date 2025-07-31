// Create a simple server using chi that responds with "Hello, World!" on the root path.
package main

import (
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
		fmt.Fprint(w, `{"Message": "Hello, World!"}`)
	})

	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(port, r)
}
