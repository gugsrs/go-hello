package main

import (
	"fmt"
)

func primos(n int, c chan int) {

	k, i := 0, 0

	for k < n {
		if i
	}

	close(c)
}

func main() {

	fmt.Println("Informe a quantidade de números a ser calculada:")
	var tam int
	fmt.Scanln(&tam)
	c := make(chan int, tam)
	go primos(tam, c)
	for i := range c {
		fmt.Println(i)
	}
}
