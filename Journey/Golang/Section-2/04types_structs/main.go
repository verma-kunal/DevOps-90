package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	Lastname    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Kunal",
		Lastname:  "Verma",
	}
	log.Println(user.FirstName)
	log.Println(user.Lastname)
}
