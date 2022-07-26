package main

import "log"

func main() {
	// If-else statement:

	//isTrue := true // a boolean type
	//
	//if isTrue {
	//	log.Println("Its true!")
	//} else {
	//	log.Println("Its false!")
	//}

	// Switch Statements:
	myAge := 21

	switch myAge {
	case 19:
		log.Println("You cannot vote - Age =", myAge)
	case 20:
		log.Println("You cannot vote - Age =", myAge)
	case 21:
		log.Println("You can vote - Age =", myAge)
	default:
		log.Println("Invalid Input!")
	}

}
