package chess

import "fmt"

type PieceType int
type PlayerType int

const (
    PieceNone PieceType = iota
    PiecePawn
    PieceBishop
    PieceKnight
    PieceRook
    PieceQueen
    PieceKing
)

const (
    PlayerNone PlayerType = iota
    PlayerWhite
    PlayerBlack
)

type square struct {
    x int
    y int
}

type Piece struct {
    pos square
    pieceType PieceType
    player PlayerType
}

func NewPiece(x, y int, pieceType PieceType, player PlayerType) *Piece {
    return &Piece {
        pos: square{x, y},
        pieceType: pieceType,
        player: player,
    }
}

func (p *Piece) String() string {
    return fmt.Sprintf("type: %v (player: %v) on: %v,%v", p.pieceType, p.player, p.pos.x, p.pos.y)
}
