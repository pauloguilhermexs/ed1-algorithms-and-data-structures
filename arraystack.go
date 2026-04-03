package main

import (
    "fmt"
    "errors"
    )

type IList interface {
    Push(value int)
    Pop() (int, error)
    Peek() (int, error)
    IsEmpty() bool
    Size() int
}

type ArrayStack struct {
    values [] int    
}

func (s *ArrayStack) Size() int{
    return len(s.values)
}

func (s *ArrayStack) IsEmpty() bool{
    return len(s.values) == 0
}

func (s *ArrayStack) Push(value int){
   s.values = append(s.values, value) 
}

func (s *ArrayStack) Pop() (int, error){
    if s.IsEmpty(){
        return -1, errors.New("stack is empty")
    }
    
    index := len(s.values) - 1
    value := s.values[index]
    s.values = s.values[:index] 
    return value, nil    
}

func (s *ArrayStack) Peek() (int, error){
    if s.IsEmpty(){
        return -1, errors.New("stack is empty")
    }
    return s.values[s.Size() - 1], nil
}


func main() {

    stack := &ArrayStack{}
	
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Tamanho da pilha:", stack.Size())
    fmt.Println("vazia?: " , stack.IsEmpty())
    //stack.Pop()
    //stack.Pop()
    //stack.Pop()

	val, _ := stack.Pop() // pega o ultimo elemento (topo)
	fmt.Println("Valor retirado do topo:", val)
    
    val, _ = stack.Peek() 
    fmt.Println("Topo atual: ", val)    
   
}