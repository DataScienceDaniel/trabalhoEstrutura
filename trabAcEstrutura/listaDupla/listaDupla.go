package main

import "fmt"


type Nodo struct {
    valor    int
    anterior *Nodo
    próximo  *Nodo
}

type ListaDuplamenteEncadeada struct {
    primeiroNodo *Nodo
    últimoNodo   *Nodo
}

func (lista *ListaDuplamenteEncadeada) ExibirNodos() {
    atual := lista.primeiroNodo
    for atual != nil {
        fmt.Printf("%d -> ", atual.valor)
        atual = atual.próximo
    }
    fmt.Println("nil")
}

func (lista *ListaDuplamenteEncadeada) BuscarNodo(valorBuscado int) *Nodo {
    atual := lista.primeiroNodo
    for atual != nil {
        if atual.valor == valorBuscado {
            return atual
        }
        atual = atual.próximo
    }
    return nil
}

func (lista *ListaDuplamenteEncadeada) InserirNodo(valorInserir int) {
    novoNodo := &Nodo{valor: valorInserir}
    if lista.primeiroNodo == nil {
        lista.primeiroNodo = novoNodo
        lista.últimoNodo = novoNodo
    } else if valorInserir < lista.primeiroNodo.valor {
        novoNodo.próximo = lista.primeiroNodo
        lista.primeiroNodo.anterior = novoNodo
        lista.primeiroNodo = novoNodo
    } else if valorInserir > lista.últimoNodo.valor {
        lista.últimoNodo.próximo = novoNodo
        novoNodo.anterior = lista.últimoNodo
        lista.últimoNodo = novoNodo
    } else {
        atual := lista.primeiroNodo
        for atual != nil && atual.valor < valorInserir {
            atual = atual.próximo
        }
        novoNodo.anterior = atual.anterior
        novoNodo.próximo = atual
        atual.anterior.próximo = novoNodo
        atual.anterior = novoNodo
    }
}

func (lista *ListaDuplamenteEncadeada) RemoverNodo(valorRemover int) {
    atual := lista.BuscarNodo(valorRemover)
    if atual != nil {
        if atual == lista.primeiroNodo {
            lista.primeiroNodo = atual.próximo
        }
        if atual == lista.últimoNodo {
            lista.últimoNodo = atual.anterior
        }
        if atual.anterior != nil {
            atual.anterior.próximo = atual.próximo
        }
        if atual.próximo != nil {
            atual.próximo.anterior = atual.anterior
        }
    }
}

func main() {
    lista := ListaDuplamenteEncadeada{}

    // Inserir elementos na lista
    lista.InserirNodo(5)
    lista.InserirNodo(10)
    lista.InserirNodo(3)
    lista.InserirNodo(8)

    // Exibir os nós na lista
    fmt.Println("Lista de nós:")
    lista.ExibirNodos()

    // Buscar um nó
    nodoBuscado := lista.BuscarNodo(3)
    if nodoBuscado != nil {
        fmt.Printf("Nó com valor 3 encontrado: %d\n", nodoBuscado.valor)
    } else {
        fmt.Println("Nó com valor 3 não encontrado.")
    }

    // Remover um nó
    lista.RemoverNodo(10)

    // Exibir os nós na lista após a remoção
    fmt.Println("Lista de nós após a remoção:")
    lista.ExibirNodos()
}
