# Lista Circular
Uma lista circular é uma estrutura de dados em que os elementos formam uma sequência contínua, onde o último elemento está ligado ao primeiro, criando um loop. Ao contrário das listas lineares tradicionais, não há um final definido na lista circular, pois o último elemento aponta de volta para o primeiro, criando uma estrutura circular infinita.

### Exibição dos Nós em uma Lista Circular:

```
PROGRAMA ExibirListaCircular(Lista):
    Se Lista for vazia, então:
        Escrever("Lista Circular Vazia")
    Senão
        NóAtual = Lista
        Repita
            Escrever(NóAtual.valor)
            NóAtual = NóAtual.proximo
        Até que NóAtual seja igual ao primeiro nó da lista
```

### Inserção de um Nó no Início da Lista Circular:
```
PROGRAMA InserirNoInicio(Lista, valor):
    NovoNó = CriarNovoNó(valor)
    Se Lista for vazia, então
        Lista = NovoNó
        Lista.proximo = Lista
    Senão
        NovoNó.proximo = Lista.proximo
        Lista.proximo = NovoNó

PROGRAMA CriarNovoNó(valor):
    Retorna um novo nó com o valor especificado e o ponteiro próximo inicializado como nulo
```

### Exclusão do Primeiro Nó da Lista Circular:
```
PROGRAMA ExcluirPrimeiroNo(Lista):
    Se Lista for vazia, então
        Escrever("Lista Circular Vazia")
    Senão
        Se Lista.proximo for igual a Lista, então
            Lista = Nulo
        Senão
            Lista.proximo = Lista.proximo.proximo
```

