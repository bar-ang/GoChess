package chess

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/require"
)

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
    t.Run("SelectRook", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(3, 2, NewPiece(PieceRook, PlayerWhite))
        _, err := board.SelectPiece(3, 2)
        require.NoError(t, err)
    })
}
