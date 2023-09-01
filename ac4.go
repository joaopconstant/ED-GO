package main

import "fmt"

func main() {
	n := 3
	A := "Origem"
	B := "Destino"
	C := "Trabalho"
	potencia := 1

	for i := 1; i <= n; i++ {
		potencia *= 2
	}
	fmt.Println("Numero mínimo de jogadas:", potencia-1)

	hanoi(n, A, B, C)
}

func hanoi(n int, A, B, C string) {
	if n > 0 {
		hanoi(n-1, A, C, B)
		fmt.Print("Mova o disco ", n, " de ", A, " até ", B, "\n")
		hanoi(n-1, C, B, A)
	}
}
