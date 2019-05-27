package note

// Calculator is responsible for finding importance of Note by it's title
type Calculator interface {
	// Calculate returns importance by title
	Calculate(title string) int
}
