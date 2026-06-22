package main

func main() {
	day := "lal shukrubar"

	// if day == "sat" {
	// 	println("Yes, sleep more")
	// } else if day == "fri" {
	// 	println("weekend so sleep more")
	// } else if day == "sun" {
	// 	println("work day")
	// } else {
	// 	println("another boring day")

	// }

	// switch day {  // tagged switch
	// case "sat":
	// 	println("Yes, sleep more")
	// case "fri":
	// 	println("Yes, sleep more its friday!!")
	// case "sun":
	// 	println("Work day")
	// default:
	// 	println("another boring day")
	// }

	switch { //  switch
	case day == "sat":
		println("Yes, sleep more")
	case day == "fri":
		println("Yes, sleep more its friday!!")
	case day == "sun":
		println("Work day")
	default:
		println("another boring day")
	}

}
