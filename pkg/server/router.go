package server

import (
	"foolslideproxy/pkg/modules/local"
	"foolslideproxy/pkg/server/routes"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// r.Use(middleware.CleanPath)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// add local module to modules.Modules to bypass dependecies cycle
	local.Init()

	r.Get("/directory/1/", routes.Directory1)
	r.Get("/latest/1/", routes.Latest1)
	r.Post("/search/", routes.Search)
	r.Post("/series/{path}", routes.Series)
	r.Get("/series/{path}", routes.SeriesRedirect)
	r.Post("/read/{path}", routes.Read)
	r.Get("/image/{path}", routes.Image)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://github.com/Minettyx/FoolslideProxy")
		w.WriteHeader(302)
		return
	})

	return r
}
