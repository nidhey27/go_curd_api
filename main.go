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

type Movies struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json"firstName"`
	LastName  string `json"lastName"`
}

var movies []Movies

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movies{ID: "1", Isbn: "3451", Title: "Movie 1", Director: &Director{FirstName: "Nidhey", LastName: "Indrukar"}})
	movies = append(movies, Movies{ID: "2", Isbn: "3451", Title: "Movie 2", Director: &Director{FirstName: "Nidhey", LastName: "Indrukar"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie", createMovie).Methods("POST")

	fmt.Println("Starting server at PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		} else {
			json.NewEncoder(w).Encode(make([]string, 0))
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// movies = append(movies[:index], movies[index+1])
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// movies = append(movies[:index], movies[index+1])
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movies
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Int())
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Int())
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}
