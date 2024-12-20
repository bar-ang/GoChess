package main

import (
	//"fmt"
	"goChess/chess"
	"goChess/printer"
)

func main() {
	board := chess.NewChessBoard()
	board.SetStartingPos()


	printer.PrintChessBoard(board)
}
