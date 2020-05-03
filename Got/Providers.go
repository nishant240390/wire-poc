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

func ProvideBeyondWall( )House.BeyondWall{
	return House.BeyondWall{Leader:"NightKing"}
}

func NewWar(h1 House.Stark , h2 House.Lannisters, h3 House.BeyondWall) War{
	return War{house1:h1,house2:h2, house3:h3}
}
