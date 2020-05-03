package House
import "fmt"

type MockBeyondWall struct {
	 Leader string
}
func (b *MockBeyondWall) Fight()string {
	return fmt.Sprintf("Mock NightKing   BeyondWall assemble")
}
