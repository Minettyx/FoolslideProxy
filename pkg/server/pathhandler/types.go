package pathhandler

type MangaPath struct {
	ModId   string
	MangaId string
}

type ChapterPath struct {
	ModId     string
	MangaId   string
	ChapterId string
}

type ImagePath struct {
	ModId     string
	MangaId   string
	ChapterId string
	ImageId   string
}

type PathHandler interface {
	MangaPath(modId string, id string) string
	ChapterPath(modId string, manga string, id string) string
	ImagePath(modId string, manga string, chapter string, id string) string

	ParseMangaPath(path string) (*MangaPath, error)
	ParseChapterPath(path string) (*ChapterPath, error)
	ParseImagePath(path string) (*ImagePath, error)
}
