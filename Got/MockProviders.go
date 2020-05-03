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

func NewMockWar(h1 House.Stark , h2 House.Lannisters, h3 House.IHouse) MockWar{
	return MockWar{house1:h1,house2:h2, house3:h3}
}