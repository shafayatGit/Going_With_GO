package main

import "fmt"

// -------- Task 1
// type Speaker interface {
// 	speak()
// }

// type Animal struct {
// 	name  string
// 	sound string
// }

// // receiver function
// func (a Animal) speak() {

// 	if a.name == "cat" {

// 		fmt.Println("Cat says: Meaw!")
// 	}
// 	if a.name == "dog" {

// 		fmt.Println("Dog says: Woof!")
// 	}
// }

// ---------> Task 2
// type product struct {
// 	Name     string
// 	Price    float64
// 	Quantity int
// }

// func (pr product) total() float64 {
// 	return pr.Price * float64(pr.Quantity)
// }

// type Cart struct {
// 	products []product
// }

// func (c Cart) GrantTotal() {
// 	totalPrice := 0.0
// 	for _, product := range c.products {
// 		totalPrice += product.total()
// 	}
// }

// func CartReceipt(cart Cart) {
// 	for _, product := range cart.products {
// 		fmt.Printf("%-10s x%-5d   $%.2f\n",
// 			product.Name,
// 			product.Quantity,
// 			product.total(),
// 		)
// 	}
// }

// ---------> Task 3
// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	Radius float64
// }

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func (cr *Circle) Area() float64 {
// 	return 3.1416 * cr.Radius * cr.Radius
// }
// func (cr *Rectangle) Area() float64 {
// 	return cr.Width * cr.Height
// }

// ---------> Task 4
// type BankAccount struct {
// 	Owner   string
// 	Balence float64
// }

// func (ba *BankAccount) Deposit(amount float64) {

// 	ba.Balence += amount
// }
// func (ba *BankAccount) WithDraw(amount float64) {
// 	if ba.Balence > amount {
// 		ba.Balence -= amount
// 	} else {
// 		fmt.Println("Dont have sufficient balence")

// 	}
// }
// func (ba *BankAccount) Display() {
// 	fmt.Println("Account Owner:", ba.Owner)
// 	fmt.Println("Account Money:", ba.Balence)

// }

// -------------> Task 5
type Student struct {
	Name   string
	Grades []int
}

func (st Student) Average() float64 {
	total := 0
	for _, grade := range st.Grades {
		total += grade
	}
	return float64(total) / float64(len(st.Grades))
}

func (s Student) LetterGrade() string {
	avg := s.Average()

	if avg >= 90 {
		return "A"
	} else if avg >= 80 {
		return "B"
	} else if avg >= 70 {
		return "C"
	} else {
		return "F"
	}
}
func main() {
	//Task 1
	// animals := []Animal{{
	// 	name: "dog",
	// }, {
	// 	name: "cat",
	// }}

	// for _, name := range animals {
	// 	name.speak()
	// }

	// --------> Task 2
	// products := []product{
	// 	{
	// 		Name: "Apple", Price: 20, Quantity: 2,
	// 	},
	// 	{
	// 		Name: "Banana", Price: 10, Quantity: 1,
	// 	},
	// }

	// cart := Cart{
	// 	products: products,
	// }

	// CartReceipt(cart)

	//---------> Task 3

	// circle := Circle{
	// 	Radius: 5,
	// }

	// rectangle := Rectangle{
	// 	Width:  10,
	// 	Height: 4,
	// }

	// shapes := []Shape{
	// 	&circle,
	// 	&rectangle,
	// }

	// // Loop using range
	// for _, shape := range shapes {
	// 	fmt.Printf("Area: %.2f\n", shape.Area())
	// }

	// ---------> Task 4
	// account := BankAccount{
	// 	Owner:   "Shafayat",
	// 	Balence: 2000,
	// }
	// account.Display()
	// account.Deposit(2000)
	// account.Display()
	// account.WithDraw(3000)
	// account.Display()

	students := []Student{
		{
			Name:   "Shafayat",
			Grades: []int{90, 85, 95},
		},
		{
			Name:   "Rahim",
			Grades: []int{75, 80, 70},
		},
		{
			Name:   "Karim",
			Grades: []int{60, 65, 55},
		},
	}

	for _, student := range students {
		fmt.Println("Student:", student.Name)
		fmt.Println("Grades:", student.Grades)
		fmt.Printf("Average: %.2f\n", student.Average())
		fmt.Println("Grade:", student.LetterGrade())
		fmt.Println("-------------------")
	}

}
