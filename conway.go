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
func createBoard(width, height int) *Board {
	board := make([][]bool, height+2)				// add extra row above and below board to avoid checking boundries
	for row := range board {
		board[row] = make([]bool, width+2) 			// add extra col left and right of board to avoid checking boundries
	}
	return &Board{board: board, width: width, height: height}
}

//function to read the startstate of the game from a file to the Board struct
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

// function to print the current gamestate to standard output 
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

//function to check if a particular cell is alive
func (board *Board) isAlive(row, col int) bool {
	return board.board[row+1][col+1]
}

//function to set the state of a particular cell (alive / dead)
func (board *Board) setPosition(row int, col int, alive bool) *Board {
	board.board[row+1][col+1] = alive
	return board
}



func main() {
	board := createBoard(20, 8)
    board.importStartStateScanner("life.txt")
    board.print()
}

















