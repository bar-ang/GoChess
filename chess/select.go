package chess

import "fmt"

var IllegalMoveError = fmt.Errorf("Illegal Move Error")
var EmptySquareSelectedError = fmt.Errorf("Empty Square Selected Error")

type square struct {
    x int
    y int
}

func (sq square) comp(x, y int) bool {
    return sq.x == x && sq.y == y
}

type Select struct {
    board *Board
    selected square
    possibleMoves []square
    threatenPieces []square
}

func (s *Select) moveSelectedPiece(toX, toY int) (*Board, error) {
    for _, sq := range s.possibleMoves {
        if sq.comp(toX, toY) {
            board := s.board.copy()
            if err := board.repositionPiece(sq.x, sq.y, toX, toY); err != nil {
                return nil, err
            }
            return board, nil
        }
    }

    return nil, IllegalMoveError
}

func (s *Select) Board() *Board {
    fmt.Println("USING UNIMPEMENTED FUNC Board()@select.go!")
    return s.board
}
