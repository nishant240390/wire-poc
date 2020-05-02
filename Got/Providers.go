package main

import (
	"github.com/google/wire"
	"wire-poc/House"
)

func ProvideStarks(name string) House.Stark {
	return House.Stark{Leader:name}
}

func ProvideLannisters(name string) House.Lannisters{
	return House.Lannisters{Leader:name}
}
func ProvideBeyondWall() House.BeyondWall{
	return  House.BeyondWall{}
}

func ProvideBeyondWallSet() wire.ProviderSet{
	return wire.NewSet(ProvideBeyondWall, wire.Bind(new(House.IHouse),new(House.BeyondWall)))
}

func NewWar(h1 House.Stark , h2 House.Lannisters, h3 House.BeyondWall) War{
	return War{house1:h1,house2:h2, house3:h3}
}
