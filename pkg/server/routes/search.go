package routes

import (
	"github.com/Minettyx/FoolslideProxy/pkg/modules"
	"github.com/Minettyx/FoolslideProxy/pkg/server/errors"
	"github.com/Minettyx/FoolslideProxy/pkg/server/formatter"
	"github.com/Minettyx/FoolslideProxy/pkg/server/pathhandler"
	"github.com/Minettyx/FoolslideProxy/pkg/server/transformer"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

func isSpecific(search string) *types.Module {
	for _, mod := range modules.Modules {
		if strings.HasPrefix(search, strings.ToLower(mod.Id)+":") {
			return mod
		}
	}

	return nil
}

func Search(w http.ResponseWriter, r *http.Request) {
	pathdlr := pathhandler.MixHandler
	trans := transformer.Transformer{
		PathHandler: &pathdlr,
	}

	err := r.ParseForm()
	if err != nil {
		errors.BadRequest(w)
		return
	}

	search := strings.TrimSpace(strings.ToLower(r.FormValue("search")))
	if len(search) == 0 {
		errors.BadRequest(w)
		return
	}

	results := []*types.SearchResult{}

	// check if the search is general or specific
	specific := isSpecific(search)

	if specific != nil {

		if specific.Search == nil {
			return
		}

		query := strings.TrimSpace(search[len(specific.Id)+1:])

		data, err := specific.Search(query, nil)

		if err != nil {
			log.Println(err)
			errors.ServerError(w)
			return
		}

		if data == nil {
			data = []types.SearchResult{}
		}

		for i := range data {
			trans.SearchResult(specific.Id, &data[i])
			results = append(results, &data[i])
		}

	} else {

		var wg sync.WaitGroup
		wg.Add(len(modules.Modules))

		for _, mod := range modules.Modules {
			go func(mod *types.Module) {
				defer wg.Done()

				if mod.Flags.Has(types.DISABLE_GLOBAL_SEARCH) || mod.Search == nil {
					return
				}

				res, err := mod.Search(search, nil)
				if err != nil {
					// TODO: send fake manga to tell that a source has failed
					log.Println(err)
					return
				}

				if res == nil {
					res = []types.SearchResult{}
				}

				for i := range res {
					trans.SearchResult(mod.Id, &res[i])
					results = append(results, &res[i])
				}
			}(mod)
		}

		wg.Wait()
	}

	w.Header().Set("Cache-Control", "max-age=3600, public")
	io.WriteString(w, formatter.Search(results))
}
