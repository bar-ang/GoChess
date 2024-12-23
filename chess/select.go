package chess

import "fmt"

var IllegalMoveError = fmt.Errorf("Illegal Move Error")
var EmptySquareSelectedError = fmt.Errorf("Empty Square Selected Error")

type square struct {
    x int
    y int
}

func (sq square) X() int {
    return sq.x
}

func (sq square) Y() int {
    return sq.y
}

func (sq square) comp(x, y int) bool {
    return sq.x == x && sq.y == y
}

func (sq square) inBounds() bool {
    return sq.x >= 0 && sq.y >= 0 && sq.x < BoardSize && sq.y < BoardSize
}

func sqr(x, y int) square {
    return square{x: x, y: y}
}

type Select struct {
    board *Board
    selected square
    possibleMoves []square
    threatenPieces []square
    checking bool
}

func (s *Select) Selected() square {
    return s.selected
}

func (s *Select) PossibleMoves() []square {
    return s.possibleMoves
}

func (s *Select) ThreatenPieces() []square {
    return s.threatenPieces
}

func (s *Select) removePossibleMovesDueToCheck() {
    possibles := make([]square, 0, len(s.possibleMoves))
    threatened := make([]square, 0, len(s.threatenPieces))

    for _, move := range s.possibleMoves {
        nboard, err := s.moveSelectedPiece(move.x, move.y)
        if err != nil {
            panic("could not look for possible checks.")
        }
        if !nboard.InCheck(s.Piece().player) {
            possibles = append(possibles, move)
        }
    }

    for _, move := range s.threatenPieces {
        nboard, err := s.moveSelectedPiece(move.x, move.y)
        if err != nil {
            panic("could not look for possible checks.")
        }
        if !nboard.InCheck(s.Piece().player) {
            threatened = append(threatened, move)
        }
    }

    s.possibleMoves = possibles
    s.threatenPieces = threatened
}

func (s *Select) moveSelectedPiece(toX, toY int) (*Board, error) {
    for _, sq := range s.possibleMoves {
        if sq.comp(toX, toY) {
            if board, err := s.board.repositionPiece(sq.x, sq.y, toX, toY); err != nil {
                return nil, err
            } else {
                return board, nil
            }
        }
    }

    return nil, IllegalMoveError
}

func (s *Select) threat(sq square) {
    piece := s.board.GetPiece(sq.x, sq.y)
    if !piece.isPiece() {
        panic("Trying to threat an empty square")
    }

    if piece.pieceType != PieceKing {
        s.threatenPieces = append(s.threatenPieces, sq)
        s.possibleMoves = append(s.possibleMoves, sq)
    } else {
        s.checking = true
    }
}

func (s *Select) Piece() Piece {
    return s.board.GetPiece(s.selected.x, s.selected.y)
}

func (s *Select) Checking() bool {
    return s.checking
}

func (s *Select) Board() *Board {
    return s.board.copy()
}
