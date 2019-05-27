package mock

// Calculator is a mock implementation of note.Calculator
type Calculator struct{}

// Calculate calculates importance by the length of title
func (c *Calculator) Calculate(title string) int {
	return len(title)
}
