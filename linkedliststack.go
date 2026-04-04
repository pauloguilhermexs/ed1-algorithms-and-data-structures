package main

import (
    "fmt"
    "errors"
    )

type IStack interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

type Node struct {
    value int
    next *Node
}

type LinkedListStack struct{
    head *Node   // vai ser o topo da pilha
    size int 
}

func (s *LinkedListStack) Size() int {
    return s.size
}

func (s *LinkedListStack) IsEmpty() bool{
    return s.head == nil
}

func (s *LinkedListStack) Push(value int){
    newNode := &Node{value: value, next: s.head} // novo node recebe no seu value o value de parametro e aponta para o antigo topo
    s.head = newNode // o antigo topo aponta para o novo node
    s.size++  
}

func (s *LinkedListStack) Pop() (int, error){
    if s.IsEmpty(){
        return -1, errors.New("Stack is empty")
    }
    aux := s.head.value
    s.head = s.head.next // proximo elemento da LinkedListStack passa a ser o novo topo
    s.size--
    return aux, nil
}

func (s *LinkedListStack) Peek() (int, error){
    if s.IsEmpty(){
        return -1, errors.New("Stack is empty")
    }
    return s.head.value, nil
}

func main() {
    
    stack := &LinkedListStack{}
    stack.Push(10)
    stack.Push (20)
    stack.Push(30)
    fmt.Println("Tamanho da pilha:", stack.Size())
    fmt.Println("vazia?: " , stack.IsEmpty())
    aux1, _ := stack.Pop()
    fmt.Println("Valor retirado do topo:", aux1)
    aux1, _ = stack.Peek()
    fmt.Println("Valor atual do topo:", aux1)
    stack.Pop()
    stack.Pop()
    stack.Pop()
    fmt.Println("vazia?: " , stack.IsEmpty())
    fmt.Println("Tamanho da pilha:", stack.Size())

	aux2, _ := stack.Pop() // pega o ultimo elemento (topo)
	fmt.Println("Valor retirado do topo:", aux2)
    
    aux2, _ = stack.Peek() 
    fmt.Println("Topo atual: ", aux2)   
}