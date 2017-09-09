package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    //"io"
)

type Board struct {
	board [][]bool
	width, height int
}

func (board *Board) importStartStateScanner() {


	file, err := os.Open("life.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)

    row := 0
    col := -1
    for scanner.Scan() {
        output_char := scanner.Text()
        
        col++

        if output_char == "*" {
        	board.setPosition(row, col, true)
        } else if output_char == "\n"{
        	row++
        	col = -1
        }

        // fmt.Print(output_char)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    // fmt.Println("\n\n")
    board.print()
    // return nil
}

func (board *Board) print() {
	for i := 0 ; i < board.height ; i++ {
		for j := 0 ; j < board.width ; j++ {
			if board.isAlive(i, j) {
				fmt.Print("*")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Print("\n")
	}
}

func (board *Board) isAlive(row, col int) bool {
	return board.board[row+1][col+1]
}

func (board *Board) setPosition(row int, col int, alive bool) *Board {
	board.board[row+1][col+1] = alive
	return board
}


func createBoard2020() *Board {
	return createBoard(20, 20)
}

func createBoard(width, height int) *Board {
	board := make([][]bool, height+2)
	for row := range board {
		board[row] = make([]bool, width+2)
	}
	return &Board{board: board, width: width, height: height}
}

func addLine(board *Board, line *[]bool, lineNumber int) *Board {
	board.board[lineNumber+1] = *line
	return board
}


// func importStartStateReader() *Board {
// 	board := createBoard2020()

// 	f, err := os.Open("life.txt")
//     if err != nil {
//             log.Fatal(err)
//     }

//     lineReader := bufio.NewReader(f)
//     for i:=0 ; ; i++ {
//         if sliceLine, err := lineReader.ReadSlice('\n'); err == nil || err == io.EOF {
//                 //fmt.Printf("%s", sliceLine)
//         		addLine(board, &sliceLine, i)
//                 if err == io.EOF {
//                         break
//                 }
//         } else {
//                 log.Fatal(err)
//         }

//     }
//     return board
// }

func main() {
	board := createBoard(20, 8)
    board.importStartStateScanner()
}

















