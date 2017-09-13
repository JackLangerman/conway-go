package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Board is a struct to hold the current gamestate
type Board struct {
	board         [][]bool
	width, height int
}

// createBoard is the constructor for the board struct.
func createBoard(width, height int) Board {
	board := make([][]bool, height+2) // add extra row above and below board to avoid checking boundries
	for row := range board {
		board[row] = make([]bool, width+2) // add extra col left and right of board to avoid checking boundries
	}
	return Board{board: board, width: width, height: height}
}

// importBoard imports the seed state for the board from the file named "filename".
func importBoard(filename string) Board {
	file, err := os.Open(filename) // open file for reading
	if err != nil {                // if the file does not exist log and exit
		log.Fatal(err)
	}
	defer file.Close() // close this file when the function returns

	var width, height int // make some vars for board width and height
	var board [][]bool    // make an array to hold the board

	scanner := bufio.NewScanner(file) // create a new scanner to read the file
	for scanner.Scan() {              // make a line scanner to read the file a line at a time
		line := scanner.Text() // read a line of the file

		width = 0 // set the line width to 0

		var row []bool           // make a slice for the current row
		row = append(row, false) // add a dead cell at the beggining of the row

		for _, c := range line { // loop through each char of the current line in the file
			b := c == '*'        // check if there is a *(to indicate life) in the cur cell
			row = append(row, b) // add the cell state to the end of the row
			width++              // increment the width
		}
		row = append(row, false) // add a dead cell at the end of the row

		board = append(board, row) // add the current row to the end of the board
		height++                   // increment the height
	}

	board = append(make([][]bool, 1), board...)
	board[0] = append(make([]bool, width+2)) // prepend a dead row the the board

	board = append(board, make([][]bool, 1)...)
	board[height+1] = append(make([]bool, width+2)) // append a dead row to the board

	return Board{board: board, width: width, height: height}
}

//  print prints the current gamestate to standard output.
func (board *Board) print() {
	for i := 0; i < board.height; i++ { // for each row
		for j := 0; j < board.width; j++ { // for each col
			if board.isAlive(i, j) { // check if the curent cell is alive
				fmt.Print("*") // if so print "*"
			} else {
				fmt.Print(" ") // else print " "
			}
		}
		fmt.Print("\n")
	}
}

// isAlive checks if the cell at (row, col) is alive.
func (board *Board) isAlive(row, col int) bool {
	return board.board[row+1][col+1]
}

// setPosition set the state of the cell at (row, col) to the bool "alive".
func (board *Board) setPosition(row int, col int, alive bool) *Board {
	board.board[row+1][col+1] = alive
	return board
}

// countNeighbors counts how many of cell at (row, col)'s eight neighbors are alive.
func (board *Board) countNeighbors(row, col int) int {
	var kern = [][]bool{{true, true, true},
		{true, false, true}, // kernal for convolution
		{true, true, true}}

	count := 0 // innitalize a var to count living  neighbors

	for i := -1; i <= 1; i++ { // check above and below
		for j := -1; j <= 1; j++ { // check left and right
			if kern[i+1][j+1] && board.isAlive(row+i, col+j) {
				count++
			}
		}
	}
	return count
}

// nextState advances the board to the next state.
func (board *Board) nextState() {
	nextboard := createBoard(board.width, board.height) // create a board to store the next state

	for i := 0; i < board.height; i++ {
		for j := 0; j < board.width; j++ { // for each cell
			count := board.countNeighbors(i, j) // get the neighbor count

			switch {
			case count < 2 || count > 3: // if the count is  0, 1, or 4 die
				nextboard.setPosition(i, j, false)
			case count == 3:
				nextboard.setPosition(i, j, true) // if count is 3 spring to life
			default:
				nextboard.setPosition(i, j, board.isAlive(i, j)) // if count is 2 remain unchanged
			}
		}
	}

	*board = nextboard
}
