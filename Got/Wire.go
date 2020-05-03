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

//func InitializeMock()(War,error){
//	wire.Build(
//		wire.NewSet(
//			ProvideMockBeyondWall,
//			ProvideMockStarks,
//			ProvideMockLannisters,
//			wire.Bind(new(House.IHouse), new(*House.MockBeyondWall)),
//			wire.Bind(new(House.IHouse), new(*House.MockStark)),
//			wire.Bind(new(House.IHouse), new(*House.MockLannisters)),
//			NewWar))
//	return War{},nil
//}

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


