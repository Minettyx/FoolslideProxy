package transformer

import (
	"github.com/Minettyx/FoolslideProxy/pkg/server/pathhandler"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"net/url"
	"strings"
)

type Transformer struct {
	PathHandler pathhandler.PathHandler
}

func (tm *Transformer) PopularResult(modid string, data *types.PopularResult) {
	data.Id = tm.PathHandler.MangaPath(modid, data.Id)
}

func (tm *Transformer) LatestResult(modid string, data *types.LatestResult) {
	data.Id = tm.PathHandler.MangaPath(modid, data.Id)
}

func (tm *Transformer) SearchResult(modid string, data *types.SearchResult) {
	data.Id = tm.PathHandler.MangaPath(modid, data.Id)
}

func (tm *Transformer) Manga(modid string, mangaid string, data *types.Manga) {
	for i := range data.Chapters {
		data.Chapters[i].Id = tm.PathHandler.ChapterPath(modid, mangaid, data.Chapters[i].Id)
	}
}

func (tm *Transformer) Images(modid string, mangaid string, chapterid string, data []string) {
	for i := range data {
		if strings.HasPrefix(data[i], "local://") {
			data[i] = "/image/" + url.QueryEscape(tm.PathHandler.ImagePath(modid, mangaid, chapterid, data[i][8:]))
		}
	}
}
