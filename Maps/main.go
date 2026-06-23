package main

import "fmt"

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {

	// myMap := make(map[int]string)

	// myMap[2] = "two"
	// myMap[3] = "three"
	// fmt.Println(myMap[2])

	// myMap := map[string]string{
	// 	"name":    "Shafayat",
	// 	"Success": "ok",
	// }

	// delete(myMap, "name")

	myMap := map[string]user{
		"data": user{
			name:       "Shafayat",
			age:        25,
			isLoggedIn: false,
		},
	}

	fmt.Println(myMap)

}

// crating map (using  make and literal)
// accessing map values
// adding new key-value pair to map
// deleting key-value pair from map
// map of struct
// iterating over map
