package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/michaelmov/go-chi-simple-server/internal"
)

// SetupRouter creates and configures the chi router with all routes
func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	
	// Define routes
	r.Get("/hello", internal.HelloHandler)
	r.Get("/user", internal.UserHandler)
	
	return r
}


