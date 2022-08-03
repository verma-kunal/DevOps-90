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

// Method-1 - Increasing efficiency of parsing the templates:

// Declaring a map:
var tempCache = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, temp string) {
	var tmpl *template.Template
	var err error

	// check if the template is already present in our cache (tempCache -> map)
	// searching the map for the key "temp":
	_, inMap := tempCache[temp]

	if !inMap {
		// => template not present. Therefore, need to create a one
		log.Println("Creating template & adding to cache")
		err = createTemplateCache(temp)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template!")
	}

	// once we have the value of the template -> storing that in our "tmpl" variable:
	tmpl = tempCache[temp]

	// Executing the template:
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(tempName string) error {
	// step-1
	templates := []string{
		fmt.Sprintf("./templates/%s", tempName),
		"./templates/base.layout.tmpl",
	}

	// step-2 -> parsing the template:
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// otherwise: add the template to the cache (map)
	tempCache[tempName] = tmpl

	return nil
}
