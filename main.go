package main

import (
	"log"
	"net/http"

	"github.com/Jk1484/restAPI/handlers"
)

func main() {
	mux := http.NewServeMux()

	h := handlers.New()

	mux.HandleFunc("/create", h.CreateBook)
	mux.HandleFunc("/getAll", h.GetAllBooks)
	mux.HandleFunc("/get", h.GetByID)
	mux.HandleFunc("/update", h.UpdateBookByID)
	mux.HandleFunc("/delete", h.DeleteByID)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
