package handlers

import (
	"encoding/json"
	"net/http"
	"restAPI/internal/books"
)

type Handler interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateBookByID(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	BookRepository books.Repository
}

func New() Handler {
	return &handler{
		BookRepository: books.New(),
	}
}

func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var b books.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	h.BookRepository.CreateBook(b)

	w.Write([]byte("Created"))
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	m, err := json.Marshal(h.BookRepository.GetAllBooks())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(m)
}

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	var b books.Book
	json.NewDecoder(r.Body).Decode(&b)
	if b.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not provided"))
		return
	}

	b = h.BookRepository.GetByID(b)

	m, _ := json.Marshal(b)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(m)
}

func (h *handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	var b books.Book
	json.NewDecoder(r.Body).Decode(&b)
	if b.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not provided"))
		return
	}

	h.BookRepository.DeleteByID(b)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}

func (h *handler) UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	var b books.Book
	json.NewDecoder(r.Body).Decode(&b)
	if b.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not provided"))
		return
	}

	h.BookRepository.UpdateBookByID(b)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("updated"))
}
