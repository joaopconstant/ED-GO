package main

import "fmt"

type Contato struct {
	Nome	string
	Email	string
}

func adicionaContato(c Contato, contatos [5]Contato) [5]Contato {
	for i, contato := range contatos {
		if (contato == Contato{}) {
			contatos[i] = c
			break
		}
	}
	return contatos
}

func removeContato(contatos [5]Contato) [5]Contato {
	for i, contato := range contatos {
		if (contato == Contato{}) {
			contatos[i-1] = contato
			fmt.Println(contatos)
			break
		}
	}
	return contatos
}

func main() {
	var contatos [5]Contato
	var opcao int
	msg := "-----------------------------\n LISTA DE CONTATOS\n\n Informe a opção desejada:\n 1- Adicionar contato.\n 2- Remover contato.\n Ou qualquer tecla para sair."

	for {
		fmt.Println(msg)
		fmt.Scan(&opcao)
		switch opcao {
		case 1:
			var inputNome string
			fmt.Println("Informe o nome:")
			fmt.Scan(&inputNome)
			var inputEmail string
			fmt.Println("Informe o email:")
			fmt.Scan(&inputEmail)
	
			var contato Contato
			contato = Contato {
				Nome: inputNome,
				Email: inputEmail,
			}
			contatos = adicionaContato(contato, contatos)
			fmt.Println(contatos)
			
		case 2:
			contatos = removeContato(contatos)
			
		case 3:
			fmt.Println(contatos)
			
		default:
			return
		}
	}
}
