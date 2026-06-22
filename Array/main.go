package Array

import "fmt"

func main() {

	var numbers [6]int

	numbers[0] = 42
	numbers[1] = 7

	numbers[5] = 100

	// fmt.Println(numbers)

	// fmt.Println("Length of numbers is:", len(numbers))
	fmt.Println(numbers[5])

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}
