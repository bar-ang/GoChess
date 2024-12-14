package chess

const BoardSize = 8

type Board struct {
    pieces [][]Piece
}

func NewChessBoard() *Board {
    pieces := make([][]Piece, BoardSize)
    for i, _ := range pieces {
        pieces[i] = make([]Piece, BoardSize)
        for j, _ := range pieces[i] {
            pieces[i][j] = NoPiece()
        }
    }

    return &Board{
        pieces: pieces,
    }
}
