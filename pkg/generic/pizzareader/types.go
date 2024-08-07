package pizzareader

type apiComics struct {
	Comics []struct {
		Url       string `json:"url"`
		Thumbnail string `json:"thumbnail"`
		Title     string `json:"title"`
	} `json:"comics"`
}

type apiComic struct {
	Comic struct {
		Thumbnail   string `json:"thumbnail"`
		Description string `json:"description"`
		Author      string `json:"author"`
		Artist      string `json:"artist"`
		Chapters    []struct {
			FullTitle   string `json:"full_title"`
			Url         string `json:"url"`
			PublishedOn string `json:"published_on"`
		} `json:"chapters"`
	} `json:"comic"`
}

type apiRead struct {
	Chapter struct {
		Pages []string `json:"pages"`
	} `json:"chapter"`
}
