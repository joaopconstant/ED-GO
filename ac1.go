package main

import "fmt"

func main() {
	// fmt.Println(e_primo(6))

	// fmt.Println(fibo(4))

	// dia_semana()

	// fmt.Println(e_bissexto(1995))
}

func e_primo(n int) bool {
	primo := true

	for i := 2; i <= n-1; i++ {
		if n % i == 0 {
			primo = false
			fmt.Println(i)
		}
	}

	return primo
}

func fibo(n int) int {
	var fibAnterior = 0
	var fibAtual = 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		for i := 2; i <= n; i++ {
			fibProximo := fibAnterior + fibAtual
			fibAnterior = fibAtual
			fibAtual = fibProximo
		}
		return fibAtual
	}
}

func dia_semana() string {
	var msg string
	var n int

	print("Informe o número do dia da semana (de 1 a 7):")
	fmt.Scanln(&n)

	switch n {
	case 1:
		msg = "Domingo"
		fmt.Println(msg)
	case 2:
		msg = "Segunda"
		fmt.Println(msg)
	case 3:
		msg = "Terça"
		fmt.Println(msg)
	case 4:
		msg = "Quarta"
		fmt.Println(msg)
	case 5:
		msg = "Quinta"
		fmt.Println(msg)
	case 6:
		msg = "Sexta"
		fmt.Println(msg)
	case 7:
		msg = "Sábado"
		fmt.Println(msg)
	default:
		msg = "Valor inválido"
		fmt.Println(msg)
	}

	return msg
}

func e_bissexto(n int) bool {
	bissexto := false

	if n % 4 == 0 && n % 100 != 0 {
		bissexto = true
	} else if n % 400 == 0 {
		bissexto = true
	}

	return bissexto
}
