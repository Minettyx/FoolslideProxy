package pathhandler

import (
	"fmt"
	"github.com/Minettyx/FoolslideProxy/pkg/utils"
	"regexp"
	"strings"
)

var hexRegex = regexp.MustCompile("^([0-9a-fA-F][0-9a-fA-F])+$")
var modRegex = regexp.MustCompile("^[a-zA-Z]+$")

type hexHandler struct{}

var HexHandler = hexHandler{}

var _ PathHandler = hexHandler{}

func (hexHandler) MangaPath(modId string, id string) string {
	return fmt.Sprintf("%v-%v", modId, utils.StrToHex(id))
}

func (hexHandler) ChapterPath(modId string, manga string, id string) string {
	return fmt.Sprintf("%v-%v-%v", modId, utils.StrToHex(manga), utils.StrToHex(id))
}

func (hexHandler) ImagePath(modId string, manga string, chapter string, id string) string {
	return fmt.Sprintf("%v-%v-%v-%v", modId, utils.StrToHex(manga), utils.StrToHex(chapter), utils.StrToHex(id))
}

func (hexHandler) ParseMangaPath(path string) (*MangaPath, error) {
	err := fmt.Errorf("Invalid path")
	p := strings.Split(path, "-")
	if len(p) < 2 {
		return nil, err
	}

	if !modRegex.MatchString(p[0]) {
		return nil, err
	}

	mangaid, err := utils.HexToStr(p[1])
	if err != nil {
		return nil, err
	}

	return &MangaPath{
		ModId:   p[0],
		MangaId: mangaid,
	}, nil
}

func (hexHandler) ParseChapterPath(path string) (*ChapterPath, error) {
	err := fmt.Errorf("Invalid path")
	p := strings.Split(path, "-")
	if len(p) < 3 {
		return nil, err
	}

	if !modRegex.MatchString(p[0]) {
		return nil, err
	}

	mangaid, err := utils.HexToStr(p[1])
	if err != nil {
		return nil, err
	}

	chapterid, err := utils.HexToStr(p[2])
	if err != nil {
		return nil, err
	}

	return &ChapterPath{
		ModId:     p[0],
		MangaId:   mangaid,
		ChapterId: chapterid,
	}, nil
}

func (hexHandler) ParseImagePath(path string) (*ImagePath, error) {
	err := fmt.Errorf("Invalid path")
	p := strings.Split(path, "-")
	if len(p) < 3 {
		return nil, err
	}

	if !modRegex.MatchString(p[0]) {
		return nil, err
	}

	mangaid, err := utils.HexToStr(p[1])
	if err != nil {
		return nil, err
	}

	chapterid, err := utils.HexToStr(p[2])
	if err != nil {
		return nil, err
	}

	imageid, err := utils.HexToStr(p[3])
	if err != nil {
		return nil, err
	}

	return &ImagePath{
		ModId:     p[0],
		MangaId:   mangaid,
		ChapterId: chapterid,
		ImageId:   imageid,
	}, nil
}
