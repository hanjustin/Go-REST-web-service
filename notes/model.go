package notes

import "errors"

var (
	ErrNoteNotFound = errors.New("no records found")
	ErrNoteNoTitle  = errors.New("note doesn't have title")
)

type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
