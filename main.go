package main

import (
	"net/http"
	"rest-api-server/notes"
)

func main() {
	mock := notes.NewMockService()
	h := &notes.NoteHandler{
		Service: mock,
	}

	routes := initializeRoutes(h)
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		panic(err)
	}
}

func initializeRoutes(h *notes.NoteHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/notes", h.Post)
	mux.HandleFunc("GET /api/notes", h.GetAll)
	mux.HandleFunc("GET /api/notes/{id}", h.Get)
	mux.HandleFunc("PUT /api/notes/{id}", h.Put)
	mux.HandleFunc("DELETE /api/notes/{id}", h.Delete)
	return mux
}
