package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Defining Port Number globally:
const portNumber = ":8080"

// HomePage is the About Page Handler:
func HomePage(w http.ResponseWriter, req *http.Request) {
	renderTemplate(w, "home.page.html")
}

// AboutPage is the About Page Handler:
func AboutPage(w http.ResponseWriter, req *http.Request) {
	renderTemplate(w, "about.page.html")
}

// renderTemplate - to parse our template & make those ready to use:
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// parsing our template here:
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	// checking for error
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {

	// creating HandleFunctions for the above handler functions:
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)

	log.Printf("Server listing on PORT %s", portNumber)

	// listening to the server:
	_ = http.ListenAndServe(portNumber, nil)

}
