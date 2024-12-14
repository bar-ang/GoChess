package main

import (
	"fmt"
	"goChess/chess"
)

func main() {
	p := chess.NewPiece(chess.PieceBishop, chess.PlayerBlack)
	fmt.Printf("P: %v\n", p)
}
