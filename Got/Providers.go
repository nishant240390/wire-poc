package got

import (
	"wire-poc/House"
)

func ProvideStarks() House.Stark {
	return House.Stark{Leader:"John"}
}

func ProvideLannisters() House.Lannisters{
	return House.Lannisters{Leader:"Cersie"}
}

func ProvideBeyondWall() House.BeyondWall{
	return House.BeyondWall{Leader:"NigthKing"}
}
func ProvideWarZone(war *War) *WarZone {
	return &WarZone{name:"CasteRock",War:war}
}

func ProvideWar(h1 House.Stark , h2 House.Lannisters, h3 House.BeyondWall, zone *WarZone) War{
	return War{House1:h1,House2:h2, House3:h3, WarZone:zone}
}
