package routes

import (
	"log"
	"net/http"

	"github.com/Minettyx/FoolslideProxy/pkg/modules"
	"github.com/Minettyx/FoolslideProxy/pkg/server/errors"
	"github.com/Minettyx/FoolslideProxy/pkg/server/pathhandler"
	"github.com/Minettyx/FoolslideProxy/pkg/types"

	"github.com/go-chi/chi/v5"
)

func Image(w http.ResponseWriter, r *http.Request) {

	pathdlr := pathhandler.MixHandler

	params, err := pathdlr.ParseImagePath(chi.URLParam(r, "path"))
	if err != nil {
		errors.NotFound(w)
		return
	}

	for _, mod := range modules.Modules {
		if mod, isImageModule := mod.(types.ImageProxyModule); isImageModule && mod.Id() == params.ModId {

			w.Header().Set("content-type", "image/jpeg") // here so it can be overwritten
			w.Header().Set("Cache-Control", "max-age=43200, public")

			err := mod.Image(params.MangaId, params.ChapterId, params.ImageId, w)
			if err != nil {
				log.Println(err)
				errors.ServerError(w)
				return
			}

			return
		}
	}

	errors.NotFound(w)
	return
}
