package main

import "fmt"

func main() {
	fmt.Println(e_primo(6))
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

func fibo(num int) {
	// cria um array com a serie de Fibonacci

	// verifica o número de índice 'num'
}

func dia_semana() string {
	// lê o valor (msg)
	var msg string
	var n int
	print("Informe o número do dia da semana (de 1 a 7):")
	fmt.Scanln(&n)
	// relaciona números com dias
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

func e_bissexto() {

}