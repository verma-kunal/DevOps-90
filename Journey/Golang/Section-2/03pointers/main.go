package main

import "log"

func main() {

	var myString string = "Kunal"
	log.Println("myString is set to:", myString)

	// Passing the reference to 'myString' variable, to access its value
	changeUsingPointer(&myString)

	log.Println("After change, myString is set to:", myString)
}

func changeUsingPointer(str *string) { // "str" is a pointer variable of type string
	log.Println("str is set to: ", str)
	newValue := "Verma"
	*str = newValue
}
