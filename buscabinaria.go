package main

import "fmt"

func busca_binaria_iterativa(val int, v []int ) int{
    ini := 0
    fim := len(v) - 1
    
  
    for ini <= fim{
        meio := (ini+fim)/2
        if v[meio] == val{
            return meio
        } else if v[meio] > val{
            fim = meio - 1
        } else {
            ini = meio + 1
        }
    }

    return -1
}

func busca_binaria_recursiva(val int, v []int, ini int, fim int) int {
    if ini <= fim{
        meio := (ini + fim)/2
        if v[meio] == val{
            return meio
        } else if v[meio] > val{
            return busca_binaria_recursiva(val, v, ini, meio-1)
        } else {
            return busca_binaria_recursiva(val, v, meio+1, fim)
        }
    } else{
        return -1
    }
}

func main() {
    // Vetor precisa obrigatoriamente estar ordenado
	vetor := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	
	valorProcurado := 23
	indice1 := busca_binaria_recursiva(valorProcurado, vetor, 0, len(vetor)-1)
    indice2 := busca_binaria_iterativa(valorProcurado, vetor)

	if indice1 != -1 {
		fmt.Printf("Valor %d encontrado no índice: %d\n", valorProcurado, indice1)
	} else {
		fmt.Printf("Valor %d não encontrado no vetor.\n", valorProcurado)
	}

    if indice2 != -1{
        fmt.Printf("Valor %d encontrado no índice: %d\n", valorProcurado, indice2)
    } else {
        fmt.Printf("Valor %d não encontrado no vetor.\n", valorProcurado)
    }
    
}