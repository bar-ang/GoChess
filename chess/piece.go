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

type Piece struct {
    pieceType PieceType
    player PlayerType
}

func NoPiece() Piece {
    return Piece {pieceType: PieceNone}
}

func (p Piece) isPiece() bool {
    return p.pieceType != PieceNone
}

func NewPiece(pieceType PieceType, player PlayerType) Piece {
    return Piece {
        pieceType: pieceType,
        player: player,
    }
}

func (p Piece) String() string {
    return fmt.Sprintf("type: %v (player: %v)", p.pieceType, p.player)
}

func (p Piece) Type() PieceType {
    return p.pieceType
}
