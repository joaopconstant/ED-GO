package main

import (
	"ac3/contatos"
	"ac3/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var contato contatos.Contato
	var contatos [5]contatos.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			utils.AdicionaContato(contato{Nome: nome, Email: email}, &contatos)
		case "2":
			utils.ExcluiContato(&contatos)
		case "3":
			fmt.Println("Informe o indice do contato que deseja editar")
			index := 0
			fmt.Scanln(&index)

			fmt.Println("Informe o novo email: ")
			novoEmail, _ := leitor.ReadString("\n")
			contato.AlterarEmail(&contatos[index], strings.TrimSpace(novoEmail))
		default:
			return
		}

		fmt.Println(contatos)
	}

}
