package note

import "time"

// Note is an instance of some note
type Note struct {
	Title      string
	Text       string
	Time       time.Time
	Importance int
}

// Noter handles logic of note manipulation
type Noter struct {
	Storage    Storage
	Calculator Calculator
}

// AddNote saves note into storage with some calculated importance
func (n *Noter) AddNote(title string, text string) (note Note, err error) {
	note = Note{
		Title: title,
		Text:  text,
		Time:  time.Now(),
	}
	note.Importance = n.Calculator.Calculate(title)

	err = n.Storage.SaveNote(note)
	if err != nil {
		return note, err
	}

	return note, nil
}
