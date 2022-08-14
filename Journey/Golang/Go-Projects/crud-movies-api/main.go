package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// defining the port number here:
const port = ":8080"

// creating structs:

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// slice of movies:
var movies []Movie

//

func main() {
	// creating a new Gorilla Mux router
	r := mux.NewRouter()

	// adding a few movies in our slice by appending:
	movies = append(movies,
		Movie{
			"1",
			"438227",
			"Movie One",
			&Director{"John", "Doe"}},
		Movie{
			"2",
			"45455",
			"Movie Two",
			&Director{"Steve", "Smith"}},
	)

	// Defining HandleFunctions for our routes + their respective methods:
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting the server at: %v", port)

	// Listening to the server (+ we are handling the error here as well):
	log.Fatal(http.ListenAndServe(port, r))

}

// getMovies - to get all the movies
func getMovies(w http.ResponseWriter, req *http.Request) {
	
}
