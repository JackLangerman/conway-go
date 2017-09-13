package main

import (
    "fmt"
)

const numGenerations = 1005

func main() {
    board := importBoard("life3.txt")

    for i := 0 ; i <= numGenerations ; i++ {			// advance numGenerations
    	fmt.Printf("\nState #%d\n", i)					// print state #
    	board.print()									// print board state
    	board.nextState()								// advance to the next generation
    }
}

















