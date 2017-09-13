// Package main implements Conway's Game of Life
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numGenerations = 100
	var filename string

	if len(os.Args) == 2 { // user wants to pass the filename via command line
		filename = os.Args[1]

	} else if len(os.Args) == 3 {
		filename = os.Args[1] // user wants to pass the filename and number gerneations
		ng, _ := strconv.Atoi(os.Args[2])
		numGenerations = ng

	} else {
		filename = "life.txt" // defaults to life.txt and 10 generations
		numGenerations = 10
	}

	board := importBoard(filename) // import the seed state from file at filename

	fmt.Println("Initial world")
	board.print() // print start state
	fmt.Println("===========================")

	for i := 1; i <= numGenerations; i++ { // advance numGenerations
		fmt.Printf("Generation: %d\n", i) // print state #
		board.nextState()                 // advance to the next generation
		board.print()                     // print board state
		fmt.Println("====================")
	}
}
