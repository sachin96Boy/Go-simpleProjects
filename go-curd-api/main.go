package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Director *Director `json:"director"`
}

type Director struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies = []Movie{
	{ID: 1, Title: "Casablanca", Year: 1942, Director: &Director{ID: 1, FirstName: "Michael", LastName: "Chabon"}},
	{ID: 2, Title: "Cool Hand Luke", Year: 1967, Director: &Director{ID: 2, FirstName: "Fyodor", LastName: "Dostoevsky"}},
	{ID: 3, Title: "Bullitt", Year: 1968, Director: &Director{ID: 3, FirstName: "Frank", LastName: "Herbert"}},
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/movies", handleGetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handleGetMovie).Methods("GET")
	r.HandleFunc("/movies", handlecreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handleUpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handleDeleteMovie).Methods("DELETE")

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleGetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleGetMovies")
}

func handleGetMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleGetMovie")
}

func handlecreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlecreateMovie")
}

func handleUpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleUpdateMovie")
}

func handleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleDeleteMovie")
}
