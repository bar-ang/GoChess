package chess

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestPieceSelectionOnEmptyBoard(t *testing.T) {
    t.Run("SelectRook", func(t *testing.T) {
        board := NewChessBoard()
        board.SetPiece(3, 2, NewPiece(PieceRook, PlayerWhite))
        _, err := board.SelectPiece(3, 2)
        require.NoError(t, err)
    })
}
