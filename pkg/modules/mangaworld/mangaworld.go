package mangaworld

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"
	"github.com/dop251/goja"

	"github.com/PuerkitoBio/goquery"
)

//go:embed aes.min.js
var aesMinJs string

type mangaWorld struct {
	baseUrl string
	name    string
	id      string
}

var Module = mangaWorld{
	baseUrl: "https://www.mangaworld.ac/",
	name:    "MangaWorld",
	id:      "mw",
}

var ModuleAdult = mangaWorld{
	baseUrl: "https://www.mangaworldadult.net/",
	name:    "MangaWorld adult",
	id:      "mwa",
}

var _ types.Module = mangaWorld{}

func (c mangaWorld) Id() string {
	return c.id
}
func (c mangaWorld) Name() string {
	return c.name
}
func (c mangaWorld) Flags() types.ModuleFlags {
	return types.ModuleFlags{}
}

func client(req *http.Request) (*http.Response, error) {

	hclient := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := hclient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyB, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	body := string(bodyB)
	body, _ = utils.StrBetweenFirst(body, "<script>", "location.href")
	body = strings.Replace(body, "document.cookie", "const finalcookie", 1)

	code := aesMinJs + "\n\n" + body

	vm := goja.New()
	_, err = vm.RunString(code)
	if err != nil {
		return nil, err
	}

	var finalcookie string
	err = vm.ExportTo(vm.Get("finalcookie"), &finalcookie)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", finalcookie)

	finalres, err := hclient.Do(req)
	if err != nil {
		return finalres, err
	}

	if finalres.StatusCode == 301 || finalres.StatusCode == 302 {
		loc, _ := finalres.Location()
		recreq, _ := http.NewRequest("GET", loc.String(), nil)
		return client(recreq)
	}

	return finalres, err
}

func (c mangaWorld) Search(query string) ([]types.SearchResult, error) {
	req, _ := http.NewRequest("GET", c.baseUrl+"archive?keyword="+query, nil)
	res, err := client(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	results := []types.SearchResult{}

	doc.Find(".entry").Each(func(i int, s *goquery.Selection) {
		mid, _ := s.Find("a").Attr("href")
		pts := strings.Split(mid, "/manga/")
		if len(pts) < 2 {
			return
		}
		mid = pts[1]

		title := s.Find(".manga-title").Text()

		image, _ := s.Find("img").Attr("src")

		results = append(results, types.SearchResult{
			Id:    mid,
			Title: title,
			Image: image,
		})
	})

	return results, nil
}

func (c mangaWorld) Manga(id string) (*types.Manga, error) {
	req, _ := http.NewRequest("GET", c.baseUrl+"manga/"+id, nil)
	res, err := client(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	synopsis := doc.Find("#noidungm").Text()
	author := ""
	artist := ""

	doc.Find("div.col-12.col-md-6").Each(func(i int, s *goquery.Selection) {
		inx := s.Find("span").Text()
		val := s.Find("a").Text()

		if inx == "Autore: " {
			author = val
		} else if inx == "Artista: " {
			artist = val
		}
	})

	img, _ := doc.Find(".rounded").Attr("src")

	chapters := []types.Chapter{}

	doc.Find(".chapter").Each(func(i int, s *goquery.Selection) {
		chtitle := s.Find("span").Text()

		chid, _ := s.Find("a").Attr("href")
		pts := strings.Split(chid, "/read/")
		if len(pts) < 2 {
			return
		}
		chid = pts[1]

		datestr := s.Find("i").Text()
		dat, err := parseDate(datestr)
		if err != nil {
			return
		}

		chapters = append(chapters, types.Chapter{
			Title: chtitle,
			Id:    chid,
			Date:  *dat,
		})
	})

	manga := types.Manga{
		Synopsis:  synopsis,
		Author:    author,
		Artist:    artist,
		Img:       img,
		Chapters:  chapters,
		Sourceurl: c.baseUrl + "manga/" + id,
	}

	return &manga, nil
}

func (c mangaWorld) Chapter(manga, id string) ([]string, error) {
	req, _ := http.NewRequest("GET", c.baseUrl+"manga/"+manga+"/read/"+id, nil)
	res, err := client(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonstr, err := utils.StrBetweenFirst(string(body), `"pages":`, "]")
	jsonstr += "]"
	if err != nil {
		return nil, err
	}

	var pages []string
	err = json.Unmarshal([]byte(jsonstr), &pages)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	firstimage, _ := doc.Find("#page").Find("img").Attr("src")
	pts := strings.Split(firstimage, "/")
	pts = pts[0 : len(pts)-1]
	baseimageurl := strings.Join(pts, "/")

	for i := range pages {
		pages[i] = baseimageurl + "/" + pages[i]
	}

	return pages, nil
}

func (c mangaWorld) Latest() ([]types.LatestResult, error) {
	return []types.LatestResult{}, nil
}

func (c mangaWorld) Popular() ([]types.PopularResult, error) {
	return []types.PopularResult{}, nil
}

func parseDate(input string) (*time.Time, error) {
	var mese time.Month
	parts := strings.Split(input, " ")

	if len(parts) < 3 {
		return nil, fmt.Errorf("Error parsing date string")
	}

	switch parts[1] {
	case "Gennaio":
		mese = time.January
		break

	case "Febbraio":
		mese = time.February
		break

	case "Marzo":
		mese = time.March
		break

	case "Aprile":
		mese = time.April
		break

	case "Maggio":
		mese = time.May
		break

	case "Giugno":
		mese = time.June
		break

	case "Luglio":
		mese = time.July
		break

	case "Agosto":
		mese = time.August
		break

	case "Settembre":
		mese = time.September
		break

	case "Ottobre":
		mese = time.October
		break

	case "Novembre":
		mese = time.November
		break

	case "Dicembre":
		mese = time.December
		break
	}

	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	dat := time.Date(year, mese, day, 0, 0, 0, 0, time.UTC)
	return &dat, nil
}
