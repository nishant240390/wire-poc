package got

import (
	"fmt"
	"wire-poc/House"
)

type War struct{
	House1 House.Stark
	House2 House.Lannisters
	House3 House.IHouse
}
func (w *War) startWar() {
	fmt.Printf("%s\n%s\n%s",w.House1.Fight(),w.House2.Fight(),w.House3.Fight())
}