package main

import "wire-poc/House"

func ProvideMockBeyondWall()*House.MockBeyondWall {
	n := new(House.MockBeyondWall)
	return n
}
