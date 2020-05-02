// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from Wire.go:

func Initialise() (War, error) {
	stark := ProvideStarks()
	lannisters := ProvideLannisters()
	war := NewWar(stark, lannisters)
	return war, nil
}
