package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func handleGetMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleGetMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func handlecreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlecreateMovie")
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = len(movies) + 1
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func handleUpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleUpdateMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if strconv.Itoa(item.ID) == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	} 
}

func handleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleDeleteMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if strconv.Itoa(item.ID) == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
