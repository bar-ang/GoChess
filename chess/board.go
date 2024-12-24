package chess

import "fmt"

const BoardSize = 8
var RepositionEmptySquareError = fmt.Errorf("No piece is selected")
var RepositionPieceToSameSquareError = fmt.Errorf("Tried to reposition piece from one position to the same.")


type Board struct {
    pieces [][]Piece
    active [][]bool
}

func NewChessBoard() *Board {
    pieces := make([][]Piece, BoardSize)
    active := make([][]bool, BoardSize)
    for i, _ := range pieces {
        pieces[i] = make([]Piece, BoardSize)
        active[i] = make([]bool, BoardSize)
        for j, _ := range pieces[i] {
            pieces[i][j] = NoPiece()
        }
    }

    return &Board{
        pieces: pieces,
        active: active,
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

func (b *Board) SetPiece(x, y int, piece Piece) {
    b.pieces[x][y] = piece
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

func (b *Board) repositionPiece(fromX, fromY, toX, toY int) (*Board, error) {
    if fromX == toX && fromY == toY {
        return nil, RepositionPieceToSameSquareError
    }

    nb := b.copy()
    p := nb.pieces[fromX][fromY]
    if !p.isPiece() {
        return nil, RepositionEmptySquareError
    }

    nb.pieces[fromX][fromY] = NoPiece()
    nb.pieces[toX][toY] = p

    nb.active[fromX][fromY] = true

    return nb, nil
}

func (b *Board) rightCastleAvailable(kingX, kingY int) bool {
    return b.castleAvailable(kingX, kingY, true)
}

func (b *Board) leftCastleAvailable(kingX, kingY int) bool {
    return b.castleAvailable(kingX, kingY, false)
}

func (b *Board) castleAvailable(kingX, kingY int, right bool) bool {
    if b.active[kingX][kingY] {
        return false
    }

    king := b.GetPiece(kingX, kingY)
    if king.pieceType != PieceKing {
        panic("Absurd board position :(")
    }

    s := 1
    t := kingY
    p := b.GetPiece(kingX, 0)
    if right {
        s = kingX + 1
        t = BoardSize - 1
        p = b.GetPiece(kingX, BoardSize - 1)
    }

    if p.pieceType != PieceRook || p.player != king.player {
        return false
    }

    for i := s; i < t; i++ {
        if b.hasPiece(kingX, i) {
            return false
        }
    }

    return true
}

func (b *Board) copy() *Board {
    nPieces := make([][]Piece, len(b.pieces))
    for i := range b.pieces {
        nPieces[i] = make([]Piece, len(b.pieces[i]))
        copy(nPieces[i], b.pieces[i])
    }
    return &Board{ pieces : nPieces }
}

func (b *Board) GetPiece(x, y int) Piece {
    return b.pieces[x][y]
}

func (b *Board) hasPiece(x, y int) bool {
    return b.pieces[x][y].isPiece()
}

func (b *Board) isActive(x, y int) bool {
    return b.active[x][y]
}

func (b *Board) promotionNeeded(x, y int) bool {
    if x > 0 && x < 7 {
        return false
    }
    p := b.GetPiece(x, y)
    return p.pieceType == PiecePawn && ((p.player == PlayerWhite && x == 0) || (p.player == PlayerBlack && x == 7))
}

func (b *Board) promote(x, y int, newType PieceType) {
    p := b.GetPiece(x, y)
    b.SetPiece(x, y, NewPiece(newType, p.player))
}

func (b *Board) InCheck(player PlayerType) bool {
    for i := 0; i < BoardSize; i++ {
        for j := 0; j < BoardSize; j++ {
            piece := b.GetPiece(i, j)
            if piece.isPiece() && piece.player != player {
                sel, err := b.SelectPieceIgnoreCheck(i, j)
                if err != nil {
                    panic("cannot verify check")
                }

                if sel.checking {
                    return true
                }
            }
        }
    }

    return false
}

func  (b *Board) SelectPiece(x, y int) (Select, error) {
    sel, err := b.SelectPieceIgnoreCheck(x, y)
    if err != nil {
        return Select{}, err
    }

    sel.removePossibleMovesDueToCheck()

    return sel, nil
}

func (b *Board) applySpecialRules(sx, sy, tx, ty int) {
    if b.promotionNeeded(tx, ty) {
        b.promote(tx, ty, PieceQueen)
    }
}

func  (b *Board) SelectPieceIgnoreCheck(x, y int) (Select, error) {
    sel := Select{}
    piece := b.GetPiece(x, y)
    switch t := piece.pieceType; t {
    case PieceQueen:
        sel = b.selectQueen(x, y)
    case PieceRook:
        sel = b.selectRook(x, y)
    case PieceBishop:
        sel = b.selectBishop(x, y)
    case PieceKing:
        sel = b.selectKing(x, y)
    case PiecePawn:
        sel = b.selectPawn(x, y)
    case PieceKnight:
        sel = b.selectKnight(x, y)
    default:
        return Select{}, EmptySquareSelectedError
    }

    return sel, nil
}

func (b *Board) Size() int {
    return len(b.pieces)
}

func (b *Board) SelectNone() Select {
    return Select{board: b, selected: sqr(-1, -1)}
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
        sqr(-1, -1),
        sqr(-1,  0),
        sqr(-1,  1),
        sqr( 0, -1),
        sqr( 0,  1),
        sqr( 1, -1),
        sqr( 1,  0),
        sqr( 1,  1),
    }

    return b.selectRookOrBishopOrQueenByDirs(x, y, dirs)
}

func (b *Board) selectKnight(x, y int) Select {
    selected := b.GetPiece(x, y)

    sel :=  Select {
        board: b,
        selected: sqr(x, y),
        possibleMoves: make([]square, 0, 8),
        threatenPieces: make([]square, 0, 8),
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

    for _, d := range dir {
        sq := sqr(x+d.x, y+d.y)
        if !sq.inBounds() {
            continue
        }
        if p := b.GetPiece(sq.x, sq.y); p.isPiece() {
            if p.player != selected.player {
                sel.threat(sq)
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

    dir := -1
    if selected.player == PlayerBlack {
        dir = 1
    }

    short := sqr(x+dir, y)
    long := sqr(x+2*dir, y)
    eatRight := sqr(x+dir, y+dir)
    eatLeft := sqr(x+dir, y-dir)

    if !b.hasPiece(short.x, short.y) {
        sel.possibleMoves = append(sel.possibleMoves, short)
        if ((selected.player == PlayerBlack && x==1) || (selected.player == PlayerWhite && x==BoardSize-2)) {
            if !b.hasPiece(long.x, long.y) {
                sel.possibleMoves = append(sel.possibleMoves, long)
            }
        }
    }

    if p := b.GetPiece(eatRight.x, eatRight.y); p.isPiece() && p.player != selected.player {
        sel.threat(eatRight)
    }

    if p := b.GetPiece(eatLeft.x, eatLeft.y); p.isPiece() && p.player != selected.player {
        sel.threat(eatLeft)
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
        possibleCastle: make([]square, 0, 2),
    }

    for i := -1; i < 2; i++ {
        for j := -1; j < 2; j++ {
            sq := sqr(x+i, y+j)
            if (i != 0 || j != 0) && sq.inBounds() {
                p := b.GetPiece(x+i, y+j)
                if p.isPiece() {
                    if p.player != selected.player {
                        sel.threat(sq)
                    }
                } else {
                    sel.possibleMoves = append(sel.possibleMoves, sq)
                }
            }
        }
    }

    if b.rightCastleAvailable(x, y) {
        sel.possibleCastle = append(sel.possibleCastle, sqr(x, y+2))
    }

    if b.leftCastleAvailable(x, y) {
        sel.possibleCastle = append(sel.possibleCastle, sqr(x, y-2))
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
                    sel.threat(sq)
                }
                break
            } else {
                sel.possibleMoves = append(sel.possibleMoves, sq)
            }
        }
    }

    return sel
}
