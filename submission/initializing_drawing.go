package main

import (
	"bufio"
	"image"
	"log"
	"os"
)

// InitBoard reads a file that represent the initial state of a board and returns a GameBoard
func InitBoard(filepath string) GameBoard {
	// open the file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create an empty GameBoard
	board := make(GameBoard, 0)

	// read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		var currentCellArr []Cell

		// if this line does not contain strategy elements (C or D), then we should not read anything from this line
		hasStrat := string(currentLine[0]) == "C" || string(currentLine[0]) == "D"
		if hasStrat {
			for i := range currentLine {
				currentCell := Cell{strategy: "", score: 0}
				val := currentLine[i]
				currentCell.strategy = string(val)
				currentCellArr = append(currentCellArr, currentCell)
			}
			board = append(board, currentCellArr)
		}
	}
	return board
}

// DrawGameBoard draws a single GameBoard
// adopted from the same function from drawing.go in CellularAutomata
// * How do I do a unit-test just for this function?
func DrawGameBoard(board GameBoard, cellWidth int) image.Image {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := CreateNewPalettedCanvas(width, height, nil)

	// declare colors
	blue := MakeColor(80, 130, 255)
	red := MakeColor(150, 40, 30)
	green := MakeColor(0, 255, 0)

	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j].strategy == "C" {
				c.SetFillColor(blue)
			} else if board[i][j].strategy == "D" {
				c.SetFillColor(red)
			} else {
				c.SetFillColor(green)
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

	return GetImage(c)
}

// DrawGameBoards returns a list of boards; should use them to produce a gif
// ref: this is the same function from drawing.go in CellularAutomata
func DrawGameBoards(boards []GameBoard, cellWidth int) []image.Image {
	numGenerations := len(boards)
	imageList := make([]image.Image, numGenerations)
	for i := range boards {
		imageList[i] = DrawGameBoard(boards[i], cellWidth)
	}
	return imageList
}

// DrawAndSaveImgPNG draws and save the board to .png
func DrawAndSaveImgPNG(board GameBoard, cellWidth int) {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := CreateNewPalettedCanvas(width, height, nil)

	// declare colors
	blue := MakeColor(80, 130, 255)
	red := MakeColor(150, 40, 30)
	green := MakeColor(0, 255, 0)

	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j].strategy == "C" {
				c.SetFillColor(blue)
			} else if board[i][j].strategy == "D" {
				c.SetFillColor(red)
			} else {
				c.SetFillColor(green)
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}
	c.SaveToPNG("final.png")
}
