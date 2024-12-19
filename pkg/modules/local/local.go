package local

import (
	"fmt"
	"github.com/Minettyx/FoolslideProxy/pkg/modules"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"time"
)

type mangaDB struct {
	types.Manga
	Id       string
	Title    string
	Chapters []chapterDB
}

type chapterDB struct {
	types.Chapter
	Images []string
}

var db = [](func() *mangaDB){
	func() *mangaDB {

		chs := []chapterDB{}

		for _, mod := range modules.Modules {
			if mod.Flags().Has(types.HIDDEN) {
				continue
			}

			chs = append(chs, chapterDB{
				Chapter: types.Chapter{
					Title: fmt.Sprintf("%v (%v)", mod.Name(), mod.Id()),
					Id:    mod.Id(),
					Date:  time.Now(),
				},
				Images: []string{},
			})
		}

		return &mangaDB{
			Id:    "supportedsources",
			Title: "Supported sources",
			Manga: types.Manga{
				Synopsis:  "Click on WebView for more infos",
				Author:    "Minettyx",
				Artist:    "",
				Img:       "",
				Sourceurl: "https://github.com/Minettyx/FoolslideProxy/wiki/Available-sources",
			},
			Chapters: chs,
		}
	},
	func() *mangaDB {
		return &mangaDB{
			Manga: types.Manga{
				Synopsis:  "Download our fork of the Foolslide extention to see cover images when searching (click on WebView)",
				Author:    "Minettyx",
				Artist:    "",
				Img:       "https://img.0kb.eu/kIq81BcY.jpg",
				Sourceurl: "https://github.com/Minettyx/foolslideproxy-extension/releases",
			},
			Id:       "fixsearchimages",
			Title:    "Fix search images",
			Chapters: []chapterDB{},
		}
	},
}

type localModule struct{}

var _ types.Module = localModule{}

func (c localModule) Id() string {
	return "internal"
}
func (c localModule) Name() string {
	return "internal"
}
func (c localModule) Flags() types.ModuleFlags {
	return types.ModuleFlags{types.HIDDEN, types.DISABLE_GLOBAL_SEARCH}
}

func (c localModule) Popular() ([]types.PopularResult, error) {

	res := []types.PopularResult{}

	for _, f := range db {
		mg := f()

		res = append(res, types.PopularResult{
			Id:    mg.Id,
			Title: mg.Title,
			Image: mg.Img,
		})
	}

	return res, nil
}

func (c localModule) Manga(id string) (*types.Manga, error) {

	for _, v := range db {
		mg := v()
		chs := []types.Chapter{}
		for _, ch := range mg.Chapters {
			chs = append(chs, ch.Chapter)
		}
		mg.Manga.Chapters = chs

		if mg.Id == id {
			return &mg.Manga, nil
		}
	}

	return nil, nil
}

func (c localModule) Chapter(manga, id string) ([]string, error) {
	for _, v := range db {
		mg := v()

		if mg.Id == manga {
			for _, ch := range mg.Chapters {
				if ch.Id == id {
					return ch.Images, nil
				}
			}
			return nil, nil
		}
	}

	return nil, nil
}

func (c localModule) Latest() ([]types.LatestResult, error) {
	return []types.LatestResult{}, nil
}

func (c localModule) Search(query string) ([]types.SearchResult, error) {
	return []types.SearchResult{}, nil
}

func Init() {
	modules.Modules[0] = localModule{}
}
