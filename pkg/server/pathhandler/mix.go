package pathhandler

type mixHandler struct{}

var MixHandler = mixHandler{}

var _ PathHandler = mixHandler{}

func (mixHandler) MangaPath(modId string, id string) string {
	return HexHandler.MangaPath(modId, id)
}

func (mixHandler) ChapterPath(modId string, manga string, id string) string {
	return HexHandler.ChapterPath(modId, manga, id)
}

func (mixHandler) ImagePath(modId string, manga string, chapter string, id string) string {
	return JWTHandler.ImagePath(modId, manga, chapter, id)
}

func (mixHandler) ParseMangaPath(path string) (*MangaPath, error) {
	return HexHandler.ParseMangaPath(path)
}

func (mixHandler) ParseChapterPath(path string) (*ChapterPath, error) {
	return HexHandler.ParseChapterPath(path)
}

func (mixHandler) ParseImagePath(path string) (*ImagePath, error) {
	return JWTHandler.ParseImagePath(path)
}
