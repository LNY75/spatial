package main

import "fmt"

// maxCoord returns the coordinate of the maximum element in the array [i, j]
func maxCoord(Nbrs [][]Cell) []int {
	maxScore := Nbrs[1][1].score
	maxCoord := []int{1, 1}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if Nbrs[i][j].score > maxScore {
				maxScore = Nbrs[i][j].score
				maxCoord[0] = i
				maxCoord[1] = j
			}
		}
	}
	return maxCoord
}

// detScore returns the new score of the current cell
func detScore(Nbrs [][]Cell, b float64) float64 {
	numCpNbrs := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if Nbrs[i][j].strategy == "C" {
				numCpNbrs++
			}
		}
	}
	// fmt.Print(numCpNbrs, " ")
	if Nbrs[1][1].strategy == "C" {
		return float64(numCpNbrs - 1)
	} else {
		return float64(numCpNbrs) * b
	}
}

// detStrat determine a new strategy for the current cell
func detStrat(Nbrs [][]Cell) string {
	maxCoord := maxCoord(Nbrs)
	return Nbrs[maxCoord[0]][maxCoord[1]].strategy
}

// store the neighbors of the current cell
func getNbrs(currentBoard GameBoard, i int, j int) [][]Cell {
	// fmt.Println(i, j)
	var Nbrs [][]Cell
	for r := i - 1; r <= i+1; r++ {
		var rowNbrs []Cell
		for c := j - 1; c <= j+1; c++ {
			// fmt.Print("r=", r, ",c=", c, "|| ")
			if r >= 0 && c >= 0 && r < len(currentBoard) && c < len(currentBoard[0]) {
				rowNbrs = append(rowNbrs, currentBoard[r][c])
				// fmt.Printf("%5.2f%1s |", currentBoard[r][c].score, currentBoard[r][c].strategy)
			} else {
				emptyCell := Cell{strategy: "N", score: -1.0}
				rowNbrs = append(rowNbrs, emptyCell)
				// fmt.Printf("%5.2f%1s |", -1.0, "N")
			}
		}
		// fmt.Println()
		Nbrs = append(Nbrs, rowNbrs)
	}
	// fmt.Println()
	return Nbrs
}

// helps debug
func PrintBoard(board GameBoard) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%5.2f%1s|", board[i][j].score, board[i][j].strategy)
		}
		fmt.Println()
	}
	fmt.Println()
}

// UpdateBoard updates the GameBoard according to the scores of each cell
func UpdateBoard(currentBoard GameBoard, b float64) GameBoard {
	var newBoard GameBoard
	// let's calculate scores first
	for i := 0; i < len(currentBoard); i++ {
		var newCells []Cell
		for j := 0; j < len(currentBoard[0]); j++ {
			Nbrs := getNbrs(currentBoard, i, j)
			var newCell Cell
			newCell.strategy = currentBoard[i][j].strategy
			newCell.score = detScore(Nbrs, b)
			newCells = append(newCells, newCell)
		}
		newBoard = append(newBoard, newCells)
	}

	// let's all change our strategies, but we can't let's our neighbors know
	var stratsSecret [][]string
	for i := 0; i < len(currentBoard); i++ {
		var rowStratsSecret []string
		for j := 0; j < len(currentBoard[0]); j++ {
			Nbrs := getNbrs(newBoard, i, j)
			strat := detStrat(Nbrs)
			rowStratsSecret = append(rowStratsSecret, strat)
		}
		stratsSecret = append(stratsSecret, rowStratsSecret)
	}

	// perform the change altogether
	for i := 0; i < len(newBoard); i++ {
		for j := 0; j < len(newBoard[0]); j++ {
			newBoard[i][j].strategy = stratsSecret[i][j]
		}
	}

	return newBoard
}

// PlaySpatial
func PlaySpatial(initialBoard GameBoard, b float64, iter int) []GameBoard {
	boards := make([]GameBoard, iter+1)
	boards[0] = initialBoard
	for i := 1; i <= iter; i++ {
		boards[i] = UpdateBoard(boards[i-1], b)
	}
	return boards
}
