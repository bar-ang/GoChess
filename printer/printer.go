package printer

import (
    "fmt"
    "goChess/chess"
    "github.com/fatih/color"
)

//Color Scheme
var BlackSquareColor = color.BgRGB(40, 40, 40)
var WhiteSquareColor = color.BgRGB(140, 140, 140)
var BlackPlayerColor = color.RGB(255, 0, 0)
var WhitePlayerColor = color.RGB(255, 255, 255)
var SelectedColor    = color.BgRGB(0, 255, 0)
var ThreatenedColor  = color.BgRGB(255, 0, 255)
var PossibleColor    = color.BgRGB(0, 0, 255)

type printUnit struct {
    piece        chess.Piece
    light        bool
    selected     bool
    threatened   bool
    possibleMove bool
}

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

func makePrintUnitsMap(sel *chess.Select) [][]printUnit {
    board := sel.Board()
    pu := make([][]printUnit, board.Size())

    for i := 0; i < board.Size(); i++ {
        pu[i] = make([]printUnit, board.Size())
        for j := 0; j < board.Size(); j++ {
            pu[i][j] = printUnit {
                piece: board.GetPiece(i, j),
                light: (i+j) % 2 != 0,
            }
        }
    }

    pu[sel.Selected().X()][sel.Selected().Y()].selected = true

    for _, v := range sel.ThreatenPieces() {
        pu[v.X()][v.Y()].threatened = true
    }

    for _, v := range sel.PossibleMoves() {
        pu[v.X()][v.Y()].possibleMove = true
    }

    return pu
}

func (pu printUnit) format() *color.Color {
    if pu.threatened {
        return ThreatenedColor
    }

    if pu.possibleMove {
        return PossibleColor
    }

    if pu.selected {
        return SelectedColor
    }

    if pu.light {
        return WhiteSquareColor
    }

    return BlackSquareColor
}

func PrintSelect(sel *chess.Select) {
    pu := makePrintUnitsMap(sel)
    for _, row := range pu {
        for _, v := range row {
            v.format().Printf(" %v ", ChessPieceToString(v.piece))
        }
        fmt.Println()
    }
}
