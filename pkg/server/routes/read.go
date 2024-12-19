package routes

import (
	"io"
	"log"
	"net/http"

	"github.com/Minettyx/FoolslideProxy/pkg/modules"
	"github.com/Minettyx/FoolslideProxy/pkg/server/errors"
	"github.com/Minettyx/FoolslideProxy/pkg/server/pathhandler"
	"github.com/Minettyx/FoolslideProxy/pkg/server/templates"
	"github.com/Minettyx/FoolslideProxy/pkg/server/transformer"

	"github.com/go-chi/chi/v5"
)

func Read(w http.ResponseWriter, r *http.Request) {
	pathdlr := pathhandler.MixHandler

	params, err := pathdlr.ParseChapterPath(chi.URLParam(r, "path"))
	if err != nil {
		errors.NotFound(w)
		return
	}

	trans := transformer.Transformer{
		PathHandler: pathdlr,
	}

	for _, mod := range modules.Modules {
		if mod.Id() == params.ModId {
			images, err := mod.Chapter(params.MangaId, params.ChapterId)
			if err != nil {
				log.Println(err)
				errors.ServerError(w)
				return
			}

			if images == nil {
				errors.NotFound(w)
				return
			}

			trans.Images(mod.Id(), params.MangaId, params.ChapterId, images)

			w.Header().Set("Cache-Control", "max-age=3600, public")
			io.WriteString(w, templates.Read(images))
			return
		}
	}

	errors.NotFound(w)
	return
}
