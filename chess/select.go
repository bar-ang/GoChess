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
    possibleCastle []square
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

func (s *Select) canCastle() bool {
    return s.possibleCastle != nil && len(s.possibleCastle) > 0
}

func (s *Select) castle(x, y int) *Board {
    if !s.canCastle() {
        panic("castling should've already been verified possible!")
    }

    rookPos := sqr(s.selected.x, 0)
    dir := 1
    if x > s.selected.x {
        rookPos = sqr(s.selected.x, BoardSize-1)
        dir = -1
    }

    nb, err := s.board.repositionPiece(s.selected.x, s.selected.y, x, y)
    if err != nil {
        panic(fmt.Errorf("cannot castle non-king: %v", err))
    }
    nb, err = nb.repositionPiece(rookPos.x, rookPos.y, rookPos.x, y+dir)
    if err != nil {
        panic(fmt.Errorf("cannot move rook during castlling: %v", err))
    }

    return nb
}

func (s *Select) removePossibleMovesDueToCheck() {
    possibles := make([]square, 0, len(s.possibleMoves))
    threatened := make([]square, 0, len(s.threatenPieces))

    for _, move := range s.possibleMoves {
        nboard, err := s.board.repositionPiece(s.selected.x, s.selected.y, move.x, move.y)
        if err != nil {
            panic("could not look for possible checks.")
        }
        if !nboard.InCheck(s.Piece().player) {
            possibles = append(possibles, move)
        }
    }

    for _, move := range s.threatenPieces {
        nboard, err := s.board.repositionPiece(s.selected.x, s.selected.y, move.x, move.y)
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
            if board, err := s.board.repositionPiece(s.selected.x, s.selected.y, toX, toY); err != nil {
                return nil, err
            } else {
                board.applySpecialRules(s.selected.x, s.selected.y, toX, toY)
                return board, nil
            }
        }
    }

    if s.possibleCastle != nil {
        for _, sq := range s.possibleCastle {
            if sq.comp(toX, toY) {
                board := s.castle(toX, toY)
                board.applySpecialRules(s.selected.x, s.selected.y, toX, toY)
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
