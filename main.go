package main

import (
	"fmt"
	"goChess/chess"
)

func main() {
	board := chess.NewChessBoard()
	board.SetStartingPos()
	fmt.Printf("P: %v\n", board)
}
