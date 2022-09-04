package main

import "fmt"

func main() {

	fmt.Println("=== Loop example 1 ===")
	i := 1
	for i < 3 {
		i++
		fmt.Println(i)
	}
}