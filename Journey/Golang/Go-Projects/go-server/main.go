package main

import (
	"fmt"
	"log"
	"net/http"
)

// Defining Port Number globally:
const portNumber = ":8080"

func main() {
	// Setting the path for our static folder here: (will look for index.html file in "static" folder)
	fileServer := http.FileServer(http.Dir("./static"))

	// creating a Handle function:
	http.Handle("/", fileServer)

	// creating a Handler function:
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	log.Printf("Server listing on PORT %s", portNumber)

	// listening to the server: ("ListenAndServe" returns as error, so we should handle that as well)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func formHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	// print when our request is successful
	fmt.Fprintf(w, "POST request successful")

	// getting the values from our form (on-submit):
	firstName := req.FormValue("fname")
	lastName := req.FormValue("lname")

	// printing the values to the screen:
	fmt.Fprintf(w, "\nFirst Name: %v\n", firstName)
	fmt.Fprintf(w, "Last Name: %v\n", lastName)

}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	// writing some basic checks:

	// 1. If the path is not "/hello":
	if req.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	// 2. If the method of our request is not "GET"
	if req.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// In an ideal case, do this:
	fmt.Fprintf(w, "Hello World!")
}
