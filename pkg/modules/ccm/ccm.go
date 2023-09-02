package ccm

import (
	"encoding/json"
	"foolslideproxy/pkg/types"
	"foolslideproxy/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var baseUrl = "https://ccmscans.in/api/"

var CCM = types.Module{
	Id:   "ccm",
	Name: "CCM Translations",

	// Latest: func() ([]types.LatestResult, error) {
	// 	var data []ccmLatestReq
	// 	err, _ := utils.GetAndJsonParse(baseUrl+"latest.json", &data)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	result := []types.LatestResult{}
	//
	// 	for _, ch := range data {
	// 		found := false
	// 		for i := range result {
	// 			if result[i].Id == ch.Manga.Id {
	// 				found = true
	// 				if time.UnixMilli(ch.Time).After(result[i].Date) {
	// 					result[i].Date = time.UnixMilli(ch.Time)
	// 				}
	// 			}
	// 		}
	//
	// 		if !found {
	// 			result = append(result, types.LatestResult{
	// 				Id:    ch.Manga.Id,
	// 				Title: ch.Manga.Title,
	// 				Image: ch.Manga.Cover,
	// 				Date:  time.UnixMilli(ch.Time),
	// 			})
	// 		}
	// 	}
	//
	// 	return result, nil
	// },

	Search: func(query string, _ *string) ([]types.SearchResult, error) {

		var data []ccmMangasReq
		err, _ := utils.GetAndJsonParse(baseUrl+"mangas.json", &data)
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
	},

	Manga: func(mangaid string) (*types.Manga, error) {
		var data ccmMangaReq
		err, notfound := getJsonParseAndNotFound(baseUrl+"manga/"+url.QueryEscape(mangaid)+".json", &data)
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
	},

	Chapter: func(manga, id string) ([]string, error) {
		var data ccmChapterReq
		err, notfound := getJsonParseAndNotFound(baseUrl+"chapter/"+url.QueryEscape(manga)+"/"+url.QueryEscape(id)+".json", &data)
		if notfound {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		return data.Images, nil
	},
}

// equivalent to GetAndJsonParse but custom because the response code is 200 when not found on ccmscans.in
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
