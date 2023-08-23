package main

import (
	"bufio"
	"fmt"
	"os"
	"ac3/contatos"
	"ac3/utils"
)

func main() {
	var contatosLista [5]contatos.Contato
	var nome, email, novoEmail, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover, (3) para alterar email, (4) para listar os contatos, ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			utils.AdicionaContato(contatos.Contato{Nome: nome, Email: email}, &contatosLista)
		case "2":
			utils.ExcluiContato(&contatosLista)
		case "3":
			utils.ExibeContatos(&contatosLista)
			fmt.Print("Informe o índice do contato que deseja alterar: ")
			var indice int
			fmt.Scanln(&indice)
			fmt.Print("Informe o novo Email: ")
			fmt.Scanln(&novoEmail)

			utils.EditaEmail(indice, novoEmail, &contatosLista)
		case "4":
			utils.ExibeContatos(&contatosLista)
		default:
			return
		}
	}
}