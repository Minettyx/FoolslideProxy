package pizzareader

import (
	"github.com/Minettyx/FoolslideProxy/pkg/types"
)

var DDTHastaTeam = pizzaReader{
	moduleId:    "ddtht",
	moduleName:  "DDT HastaTeam",
	moduleFlags: []types.ModuleFlag{},
	baseUrl:     "https://ddt.hastateam.com",
}

var HastaTeam = pizzaReader{
	moduleId:    "ht",
	moduleName:  "HastaTeam",
	moduleFlags: []types.ModuleFlag{},
	baseUrl:     "https://reader.hastateam.com",
}

var TuttoAnimeManga = pizzaReader{
	moduleId:    "tam",
	moduleName:  "TuttoAnimeManga",
	moduleFlags: []types.ModuleFlag{},
	baseUrl:     "https://tuttoanimemanga.net",
}

var PhoenixScans = pizzaReader{
	moduleId:    "ps",
	moduleName:  "Phoenix Scans",
	moduleFlags: []types.ModuleFlag{},
	baseUrl:     "https://www.phoenixscans.com",
}
