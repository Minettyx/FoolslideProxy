package onepiecepower

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
)

type onePiecePower struct {
	baseUrl string
}

var Module = onePiecePower{
	baseUrl: "https://onepiecepower.com/",
}

var _ types.Module = onePiecePower{}

func (c onePiecePower) Id() string {
	return "opp"
}
func (c onePiecePower) Name() string {
	return "One Piece Power"
}
func (c onePiecePower) Flags() types.ModuleFlags {
	return []types.ModuleFlag{}
}

func (c onePiecePower) Search(query string) ([]types.SearchResult, error) {
	res, err := http.Get(c.baseUrl + "manga8/lista-manga")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	results := []types.SearchResult{}

	doc.Find("#allList > a").Each(func(i int, s *goquery.Selection) {
		idd, _ := s.Attr("href")

		pts := strings.Split(idd, "/")
		pts[len(pts)-1] = "images/cover.jpg"
		image := c.baseUrl + "manga8/" + strings.Join(pts, "/")

		// println(image)

		if utils.StrConaintsIgnoreCase(s.Text(), query) {
			results = append(results, types.SearchResult{
				Id:    idd,
				Title: s.Text(),
				Image: image,
			})
		}
	})

	return results, nil
}

func (c onePiecePower) Manga(id string) (*types.Manga, error) {
	res, err := http.Get(c.baseUrl + "manga8/" + id)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	key := ""
	author := ""
	artist := ""
	desc := ""
	doc.Find("body > table > tbody > tr:nth-child(3) > td > *").Each(func(i int, s *goquery.Selection) {
		if s.Is("span") {
			key = s.Text()
		} else if s.Is("em") {
			if key == "Autore:" {
				author = s.Text()
			} else if key == "Descrizione:" {
				desc = s.Text()
			} else if key == "Artista:" {
				artist = s.Text()
			}
		}
	})

	pts := strings.Split(id, "/")
	pts[len(pts)-1] = "images/cover.jpg"
	image := c.baseUrl + "manga8/" + strings.Join(pts, "/")

	chapters := []types.Chapter{}
	doc.Find("tbody > tr:nth-child(5) > td > a").Each(func(i int, s *goquery.Selection) {
		chid, _ := s.Attr("href")
		chname := s.Text()

		if strings.Contains(chname, "(Disponibile") || strings.Contains(chname, "(Available") {
			return
		}

		chapters = append(chapters, types.Chapter{
			Title: chname,
			Id:    chid,
		})
	})

	for i, j := 0, len(chapters)-1; i < j; i, j = i+1, j-1 {
		chapters[i], chapters[j] = chapters[j], chapters[i]
	}

	return &types.Manga{
		Synopsis:  desc,
		Author:    author,
		Artist:    artist,
		Img:       image,
		Sourceurl: c.baseUrl + "manga8/" + id,
		Chapters:  chapters,
	}, nil
}

func (c onePiecePower) Chapter(manga, id string) ([]string, error) {
	pts := strings.Split(manga, "/")
	pts[len(pts)-1] = id
	urls := c.baseUrl + "manga8/" + strings.Join(pts, "/")
	// urlobj, err := url.Parse(urls)
	// if err != nil {
	// 	return nil, err
	// }

	resp, err := http.Get(urls)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	base, err := pageBaseUrl(string(html), manga, id)

	if err != nil {
		return nil, err
	}

	// println(base)

	// calculate capacity
	capacity := 20
	for {
		exis, err := pageExist(base, capacity+1)
		if err != nil {
			return nil, err
		}

		if exis {
			capacity *= 2
		} else {
			break
		}
	}

	// println("capacity: ", capacity)

	// binary search
	start := 1
	end := capacity
	size := capacity

	for start <= end {
		middle := (start + end) / 2
		// println("middle:", middle)

		exis, err := pageExist(base, middle)
		if err != nil {
			return nil, err
		}

		if exis {
			if end-start <= 1 {
				size = middle
				break
			}
			start = middle + 1
		} else {
			if end-start <= 1 {
				size = middle - 1
				break
			}
			end = middle - 1
		}
	}

	// println("capacity: ", size)

	results := []string{}
	for i := 1; i <= size; i++ {
		results = append(results, pageUrl(base, i))
	}

	return results, nil
}

func (c onePiecePower) Latest() ([]types.LatestResult, error) {
	return []types.LatestResult{}, nil
}

func (c onePiecePower) Popular() ([]types.PopularResult, error) {
	return []types.PopularResult{}, nil
}

func pageBaseUrl(html string, manga string, chapter string) (string, error) {
	pts := strings.Split(manga, "/")
	pts[len(pts)-1] = chapter
	urls := Module.baseUrl + "manga8/" + strings.Join(pts, "/")
	urlobj, err := url.Parse(urls)
	if err != nil {
		return "", err
	}

	code0, err := utils.StrBetween(html, `<script type="text/javascript">`, `link.concat(".jpg")`)
	if err != nil {
		return "", err
	}

	codelns := strings.Split(code0, "\n")
	codelns2 := []string{}
	for i := range codelns {
		if i == len(codelns)-1 {
			continue
		}

		p := strings.TrimSpace(codelns[i])

		if strings.HasPrefix(p, "$") || strings.Contains(p, "XMLHttpRequest") {
			continue
		}

		if strings.Contains(p, "window.location.href") {
			p = strings.ReplaceAll(p, "window.location.href", `"`+urls+`"`)
		}
		if strings.Contains(p, "location.pathname") {
			p = strings.ReplaceAll(p, "location.pathname", `"`+urlobj.Path+`"`)
		}
		if strings.Contains(p, "location.search") {
			p = strings.ReplaceAll(p, "location.search", `""`)
		}

		codelns2 = append(codelns2, p)
	}

	code := strings.Join(codelns2, "\n")

	// println(code)

	vm := goja.New()
	_, err = vm.RunString(code)
	if err != nil {
		return "", err
	}

	var res string
	err = vm.ExportTo(vm.Get("link"), &res)
	if err != nil {
		return "", err
	}

	a := strings.Split(res, "/")
	res = strings.Join(a[:len(a)-1], "/")

	// println("urls:", urls, "res: ", res)
	resobj, err := url.Parse(res)
	if err != nil {
		return "", err
	}
	if resobj.IsAbs() {
		return res + "/", nil
	} else {
		rr := urlobj.JoinPath("../", res)
		return rr.String() + "/", nil
	}
}

func pageUrl(base string, page int) string {
	if page < 10 {
		base += "0"
	}

	base += fmt.Sprint(page)
	base += ".jpg"

	return base
}

func pageExist(base string, page int) (bool, error) {
	res, err := http.Get(pageUrl(base, page))
	if err != nil {
		return false, err
	}
	return res.StatusCode == 200, nil
}
