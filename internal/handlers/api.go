package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/AidanLDev/go-api-tutorial/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", function (router, chi.Router)) {
		// Middleware for /account rout
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	}
}


