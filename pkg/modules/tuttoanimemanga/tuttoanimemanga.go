package tuttoanimemanga

import (
	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"
	"net/url"
	"time"
)

var baseurl = "https://tuttoanimemanga.net/"

var TuttoAnimeManga = types.Module{
	Id:    "tam",
	Name:  "TuttoAnimeManga",
	Flags: types.ModuleFlags{},

	Search: func(query string, language *string) ([]types.SearchResult, error) {
		var data searchRes
		err, _ := utils.GetAndJsonParse(baseurl+"api/search/"+url.QueryEscape(query), &data)

		results := []types.SearchResult{}

		for _, com := range data.Comics {
			results = append(results, types.SearchResult{
				Id:    com.Url,
				Title: com.Title,
				Image: com.Thumbnail,
			})

			println(com.Thumbnail)
		}

		return results, err
	},

	Manga: func(id string) (*types.Manga, error) {
		var data mangaRes
		err, _ := utils.GetAndJsonParse(baseurl+"api"+id, &data)
		if err != nil {
			return nil, err
		}

		chapters := []types.Chapter{}

		for _, ch := range data.Comic.Chapters {
			date, err := time.Parse(time.RFC3339Nano, ch.PublishedOn)
			if err != nil {
				return nil, err
			}

			chapters = append(chapters, types.Chapter{
				Title: ch.FullTitle,
				Id:    ch.Url,
				Date:  date,
			})
		}

		return &types.Manga{
			Synopsis:  data.Comic.Description,
			Author:    data.Comic.Author,
			Artist:    data.Comic.Artist,
			Img:       data.Comic.Thumbnail,
			Chapters:  chapters,
			Sourceurl: "https://tuttoanimemanga.net" + id,
		}, nil
	},

	Chapter: func(manga, id string) ([]string, error) {
		var data chapterRes

		println(baseurl + "api" + id)

		err, _ := utils.GetAndJsonParse(baseurl+"api"+id, &data)
		if err != nil {
			return nil, err
		}

		return data.Chapter.Pages, nil
	},
}
