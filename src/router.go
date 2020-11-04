package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/lgranade/minesweeperapi/controller"
)

func createRoutes() (chi.Router, error) {
	r := chi.NewRouter()

	// make it work from any domain on the browser
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
		MaxAge:           600,
	})

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RedirectSlashes,
		cors.Handler,
	)

	r.Post("/auth", controller.Authenticate)
	r.Post("/users", controller.CreateUser)

	r.Group(func(r chi.Router) {
		// TODO: add auth middleware on this router group.
		// That middleware will validate access token and leave user id in context.

		r.Get("/users/{userID}", controller.ReadUser)
		r.Post("/games", controller.CreateGame)
		r.Get("/games/{gameID}", controller.ReadGame)
		r.Post("/games/{gameID}/play", controller.Play)
		r.Post("/games/{gameID}/pause", controller.Pause)
	})

	return r, nil
}
