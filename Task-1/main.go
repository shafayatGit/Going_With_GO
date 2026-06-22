package main

import "fmt"

/*

Task-1: Grade Calculator with Menu

Create a program that:
Shows a menu with options: 1) Calculate grade, 2) Check pass/fail status, 3) Exit
Use a switch statement for menu selection
For option 1: Take a score (0-100) and use if-else to assign letter grades (A: 90-100, B: 80-89, C: 70-79, D: 60-69, F: below 60)
For option 2: Use a switch to check if a score is pass (≥60) or fail
Use a for loop to keep showing the menu until user chooses exit
Include at least one anonymous function to display the menu

Example Output:
Menu:
1) Calculate grade
2) Check pass/fail status
3) Exit
Enter your choice: 1
Enter score (0-100): 85
Grade: B

Menu:
1) Calculate grade
2) Check pass/fail status
3) Exit
Enter your choice: 2
Enter score (0-100): 55
Status: Fail

Menu:
1) Calculate grade
2) Check pass/fail status
3) Exit
Enter your choice: 3
Exiting program. Goodbye!

*/

func main() {

	// logical and && 	//logical or ||
	// age := 15
	// nid := true

	// if age >= 18 || nid == true {
	// 	fmt.Println("You can vote")
	// } else {
	// 	fmt.Println("You can not vote")

	// }

	// Shows a menu with options: 1) Calculate grade, 2) Check pass/fail status, 3) Exit

	displayMenu := func() {
		fmt.Println("Welcome to grade calculator")
		fmt.Println("1) Calculate grade")
		fmt.Println("2) Check pass/fail status")
		fmt.Println("3) Exit the program")
		fmt.Print("Choose an option: ")
	}

	var choice int
	var score int
	running := true

	for running {
		displayMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter a score (0-100):")
			fmt.Scan(&score)
			result := calculateGrade(score)
			if result == "Invalid score" {
				fmt.Println("Choose a valid number from 0 - 100")
			}
			fmt.Println("You got", result)

		case 2:
			fmt.Println("check pass fail")

		case 3:
			fmt.Println("exiting program ...")
			running = false

		default:
			fmt.Println("Choose a valid option from the menu")

		}

	}

}

func calculateGrade(score int) string {
	if score >= 90 && score <= 100 {
		return "A"
	} else if score >= 80 && score <= 89 {
		return "B"
	} else if score >= 70 && score <= 79 {
		return "C"
	} else if score >= 60 && score <= 69 {
		return "D"
	} else if score >= 0 && score < 60 {
		return "F"
	} else {
		return "Invalid score"
	}
}

func checkPassFail(score int) string {
	switch {
	case score >= 60 && score <= 100:
		return "Pass"
	case score >= 0 && score < 60:
		return "Fail"
	default:
		return "Invalid score"
	}
}
