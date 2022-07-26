package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	// 1st method:
	myMap := make(map[string]User)

	first := User{
		FirstName: "Kunal",
		LastName:  "Verma",
	}
	myMap["first"] = first

	//log.Println(myMap["first"].FirstName)
	//log.Println(myMap["first"].LastName)

	// 2nd method:
	//myMap2 := map[string]User{
	//	"first": {
	//		FirstName: "Kunal",
	//		LastName:  "Verma",
	//	},
	//	"second": {
	//		FirstName: "Saiyam",
	//		LastName:  "Pathak",
	//	},
	//}
	//log.Println(myMap2["second"].FirstName)
	//log.Println(myMap2["second"].LastName)

	// _________________________________________________________________
	// Slices

	var names []string
	names = append(names, "kunal", "Brad")
	//log.Println(names)

	var nums = []int{4, 1, 6, 0}

	//log.Println("Before sorting: ", nums)
	sort.Ints(nums)
	//log.Println("After sorting: ", nums)

	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// printing in a range:
	log.Println(nums2[1:4])

}
