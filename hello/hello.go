package main

import "fmt"

func main() {
	fmt.Printf("Hello, world! Welcome to GoLang!\n")
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	for i := 0; i < 10; i++ {
		fmt.Println("Welcome to GoLang! Hope you like it ")
		fmt.Println(i)
	}
}
