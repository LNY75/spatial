package main

import (
	"fmt"
	"gifhelper"
	"os"
	"strconv"
)

// The data stored in a single cell of a field
type Cell struct {
	strategy string  //represents "C" or "D" corresponding to the type of prisoner in the cell
	score    float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}

// The game board is a 2D slice of Cell objects
type GameBoard [][]Cell

func main() {
	fmt.Println("Spatial!")
	// collect input params
	filePath := os.Args[1]
	b, err1 := strconv.ParseFloat(os.Args[2], 64)
	if err1 != nil {
		panic("cannot convert input b to float64")
	}
	steps, err2 := strconv.Atoi(os.Args[3])
	if err2 != nil {
		panic("cannot convert input step to int")
	}
	cellWidth := 3

	// initialize the board
	board := InitBoard(filePath)
	fmt.Println("Start simulating prisoners...")

	// get a list of boards to-be-drawn
	boards := PlaySpatial(board, b, steps)
	fmt.Println("Hold still. I'm drawing prisoners...")

	// we need a slice of image objects
	imglist := DrawGameBoards(boards, cellWidth)
	fmt.Println("Finished drawing, now making animated GIF...")

	// convert images to a GIF
	gifhelper.ImagesToGIF(imglist, "out")

	// save the image of the last board drawn as .png output
	DrawAndSaveImgPNG(boards[len(boards)-1], cellWidth)

	fmt.Println("Behold! The prisoners!")
}
