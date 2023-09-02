package juinjutsu

import (
	"encoding/json"
	"fmt"
	"foolslideproxy/pkg/types"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var baseurl = "https://juinjutsureader.ovh/"

var JuinJutsu = types.Module{
	Id:    "jj",
	Name:  "JuinJutsu",
	Flags: []types.ModuleFlag{},

	Search: func(query string, language *string) ([]types.SearchResult, error) {
		res, err := http.PostForm(baseurl+"search/", url.Values{
			"search": {query},
		})
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}

		results := []types.SearchResult{}

		doc.Find("div .series_element").Each(func(i int, s *goquery.Selection) {

			id, _ := s.Find("a").First().Attr("href")
			par := strings.Split(id, "series/")
			if len(par) < 2 {
				return
			}

			id = par[1]

			title := s.Find(".title > a").Text()

			image, _ := s.Find("img").First().Attr("src")

			results = append(results, types.SearchResult{
				Id:    id,
				Title: title,
				Image: image,
			})
		})

		return results, nil
	},

	Manga: func(id string) (*types.Manga, error) {
		res, err := http.Get(baseurl + "series/" + id)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}

		synopsis := doc.Find(".trama").Text()
		if len(synopsis) > 7 {
			synopsis = synopsis[7:]
		}

		author := doc.Find(".autore").Text()
		if len(author) > 8 {
			author = author[8:]
		}

		artist := doc.Find(".artista").Text()
		if len(artist) > 9 {
			artist = artist[9:]
		}

		img, _ := doc.Find(".thumb").Attr("src")

		chapters := []types.Chapter{}

		doc.Find(".element").Each(func(i int, s *goquery.Selection) {

			chtitle := s.Find("a").Text()
			chid, _ := s.Find("a").Attr("href")
			pts := strings.Split(chid, id)
			if len(pts) < 2 {
				return
			}
			chid = pts[1]

			datestr := s.Find(".meta_r").Text()
			date, err := parseDate(datestr)
			if err != nil {
				return
			}

			chapters = append(chapters, types.Chapter{
				Title: chtitle,
				Id:    chid,
				Date:  date,
			})
		})

		result := types.Manga{
			Synopsis:  synopsis,
			Author:    author,
			Artist:    artist,
			Img:       img,
			Sourceurl: "https://juinjutsureader.ovh/series/" + id,
			Chapters:  chapters,
		}

		return &result, nil
	},

	Chapter: func(manga, id string) ([]string, error) {
		res, err := http.Get(baseurl + "read/" + manga + id)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		pts := strings.Split(string(body), "var pages = ")
		if len(pts) < 2 {
			return nil, fmt.Errorf("Error getting pages for juinjutsu chapter")
		}

		jsonstr := strings.Split(pts[1], ";")[0]

		type pagesJson struct {
			Url string `json:"url"`
		}

		var jsonv []pagesJson

		err = json.Unmarshal([]byte(jsonstr), &jsonv)
		if err != nil {
			return nil, err
		}

		result := []string{}

		for _, v := range jsonv {
			result = append(result, v.Url)
		}

		return result, nil
	},
}

func parseDate(date string) (time.Time, error) {
	switch date {
	case "Oggi":
		return time.Now(), nil
	case "Ieri":
		return time.Now().Add(time.Hour * -24), nil
	default:
		return time.Parse("2006.1.02", date)
	}
}
