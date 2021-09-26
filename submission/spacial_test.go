package main

import (
	"fmt"
	"gifhelper"
	"testing"
)

// Tests the InitBoard function in spacial.go; the testing functions take a pointer to the testing packages's testing.T type as a parameter
// func TestInitBoard(t *testing.T) {
// 	filePath := "f99.txt"
// 	board := InitBoard(filePath)
// 	// I don't wanna print the entire board, it's just gonna flood the terminal, making my eyes sore
// 	fmt.Println(board[0])
// }

// func TestDrawBoards(t *testing.T) {
// 	outputFile := os.Args[1]
// 	board := InitBoard("./f99.txt")
// 	var boards []GameBoard
// 	boards = append(boards, board)
// 	imgList := DrawGameBoards(boards, 10)
// 	gifhelper.ImagesToGIF(imgList, outputFile)
// }

func TestDrawTwoBoards(t *testing.T) {
	fmt.Println("Spatial!")
	// collect input params
	filePath := "tinyfield.txt"
	b := 1.65
	steps := 1

	// initialize the board
	board := InitBoard(filePath)
	fmt.Println("Start simulating prisoners...")

	// get a list of boards to-be-drawn
	boards := PlaySpatial(board, b, steps)
	fmt.Println("Hold still. I'm drawing prisoners...")

	// we need a slice of image objects
	imglist := DrawGameBoards(boards, 10)
	fmt.Println("Finished drawing, now making animated GIF...")

	// convert images to a GIF
	gifhelper.ImagesToGIF(imglist, "out")

	fmt.Println("Behold! The prisoners!")
}
