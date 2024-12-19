package ccm

import (
	"encoding/json"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ccm struct {
	baseUrl string
}

var Module = ccm{
	baseUrl: "https://ccm.0kb.eu/api/",
}

var _ types.Module = ccm{}

func (c ccm) Id() string {
	return "ccm"
}
func (c ccm) Name() string {
	return "CCM Translations"
}
func (c ccm) Flags() types.ModuleFlags {
	return types.ModuleFlags{}
}

func (c ccm) Search(query string) ([]types.SearchResult, error) {

	var data []ccmMangasReq
	err, _ := utils.GetAndJsonParse(c.baseUrl+"mangas.json", &data)
	if err != nil {
		return nil, err
	}

	result := []types.SearchResult{}

	for _, manga := range data {

		if utils.StrConaintsIgnoreCase(manga.Title, strings.ToLower(query)) {
			result = append(result, types.SearchResult{
				Id:    manga.Id,
				Title: manga.Title,
				Image: manga.Cover,
			})
		}
	}

	return result, nil
}

func (c ccm) Manga(mangaid string) (*types.Manga, error) {
	var data ccmMangaReq
	err, notfound := getJsonParseAndNotFound(c.baseUrl+"manga/"+url.QueryEscape(mangaid)+".json", &data)
	if notfound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	chapters := []types.Chapter{}
	for i := len(data.Chapters) - 1; i >= 0; i-- {
		c := data.Chapters[i]
		chapters = append(chapters, types.Chapter{
			Id:    c.Chapter,
			Date:  time.UnixMilli(c.Time),
			Title: utils.GenTitle(c.Volume, c.Chapter, c.Title),
		})
	}

	return &types.Manga{
		Synopsis:  "",
		Author:    data.Author,
		Artist:    data.Artist,
		Img:       data.Cover,
		Chapters:  chapters,
		Sourceurl: "https://ccmscans.in/manga/" + data.Id,
	}, nil
}

func (c ccm) Chapter(manga string, id string) ([]string, error) {
	var data ccmChapterReq
	err, notfound := getJsonParseAndNotFound(c.baseUrl+"chapter/"+url.QueryEscape(manga)+"/"+url.QueryEscape(id)+".json", &data)
	if notfound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return data.Images, nil
}

func (c ccm) Latest() ([]types.LatestResult, error) {
	return []types.LatestResult{}, nil
}

func (c ccm) Popular() ([]types.PopularResult, error) {
	return []types.PopularResult{}, nil
}

// equivalent to GetAndJsonParse but custom because the response code is 200 when not found on ccmscans website
func getJsonParseAndNotFound(rurl string, v any) (error, bool) {
	is404 := false

	res, err := http.Get(rurl)
	if err != nil {
		return err, false
	}
	defer res.Body.Close()

	if res.StatusCode == 404 || res.Header.Get("content-type") != "application/json" {
		is404 = true
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err, is404
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err, is404
	}

	return nil, is404
}
