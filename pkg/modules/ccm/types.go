package ccm

// Latest request

type ccmLatestReq struct {
	// Chapter string
	// Volume  string
	// Title   string
	Time  int64
	Manga ccmLatestReq_manga
}

type ccmLatestReq_manga struct {
	Id    string
	Title string
	Cover string
}

// Mangas request

type ccmMangasReq struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
	Id    string `json:"id"`
}

// Manga request

type ccmMangaReq struct {
	Id string `json:"id"`
	// Title    string                `json:"title"`
	Cover string `json:"cover"`
	// Status   int                   `json:"status"`
	Author   string                `json:"author"`
	Artist   string                `json:"artist"`
	Chapters []ccmMangaReq_chapter `json:"chapters"`
}

type ccmMangaReq_chapter struct {
	Chapter string `json:"chapter"`
	Volume  string `json:"volume"`
	Title   string `json:"title"`
	Time    int64  `json:"time"`
}

// Chapter request

type ccmChapterReq struct {
	// Chaper  string   `json:"chapter"`
	// Volume  string   `json:"volume"`
	// Title   string   `json:"title"`
	// Webtoon bool     `json:"webtoon"`
	Images []string `json:"images"`
}

// type ccmChapterReq_manga struct {
// 	Id       string                `json:"id"`
// 	Title    string                `json:"title"`
// 	Cover    string                `json:"cover"`
// 	Chapters []ccmMangaReq_chapter `json:"chapters"`
// }
//
// type ccmChapterReq_manga_chapter struct {
// 	Chapter string `json:"chapter"`
// 	Volume  string `json:"volume"`
// 	Title   string `json:"title"`
// }
