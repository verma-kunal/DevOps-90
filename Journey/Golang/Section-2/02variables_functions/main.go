package main

import "fmt"

var global string

func main() {
	fmt.Println("Hello!")

	// Declaring a variable:
	var whatToSay string = "GoodBye!"

	// Printing the value:
	fmt.Println(whatToSay)

	var num int = 10
	fmt.Println("Num is set to:", num)

	whatWasSaid, otherThingSaid := something()
	fmt.Println(whatWasSaid+","+otherThingSaid)

}

func something() (string, string) {
	return "Kunal", "Other String"
}