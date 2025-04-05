package main

import (
	"fmt" // Importa o pacote fmt para entrada e saída formatada
)

// Funcionario representa um nó do grafo com nome, id, cor e lista de
// vizinhos representados por seus IDs (inteiros)
type Funcionario struct {
	id       int    // ID único do funcionário
	cor      int    // -1 = não visitado, 0 ou 1 = cor atribuída durante a BFS
	nome     string // Nome do funcionário
	vizinhos []int  // Lista de IDs dos funcionários vizinhos
}

// isBipartite verifica se o grafo é bipartido e atribui cores aos vertices usando BFS
func isBipartite(funcionarios []*Funcionario) bool {
	for _, f := range funcionarios {
		if f.cor == -1 { // Se o nó ainda não foi visitado
			queue := []*Funcionario{f} // Inicia a fila com o nó atual
			f.cor = 0                  // Atribui a primeira cor

			for len(queue) > 0 { // Enquanto houver nós na fila
				curr := queue[0]  // Pega o primeiro da fila
				queue = queue[1:] // Remove o primeiro da fila

				for _, vizID := range curr.vizinhos { // Itera sobre os vizinhos do nó atual
					viz := funcionarios[vizID] // Acessa o vizinho via ID
					if viz.cor == -1 {         // Se o vizinho ainda não foi visitado
						viz.cor = 1 - curr.cor     // Atribui a cor oposta ao vizinho
						queue = append(queue, viz) // Adiciona o vizinho à fila
					} else if viz.cor == curr.cor {
						return false // Conflito: dois vizinhos com a mesma cor => não bipartido
					}
				}
			}
		}
	}
	return true // Se nenhum conflito foi encontrado, o grafo é bipartido
}

func main() {
	// Lista de nomes dos funcionários, ordenados pelo índice (ID)
	nomes := []string{
		"Diogo",             // ID 0
		"Paixao",            // ID 1
		"Virna",             // ID 2
		"Caio",              // ID 3
		"Pedro",             // ID 4
		"Carlos",            // ID 5
		"Wiener",            // ID 6
		"Peter",             // ID 7
		"Cristiano Ronaldo", // ID 8
		"Sam",               // ID 9
	}

	// Representação do grafo onde a chave é o ID de um funcionário
	// e o valor é a lista de IDs com quem ele tem conflito (vizinhos)
	grafo := map[int][]int{
		0: {1, 4}, // Diogo é vizinho de Paixao e Pedro
		1: {0, 2}, // Paixao é vizinho de Diogo e Virna
		2: {1, 5}, // Virna é vizinha de Paixao e Carlos
		3: {6, 7}, // Caio é vizinho de Wiener e Peter
		4: {0, 8}, // Pedro é vizinho de Diogo e Cristiano Ronaldo
		5: {2, 9}, // Carlos é vizinho de Virna e Sam
		6: {3, 8}, // Wiener é vizinho de Caio e Cristiano Ronaldo
		7: {3, 9}, // Peter é vizinho de Caio e Sam
		8: {4, 6}, // Cristiano Ronaldo é vizinho de Pedro e Wiener
		9: {5, 7}, // Sam é vizinho de Carlos e Peter
	}

	// Cria o slice de funcionários e inicializa cada funcionário com seus dados
	var funcionarios []*Funcionario
	for id, nome := range nomes {
		funcionarios = append(funcionarios, &Funcionario{
			id:       id,        // ID do funcionário
			cor:      -1,        // Cor inicial indefinida
			nome:     nome,      // Nome do funcionário
			vizinhos: grafo[id], // IDs dos vizinhos do grafo
		})
	}

	// Verifica se é possível dividir os funcionários em dois grupos (grafo bipartido)
	if isBipartite(funcionarios) {
		fmt.Println("É possível dividir os funcionários em dois grupos sem conflitos:")
		fmt.Println("Grupo 1:")
		for _, f := range funcionarios {
			if f.cor == 0 { // Grupo com cor 0
				fmt.Println("-", f.nome)
			}
		}
		fmt.Println("Grupo 2:")
		for _, f := range funcionarios {
			if f.cor == 1 { // Grupo com cor 1
				fmt.Println("-", f.nome)
			}
		}
	} else {
		fmt.Println("Não é possível dividir os funcionários em dois grupos sem conflitos.")
	}
}
