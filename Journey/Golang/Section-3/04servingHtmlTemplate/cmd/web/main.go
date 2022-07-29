package main

import (
	"github.com/verma-kunal/servingHTML/pkg/handlers"
	"log"
	"net/http"
)

// Defining Port Number globally:
const portNumber = ":8080"

func main() {

	// creating HandleFunctions for the above handler functions:
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	log.Printf("Server listing on PORT %s", portNumber)

	// listening to the server:
	_ = http.ListenAndServe(portNumber, nil)

}
