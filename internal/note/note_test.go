package note_test

import (
	"testing"

	"github.com/savolro/awsome/internal/note/artificialinteligencesuperduper"
	"github.com/savolro/awsome/internal/note/somefancydb"
	"github.com/savolro/awsome/test/mock"

	"github.com/savolro/awsome/internal/note"
	"github.com/stretchr/testify/assert"
)

// This is a mocked unit test that uses mock implementations
func TestNoter_AddNote(t *testing.T) {
	DoTestNoterAddNote(t, &mock.Storage{}, &mock.Calculator{})
}

// This is an integration test that uses our special implementations
func TestNoter_AddNote_Integration(t *testing.T) {
	// You can also add some setup here to ignore this on unit test execution
	DoTestNoterAddNote(t, &somefancydb.Storage{}, &artificialinteligencesuperduper.Calculator{})
}

// This is an integration test which uses our fancy, long-running implementations

func DoTestNoterAddNote(t *testing.T, storage note.Storage, calculator note.Calculator) {
	// You could also create a constructor for Noter
	noter := note.Noter{
		Storage:    storage,
		Calculator: calculator,
	}

	title := "good title"
	text := "bad content"

	n, err := noter.AddNote(title, text)
	assert.NoError(t, err)
	assert.Equal(t, title, n.Title)
	assert.Equal(t, text, n.Text)
}
