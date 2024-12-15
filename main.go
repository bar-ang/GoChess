package main

import (
	"fmt"
	"goChess/chess"
	"goChess/printer"
)

func main() {
	board := chess.NewChessBoard()
	board.SetStartingPos()

	sel, err := board.SelectPiece(7, 3)
	// move selected piece somehow

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printer.PrintSelect(&sel)
}
