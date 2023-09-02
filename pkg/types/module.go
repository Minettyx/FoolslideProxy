package types

import "io"

type Module struct {
	Id    string
	Name  string
	Flags ModuleFlags

	Popular func() ([]PopularResult, error)
	Latest  func() ([]LatestResult, error)
	Search  func(query string, language *string) ([]SearchResult, error)
	Manga   func(id string) (*Manga, error)
	Chapter func(manga string, id string) ([]string, error)
	Image   func(manga string, chapter string, id string, w io.Writer) error
}
