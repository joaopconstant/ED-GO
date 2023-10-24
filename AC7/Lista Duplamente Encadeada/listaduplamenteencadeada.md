# Lista Duplamente Encadeada
Uma lista duplamente encadeada é uma estrutura de dados na qual cada elemento (nó) contém um valor e dois ponteiros: um para o nó anterior e outro para o próximo nó na sequência. Isso permite a navegação bidirecional na lista, facilitando operações de inserção, exclusão e busca em ambas as direções. Essa versatilidade torna a lista duplamente encadeada eficiente para várias aplicações, proporcionando acesso flexível e rápido aos elementos da lista.

### Exibição dos Nós em uma Lista Duplamente Encadeada:
```
PROGRAMA ExibirListaDuplamenteEncadeada(Lista):
    Se Lista for vazia, então
        Escrever("Lista Duplamente Encadeada Vazia")
    Senão
        NóAtual = Lista.inicio
        Enquanto NóAtual não for nulo faça
            Escrever(NóAtual.valor)
            NóAtual = NóAtual.proximo

```

### Busca de um Nó em uma Lista Duplamente Encadeada:
```
PROGRAMA BuscarNo(Lista, valorBuscado):
    NóAtual = Lista.inicio
    Enquanto NóAtual não for nulo faça
        Se NóAtual.valor for igual a valorBuscado, então
            Retornar NóAtual
        Fim Se
        NóAtual = NóAtual.proximo
    Retornar Nulo
```

### Inserção de um Nó em uma Lista Duplamente Encadeada:
```
PROGRAMA InserirNo(Lista, valor):
    NovoNó = CriarNovoNó(valor)
    Se Lista for vazia, então
        Lista.inicio = NovoNó
    Senão
        NovoNó.proximo = Lista.inicio
        Lista.inicio.anterior = NovoNó
        Lista.inicio = NovoNó

PROGRAMA CriarNovoNó(valor):
    NovoNó = Estrutura Nó
    NovoNó.valor = valor
    NovoNó.proximo = Nulo
    NovoNó.anterior = Nulo
    Retornar NovoNó
```

### Remoção de um Nó em uma Lista Duplamente Encadeada:
```
Procedimento RemoverNo(Lista, valor):
    NóAtual = Lista.inicio
    Enquanto NóAtual não for nulo faça
        Se NóAtual.valor for igual a valor, então
            Se NóAtual.anterior não for nulo, então
                NóAtual.anterior.proximo = NóAtual.proximo
            Senão
                Lista.inicio = NóAtual.proximo
            Se NóAtual.proximo não for nulo, então
                NóAtual.proximo.anterior = NóAtual.anterior
            Retornar
        NóAtual = NóAtual.proximo
```