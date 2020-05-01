package main

import "wire-poc/House"

type War struct{
	house1 House.Stark
	house2 House.Lannisters
}
func (w *War) startWar() {
	w.house1.Fight()
	w.house2.Fight()
}