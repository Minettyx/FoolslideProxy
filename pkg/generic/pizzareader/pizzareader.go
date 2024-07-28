package pizzareader

import (
	"net/url"
	"time"

	"github.com/Minettyx/FoolslideProxy/pkg/generic"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"
)



func PizzaReader(config generic.GenericConfig) *types.Module {
	return &types.Module{
		Id:    config.Id,
		Name:  config.Name,
		Flags: config.Flags,

		Search: func(query string, language *string) ([]types.SearchResult, error) {

      if len(query) < 3 {
        return nil, nil
      }

			var data apiComics
			err, _ := utils.GetAndJsonParse(config.BaseUrl+"/api/search/"+url.QueryEscape(query), &data)

			results := []types.SearchResult{}

			for _, com := range data.Comics {
				results = append(results, types.SearchResult{
					Id:    com.Url,
					Title: com.Title,
					Image: com.Thumbnail,
				})
			}

			return results, err
		},

		Manga: func(id string) (*types.Manga, error) {
			var data apiComic
			err, _ := utils.GetAndJsonParse(config.BaseUrl+"/api"+id, &data)
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
				Sourceurl: config.BaseUrl + id,
			}, nil
		},

		Chapter: func(manga, id string) ([]string, error) {
			var data apiRead

			err, _ := utils.GetAndJsonParse(config.BaseUrl+"/api"+id, &data)
			if err != nil {
				return nil, err
			}

			return data.Chapter.Pages, nil
		},
	}
}
