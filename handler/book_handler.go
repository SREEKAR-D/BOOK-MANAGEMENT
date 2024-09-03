package handler

import (
	"encoding/json"
	"golang2/service"
	"net/http"

	uuid "github.com/google/uuid"

	"github.com/go-chi/chi"
)

type BookHandler struct {
	service service.Bookservice
}

func NewBookHandler(s service.Bookservice) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAllBooks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idstr)

	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book service.BookDTO

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Cannot Decode", http.StatusInternalServerError)
		return
	}

	err = h.service.AddBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New Book Added"))
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book service.BookDTO

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Error Decoding", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateBook(id, book)
	if err != nil {
		http.Error(w, "Error Updating", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book Data Updated Successfully"))
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBook(id)
	if err != nil {
		http.Error(w, "Error Deleting", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book Data Deleted Successfully"))
}
