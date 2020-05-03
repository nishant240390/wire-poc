package main

import (
	"wire-poc/House"
)

func ProvideStarks() House.Stark {
	return House.Stark{Leader:"John"}
}

func ProvideLannisters() House.Lannisters{
	return House.Lannisters{Leader:"Cersie"}
}

func ProvideBeyondWall( )*House.BeyondWall{
	b := new(House.BeyondWall)
	b.Leader = "Real knight king"
	return b
}

func NewWar(h1 House.Stark , h2 House.Lannisters, h3 House.IHouse) War{
	return War{House1:h1,House2:h2, House3:h3}
}
