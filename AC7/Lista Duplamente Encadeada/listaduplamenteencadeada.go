package main

import "fmt"

type No struct {
    valor    int
    anterior *No
    proximo  *No
}

type ListaDuplamenteEncadeada struct {
    inicio *No
}

func (lista *ListaDuplamenteEncadeada) InserirNo(valor int) {
    novoNo := &No{valor: valor}
    if lista.inicio == nil {
        lista.inicio = novoNo
    } else {
        atual := lista.inicio
        for atual.proximo != nil {
            atual = atual.proximo
        }
        novoNo.anterior = atual
        atual.proximo = novoNo
    }
}

func (lista *ListaDuplamenteEncadeada) RemoverNo(valor int) {
    if lista.inicio == nil {
        return
    }
    if lista.inicio.valor == valor {
        lista.inicio = lista.inicio.proximo
        if lista.inicio != nil {
            lista.inicio.anterior = nil
        }
        return
    }
    atual := lista.inicio
    for atual != nil && atual.valor != valor {
        atual = atual.proximo
    }
    if atual == nil {
        return
    }
    if atual.anterior != nil {
        atual.anterior.proximo = atual.proximo
    }
    if atual.proximo != nil {
        atual.proximo.anterior = atual.anterior
    }
}

func (lista *ListaDuplamenteEncadeada) BuscarNo(valor int) *No {
    atual := lista.inicio
    for atual != nil {
        if atual.valor == valor {
            return atual
        }
        atual = atual.proximo
    }
    return nil
}

func (lista *ListaDuplamenteEncadeada) Exibir() {
    atual := lista.inicio
    for atual != nil {
        fmt.Print(atual.valor, " <-> ")
        atual = atual.proximo
    }
    fmt.Println("nil")
}

func main() {
    // Exemplo de uso da lista duplamente encadeada
    lista := &ListaDuplamenteEncadeada{}
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
