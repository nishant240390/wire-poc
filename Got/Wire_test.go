package got

import (
	"testing"
	"wire-poc/House"
)
//
//func TestInitialiseUsingWire(t *testing.T) {
//	tests := []struct {
//		name    string
//	}{
//		{
//			name: "name",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := Initialize()
//			if  err != nil{
//				t.Errorf("Initialise() error ")
//				return
//			}
//			got.startWar()
//		})
//	}
//}
//func TestInitialiseWithoutUsingWireForError(t *testing.T) {
//	tests := []struct {
//		name    string
//	}{
//		{
//			name: "name",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			stark := House.Stark{Leader:"john"}
//			lannister := House.Lannisters{"cersie"}
//			beyond := House.BeyondWall{"NightKing"}
//			war := War{House1:stark,House2:lannister,House3:beyond }
//			war.startWar()
//		})
//	}
//}
func TestInitialiseWithoutUsingWireForCircular(t *testing.T) {
	tests := []struct {
		name    string
	}{
		{
			name: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stark := House.Stark{Leader:"john"}
			lannister := House.Lannisters{"cersie"}
			beyond := House.BeyondWall{"NightKing"}
			zone :=  WarZone{name:"casterock", War:nil}
			war := War{House1:stark,House2:lannister,House3:beyond,WarZone:&zone }
			war.startWar()
		})
	}
}
