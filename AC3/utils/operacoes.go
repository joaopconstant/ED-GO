package utils

import (
	"fmt"
	"ac3/contatos"
)

func AdicionaContato(c contatos.Contato, a *[5]contatos.Contato) {
	for ind, contato := range *a {
		if (contato == contatos.Contato{}) {
			(*a)[ind] = c
			break
		}
	}
}

func ExcluiContato(a *[5]contatos.Contato) {
	if ((*a)[0] == contatos.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
		return
	}

	for ind, contato := range *a {
		if (contato == contatos.Contato{}) {
			(*a)[ind-1] = contatos.Contato{}
		}
	}
}

func EditaEmail(indice int, novoEmail string, contatosLista*[5]contatos.Contato) {
	if indice < 0 && indice > 4 {
		fmt.Print("Valor inválido.")
		return
	}

	contatosLista[indice].AlterarEmail(novoEmail)
}

func ExibeContatos(a*[5]contatos.Contato) {
	for ind, contato := range *a {
		if contato != (contatos.Contato{}) {
			fmt.Print("------------------\n" + "Nome: " + contato.Nome + "Email: " + contato.Email + "\nÍndice: ") 
			fmt.Println(ind)
		}
	}
}
