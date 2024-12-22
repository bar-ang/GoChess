package chess

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/require"
)

func TestBoardCopy(t *testing.T) {
    board := NewChessBoard()
    board.SetPiece(0, 0, NewPiece(PieceKing, PlayerBlack))
    board.SetPiece(1, 1, NewPiece(PieceQueen, PlayerWhite))
    board.SetPiece(6, 5, NewPiece(PiecePawn, PlayerBlack))

    nBoard := board.copy()

    for i, row := range board.pieces {
        for j, v := range row {
            require.Equal(t, nBoard.pieces[i][j], v)
        }
    }

    nBoard.pieces[6][5] = NoPiece()
    nBoard.pieces[6][6] = NewPiece(PieceBishop, PlayerWhite)

    require.Equal(t, board.pieces[6][5], NewPiece(PiecePawn, PlayerBlack))
    require.Equal(t, board.pieces[6][6], NoPiece())
}

func TestRepositionPiece(t *testing.T) {
    t.Run("NoError", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(0, 0, NewPiece(PieceKing, PlayerBlack))
        board.SetPiece(1, 1, NewPiece(PieceQueen, PlayerWhite))
        board.SetPiece(6, 5, NewPiece(PiecePawn, PlayerBlack))

        nb, err := board.repositionPiece(6, 5, 3, 4)
        require.NoError(t, err)

        require.Equal(t, nb.pieces[6][5], NoPiece())
        require.Equal(t, nb.pieces[3][4], NewPiece(PiecePawn, PlayerBlack))
        require.Equal(t, nb.pieces[1][1], NewPiece(PieceQueen, PlayerWhite))
        require.Equal(t, nb.pieces[0][0], NewPiece(PieceKing, PlayerBlack))

        nb2, err := nb.repositionPiece(3, 4, 1, 1)
        require.NoError(t, err)

        require.Equal(t, nb2.pieces[3][4], NoPiece())
        require.Equal(t, nb2.pieces[1][1], NewPiece(PiecePawn, PlayerBlack))
        require.Equal(t, nb2.pieces[0][0], NewPiece(PieceKing, PlayerBlack))
    })
    t.Run("RepositionEmptySquareError", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(0, 0, NewPiece(PieceKing, PlayerBlack))
        board.SetPiece(1, 1, NewPiece(PieceQueen, PlayerWhite))
        board.SetPiece(6, 5, NewPiece(PiecePawn, PlayerBlack))

        _, err := board.repositionPiece(6, 6, 7, 7)
        require.ErrorIs(t, err, RepositionEmptySquareError)
    })
}

func TestSelectBasic(t *testing.T) {
    pieces := []Piece {
        NewPiece(PieceRook, PlayerWhite),
        NewPiece(PieceRook, PlayerBlack),
        NewPiece(PieceBishop, PlayerWhite),
        NewPiece(PieceBishop, PlayerBlack),
        NewPiece(PieceKnight, PlayerWhite),
        NewPiece(PieceKnight, PlayerBlack),
        NewPiece(PieceQueen, PlayerWhite),
        NewPiece(PieceQueen, PlayerBlack),
        NewPiece(PieceKing, PlayerWhite),
        NewPiece(PieceKing, PlayerBlack),
        NewPiece(PiecePawn, PlayerWhite),
        NewPiece(PiecePawn, PlayerBlack),
    }
    for _, piece := range pieces {
        t.Run(fmt.Sprintf("Select%s%s", piece.player, piece.pieceType), func(t *testing.T) {
            board := NewChessBoard()
            board.SetPiece(3, 2, piece)
            sel, err := board.SelectPiece(3, 2)
            require.NoError(t, err)

            require.Equal(t, sel.selected.x, 3)
            require.Equal(t, sel.selected.y, 2)
            require.Equal(t, sel.board, board)
            require.False(t, sel.checking)
        })
    }
}

func TestPieceSelectionOnEmptyBoard(t *testing.T) {
    t.Run("SelectRookInCenter", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(3, 2, NewPiece(PieceRook, PlayerWhite))
        sel, err := board.SelectPiece(3, 2)
        require.NoError(t, err)
        require.NotEmpty(t, sel.possibleMoves)
        require.Empty(t, sel.threatenPieces)

        require.Len(t, sel.possibleMoves, 14)

        require.Contains(t, sel.possibleMoves, sqr(3, 0))
        require.Contains(t, sel.possibleMoves, sqr(3, 1))
        require.Contains(t, sel.possibleMoves, sqr(3, 3))
        require.Contains(t, sel.possibleMoves, sqr(3, 4))
        require.Contains(t, sel.possibleMoves, sqr(3, 5))
        require.Contains(t, sel.possibleMoves, sqr(3, 6))
        require.Contains(t, sel.possibleMoves, sqr(3, 7))

        require.Contains(t, sel.possibleMoves, sqr(0, 2))
        require.Contains(t, sel.possibleMoves, sqr(1, 2))
        require.Contains(t, sel.possibleMoves, sqr(2, 2))
        require.Contains(t, sel.possibleMoves, sqr(4, 2))
        require.Contains(t, sel.possibleMoves, sqr(5, 2))
        require.Contains(t, sel.possibleMoves, sqr(6, 2))
        require.Contains(t, sel.possibleMoves, sqr(7, 2))
    })
    t.Run("SelectRookInCorner", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(7, 7, NewPiece(PieceRook, PlayerWhite))
        sel, err := board.SelectPiece(7, 7)
        require.NoError(t, err)
        require.NotEmpty(t, sel.possibleMoves)
        require.Empty(t, sel.threatenPieces)

        require.Len(t, sel.possibleMoves, 14)

        require.Contains(t, sel.possibleMoves, sqr(7, 0))
        require.Contains(t, sel.possibleMoves, sqr(7, 1))
        require.Contains(t, sel.possibleMoves, sqr(7, 2))
        require.Contains(t, sel.possibleMoves, sqr(7, 3))
        require.Contains(t, sel.possibleMoves, sqr(7, 4))
        require.Contains(t, sel.possibleMoves, sqr(7, 5))
        require.Contains(t, sel.possibleMoves, sqr(7, 6))

        require.Contains(t, sel.possibleMoves, sqr(0, 7))
        require.Contains(t, sel.possibleMoves, sqr(1, 7))
        require.Contains(t, sel.possibleMoves, sqr(3, 7))
        require.Contains(t, sel.possibleMoves, sqr(2, 7))
        require.Contains(t, sel.possibleMoves, sqr(4, 7))
        require.Contains(t, sel.possibleMoves, sqr(5, 7))
        require.Contains(t, sel.possibleMoves, sqr(6, 7))
    })
    t.Run("SelectBishopInCenter", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(2, 3, NewPiece(PieceBishop, PlayerBlack))
        sel, err := board.SelectPiece(2, 3)
        require.NoError(t, err)
        require.NotEmpty(t, sel.possibleMoves)
        require.Empty(t, sel.threatenPieces)

        require.Len(t, sel.possibleMoves, 11)

        require.Contains(t, sel.possibleMoves, sqr(0, 1))
        require.Contains(t, sel.possibleMoves, sqr(1, 2))
        require.Contains(t, sel.possibleMoves, sqr(3, 4))
        require.Contains(t, sel.possibleMoves, sqr(4, 5))
        require.Contains(t, sel.possibleMoves, sqr(5, 6))
        require.Contains(t, sel.possibleMoves, sqr(6, 7))

        require.Contains(t, sel.possibleMoves, sqr(0, 5))
        require.Contains(t, sel.possibleMoves, sqr(1, 4))
        require.Contains(t, sel.possibleMoves, sqr(3, 2))
        require.Contains(t, sel.possibleMoves, sqr(4, 1))
        require.Contains(t, sel.possibleMoves, sqr(5, 0))
    })
    t.Run("SelectBishopInCorner", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(0, 7, NewPiece(PieceBishop, PlayerBlack))
        sel, err := board.SelectPiece(0, 7)
        require.NoError(t, err)
        require.NotEmpty(t, sel.possibleMoves)
        require.Empty(t, sel.threatenPieces)

        require.Len(t, sel.possibleMoves, 7)

        require.Contains(t, sel.possibleMoves, sqr(1, 6))
        require.Contains(t, sel.possibleMoves, sqr(2, 5))
        require.Contains(t, sel.possibleMoves, sqr(3, 4))
        require.Contains(t, sel.possibleMoves, sqr(4, 3))
        require.Contains(t, sel.possibleMoves, sqr(5, 2))
        require.Contains(t, sel.possibleMoves, sqr(6, 1))
        require.Contains(t, sel.possibleMoves, sqr(7, 0))
    })
    t.Run("SelectQueenInCenter", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(2, 3, NewPiece(PieceQueen, PlayerWhite))
        sel, err := board.SelectPiece(2, 3)
        require.NoError(t, err)
        require.NotEmpty(t, sel.possibleMoves)
        require.Empty(t, sel.threatenPieces)

        require.Contains(t, sel.possibleMoves, sqr(0, 1))
        require.Contains(t, sel.possibleMoves, sqr(1, 2))
        require.Contains(t, sel.possibleMoves, sqr(3, 4))
        require.Contains(t, sel.possibleMoves, sqr(4, 5))
        require.Contains(t, sel.possibleMoves, sqr(5, 6))
        require.Contains(t, sel.possibleMoves, sqr(6, 7))

        require.Contains(t, sel.possibleMoves, sqr(0, 5))
        require.Contains(t, sel.possibleMoves, sqr(1, 4))
        require.Contains(t, sel.possibleMoves, sqr(3, 2))
        require.Contains(t, sel.possibleMoves, sqr(4, 1))
        require.Contains(t, sel.possibleMoves, sqr(5, 0))

        require.Contains(t, sel.possibleMoves, sqr(2, 0))
        require.Contains(t, sel.possibleMoves, sqr(2, 1))
        require.Contains(t, sel.possibleMoves, sqr(2, 2))
        require.Contains(t, sel.possibleMoves, sqr(2, 4))
        require.Contains(t, sel.possibleMoves, sqr(2, 5))
        require.Contains(t, sel.possibleMoves, sqr(2, 6))
        require.Contains(t, sel.possibleMoves, sqr(2, 7))

        require.Contains(t, sel.possibleMoves, sqr(0, 3))
        require.Contains(t, sel.possibleMoves, sqr(1, 3))
        require.Contains(t, sel.possibleMoves, sqr(3, 3))
        require.Contains(t, sel.possibleMoves, sqr(4, 3))
        require.Contains(t, sel.possibleMoves, sqr(5, 3))
        require.Contains(t, sel.possibleMoves, sqr(6, 3))
        require.Contains(t, sel.possibleMoves, sqr(7, 3))

        require.Len(t, sel.possibleMoves, 7*2 + 11)
    })
}
