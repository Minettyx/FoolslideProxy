package pathhandler

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type jWTHandler struct{}

var JWTHandler = jWTHandler{}

var SIGN_TOKEN = os.Getenv("SIGN_TOKEN")

func (*jWTHandler) MangaPath(modId string, id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mod":   modId,
		"manga": id,
	})

	tokenString, err := token.SignedString([]byte(SIGN_TOKEN + "manga"))
	if err != nil {
		panic(err)
	}

	return "jwt-" + tokenString
}

func (*jWTHandler) ChapterPath(modId string, manga string, id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mod":     modId,
		"manga":   manga,
		"chapter": id,
	})

	tokenString, _ := token.SignedString([]byte(SIGN_TOKEN + "chapter"))

	return "jwt-" + tokenString
}

func (*jWTHandler) ImagePath(modId string, manga string, chapter string, id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mod":     modId,
		"manga":   manga,
		"chapter": chapter,
		"image":   id,
	})

	tokenString, _ := token.SignedString([]byte(SIGN_TOKEN + "image"))

	return "jwt-" + tokenString
}

func (*jWTHandler) ParseMangaPath(path string) (*MangaPath, error) {

	if !strings.HasPrefix(path, "jwt-") {
		return nil, fmt.Errorf("Invalid path")
	}
	path = path[4:]

	token, err := jwt.Parse(path, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SIGN_TOKEN + "manga"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return &MangaPath{
			ModId:   claims["mod"].(string),
			MangaId: claims["manga"].(string),
		}, nil

	} else {
		return nil, err
	}
}

func (*jWTHandler) ParseChapterPath(path string) (*ChapterPath, error) {
	if !strings.HasPrefix(path, "jwt-") {
		return nil, fmt.Errorf("Invalid path")
	}
	path = path[4:]

	token, err := jwt.Parse(path, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SIGN_TOKEN + "chapter"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return &ChapterPath{
			ModId:     claims["mod"].(string),
			MangaId:   claims["manga"].(string),
			ChapterId: claims["chapter"].(string),
		}, nil

	} else {
		return nil, err
	}
}

func (*jWTHandler) ParseImagePath(path string) (*ImagePath, error) {
	if !strings.HasPrefix(path, "jwt-") {
		return nil, fmt.Errorf("Invalid path")
	}
	path = path[4:]

	token, err := jwt.Parse(path, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SIGN_TOKEN + "image"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return &ImagePath{
			ModId:     claims["mod"].(string),
			MangaId:   claims["manga"].(string),
			ChapterId: claims["chapter"].(string),
			ImageId:   claims["image"].(string),
		}, nil

	} else {
		return nil, err
	}
}
