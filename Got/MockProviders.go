package main

import "wire-poc/House"

func ProvideMockBeyondWall()*House.MockBeyondWall {
	n := new(House.MockBeyondWall)
	return n
}

func ProvideMockStarks() *House.MockStark {
	n := new(House.MockStark)
	return n
}

func ProvideMockLannisters() *House.MockLannisters{
	n := new(House.MockLannisters)
	return n
}