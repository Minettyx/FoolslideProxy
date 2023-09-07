package formatter

import (
	"fmt"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"
)

func Directory(data []*types.PopularResult) string {
	res := ""

	for _, v := range data {
		res += fmt.Sprintf(`<div class="group"><img class="preview" src="%v" /><div class="title"><a href="/series/%v" title="%v">%v</a></div></div>`,
			v.Image,
			v.Id,
			v.Title,
			v.Title,
		)
	}

	return res
}

func Latest(data []*types.LatestResult) string {
	res := ""

	for _, v := range data {
		res += fmt.Sprintf(`<div class="group"><img class="preview" src="%v" /><div class="title"><a href="/series/%v" title="%v">%v</a></div></div>`,
			v.Image,
			v.Id,
			v.Title,
			v.Title,
		)
	}

	return res
}

func Search(data []*types.SearchResult) string {
	res := ""

	for _, v := range data {
		res += fmt.Sprintf(`<div class="group"><img class="preview" src="%v" /><div class="title"><a href="/series/%v" title="%v">%v</a></div></div>`,
			v.Image,
			v.Id,
			v.Title,
			v.Title,
		)
	}

	return res
}

func Series(mod *types.Module, data *types.Manga) string {
	res := fmt.Sprintf(`<html><head></head><body><div id="wrapper"><article id="content"><div class="panel"><div class="comic info"><div class="thumbnail"><img src="%v" /></div><div class="large comic"><h1 class="title"></h1><div class="info"><b>Author</b>: %v<br><b>Artist</b>: %v<br><b>Synopsis</b>: %v</div></div></div><div class="list"><div class="group"><div class="title">Volume</div>`,
		data.Img,
		utils.AuthorArtist(data.Author, data.Artist),
		mod.Name,
		data.Synopsis,
	)

	for i := range data.Chapters {
		chapter := &data.Chapters[i]

		res += fmt.Sprintf(`<div class="element"><div class="title"><a href="/read/%v" title="%v">%v</a></div><div class="meta_r">by <a href="" title="" ></a>, %v</div></div>`,
			chapter.Id,
			chapter.Title,
			chapter.Title,
			chapter.Date.Format("2006.1.02"),
		)
	}

	res += `</div></div></div></article></div></body></html>`
	return res
}

func Read(data []string) string {
	res := "<script>var pages = ["

	if len(data) > 0 {
		for _, v := range data {
			res += `{"url":"` + v + `"},`
		}

		res = res[:len(res)-1] // remove last comma
	}

	res += "];</script>"
	return res
}
