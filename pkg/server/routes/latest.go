package routes

import (
	"foolslideproxy/pkg/modules"
	"foolslideproxy/pkg/server/formatter"
	"foolslideproxy/pkg/server/pathhandler"
	"foolslideproxy/pkg/server/transformer"
	"foolslideproxy/pkg/types"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
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
		go func(mod *types.Module) {
			defer wg.Done()

			if mod.Latest == nil {
				return
			}

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
				trans.LatestResult(mod.Id, &res[i])
				results = append(results, &res[i])
			}

		}(mod)
	}

	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		return results[j].Date.Before(results[i].Date)
	})

	w.Header().Set("Cache-Control", "max-age=3600, public")
	io.WriteString(w, formatter.Latest(results))
}
