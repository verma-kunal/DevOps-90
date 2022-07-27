package main

// continuing with the "Hello World" Application
import (
	"fmt"
	"log"
	"net/http"
)

// Defining Port Number globally:
const portNumber = ":8080"

// creating a Handler Function:

// HomePage is the About Page Handler:
func HomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the Home Page!")
}

// AboutPage is the About Page Handler:
func AboutPage(w http.ResponseWriter, req *http.Request) {
	sum, _ := addValue(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About Page & (2+2) = %d", sum))
}

// creating a normal function:
// good practice a return an error as well! ('nil' here because we don't have any to check for)

func addValue(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil
}

func main() {

	// creating HandleFunctions for the above handler functions:
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)

	log.Printf("Server listing on PORT %s", portNumber)

	// listening to the server:
	_ = http.ListenAndServe(portNumber, nil)

}
