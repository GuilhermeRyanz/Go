package main

import "fmt"

func quickSort(lista []int) []int {
	if len(lista) <= 1 {
		return lista
	}

	pivor := lista[len(lista)-1]
	var esq, dir []int

	for _, v := range lista[:len(lista)-1] {
		if v <= pivor {
			esq = append(esq, v)
		} else {
			dir = append(dir, v)
		}
	}

	esq = quickSort(esq)
	dir = quickSort(dir)

	return append(append(esq, pivor), dir...)
}

func main() {
	lista := []int{12, 11, 13, 5, 6, 7}
	fmt.Println(lista)

	lista = quickSort(lista)
	fmt.Println(lista)
}
