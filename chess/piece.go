package chess

type PieceType int
type PlayerType int

const (
    None PieceType = iota
    Pawn
    Bishop
    Knight
    Rook
    Queen
    King
)

const (
    None PlayerType = iota
    White
    Black
)

type Square struct {
    x int
    y int
}

type Piece struct {
    pos Square
    type PieceType
    player PlayerType
}
