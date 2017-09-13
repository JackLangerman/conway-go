// Package main implements Conway's Game of Life
package main

import (
    "fmt"
    "os"
)

const numGenerations = 1005

func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "life.txt"
	}

    board := importBoard(filename)

    for i := 0 ; i <= numGenerations ; i++ {			// advance numGenerations
    	fmt.Printf("\nState #%d\n", i)					// print state #
    	board.print()									// print board state
    	board.nextState()								// advance to the next generation
    }
}

















