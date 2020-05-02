package House
import "fmt"

type BeyondWall struct {
	Leader string
}
func (b *BeyondWall) Fight()string {
	return fmt.Sprintf("%s : BeyondWall assemble",b.Leader)
}
