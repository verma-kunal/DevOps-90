package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// creating a type:

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FavColor  string `json:"fav_color"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Kunal",
		"last_name": "Verma",
		"fav_colour": "Black"
	},	
	{
		"first_name": "Saiyam",
		"last_name": "Pathak",
		"fav_colour": "Blue"
	}	
]`

	var unMarshelled []Person

	// converting the json data into byte-data
	jsonInBytes := []byte(myJson)

	err := json.Unmarshal(jsonInBytes, &unMarshelled)
	if err != nil {
		log.Println("Error unmarshelling json", err)
	}

	log.Printf("Unmarshelled: %v", unMarshelled)

	// Write json from a struct:
	var mySlice []Person

	// creating some vars of the type "Person"
	var m1 Person
	m1.FirstName = "Chris"
	m1.LastName = "Evans"
	m1.FavColor = "Dark-Blue"
	
	var m2 Person
	m2.FirstName = "Barry"
	m2.LastName = "Alan"
	m2.FavColor = "Red"

	// adding or appending the "m1" & "m2" variables to the slice:
	mySlice = append(mySlice, m1, m2)

	// calling the "marhshall()" function now:
	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	// checking for error:
	if err != nil {
		log.Println("Error marshelling json", err)
	}

	// converting our slice of bytes into a string & printing out:
	fmt.Println(string(newJson))
}
