package main

import (
	"fmt"
)

func iter(s []int) func() (int, bool) {
	i := 0
	return func() (int, bool) {
		i += 1
		if i >= len(s) {
			return 0, false
		} else {
			return s[i], true
		}
	}
}

func main() {
	a := iter([]int{1, 2, 3, 5, 7, 7})
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
