package main

import (
    "github.com/IzcoatlRam/puzzle8/node"
    "github.com/IzcoatlRam/puzzle8/queue"
    "fmt"
    "os"
    "os/exec"
    "time"
)

func main() {
    fmt.Println("Estado inicial:")

	// s := node.RandomState()

	s := node.State{
		{1, 3, 6},
		{5, 2, 8},
		{4, 7, 0},
	}

	 // Convertir la matriz 2D en una lista de números
	 puzzle := make([]int, 0, len(s)*len(s[0]))
	 for _, row := range s {
		 for _, num := range row {
			 puzzle = append(puzzle, num)
		 }
	 }
 
	 // Verificar si el puzzle tiene solución
	 if !isSolvable(puzzle) {
		 fmt.Println("Este puzzle no tiene solución.")
		 return
	 }

    initialNode := node.NewNode(s, "", nil)
    initialNode.PrintState()
    time.Sleep(500 * time.Millisecond) // Espera más tiempo
    ClearScreen()

    solutionNode := BreadthFirstSearch(s)
    if solutionNode != nil {
        fmt.Println("Solución encontrada:")
        printSolution(solutionNode)
    } else {
        fmt.Println("No se encontró solución.")
    }
}

func BreadthFirstSearch(initialState node.State) *node.Node {
    initialNode := node.NewNode(initialState, "", nil)
    queue := queue.NewQueue()
    queue.Encolar(initialNode)

    for queue.Len() > 0 {
        currentNode := queue.Desencolar()

        // Verificar si el estado actual es la solución
        if node.IsGoalState(currentNode.State) {
            return currentNode // Has encontrado la solución
        }

        // Generar sucesores y encolarlos
        successors := node.GenerateSuccessors(currentNode)
        for _, successor := range successors {
            queue.Encolar(successor)
        }
    }

    return nil // No se encontró solución
}

func printSolution(n *node.Node) {
    if n == nil {
        return
    }
    printSolution(n.Parent)
    ClearScreen()
    n.PrintState()
    fmt.Println()
    time.Sleep(1000 * time.Millisecond) // Espera más tiempo
}

func ClearScreen() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func isSolvable(puzzle []int) bool {
    inversions := 0
    for i := 0; i < len(puzzle); i++ {
        for j := i + 1; j < len(puzzle); j++ {
            if puzzle[i] != 0 && puzzle[j] != 0 && puzzle[i] > puzzle[j] {
                inversions++
            }
        }
    }
    return inversions%2 == 0
}