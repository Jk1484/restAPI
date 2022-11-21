package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jk1484/restAPI/handlers"
)

func main() {
	mux := http.NewServeMux()

	h := handlers.New()

	mux.HandleFunc("/", greet)
	mux.HandleFunc("/create", h.CreateBook)
	mux.HandleFunc("/getAll", h.GetAllBooks)
	mux.HandleFunc("/get", h.GetByID)
	mux.HandleFunc("/update", h.UpdateBookByID)
	mux.HandleFunc("/delete", h.DeleteByID)

	fmt.Println("server is running on port: 8081")

	log.Fatalln(http.ListenAndServe(":8081", mux))
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In memory rest api!!! :D"))
}
