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

func Btoi(b bool) int {
    if b {
        return 1
    }
    return 0
 }

func (board *Board) countNeighbors4(row, col int) int {
	return Btoi( board.isAlive(row-1, col) ) + 
		Btoi( board.isAlive(row+1, col) ) + 
		Btoi( board.isAlive(row, col-1) ) + 
		Btoi( board.isAlive(row, col+1) ) 
}

func (board *Board) countNeighbors(row, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j:= -1; j <= 1; j++ {
			// fmt.Printf("|%d,%d-", row+i, col+j)
			if !(i==0 && j==0) && board.isAlive(row+i, col+j) {
				count++
			}
			// fmt.Printf("%d|", count)
		}
	}
	// fmt.Println()
	return count
}

func (board *Board) clone() *Board {
	newBoard := createBoard(board.width, board.height)
	for i:=0;i<board.height;i++ {
		for j:=0;j<board.width;j++ {
			newBoard.setPosition(i, j, board.isAlive(i,j) )
		}
	}
	return newBoard
}
func (oldboard *Board) update(newBoard *Board) {
	for i:=0;i<oldboard.height;i++ {
		for j:=0;j<oldboard.width;j++ {
			oldboard.setPosition(i, j, newBoard.isAlive(i,j) )
		}
	}
}

func (board *Board) nextState() {
	nextboard := board.clone()

	// fmt.Println("old board\n")
	// board.print()
	// fmt.Println("\nnew board\n")
	// nextboard.print()
	// fmt.Println("\nneighborcounts\n")
	for i:=0; i<board.height; i++ {
		for j:=0; j<board.width; j++ {
			count := board.countNeighbors(i, j)
			// fmt.Print(count)
			if count < 2 || count > 3 {
			// if (count == 0 || count == 1 || count == 4) {
				nextboard.setPosition(i, j, false)
			} else if count == 3 {
				nextboard.setPosition(i, j, true)
			} // else if count == 2 {
			  //		noop()
			  // }
		}
		// fmt.Println()
	}

	// fmt.Println("old board\n")
	// board.print()
	// fmt.Println("\nnew board\n")
	// nextboard.print()


	board.update(nextboard)

	// fmt.Println("\nperform update\n\n")
	// board.print()
	// fmt.Println("\nend round....\n\n\n\n")

}

func main() {
	board := createBoard(20, 8)
    board.importStartStateScanner("life.txt")
    

    for i := 0 ; i <= 10 ; i++ {
    	fmt.Printf("\nState #%d\n", i)
    	board.print()
    	board.nextState()
    }
}

















