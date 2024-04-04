package notes

import (
	"encoding/json"
	"errors"
	"net/http"
)

type NoteHandler struct {
	Service NoteService
}

type NoteService interface {
	Create(note Note) (Note, error)
	ReadAll() ([]Note, error)
	Read(id string) (Note, error)
	Update(id string, note Note) error
	Delete(id string) error
}

func (h *NoteHandler) Post(w http.ResponseWriter, r *http.Request) {
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newNote, err := h.Service.Create(note)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, ErrNoteNoTitle) {
			statusCode = http.StatusBadRequest
		}

		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	noteJSON, err := json.Marshal(newNote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(noteJSON)
}

func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	notes, err := h.Service.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	notesJSON, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(notesJSON)
}

func (h *NoteHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	note, err := h.Service.Read(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, ErrNoteNotFound) {
			statusCode = http.StatusBadRequest
		}

		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	noteJSON, err := json.Marshal(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(noteJSON)
}

func (h *NoteHandler) Put(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var note Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Service.Update(id, note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
