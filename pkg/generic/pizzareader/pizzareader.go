package pizzareader

import (
	"net/url"
	"time"

	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"
)

type pizzaReader struct {
	moduleId    string
	moduleName  string
	moduleFlags types.ModuleFlags
	baseUrl     string
}

var _ types.Module = pizzaReader{}

func (c pizzaReader) Id() string {
	return c.moduleId
}
func (c pizzaReader) Name() string {
	return c.moduleName
}
func (c pizzaReader) Flags() types.ModuleFlags {
	return c.moduleFlags
}

func (config pizzaReader) Search(query string) ([]types.SearchResult, error) {
	if len(query) < 3 {
		return nil, nil
	}

	var data apiComics
	err, _ := utils.GetAndJsonParse(config.baseUrl+"/api/search/"+url.QueryEscape(query), &data)

	results := []types.SearchResult{}

	for _, com := range data.Comics {
		results = append(results, types.SearchResult{
			Id:    com.Url,
			Title: com.Title,
			Image: com.Thumbnail,
		})
	}

	return results, err
}

func (config pizzaReader) Manga(id string) (*types.Manga, error) {
	var data apiComic
	err, _ := utils.GetAndJsonParse(config.baseUrl+"/api"+id, &data)

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
		Sourceurl: config.baseUrl + id,
	}, nil
}

func (config pizzaReader) Chapter(manga string, id string) ([]string, error) {
	var data apiRead

	err, _ := utils.GetAndJsonParse(config.baseUrl+"/api"+id, &data)
	if err != nil {
		return nil, err
	}

	return data.Chapter.Pages, nil
}

func (p pizzaReader) Latest() ([]types.LatestResult, error) {
	return []types.LatestResult{}, nil
}

func (p pizzaReader) Popular() ([]types.PopularResult, error) {
	return []types.PopularResult{}, nil
}
