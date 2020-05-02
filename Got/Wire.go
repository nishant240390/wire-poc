
// +build wireinject

package main

import "github.com/google/wire"

func Initialise()(War,error){
	wire.Build(NewWar, ProvideStarks, ProvideLannisters)
	return War{},nil
}

