package main

import "fmt"

// type user struct {
// 	name       string
// 	age        int
// 	isLoggedIn bool
// 	greet      func()
// }

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {

	// user1 := user{
	// 	name:       "John",
	// 	age:        25,
	// 	isLoggedIn: false,
	// }
	// user1.greet = func() {
	// 	fmt.Println("Hello", user1.name)
	// }
	// user1.greet()

	user1 := user{
		name:       "John",
		age:        25,
		isLoggedIn: false,
	}

	user1.greet()
	// pointerUser1 := &user1
	user1.login()

	fmt.Printf("%+v", user1)

}

func (u *user) login() { // if i take pointer on a struct then it will autometically take the address
	fmt.Println("Login called")
	u.isLoggedIn = true
}

func (u user) greet() {
	fmt.Println("Hello!", u.name)
}
