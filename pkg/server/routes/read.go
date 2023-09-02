package routes

import (
	"foolslideproxy/pkg/modules"
	"foolslideproxy/pkg/server/errors"
	"foolslideproxy/pkg/server/formatter"
	"foolslideproxy/pkg/server/pathhandler"
	"foolslideproxy/pkg/server/transformer"
	"io"
	"log"
	"net/http"

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
		PathHandler: &pathdlr,
	}

	for _, mod := range modules.Modules {
		if mod.Id == params.ModId {
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

			trans.Images(mod.Id, params.MangaId, params.ChapterId, images)

			w.Header().Set("Cache-Control", "max-age=3600, public")
			io.WriteString(w, formatter.Read(images))
			return
		}
	}

	errors.NotFound(w)
	return
}
