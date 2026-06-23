package main

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
type product struct {
	Name     string
	Price    float64
	Quantity int
}

func (pr product) total() float64 {
	return pr.Price * float64(pr.Quantity)
}

type Cart struct {
	products []product
}

func (c Cart) GrantTotal() {
	totalPrice := 0.0
	for _, product := range c.products {
		totalPrice += product.total()
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

}
