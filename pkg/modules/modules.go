package modules

import (
	"github.com/Minettyx/FoolslideProxy/pkg/generic/pizzareader"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/ccm"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/juinjutsu"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/mangareader"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/mangaworld"
	"github.com/Minettyx/FoolslideProxy/pkg/modules/onepiecepower"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
)

var Modules = [...]types.Module{
	nil, // local module to bypass dependencies cycle, initialized in ./local/local.go
	ccm.Module,
	juinjutsu.Module,
	mangaworld.Module,
	onepiecepower.Module,
	mangareader.Module,

	pizzareader.TuttoAnimeManga,
	pizzareader.HastaTeam,
	pizzareader.DDTHastaTeam,
	pizzareader.PhoenixScans,
}
