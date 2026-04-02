package main

import "fmt"

// struct LinkedList com a cabeca(head) ponteiro para a struct Node e o tamanho size da linkedList
type LinkedList struct {
   head *Node
   size int
}

// estrutura noh
type Node struct {
   value int
   next *Node
}

func (list *LinkedList) reverse() {
    var prev *Node = nil // variavel ponteiro para anterior do atual
    var next *Node =  nil // variavel ponteiro para proximo do atual
    current := list.head // variavel ponteiro para a cabeca da linkedList

    for current != nil {
        next = current.next // faco o ponteiro next apontar para o proximo de current
        current.next = prev // faco o ponteiro de current apontar para o anterior
        prev = current // faco o ponteiro prev apontar para o atual
        current = next // faco o ponteiro atual apontar para o proximo
    }

    list.head = prev // o ponteiro anterior a current(que vai ser nil, vai parar o for) passa a ser a cabeca
    
}

func (list *LinkedList) display() { // pedi pra o gemini fazer essa funcao para mostra todos os valores invertidos da LinkedList
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next // Move para o próximo nó
	}
	fmt.Println("nil") // Indica o fim da lista
}

func main() {

    
    // pedi pra o gemini fazer esses comandos para visualizar a LinkedList reversa
    
    // Exemplo de uso para testar sua função
    lista := &LinkedList{}
    
    // Adicionando alguns nós manualmente para testar
    lista.head = &Node{value: 10, next: &Node{value: 20, next: &Node{value: 30, next: nil}}}
    
    fmt.Println("Invertendo a lista...")
    lista.reverse()

    lista.display()
    
    // Verificando o novo primeiro elemento (deve ser 30)
    if lista.head != nil {
        fmt.Printf("A nova cabeça da lista é: %d\n", lista.head.value)
     }   
}