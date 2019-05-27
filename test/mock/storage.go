package mock

import (
	"github.com/savolro/awsome/internal/note"
)

// Storage is a mock implementation of note.Storage
type Storage struct{}

// SaveNote makes you think it saved something
func (s *Storage) SaveNote(note note.Note) (err error) {
	return nil
}
