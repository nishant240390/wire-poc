package main

import (
	"fmt"
	"wire-poc/House"
)

type War struct{
	house1 House.Stark
	house2 House.Lannisters
}
func (w *War) startWar() {
	fmt.Printf("%s\n%s",w.house1.Fight(),w.house2.Fight())
}