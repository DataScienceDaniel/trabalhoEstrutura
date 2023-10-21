package main

import "fmt"

type Nodo struct {
    valor   int
    próximo *Nodo
}

type ListaCircular struct {
    primeiroNodo *Nodo
}

func (lista *ListaCircular) ExibirNodos() {
    if lista.primeiroNodo == nil {
        fmt.Println("A lista está vazia.")
        return
    }

    atual := lista.primeiroNodo
    for {
        fmt.Printf("%d -> ", atual.valor)
        atual = atual.próximo
        if atual == lista.primeiroNodo {
            break
        }
    }
    fmt.Println()
}

func (lista *ListaCircular) InserirNoInicio(valorInserir int) {
    novoNodo := &Nodo{valor: valorInserir}
    if lista.primeiroNodo == nil {
        lista.primeiroNodo = novoNodo
        novoNodo.próximo = novoNodo
    } else {
        novoNodo.próximo = lista.primeiroNodo
        atual := lista.primeiroNodo
        for atual.próximo != lista.primeiroNodo {
            atual = atual.próximo
        }
        atual.próximo = novoNodo
        lista.primeiroNodo = novoNodo
    }
}

func (lista *ListaCircular) ExcluirPrimeiroNodo() {
    if lista.primeiroNodo != nil {
        if lista.primeiroNodo.próximo == lista.primeiroNodo {
            lista.primeiroNodo = nil
        } else {
            atual := lista.primeiroNodo
            for atual.próximo != lista.primeiroNodo {
                atual = atual.próximo
            }
            atual.próximo = lista.primeiroNodo.próximo
            lista.primeiroNodo = lista.primeiroNodo.próximo
        }
    }
}

func main() {
    lista := ListaCircular{}

    // Inserir elementos no início da lista
    lista.InserirNoInicio(5)
    lista.InserirNoInicio(10)
    lista.InserirNoInicio(3)
    lista.InserirNoInicio(8)

    // Exibir os nós na lista
    fmt.Println("Lista de nós:")
    lista.ExibirNodos()

    // Excluir o primeiro nó
    lista.ExcluirPrimeiroNodo()

    // Exibir os nós na lista após a exclusão
    fmt.Println("Lista de nós após a exclusão do primeiro nó:")
    lista.ExibirNodos()
}
