package somefancydb

import (
	"time"

	"github.com/savolro/awsome/internal/note"
)

// Storage is a note storage implementation which uses very fancy db
type Storage struct {
	Notes []note.Note
}

// SaveNote saves node into slice of notes
func (s *Storage) SaveNote(note note.Note) (err error) {
	time.Sleep(time.Second * 2)
	s.Notes = append(s.Notes, note)
	return nil
}
