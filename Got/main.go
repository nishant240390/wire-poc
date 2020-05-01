package main

import (
	"log"
	"wire-poc/House"
)

func main(){

     war,err := Initialise()
     if err != nil {
     	log.Fatal(err)
	 }
	 war.startWar()
}
func provideStarks () House.Stark {
	return House.Stark{}
}
func provideLannisters() House.Lannisters{
	return House.Lannisters{}
}
func NewWar (h1 House.Stark , h2 House.Lannisters) War{
	return War{house1:h1,house2:h2}
}
