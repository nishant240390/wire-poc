//+build wireinject


package main

import (
	"github.com/google/wire"
	"wire-poc/House"
)

func Initialize()(War,error){
	wire.Build(
			ProvideBeyondWall,
			wire.Bind(new(House.IHouse), new(*House.BeyondWall)),
			NewStarkProvider,
			NewLannisterProvider,
			NewWar)
	return War{},nil
}

func InitializeMock()(War,error){
	wire.Build(
		wire.NewSet(
			ProvideMockBeyondWall,
			wire.Bind(new(House.IHouse), new(*House.MockBeyondWall)),
			NewStarkProvider,
			NewLannisterProvider,
			NewWar))
	return War{},nil
}

func NewStarkProvider()(House.Stark,error){
	wire.Build(ProvideStarks)
	return House.Stark{},nil
}

func NewLannisterProvider()(House.Lannisters,error){
	wire.Build(ProvideLannisters)
	return House.Lannisters{},nil
}


