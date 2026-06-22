package main

import "fmt" //standard library

var isAdmin bool = true

// isAdmin:=true ==> will not work

//------------>Without Return function
// func makeCoffee(kind string) {
// 	fmt.Printf("Making %s coffee", kind)
// }

//-------------> Return
// func makeCoffee(kind string) string {
// 	fmt.Printf("Making %s coffee \n", kind)
// 	return kind
// }

// --------->Multiple return value
//
//	func makeCoffee(kind string) (string, int) {
//		price := 100
//		fmt.Printf("Making %s coffee \n", kind)
//		return kind, price
//	}

// -------> Multiple named return
func makeCoffee(kind string) (coffee string, price int) {
	price = price
	coffee = coffee
	fmt.Printf("Making %s coffee \n", kind)
	return // no need to tell anything
}
func main() {
	// var name string;
	// name = "King"
	// var name string = "Shafayat"
	name := "Shafayat" // Go took type by its own => Short Way and mostly used but will not work outside of the function
	// name string :="king" //cant do this

	//Variable Grouping
	var (
		age         int    = 23
		email       string = "king@gmai"
		isAvailable bool   = true
	)

	// => Multiple variable declaration
	// var x,y int;
	// x=10
	// y=20

	const pi = 3.1416 // constant and never going to change and cant re-asign

	var x, y int = 20, 30
	fmt.Println("Hello World!", name, age, email, isAvailable)
	fmt.Println("Hello World!", x, y)

	myCoffee, myBill := makeCoffee("Black")
	fmt.Println(myCoffee, myBill)
}

//no need to call main file/function
