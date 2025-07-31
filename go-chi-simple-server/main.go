// Create a simple server using chi that responds with "Hello, World!" on the root path.
package main

import (
	"fmt"
	"net/http"

	"github.com/michaelmov/go-chi-simple-server/internal/router"
)

var port = ":3000"

func main() {
	r := router.SetupRouter()

	fmt.Printf("Server is running on http://localhost%s", port)
	http.ListenAndServe(port, r)
}
