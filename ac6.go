package main

import (
	"fmt"
)

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	if *n >= M {
		fmt.Println("Overflow")
		return
	}

	pos := *n
	for pos > 0 && v[pos-1] > novoValor {
		v[pos] = v[pos-1]
		pos--
	}

	v[pos] = novoValor
	*n++
	fmt.Printf("Inserido %d\n", novoValor)
}
