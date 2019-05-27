package note

// Storage is responsible for manipulating data of Notes
type Storage interface {
	// SaveNote saves Note into storage
	SaveNote(note Note) (err error)
}
