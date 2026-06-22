package main

import "fmt"

func makeCoffee(coffeeNo int) {
	fmt.Println("making coffee....", coffeeNo)
}

func main() {

	// for i := 1; i <= 5; i++ { // for initialization; condition; increment/decrement
	// 	makeCoffee(i)
	// }

	// while styled loop

	// i := 1

	// for i <= 5 {
	// 	makeCoffee(i)
	// 	i++
	// }

	// infinite loop
	// for {
	// 	makeCoffee(2)
	// }

	// break (break the loop), continue (skip the iteration)

	// for i := 1; i <= 10; i++ {
	// 	if i == 6 {
	// 		break // stops the loop
	// 	}
	// 	makeCoffee(i)

	// }

	for i := 1; i <= 10; i++ {

		if i%2 != 0 {
			continue // skips the iteration
		}

		makeCoffee(i)

	}

}

// i := 1, true, run the body, increment
// i = 2, true, run the body, increment
// i = 3, true, run the body, increment
// i = 4, true, run the body, increment
// i = 5, true, run the body, increment
// i = 6, false, =======
