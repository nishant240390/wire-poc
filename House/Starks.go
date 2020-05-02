package House

import "fmt"

type Stark struct {
   Leader string
}
func (s *Stark) Fight()string {
   return fmt.Sprintf("%s : stark assemble",s.Leader)
}