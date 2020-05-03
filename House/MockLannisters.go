package House

import "fmt"

type MockLannisters struct {
	Leader string
}
func (l *MockLannisters) Fight()string {
	return fmt.Sprintf("Mock Lannisters : Lannisters assemble")
}