package main

import (
	"log"
)

func main(){

     war,err := Initialise()
     if err != nil {
     	log.Fatal(err)
	 }
	 war.startWar()
}
