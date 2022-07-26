package main

import "log"

// Creating a struct:
type myStruct struct {
	FirstName string
}

func (s *myStruct) printFirstName() string {
	return s.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Kunal"

	// Another way of initialising & assigning values
	myVar2 := myStruct{
		FirstName: "Brad",
	}

	log.Println("myVar is set to:", myVar.printFirstName())
	log.Println("myVar2 is set to:", myVar2.printFirstName())
}
