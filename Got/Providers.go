package main

import (
	"wire-poc/House"
)

func ProvideStarks(name string) House.Stark {
	return House.Stark{Leader:name}
}

func ProvideLannisters() House.Lannisters{
	return House.Lannisters{Leader:"cersie"}
}

func NewWar(h1 House.Stark , h2 House.Lannisters) War{
	return War{house1:h1,house2:h2}
}
