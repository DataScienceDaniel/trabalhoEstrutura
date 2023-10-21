Uma **lista circular** é uma estrutura de dados em que cada elemento (nó) aponta para o próximo elemento, e o último elemento aponta de volta para o primeiro, criando um ciclo. Isso significa que não há um "fim" da lista, e você pode percorrer os elementos continuamente.

**Exibição dos nós em uma lista circular**:
```
Se lista_vazia:
    Exibir "A lista está vazia."
Senão:
    Nodo atual = primeiro_nodo
    Repita:
        Exibir valor de atual
        atual = atual.próximo
    Enquanto atual ≠ primeiro_nodo
```
----------------------------------

**Inserção de um nó no início da lista**:
```
Novo_nodo = CriarNodo(valor)

Se lista_vazia:
    primeiro_nodo = Novo_nodo
    Novo_nodo.próximo = Novo_nodo
Senão:
    Novo_nodo.próximo = primeiro_nodo
    Nodo último = primeiro_nodo
    Enquanto último.próximo ≠ primeiro_nodo:
        último = último.próximo
    último.próximo = Novo_nodo
    primeiro_nodo = Novo_nodo
```

----------------------------------

**Exclusão do primeiro nó da lista**:
```
Se lista_vazia:
    Exibir "A lista está vazia. Não há nada para excluir."
Senão Se primeiro_nodo.próximo = primeiro_nodo:
    // A lista possui apenas um nó
    Remover primeiro_nodo
    primeiro_nodo = Nulo
Senão:
    Nodo último = primeiro_nodo
    Enquanto último.próximo ≠ primeiro_nodo:
        último = último.próximo
    último.próximo = primeiro_nodo.próximo
    Remover primeiro_nodo
    primeiro_nodo = último.próximo
```
