package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(strings.NewReader(arquivo))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("	%d	%d	%d	%s\n", strings.Count(arquivo, "\n"), count, strings.Count(arquivo, "")-1, x)
}
