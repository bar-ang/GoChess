package chess

import "fmt"

const BoardSize = 8
var RepositionEmptySquareError = fmt.Errorf("No piece is selected")


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

func (b *Board) setKingsInStartingPos() {
    b.pieces[0][4] = NewPiece(PieceKing, PlayerBlack)
    b.pieces[BoardSize-1][3] = NewPiece(PieceKing, PlayerWhite)
}

func (b *Board) setQueensInStartingPos() {
    b.pieces[0][3] = NewPiece(PieceQueen, PlayerBlack)
    b.pieces[BoardSize-1][4] = NewPiece(PieceQueen, PlayerWhite)
}

func (b *Board) setBishopsInStartingPos() {
    b.pieces[0][2] = NewPiece(PieceBishop, PlayerBlack)
    b.pieces[0][5] = NewPiece(PieceBishop, PlayerBlack)
    b.pieces[BoardSize-1][2] = NewPiece(PieceBishop, PlayerWhite)
    b.pieces[BoardSize-1][5] = NewPiece(PieceBishop, PlayerWhite)
}

func (b *Board) setKnightsInStartingPos() {
    b.pieces[0][1] = NewPiece(PieceKnight, PlayerBlack)
    b.pieces[0][6] = NewPiece(PieceKnight, PlayerBlack)
    b.pieces[BoardSize-1][1] = NewPiece(PieceKnight, PlayerWhite)
    b.pieces[BoardSize-1][6] = NewPiece(PieceKnight, PlayerWhite)
}

func (b *Board) setRooksInStartingPos() {
    b.pieces[0][0] = NewPiece(PieceRook, PlayerBlack)
    b.pieces[0][7] = NewPiece(PieceRook, PlayerBlack)
    b.pieces[BoardSize-1][0] = NewPiece(PieceRook, PlayerWhite)
    b.pieces[BoardSize-1][7] = NewPiece(PieceRook, PlayerWhite)
}

func (b *Board) setPawnsInStartingPos() {
    for i := 0; i < 8; i++ {
        b.pieces[1][i] = NewPiece(PiecePawn, PlayerBlack)
        b.pieces[6][i] = NewPiece(PiecePawn, PlayerWhite)
    }
}

func (b *Board) SetStartingPos() {
    b.setKingsInStartingPos()
    b.setQueensInStartingPos()
    b.setBishopsInStartingPos()
    b.setKnightsInStartingPos()
    b.setRooksInStartingPos()
    b.setPawnsInStartingPos()
}

func (b *Board) repositionPiece(fromX, fromY, toX, toY int) error {
    p := b.pieces[fromX][fromY]
    if !p.isPiece() {
        return RepositionEmptySquareError
    }

    b.pieces[fromX][fromY] = NoPiece()
    b.pieces[toX][toY] = p

    return nil
}

func (b *Board) copy() *Board {
    // TODO: MUST COMPLETE THIS FUNC!
    fmt.Printf("USING NON IMPLEMENTED FUNCTION 'copy()@board.go'!\n")
    return b
}

func (b *Board) GetPiece(x, y int) Piece {
    return b.pieces[x][y]
}

func (b *Board) hasPiece(x, y int) bool {
    return b.pieces[x][y].isPiece()
}

func  (b *Board) SelectPiece(x, y int) (Select, error) {
    piece := b.GetPiece(x, y)
    switch t := piece.pieceType; t {
    case PieceQueen:
        return b.selectQueen(x, y), nil
    case PieceRook:
        return b.selectRook(x, y), nil
    case PieceBishop:
        return b.selectBishop(x, y), nil
    default:
        return Select{}, EmptySquareSelectedError
    }
}

func (b *Board) selectRook(x, y int) Select {
    dirs := []square{
        sqr(-1,  0),
        sqr( 1,  0),
        sqr( 0, -1),
        sqr( 0,  1),
    }

    return b.selectRookOrBishopOrQueenByDirs(x, y, dirs)
}

func (b *Board) selectBishop(x, y int) Select {
    dirs := []square{
        sqr(-1,  1),
        sqr( 1,  1),
        sqr( 1, -1),
        sqr(-1, -1),
    }

    return b.selectRookOrBishopOrQueenByDirs(x, y, dirs)
}

func (b *Board) selectQueen(x, y int) Select {
    dirs := []square{
        sqr(-1,  0),
        sqr( 1,  0),
        sqr( 0, -1),
        sqr( 0,  1),
        sqr(-1,  1),
        sqr( 1,  1),
        sqr( 1, -1),
        sqr( 1,  1),
    }

    return b.selectRookOrBishopOrQueenByDirs(x, y, dirs)
}

func (b *Board) selectRookOrBishopOrQueenByDirs(x, y int, dirs []square) Select {
    selected := b.GetPiece(x, y)
    possible := make([]square, 0, 15)
    threaten := make([]square, 0, 4)

    for _, dir := range dirs {
        for i := 1; i < BoardSize; i++ {
            sq := sqr(x+i*dir.x, y+i*dir.y)
            if !sq.inBounds() {
                continue
            }
            pc := b.GetPiece(sq.x, sq.y)
            if pc.isPiece() {
                if pc.player != selected.player {
                    possible = append(possible, sq)
                    threaten = append(threaten, sq)
                }
                break
            } else {
                possible = append(possible, sq)
            }
        }
    }

    return Select {
        board: b,
        selected: sqr(x, y),
        possibleMoves: possible,
        threatenPieces: threaten,
    }
}
