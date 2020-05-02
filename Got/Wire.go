// +build wireinject

package main

func Initialise()(War,error){
	wire.Build(NewWar, ProvideStarks, ProvideLannisters)
	return War{},nil
}
func ProvideStarks() House.Stark {
	return House.Stark{}
}
func ProvideLannisters() House.Lannisters{
	return House.Lannisters{}
}

func NewWar(h1 House.Stark , h2 House.Lannisters) War{
	return War{house1:h1,house2:h2}
}
