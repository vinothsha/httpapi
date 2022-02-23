package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type book struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

var Books []book

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST GET DELETE UPDATE PUT")
		next.ServeHTTP(w, r)
		return
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to bookshop")
}

func allBooks(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Books)
}
func delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	delId := vars["id"]
	for index, value := range Books {
		if value.ID == delId {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Books)

}
func one(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	getId := vars["id"]
	for _, value := range Books {
		if value.ID == getId {
			json.NewEncoder(w).Encode(value)
			break
		}
	}

}

func create(w http.ResponseWriter, r *http.Request) {
	var Book book
	_ = json.NewDecoder(r.Body).Decode(&Book)
	Book.ID = strconv.Itoa(rand.Intn(1000))
	Books = append(Books, Book)
	json.NewEncoder(w).Encode(Book)
}
func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update books")
}

func main() {
	Books = append(Books, book{"1", "shaBook", "2100"})
	Books = append(Books, book{"2", "lakBook", "1100"})
	Books = append(Books, book{"3", "vickyBook", "4100"})
	Books = append(Books, book{"4", "naguBook", "1600"})
	Books = append(Books, book{"5", "mohanBook", "5100"})
	router := mux.NewRouter()
	router.Use(CORS)
	router.HandleFunc("/", home)
	router.HandleFunc("/all", allBooks).Methods("GET")
	router.HandleFunc("/delete/{id}", delete).Methods("DELETE")
	router.HandleFunc("/one/{id}", one).Methods("GET")
	router.HandleFunc("/create", create).Methods("POST")
	router.HandleFunc("/update/{id}", update).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5000", router))
}
