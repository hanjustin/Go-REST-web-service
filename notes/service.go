package notes

import (
	"sync"

	"github.com/google/uuid"
)

type mockService struct {
	sync.Mutex
	memoryStore map[string]Note
}

func NewMockService() *mockService {
	return &mockService{
		memoryStore: make(map[string]Note),
	}
}

func (s *mockService) Create(note Note) (Note, error) {
	s.Lock()
	defer s.Unlock()

	if note.Title == "" {
		return Note{}, ErrNoteNoTitle
	}

	uid, _ := uuid.NewV7()
	note.ID = uid.String()
	s.memoryStore[note.ID] = note
	return note, nil
}

func (s *mockService) ReadAll() ([]Note, error) {
	s.Lock()
	defer s.Unlock()

	notes := make([]Note, 0, len(s.memoryStore))
	for _, note := range s.memoryStore {
		notes = append(notes, note)
	}

	return notes, nil
}

func (s *mockService) Read(id string) (Note, error) {
	s.Lock()
	defer s.Unlock()

	if note, ok := s.memoryStore[id]; !ok {
		return Note{}, ErrNoteNotFound
	} else {
		return note, nil
	}
}

func (s *mockService) Update(id string, note Note) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.memoryStore[id]; !ok {
		return ErrNoteNotFound
	}

	s.memoryStore[id] = note
	return nil
}

func (s *mockService) Delete(id string) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.memoryStore[id]; !ok {
		return ErrNoteNotFound
	}

	delete(s.memoryStore, id)
	return nil
}
