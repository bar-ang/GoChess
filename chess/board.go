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
    case PieceKing:
        return b.selectKing(x, y), nil
    case PiecePawn:
        return b.selectPawn(x, y), nil
    case PieceKnight:
        return b.selectKnight(x, y), nil
    default:
        return Select{}, EmptySquareSelectedError
    }
}

func (b *Board) Size() int {
    return len(b.pieces)
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

func (b *Board) selectKnight(x, y int) Select {
    selected := b.GetPiece(x, y)

    sel :=  Select {
        board: b,
        selected: sqr(x, y),
        possibleMoves: make([]square, 0, 4),
        threatenPieces: make([]square, 0, 2),
    }

    dir := []square {
        sqr( 1, 2),
        sqr(-1, 2),
        sqr( 1,-2),
        sqr(-1,-2),
        sqr( 2, 1),
        sqr( 2,-1),
        sqr(-2, 1),
        sqr(-2,-1),
    }

    for _, sq := range dir {
        if p := b.GetPiece(sq.x, sq.y); p.isPiece() {
            if p.player != selected.player {
                sel.possibleMoves = append(sel.possibleMoves, sq)
                sel.threatenPieces = append(sel.threatenPieces, sq)
            }
        } else {
            sel.possibleMoves = append(sel.possibleMoves, sq)
        }
    }

    return sel
}

func (b *Board) selectPawn(x, y int) Select {
    selected := b.GetPiece(x, y)

    sel := Select {
        board: b,
        selected: sqr(x, y),
        possibleMoves: make([]square, 0, 4),
        threatenPieces: make([]square, 0, 2),
    }

    dir := 1
    if selected.player == PlayerBlack {
        dir = -1
    }

    short := sqr(x, y+dir)
    long := sqr(x, y+2*dir)
    eatRight := sqr(x+dir, y+dir)
    eatLeft := sqr(x-dir, y+dir)

    if !b.hasPiece(short.x, short.y) {
        sel.possibleMoves = append(sel.possibleMoves, short)
        if ((selected.player == PlayerBlack && y==1) || (selected.player == PlayerWhite && y==BoardSize-2)) {
            if !b.hasPiece(long.x, long.y) {
                sel.possibleMoves = append(sel.possibleMoves, long)
            }
        }
    }

    if p := b.GetPiece(eatRight.x, eatRight.y); p.isPiece() && p.player != selected.player {
        sel.possibleMoves = append(sel.possibleMoves, eatRight)
        sel.threatenPieces = append(sel.threatenPieces, eatRight)
    }

    if p := b.GetPiece(eatLeft.x, eatLeft.y); p.isPiece() && p.player != selected.player {
        sel.possibleMoves = append(sel.possibleMoves, eatLeft)
        sel.threatenPieces = append(sel.threatenPieces, eatLeft)
    }

    return sel
}

func (b *Board) selectKing(x, y int) Select {
    selected := b.GetPiece(x, y)

    sel := Select {
        board: b,
        selected: sqr(x, y),
        possibleMoves: make([]square, 0, 8),
        threatenPieces: make([]square, 0, 8),
    }

    for i := -1; i < 2; i++ {
        for j := -1; j < 2; j++ {
            sq := sqr(x+i, y+j)
            if (i != 0 || j != 0) && sq.inBounds() {
                p := b.GetPiece(x+i, y+j)
                if p.isPiece() {
                    if p.player != selected.player {
                        sel.possibleMoves = append(sel.possibleMoves, sq)
                        sel.threatenPieces = append(sel.threatenPieces, sq)
                    }
                } else {
                    sel.possibleMoves = append(sel.possibleMoves, sq)
                }
            }
        }
    }

    return sel
}

func (b *Board) selectRookOrBishopOrQueenByDirs(x, y int, dirs []square) Select {
    selected := b.GetPiece(x, y)

    sel := Select {
        board: b,
        selected: sqr(x, y),
        possibleMoves: make([]square, 0, 15),
        threatenPieces: make([]square, 0, 4),
    }

    for _, dir := range dirs {
        for i := 1; i < BoardSize; i++ {
            sq := sqr(x+i*dir.x, y+i*dir.y)
            if !sq.inBounds() {
                continue
            }
            pc := b.GetPiece(sq.x, sq.y)
            if pc.isPiece() {
                if pc.player != selected.player {
                    sel.possibleMoves = append(sel.possibleMoves, sq)
                    sel.threatenPieces = append(sel.threatenPieces, sq)
                }
                break
            } else {
                sel.possibleMoves = append(sel.possibleMoves, sq)
            }
        }
    }

    return sel
}
