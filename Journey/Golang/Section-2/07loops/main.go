package main

import "log"

func main() {

	// For Loop:
	//for i := 0; i <= 10; i++ {
	//	log.Println(i)
	//}

	// Slice of strings:
	//animals := []string{"Cat", "Dog", "Lion", "Hippo", "Mouse"}

	// Ranging over the slice:
	//for i, animal := range animals {
	//	log.Println(i, animal)
	//}
	//for _, animal := range animals {
	//	log.Println(animal)
	//}

	// Ranging over a Map
	//myMap := map[string]int{
	//	"first":  1,
	//	"second": 2,
	//	"third":  3,
	//}
	//for key, value := range myMap {
	//	log.Println(key, value)
	//}

	// Ranging over a String:
	//name := "Kunal"
	//
	//for index, letter := range name {
	//	log.Println(index, ":", letter)
	//}

	// Ranging over a custom type:
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{"Kunal", "Verma", "vkunal321@gmail.com", 19})
	users = append(users, User{"Saiyam", "Pathak", "saiyam911@gmail.com", 30})
	users = append(users, User{"Kunal", "Kushwaha", "commclassroom@gmail.com", 23})

	for _, info := range users {
		log.Println(info.FirstName, info.LastName, info.Email, info.Age)
	}
}
