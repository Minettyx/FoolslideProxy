package types

import "io"

type Module interface {
	Id() string
	Name() string
	Flags() ModuleFlags

	Popular() ([]PopularResult, error)
	Latest() ([]LatestResult, error)
	Search(query string) ([]SearchResult, error)
	Manga(id string) (*Manga, error)
	Chapter(manga string, id string) ([]string, error)
}

type ImageProxyModule interface {
	Module
	Image(manga string, chapter string, id string, w io.Writer) error
}
