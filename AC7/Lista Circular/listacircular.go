package main

import "fmt"

type No struct {
    valor int
    proximo *No
}

type ListaSimplesmenteEncadeada struct {
    inicio *No
}

func (lista *ListaSimplesmenteEncadeada) InserirNo(valor int) {
    novoNo := &No{valor: valor}
    if lista.inicio == nil {
        lista.inicio = novoNo
    } else {
        atual := lista.inicio
        for atual.proximo != nil {
            atual = atual.proximo
        }
        atual.proximo = novoNo
    }
}

func (lista *ListaSimplesmenteEncadeada) RemoverNo(valor int) {
    if lista.inicio == nil {
        return
    }
    if lista.inicio.valor == valor {
        lista.inicio = lista.inicio.proximo
        return
    }
    atual := lista.inicio
    for atual.proximo != nil && atual.proximo.valor != valor {
        atual = atual.proximo
    }
    if atual.proximo != nil {
        atual.proximo = atual.proximo.proximo
    }
}

func (lista *ListaSimplesmenteEncadeada) BuscarNo(valor int) *No {
    atual := lista.inicio
    for atual != nil {
        if atual.valor == valor {
            return atual
        }
        atual = atual.proximo
    }
    return nil
}

func (lista *ListaSimplesmenteEncadeada) Exibir() {
    atual := lista.inicio
    for atual != nil {
        fmt.Print(atual.valor, " -> ")
        atual = atual.proximo
    }
    fmt.Println("nil")
}

func main() {
    // Exemplo de uso da lista simplesmente encadeada
    lista := &ListaSimplesmenteEncadeada{}
    lista.InserirNo(1)
    lista.InserirNo(2)
    lista.InserirNo(3)
    lista.Exibir()

    lista.RemoverNo(2)
    lista.Exibir()

    noEncontrado := lista.BuscarNo(3)
    if noEncontrado != nil {
        fmt.Println("Nó encontrado:", noEncontrado.valor)
    } else {
        fmt.Println("Nó não encontrado")
    }
}
