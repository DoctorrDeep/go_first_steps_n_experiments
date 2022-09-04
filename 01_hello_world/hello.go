package main

import (
	"fmt"
	"math"
)

const welcome string = "home"

func main() {

	// Print stuff
	fmt.Println("=== Print stuff ===")
	fmt.Println("hello world!")

	fmt.Println("Well" + "hello" + "there") // Won't have spaces
	fmt.Println("Well", "hello", "there")   // Will have spaces

	// Math and bool
	fmt.Println("=== Math and bool ===")
	fmt.Println(1 + 1)
	fmt.Println(1.0 + 1.0)
	fmt.Println(7 / 3)
	fmt.Println(7.0 / 3.0)

	fmt.Println(true && false)

	// variables
	fmt.Println("=== Variables ===")
	var a, b int = 1, 2
	fmt.Println(a, b)

	var c, d float64 = 3.0, 4.0
	fmt.Println(c, d)

	var e string = "apples?"
	f := "Hungry for"
	fmt.Println(f, e)

	// constant (.contd)
	fmt.Println("=== Constants ===")
	const n1 int64 = 10

	fmt.Println(welcome, n1)

	const n2 = 34 + n1
	//const n3 = 42 + c // fails because c is not const

	fmt.Println(n2)
	fmt.Println(math.Sin(c))

}
