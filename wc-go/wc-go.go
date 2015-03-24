package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	x := os.Args[1]
	dat, err := ioutil.ReadFile(x)
	check(err)
	var arquivo string = string(dat)
	fmt.Printf("	%d	%d	%d	%s\n", strings.Count(arquivo, "\n"), len(strings.Fields(arquivo)), strings.Count(arquivo, "")-1, x)
}
