package mangareader

import (
	"fmt"
	"image/jpeg"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"

	mangareader_unshuffle "github.com/Minettyx/mangareader.to-image-unshuffle"
	"github.com/PuerkitoBio/goquery"
)

type mangaReader struct {
	baseUrl string
}

var Module = mangaReader{
	baseUrl: "https://mangareader.to",
}

var _ types.Module = mangaReader{}

func (c mangaReader) Id() string {
	return "mr"
}
func (c mangaReader) Name() string {
	return "MangaReader"
}
func (c mangaReader) Flags() types.ModuleFlags {
	return types.ModuleFlags{types.DISABLE_GLOBAL_SEARCH}
}

func (c mangaReader) Search(query string) ([]types.SearchResult, error) {
	if len(query) < 1 {
		return []types.SearchResult{}, nil
	}

	res, err := http.Get(c.baseUrl + "/search?keyword=" + url.QueryEscape(query))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	results := []types.SearchResult{}

	doc.Find("div.item-spc").Each(func(i int, s *goquery.Selection) {
		a := s.Find(".manga-name > a")
		title := a.Text()

		mid, _ := a.Attr("href")
		img, _ := s.Find("img").Attr("src")

		langs := strings.Split(s.Find(".manga-poster > span").Text(), "/")

		for _, v := range langs {
			results = append(results, types.SearchResult{
				Id:    mid + "|" + v,
				Title: "[" + v + "] " + title,
				Image: img,
			})
		}

	})

	return results, nil
}

func (c mangaReader) Manga(id string) (*types.Manga, error) {
	p := strings.Split(id, "|")
	if len(p) < 2 {
		return nil, fmt.Errorf("Language not found in manga id")
	}

	mangaid := p[0]
	lang := strings.ToLower(p[1])

	res, err := http.Get(c.baseUrl + mangaid)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	author := ""

	doc.Find(".anisc-info > .item").Each(func(i int, s *goquery.Selection) {
		switch s.Find(".item-head").Text() {
		case "Authors:":
			s.Find("a").Each(func(i int, s *goquery.Selection) {
				author += strings.Replace(s.Text(), ",", "", 1) + ", "
			})
		}
	})
	if len(author) > 2 {
		author = author[:len(author)-2]
	}

	img, _ := doc.Find("img.manga-poster-img").Attr("src")

	chapters := []types.Chapter{}

	doc.Find("#" + lang + "-chapters > li").Each(func(i int, s *goquery.Selection) {
		chtit := s.Find(".name").Text()
		chid, _ := s.Find("a").Attr("href")

		chapters = append(chapters, types.Chapter{
			Title: chtit,
			Id:    chid,
		})
	})

	return &types.Manga{
		Synopsis:  doc.Find(".description").Text(),
		Author:    author,
		Artist:    "",
		Img:       img,
		Sourceurl: c.baseUrl + mangaid,
		Chapters:  chapters,
	}, nil
}

func (c mangaReader) Chapter(manga, id string) ([]string, error) {
	doc, err := utils.GetAndGoquery(c.baseUrl + id)
	if err != nil {
		return nil, err
	}

	chapid, _ := doc.Find("#wrapper").Attr("data-reading-id")

	type chResp struct {
		Html string `json:"html"`
	}

	var resp chResp
	err, _ = utils.GetAndJsonParse(c.baseUrl+"/ajax/image/list/chap/"+chapid+"?mode=horizontal&quality=high&hozPageSize=1", &resp)
	if err != nil {
		return nil, err
	}

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(resp.Html))

	res := []string{}

	doc.Find(".ds-image").Each(func(i int, s *goquery.Selection) {
		imgurl, ok := s.Attr("data-url")

		if !ok {
			return
		}

		if s.HasClass("shuffled") {
			res = append(res, "local://"+imgurl)
		} else {
			res = append(res, imgurl)
		}
	})

	return res, nil
}

func (c mangaReader) Image(manga, chapter, id string, w io.Writer) error {
	println(manga, chapter, id)

	origin, err := http.Get(id)
	if err != nil {
		return err
	}

	defer origin.Body.Close()

	input_image, err := jpeg.Decode(origin.Body)
	if err != nil {
		panic(err)
	}

	output_image := mangareader_unshuffle.Unshuffle(input_image)

	err = jpeg.Encode(w, output_image, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c mangaReader) Latest() ([]types.LatestResult, error) {
	return []types.LatestResult{}, nil
}

func (c mangaReader) Popular() ([]types.PopularResult, error) {
	return []types.PopularResult{}, nil
}
