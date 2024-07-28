package pizzareader

import (
	"github.com/Minettyx/FoolslideProxy/pkg/generic"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
)

var DDTHastaTeam = PizzaReader(generic.GenericConfig{
	Id:      "ddtht",
	Name:    "DDT HastaTeam",
	Flags:   []types.ModuleFlag{},
	BaseUrl: "https://ddt.hastateam.com",
})

var HastaTeam = PizzaReader(generic.GenericConfig{
	Id:      "ht",
	Name:    "HastaTeam",
	Flags:   []types.ModuleFlag{},
	BaseUrl: "https://reader.hastateam.com",
})

var TuttoAnimeManga = PizzaReader(generic.GenericConfig{
	Id:      "tam",
	Name:    "TuttoAnimeManga",
	Flags:   types.ModuleFlags{},
	BaseUrl: "https://tuttoanimemanga.net",
})

var PhoenixScans = PizzaReader(generic.GenericConfig{
	Id:      "ps",
	Name:    "Phoenix Scans",
	Flags:   types.ModuleFlags{},
	BaseUrl: "https://www.phoenixscans.com",
})
