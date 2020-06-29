
// +build wireinject

package got

import (
	"github.com/google/wire"
	"wire-poc/House"
)

func Initialize()(War,error){
	wire.Build(
		NewStarkProvider,
		NewLannisterProvider,
		NewBeyondWallProvider,
		NewWarZoneProvder,
	    wire.Struct(new(War),"*"))
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
func NewBeyondWallProvider()(House.BeyondWall,error){
	wire.Build(ProvideBeyondWall)
	return House.BeyondWall{},nil
}
func NewWarZoneProvder(war *War)(*WarZone,error){
	wire.Build(ProvideWarZone)
	return &WarZone{},nil
}



