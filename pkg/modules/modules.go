package modules

import (
	"foolslideproxy/pkg/modules/ccm"
	"foolslideproxy/pkg/modules/juinjutsu"
	"foolslideproxy/pkg/modules/mangareader"
	"foolslideproxy/pkg/modules/mangaworld"
	"foolslideproxy/pkg/modules/onepiecepower"
	"foolslideproxy/pkg/modules/tuttoanimemanga"
	"foolslideproxy/pkg/types"
)

var Modules = [...]*types.Module{
	nil, // local module to bypass dependencies cycle, initialized in server/router.go
	&ccm.CCM,
	&juinjutsu.JuinJutsu,
	&mangaworld.MangaWorld,
	&tuttoanimemanga.TuttoAnimeManga,
	&onepiecepower.OnePiecePower,
	&mangareader.MangaReader,
}
