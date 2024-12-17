package routes

import (
	"log"
	"net/http"
	"sort"
	"sync"

	"github.com/Minettyx/FoolslideProxy/pkg/modules"
	"github.com/Minettyx/FoolslideProxy/pkg/server/pathhandler"
	"github.com/Minettyx/FoolslideProxy/pkg/server/templates"
	"github.com/Minettyx/FoolslideProxy/pkg/server/transformer"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
)

func Latest1(w http.ResponseWriter, r *http.Request) {

	pathdlr := pathhandler.MixHandler
	trans := transformer.Transformer{
		PathHandler: &pathdlr,
	}

	var wg sync.WaitGroup
	wg.Add(len(modules.Modules))

	results := []*types.LatestResult{}

	for _, mod := range modules.Modules {
		go func(mod types.Module) {
			defer wg.Done()

			res, err := mod.Latest()
			if err != nil {
				// TODO: send fake manga to tell that a source has failed
				log.Println(err)
				return
			}

			if res == nil {
				res = []types.LatestResult{}
			}

			for i := range res {
				trans.LatestResult(mod.Id(), &res[i])
				results = append(results, &res[i])
			}

		}(mod)
	}

	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		return results[j].Date.Before(results[i].Date)
	})

	w.Header().Set("Cache-Control", "max-age=3600, public")
	templates.Latest(results).Render(r.Context(), w)
}
