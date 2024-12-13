package main

import (
	"fmt"
	"goChess/chess"
)

func main() {
	p := chess.NewPiece(2, 4, chess.PieceBishop, chess.PlayerBlack)
	fmt.Printf("P: %v\n", p)
}
