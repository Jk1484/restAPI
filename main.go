package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var Books struct {
	Books    []Book
	IDsCount int
}

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", CreateBook)
	mux.HandleFunc("/getAll", GetAllBooks)
	mux.HandleFunc("/get", GetByID)
	mux.HandleFunc("/update", UpdateBookByID)
	mux.HandleFunc("/delete", DeleteByID)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	Books.IDsCount++
	b.ID = Books.IDsCount
	Books.Books = append(Books.Books, b)

	w.Write([]byte("Created"))
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	m, err := json.Marshal(Books.Books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(m)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	var b Book
	json.NewDecoder(r.Body).Decode(&b)
	if b.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not provided"))
		return
	}

	found := false
	for _, v := range Books.Books {
		if v.ID == b.ID {
			found = true
			b = v
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not found"))
		return
	}

	m, _ := json.Marshal(b)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(m)
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	var b Book
	json.NewDecoder(r.Body).Decode(&b)
	if b.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not provided"))
		return
	}

	found := false
	for i, v := range Books.Books {
		if v.ID == b.ID {
			found = true

			Books.Books = append(Books.Books[:i], Books.Books[i+1:]...)
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	var b Book
	json.NewDecoder(r.Body).Decode(&b)
	if b.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not provided"))
		return
	}

	found := false
	for i, v := range Books.Books {
		if v.ID == b.ID {
			found = true

			Books.Books[i] = b
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("updated"))
}
