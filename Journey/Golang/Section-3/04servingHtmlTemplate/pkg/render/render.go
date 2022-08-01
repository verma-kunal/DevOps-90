package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate - to parse our template & make those ready to use (render's the template):
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// parsing our template here:
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	// checking for error
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// Creating a map, that can store the values from "template.ParseFiles" function:
var templateCache = make(map[string]*template.Template)

// RenderTemplateTest - taking args as the response writer & our template name as a string
func RenderTemplateTest(w http.ResponseWriter, temp string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache (checking the templateCache map - whether the template exists in there or not)
	_, inMap := templateCache[temp]
	// inMap stores true/false for if the template is there in the map or not

	if !inMap {
		// need to create the template

	} else {
		// the template is present in the cache
		log.Println("Using cached template")
	}

	tmpl = templateCache[temp]

	// executing the template:
	err = tmpl.Execute(w, nil)
}