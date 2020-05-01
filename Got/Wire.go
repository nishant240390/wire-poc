package main

import "github.com/google/wire"

// +build wireinject

func Initialise()(War,error){
	wire.Build(NewWar,provideStarks,provideLannisters)
	return War{},nil
}