package hastateam

import (
	"strings"
	"time"

	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"
)

var baseUrl = "https://reader.hastateam.com/api"

var HastaTeam = types.Module{
	Id:    "ht",
	Name:  "HastaTeam",
	Flags: []types.ModuleFlag{},

	Search: func(query string, language *string) ([]types.SearchResult, error) {
		var data apiComics
		err, _ := utils.GetAndJsonParse(baseUrl+"/comics", &data)
		if err != nil {
			return nil, err
		}

		result := []types.SearchResult{}

		for _, manga := range data.Comics {

			if utils.StrConaintsIgnoreCase(manga.Title, strings.ToLower(query)) {
				result = append(result, types.SearchResult{
					Id:    manga.Url,
					Title: manga.Title,
					Image: manga.Thumbnail,
				})
			}
		}

		return result, nil
	},

	Manga: func(id string) (*types.Manga, error) {

		var data apiComic
		err, notfound := utils.GetAndJsonParse(baseUrl+id, &data)
		if notfound {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		chapters := []types.Chapter{}

		for _, c := range data.Comic.Chapters {
			updatedat, _ := time.Parse(time.RFC3339Nano, c.UpdatedAt)

			chapters = append(chapters, types.Chapter{
				Id:    c.Url,
				Date:  updatedat,
				Title: c.FullTitle,
			})
		}

		return &types.Manga{
			Synopsis:  "",
			Author:    data.Comic.Author,
			Artist:    data.Comic.Artist,
			Img:       data.Comic.Thumbnail,
			Chapters:  chapters,
			Sourceurl: "https://reader.hastateam.com" + id,
		}, nil
	},

	Chapter: func(manga, id string) ([]string, error) {
		var data apiRead
		err, notfound := utils.GetAndJsonParse(baseUrl+id, &data)
		if notfound {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		return data.Chapter.Pages, nil
	},
}
