package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// struct to hold the current gamestate
type Board struct {
	board [][]bool
	width, height int
}

// constructor for the board struct
func createBoard(width, height int) Board {
	board := make([][]bool, height+2)				// add extra row above and below board to avoid checking boundries
	for row := range board {
		board[row] = make([]bool, width+2) 			// add extra col left and right of board to avoid checking boundries
	}
	return Board{board: board, width: width, height: height}
}

// read the startstate of the game from a file to the Board struct
func (board *Board) importStartStateScanner(filename string) {
	file, err := os.Open(filename) 		// open file for reading
    if err != nil { 			   		// if the file does not exist log and exit
        log.Fatal(err)
    }
    defer file.Close()			   		// close this file when the function returns

    scanner := bufio.NewScanner(file)	// create a new scanner to read the file
    scanner.Split(bufio.ScanRunes)    	// tell the scanner to split on characters

    row := 0							// make some variable for looping through rows/cols
    col := -1
    for scanner.Scan() {				// while there are more characters in the file 
        output_char := scanner.Text()	// get the next character
        
        col++							// increment the col variable

        if output_char == "*" {			// if the curent char is an astrix mark the cell as alive
        	board.setPosition(row, col, true)
        } else if output_char == "\n" {	// if the cur char is a newline increment the row and reset the col
        	row++
        	col = -1
        }
    }

    if err := scanner.Err(); err != nil { // check for errors in the scanner
        log.Fatal(err)
    }

}

//  print the current gamestate to standard output 
func (board *Board) print() {
	for i := 0 ; i < board.height ; i++ {		// for each row
		for j := 0 ; j < board.width ; j++ {	// for each col
			if board.isAlive(i, j) {			// check if the curent cell is alive
				fmt.Print("*")					// if so print "*"
			} else {
				fmt.Print(" ")					// else print " "
			}
		}
		fmt.Print("\n")
	}
}

// check if a particular cell is alive
func (board *Board) isAlive(row, col int) bool {
	return board.board[row+1][col+1]
}

// set the state of a particular cell (alive / dead)
func (board *Board) setPosition(row int, col int, alive bool) *Board {
	board.board[row+1][col+1] = alive
	return board
}

// count how many of cell at (row, col)'s eight neighbors are alive
func (board *Board) countNeighbors(row, col int) int {
	count := 0						// innitalize a var to count living  neighbors
	for i := -1; i <= 1; i++ {		// check above and below
		for j:= -1; j <= 1; j++ {	// check left and right
			//check the eight neighbors (but don't count self)
			if !(i==0 && j==0) && board.isAlive(row+i, col+j) { 
				count++
			}
		}
	}
	return count
}


// advance the board to the next state
func (board *Board) nextState() {
	nextboard := createBoard(board.width, board.height)		// create a board to store the next state

	for i:=0; i<board.height; i++ {
		for j:=0; j<board.width; j++ {   					// for each cell
			count := board.countNeighbors(i, j)		    		 // get the neighbor count

			if count < 2 || count > 3 {							 // if the count is  0, 1, or 4 die
				nextboard.setPosition(i, j, false)
			} else if count == 3 {
				nextboard.setPosition(i, j, true)				 // if count is 3 spring to life
			} else {
				nextboard.setPosition(i, j, board.isAlive(i,j) ) // if count is 2 remain unchanged
			}
		}
	}

	// board.update(nextboard)										 // update the board with the new state
	*board = nextboard
}