package main

import (
    "fmt"
)

const numGenerations = 1000

func main() {
	// define the board size and number generations to simulate
	// boardWidth := 20
	// boardHeight := 8
	// numGenerations := 10

	// board := createBoard(boardWidth, boardHeight)		// create the board
    // board.importStartStateScanner("life.txt")		    // read the seed state from life.txt
    board := importBoard("life3.txt")

    for i := 0 ; i <= numGenerations ; i++ {			// advance numGenerations
    	fmt.Printf("\nState #%d\n", i)					// print state #
    	board.print()									// print board state
    	board.nextState()								// advance to the next generation
    }
}

















