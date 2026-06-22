package main

import "fmt"

// func ageStatus(age int){

// }

func main() {
	// age := 20
	// if age >= 18 {
	// 	fmt.Println("Adult!")
	// }

	// score := 84
	// fmt.Printf("outside if %d\n", score)

	if score := 60; score >= 80 {
		prizeMoney := 1000
		fmt.Println("You have won Gold Medal and prize money is", prizeMoney)
		fmt.Println("You Got Gold Medal an your score is", score)
	} else if score >= 70 {

		fmt.Println("You Got Silver Medal and your score is", score)
	} else {
		fmt.Println("You got participation certificate and your score is", score)
	}

	// if err := saveToDb(); err != nil {
	// 	fmt.Println("Error:", err)
	// }

}
