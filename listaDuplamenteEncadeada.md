Uma **lista duplamente encadeada** é uma estrutura de dados na qual cada nó possui uma referência tanto para o nó anterior quanto para o próximo nó. Isso permite a travessia da lista em ambas as direções. 

**Exibição dos Nós em uma Lista Duplamente Encadeada**
Nodo atual = primeiro_nodo
Enquanto atual ≠ Nulo:
    Exibir valor de atual
    atual = atual.próximo

-------------

**Busca de um Nó em uma Lista Duplamente Encadeada**
Função BuscarNodo(valor_buscado):
    Nodo atual = primeiro_nodo
    Enquanto atual ≠ Nulo:
        Se atual.valor = valor_buscado:
            Retornar atual
        atual = atual.próximo
    Retornar Nulo

--------------

**Inserção de um Nó em uma Lista Duplamente Encadeada**
Função InserirNodo(valor_inserir):
    Novo_nodo = CriarNodo(valor_inserir)
    Se lista_vazia:
        primeiro_nodo = Novo_nodo
        último_nodo = Novo_nodo
    Senão Se valor_inserir < primeiro_nodo.valor:
        Novo_nodo.próximo = primeiro_nodo
        primeiro_nodo.anterior = Novo_nodo
        primeiro_nodo = Novo_nodo
    Senão Se valor_inserir > último_nodo.valor:
        último_nodo.próximo = Novo_nodo
        Novo_nodo.anterior = último_nodo
        último_nodo = Novo_nodo
    Senão:
        Nodo atual = primeiro_nodo
        Enquanto atual ≠ Nulo E atual.valor < valor_inserir:
            atual = atual.próximo
        Novo_nodo.anterior = atual.anterior
        Novo_nodo.próximo = atual
        atual.anterior.próximo = Novo_nodo
        atual.anterior = Novo_nodo

---------------

**Remoção de um Nó em uma Lista Duplamente Encadeada**
Função RemoverNodo(valor_remover):
    Nodo atual = BuscarNodo(valor_remover)
    Se atual ≠ Nulo:
        Se atual = primeiro_nodo:
            primeiro_nodo = atual.próximo
        Senão Se atual = último_nodo:
            último_nodo = atual.anterior
        Senão:
            atual.anterior.próximo = atual.próximo
            atual.próximo.anterior = atual.anterior
        Remover atual








