package printer

import (
    "fmt"
    "goChess/chess"
)

func ChessPieceToString(piece chess.Piece) string {
    switch t := piece.Type(); t {
    case chess.PieceKing:
        return "K"
    case chess.PieceQueen:
        return "Q"
    case chess.PieceKnight:
        return "N"
    case chess.PieceRook:
        return "R"
    case chess.PieceBishop:
        return "B"
    case chess.PiecePawn:
        return "P"
    default:
        return " "
    }
}

func PrintChessBoard(board *chess.Board) {
    for i := 0; i < chess.BoardSize; i++ {
        for j := 0; j < chess.BoardSize; j++ {
            p := board.GetPiece(i, j)
            fmt.Printf("%v", ChessPieceToString(p))
        }
        fmt.Printf("\n")
    }
}

func PrintSelect(sel *chess.Select) {
    PrintChessBoard(sel.Board())
}
