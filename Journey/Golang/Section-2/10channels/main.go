package main

import (
	"Golang/10channels/testDir"
)

// creating a function:

const numPool = 10

func CalculateValue(intChan chan int) {

	randomNumber := testDir.RandomNumber(numPool)

	// passing the number to our channel:
	intChan <- randomNumber
}

func main() {

	// creating the channel => "intChan"
	intChan := make(chan int)
	defer close(intChan)
	// good to have a habit to doing this
	// "defer" -> delay the execution of the function or method until the nearby functions get executed!

}
