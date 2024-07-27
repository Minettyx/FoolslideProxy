package modules

import (
	"github.com/Minettyx/FoolslideProxy/pkg/modules/ccm"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/hastateam"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/juinjutsu"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/mangareader"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/mangaworld"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/onepiecepower"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/tuttoanimemanga"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
)

var Modules = [...]*types.Module{
	nil, // local module to bypass dependencies cycle, initialized in server/router.go
	&ccm.CCM,
	&juinjutsu.JuinJutsu,
	&mangaworld.MangaWorld,
	tuttoanimemanga.TuttoAnimeManga,
	&onepiecepower.OnePiecePower,
	&mangareader.MangaReader,
	hastateam.HastaTeam,
	hastateam.DDTHastaTeam,
}
