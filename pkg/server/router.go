package server

import (
	"net/http"
	"time"

	"github.com/Minettyx/FoolslideProxy/pkg/modules/local"
	"github.com/Minettyx/FoolslideProxy/pkg/server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// r.Use(middleware.CleanPath)
	r.Use(middleware.Timeout(60 * time.Second))

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
