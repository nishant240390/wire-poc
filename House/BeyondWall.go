package House
import "fmt"

type BeyondWall struct {
}
func (l *BeyondWall) Fight()string {
	return fmt.Sprintf("%s : BeyondWall assemble")
}
