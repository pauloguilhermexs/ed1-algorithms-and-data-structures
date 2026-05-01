package main

import (
	"fmt"
	"math"
)

func selection_sortOP(v []int) []int {
	ord := make([]int, len(v))
	for varredura := 0; varredura < len(v); varredura++ {
		imin := 0
		for i := 0; i < len(v); i++ {
			if v[i] < v[imin] {
				imin = i
			}
		}
		ord[varredura] = v[imin]
		v[imin] = math.MaxInt
	}
	return ord
}

func selection_sortIP(v []int) {
	for varredura := 0; varredura < len(v)-1; varredura++ {
		imin := varredura
		for i := varredura + 1; i < len(v); i++ {
			if v[i] < v[imin] {
				imin = i
			}
		}
		v[varredura], v[imin] = v[imin], v[varredura]
	}
}

func main() {

	numerosOP := []int{765, 431, 567, 876, 123, 1324, 342, 0, 1, 24, 12}
	ordenadosOP := selection_sortOP(numerosOP)
	fmt.Println(ordenadosOP)
	numerosIP := []int{12, 3, 4, 5, 8980, 32, 24, 0, 1, 1000, 100000000}
	selection_sortIP(numerosIP)
	fmt.Println(numerosIP)

}
