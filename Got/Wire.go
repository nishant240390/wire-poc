//+build wireinject


package main

import (
	"github.com/google/wire"
	"wire-poc/House"
)

func Initialize(stark House.Stark, lannisters House.Lannisters)(War,error){
	wire.Build(NewWar)
	return War{},nil
}
func NewStarkProvider(name string)(House.Stark,error){
	wire.Build(ProvideStarks)
	return House.Stark{},nil
}

func NewLannisterProvider(name string)(House.Lannisters,error){
	wire.Build(ProvideLannisters)
	return House.Lannisters{},nil
}
