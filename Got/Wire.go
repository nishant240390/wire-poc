//+build wireinject

package main

import (
	"github.com/google/wire"
	"wire-poc/House"
)

func Initialize()(War,error){
	wire.Build(
		    NewBeyondWallProvider,
			NewStarkProvider,
			NewLannisterProvider,
			NewWar)
	return War{},nil
}

func InitializeMock()(MockWar,error){
	wire.Build(
		wire.NewSet(
			ProvideMockBeyondWall,
			wire.Bind(new(House.IHouse), new(*House.MockBeyondWall)),
			NewStarkProvider,
			NewLannisterProvider,
			NewMockWar))
	return MockWar{},nil
}

func NewStarkProvider()(House.Stark,error){
	wire.Build(ProvideStarks)
	return House.Stark{},nil
}

func NewLannisterProvider()(House.Lannisters,error){
	wire.Build(ProvideLannisters)
	return House.Lannisters{},nil
}
func NewBeyondWallProvider()(House.BeyondWall,error){
	wire.Build(ProvideBeyondWall)
	return House.BeyondWall{},nil
}


