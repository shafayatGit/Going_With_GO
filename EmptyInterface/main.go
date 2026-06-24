package main

import "fmt"

// any = interface{}
// type assertion
// ok idiom

func Process(data any) {

	strData, ok := data.(string)
	if ok {
		fmt.Println(len(strData))
	}

	intData, ok := data.(int)
	if ok {
		fmt.Println(intData + 100)
	}

}

func main() {
	// Print([]int{6, 7, 2, 9})
	Process(100)
}
