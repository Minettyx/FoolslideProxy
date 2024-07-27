package hastateam

import (
	"github.com/Minettyx/FoolslideProxy/pkg/generic"
	"github.com/Minettyx/FoolslideProxy/pkg/generic/type0"
	"github.com/Minettyx/FoolslideProxy/pkg/types"
)

var DDTHastaTeam = type0.Type0(generic.GenericConfig{
	Id:      "ddtht",
	Name:    "DDT HastaTeam",
	Flags:   []types.ModuleFlag{},
	BaseUrl: "https://ddt.hastateam.com",
})

var HastaTeam = type0.Type0(generic.GenericConfig{
	Id:      "ht",
	Name:    "HastaTeam",
	Flags:   []types.ModuleFlag{},
	BaseUrl: "https://reader.hastateam.com",
})
