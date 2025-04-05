package main

import (
	"fmt"
)

// Verifica se o grafo é bipartido e separa os funcionários em dois grupos
func isBipartiteWithNames(graph [][]int, names []string) (bool, []int) {
	V := len(graph)
	color := make([]int, V)
	for i := range color {
		color[i] = -1 // Nenhum funcionário foi colorido ainda
	}

	for start := 0; start < V; start++ {
		if color[start] == -1 {
			queue := []int{start}
			color[start] = 0

			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]

				for _, neighbor := range graph[node] {
					if color[neighbor] == -1 {
						color[neighbor] = 1 - color[node]
						queue = append(queue, neighbor)
					} else if color[neighbor] == color[node] {
						return false, nil // Conflito: não bipartido
					}
				}
			}
		}
	}

	return true, color
}

func main() {
	// Lista de funcionários
	names := []string{"Diogo", "Paixao", "Carlos", "Caio", "Pedro", "Virna", "Wiener", "Peter", "Cristiano Ronaldo", "Sam"}

	// Conflitos entre funcionários (grafo)
	graph := [][]int{
		{1, 4},       
		{0, 2},       
		{1, 5},       
		{6, 7},       
		{0, 8},       
		{2, 9},       
		{3, 8},       
		{3, 9},       
		{4, 6},       
		{5, 7},       
	}

	isBipartite, color := isBipartiteWithNames(graph, names)

	if isBipartite {
		fmt.Println("É possível dividir os funcionários em dois grupos sem conflitos:")
		fmt.Println("Grupo 1:")
		for i, c := range color {
			if c == 0 {
				fmt.Println("-", names[i])
			}
		}
		fmt.Println("Grupo 2:")
		for i, c := range color {
			if c == 1 {
				fmt.Println("-", names[i])
			}
		}
	} else {
		fmt.Println("Não é possível dividir os funcionários em dois grupos sem conflitos.")
	}
}
