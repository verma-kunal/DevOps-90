package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate - to parse our template & make those ready to use:
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// parsing our template here:
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	// checking for error
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
