package main

import "fmt" //standard library

var isAdmin bool = true

// isAdmin:=true ==> will not work

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
}

//no need to call main file/function
