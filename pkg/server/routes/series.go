package routes

import (
	"log"
	"net/http"

	"github.com/Minettyx/FoolslideProxy/pkg/modules"
	"github.com/Minettyx/FoolslideProxy/pkg/server/errors"
	"github.com/Minettyx/FoolslideProxy/pkg/server/pathhandler"
	"github.com/Minettyx/FoolslideProxy/pkg/server/templates"
	"github.com/Minettyx/FoolslideProxy/pkg/server/transformer"

	"github.com/go-chi/chi/v5"
)

func Series(w http.ResponseWriter, r *http.Request) {
	pathdlr := pathhandler.MixHandler

	params, err := pathdlr.ParseMangaPath(chi.URLParam(r, "path"))
	if err != nil {
		errors.NotFound(w)
		return
	}

	trans := transformer.Transformer{
		PathHandler: pathdlr,
	}

	for _, mod := range modules.Modules {
		if mod.Id() == params.ModId {
			data, err := mod.Manga(params.MangaId)

			if err != nil {
				log.Println(err)
				errors.ServerError(w)
				return
			}

			if data == nil {
				errors.NotFound(w)
				return
			}

			trans.Manga(mod.Id(), params.MangaId, data)

			w.Header().Set("Cache-Control", "max-age=3600, public")
			templates.Series(mod, data).Render(r.Context(), w)
			return
		}
	}

	errors.NotFound(w)
	return
}

func SeriesRedirect(w http.ResponseWriter, r *http.Request) {
	pathdlr := pathhandler.MixHandler

	params, err := pathdlr.ParseMangaPath(chi.URLParam(r, "path"))
	if err != nil {
		errors.NotFound(w)
		return
	}

	for _, mod := range modules.Modules {
		if mod.Id() == params.ModId {
			data, err := mod.Manga(params.MangaId)
			if err != nil {
				log.Println(err)
				errors.ServerError(w)
				return
			}

			if data == nil {
				errors.NotFound(w)
				return
			}

			w.Header().Set("Cache-Control", "max-age=3600, public")
			w.Header().Set("Location", data.Sourceurl)
			w.WriteHeader(302)
			return
		}
	}

	errors.NotFound(w)
	return
}
