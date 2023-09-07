package routes

import (
	"github.com/Minettyx/FoolslideProxy/pkg/modules"
	"github.com/Minettyx/FoolslideProxy/pkg/server/formatter"
	"github.com/Minettyx/FoolslideProxy/pkg/server/pathhandler"
	"github.com/Minettyx/FoolslideProxy/pkg/server/transformer"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
)

func Directory1(w http.ResponseWriter, r *http.Request) {

	pathdlr := pathhandler.MixHandler
	trans := transformer.Transformer{
		PathHandler: &pathdlr,
	}

	var wg sync.WaitGroup
	wg.Add(len(modules.Modules))

	results := []*types.PopularResult{}

	for _, mod := range modules.Modules {
		go func(mod *types.Module) {
			defer wg.Done()

			if mod.Popular == nil {
				return
			}

			res, err := mod.Popular()
			if err != nil {
				// TODO: send fake manga to tell that a source has failed
				log.Println(err)
				return
			}

			if res == nil {
				res = []types.PopularResult{}
			}

			for i := range res {
				trans.PopularResult(mod.Id, &res[i])
				results = append(results, &res[i])
			}

		}(mod)
	}

	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		return results[j].Popularity < results[i].Popularity
	})

	w.Header().Set("Cache-Control", "max-age=3600, public")
	io.WriteString(w, formatter.Directory(results))
}
