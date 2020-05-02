package House

import "fmt"

type Lannisters struct {
	Leader string
}
func (l *Lannisters) Fight()string {
	return fmt.Sprintf("%s : Lannisters assemble",l.Leader)
}