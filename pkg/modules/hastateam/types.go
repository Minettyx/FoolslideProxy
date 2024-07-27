package hastateam

type apiComics struct {
	Comics []struct {
		Title     string `json:"title"`
		Thumbnail string `json:"thumbnail"`
		Url       string `json:"url"`
	} `json:"comics"`
}

type apiComic struct {
	Comic struct {
		Description string `json:"description"`
		Thumbnail   string `json:"thumbnail"`
		Author      string `json:"author"`
		Artist      string `json:"artist"`
		Url         string `json:"url"`
		Chapters    []struct {
			Url       string `json:"url"`
			FullTitle string `json:"full_title"`
			UpdatedAt string `json:"updated_at"`
		} `json:"chapters"`
	} `json:"comic"`
}

type apiRead struct {
	Chapter struct {
		Pages []string `json:"pages"`
	} `json:"chapter"`
}
