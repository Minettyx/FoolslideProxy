package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func StrToHex(str string) string {
	return hex.EncodeToString([]byte(str))
}

func HexToStr(str string) (string, error) {
	bs, err := hex.DecodeString(str)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func AuthorArtist(author string, artist string) string {
	if author == "" && artist == "" {
		return ""
	} else if artist == "" {
		return author
	} else if author == "" {
		return artist
	} else {
		if author == artist {
			return author
		}
		return fmt.Sprintf("%v, %v", author, artist)
	}
}

// func IsHex(str string) bool {
// 	re := regexp.MustCompile(`^[0-9a-fA-F]+$`)
// 	return re.MatchString(str)
// }

// func ImageProxyUrl(module_id string, manga string, chapter string, id string) string {
// 	var manga_hex string
// 	if manga == "" {
// 		manga_hex = "00"
// 	} else {
// 		manga_hex = StrToHex(manga)
// 	}
//
// 	var chapter_hex string
// 	if chapter == "" {
// 		chapter_hex = "00"
// 	} else {
// 		chapter_hex = StrToHex(chapter)
// 	}
//
// 	var id_hex string
// 	if id == "" {
// 		id_hex = "00"
// 	} else {
// 		id_hex = StrToHex(id)
// 	}
//
// 	return fmt.Sprintf("/image/%v-%v-%v-%v", module_id, manga_hex, chapter_hex, id_hex)
// }

// does a get request to the url, than calls json.Unmarshal with the body and the second argument
// returns possible errors and a bool that is true if the response status code is 404
func GetAndJsonParse(rurl string, v any) (error, bool) {
	is404 := false

	res, err := http.Get(rurl)
	if err != nil {
		return err, false
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
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

func StrConaintsIgnoreCase(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}

func GenTitle(volume string, chapter string, title string) string {
	vol := ""
	if volume != "" {
		vol = "Vol." + volume + " "
	}

	ch := "Ch." + chapter

	tit := ""
	if title != "" {
		tit = " - " + title
	}

	return vol + ch + tit
}

func StrBetween(str string, left string, right string) (string, error) {
	pts := strings.Split(str, left)

	if len(pts) < 2 {
		return "", fmt.Errorf("Left not found")
	}

	res := strings.Split(pts[len(pts)-1], right)[0]

	return res, nil
}

func GetAndGoquery(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
