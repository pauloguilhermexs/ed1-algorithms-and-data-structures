package main

import (
    "fmt"
    "errors"
    )

type IList interface {
    Add(value int)
    AddOnIndex(value int, index int) error
    RemoveOnIndex(index int) error
    Get(index int) (int, error)
    Set(value int, index int) error
    Size() int
}

type ArrayStack struct {
    values [] int    
}

func (s *ArrayStack) Size() int{
    return len(s.values)
}
// essa operacao nao pode ocorrer numa pilha
func (s *ArrayStack) Set(value int, index int) error{
    return errors.New("operation not permitted")
}
// essa operacao nao pode ocorrer numa pilha
func(s *ArrayStack) RemoveOnIndex(index int) error{
    return errors.New("operation not permitted")
}
// essa operacao nao pode ocorrer numa pilha
func(s *ArrayStack) AddOnIndex(value int, index int) error{
    return errors.New("operation not permitted")
}

func (s *ArrayStack) Add(value int){
   s.values = append(s.values, value) 
}

func (s *ArrayStack) Get(index int) (int, error){
    if s.Size() == 0 {
        return -1, errors.New("stack is empty")
    }
    
    if index < 0 || index >= s.Size(){
        return -1, errors.New("operation not permitted")
    } 

    return s.values[index], nil
    
}


func main() {

    stack := &ArrayStack{}
	
	stack.Add(10)
	stack.Add(20)
	stack.Add(30)

	fmt.Println("Tamanho da pilha:", stack.Size())

	val, err := stack.Get(2) // pega o ultimo elemento (topo)
	if err == nil {
		fmt.Println("Valor no índice 2:", val)
	} else {
		fmt.Println("Erro:", err)
	}

	// testando operação nao permitida
	err = stack.Set(50, 0)
	fmt.Println("Tentativa de Set:", err)
 
}