package chess

import "fmt"

type PieceType string
type PlayerType string

const (
    PieceNone   PieceType = "None"
    PiecePawn   PieceType = "Pawn"
    PieceBishop PieceType = "Bishop"
    PieceKnight PieceType = "Knight"
    PieceRook   PieceType = "Rook"
    PieceQueen  PieceType = "Queen"
    PieceKing   PieceType = "King"
)

const (
    PlayerNone  PlayerType = "None"
    PlayerWhite PlayerType = "White"
    PlayerBlack PlayerType = "Black"
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

func NoPiece() Piece {
    return Piece {pieceType: PieceNone}
}

func (p Piece) isPiece() bool {
    return p.pieceType != PieceNone
}

func NewPiece(x, y int, pieceType PieceType, player PlayerType) Piece {
    return Piece {
        pos: square{x, y},
        pieceType: pieceType,
        player: player,
    }
}

func (p Piece) String() string {
    return fmt.Sprintf("type: %v (player: %v) on: %v,%v", p.pieceType, p.player, p.pos.x, p.pos.y)
}
