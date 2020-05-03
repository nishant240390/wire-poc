package House

import "fmt"

type MockStark struct {
	Leader string
}
func (s *MockStark) Fight()string {
	return fmt.Sprintf("Mock John SNow : stark assemble")
}