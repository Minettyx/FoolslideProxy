package types

import "time"

type PopularResult struct {
	Id         string
	Title      string
	Image      string
	Popularity int
}

type LatestResult struct {
	Id    string
	Title string
	Image string
	Date  time.Time
}

type SearchResult struct {
	Id    string
	Title string
	Image string
}

type Manga struct {
	Synopsis  string
	Author    string
	Artist    string
	Img       string
	Chapters  []Chapter
	Sourceurl string
}

type Chapter struct {
	Title string
	Id    string
	Date  time.Time
}
