package main

import (
	"fmt"
)

type Matriz struct {
	lin, col int
	m        [][]int
}

//Fazer 3 casos de teste, não precisa ler do teclado

func multiplica(a, b Matriz) (Matriz, bool) {
	if a.col != b.lin {
		return a, false
	}
	//matriz := Matriz{a.lin, b.col, [][]int{{0}}}
	var matriz Matriz
	matriz.lin = a.lin
	matriz.col = b.col
	matriz.m = make([][]int, a.lin)
	for i := range matriz.m {
		matriz.m[i] = make([]int, matriz.col)
	}
	for i := 0; i < a.lin; i++ {
		for j := 0; j < b.col; j++ {
			soma := 0
			for k := 0; k < a.col; k++ {
				soma += a.m[i][k] * b.m[k][j]
			}
			matriz.m[i][j] = soma
		}
	}

	return matriz, true
}

func main() {
	a := Matriz{3, 3, [][]int{{7, 2, 1}, {5, 0, 2}, {1, 1, 2}}}
	b := Matriz{3, 3, [][]int{{3, 2, 3}, {4, 1, 8}, {3, 4, 5}}}
	c := Matriz{3, 1, [][]int{{2}, {3}, {4}}}
	d := Matriz{1, 3, [][]int{{5, 3, 2}}}
	e := Matriz{2, 3, [][]int{{5, 3, 2}, {1, 2, 3}}}
	f := Matriz{2, 3, [][]int{{5, 3, 2}, {3, 6, 5}}}
	r, err := multiplica(a, b)
	if !err {
		fmt.Println("Nao foi possivel calcular a multiplicacao das matrizes")
	} else {
		fmt.Println(r)
	}
	t, err2 := multiplica(c, d)
	if !err2 {
		fmt.Println("Nao foi possivel calcular a multiplicacao das matrizes")
	} else {
		fmt.Println(t)
	}
	u, err3 := multiplica(e, f)
	if !err3 {
		fmt.Println("Nao foi possivel calcular a multiplicacao das matrizes")
	} else {
		fmt.Println(u)
	}
}
