package main

import (
	"testing"
	"goChess/chess"
	"goChess/printer"
	"github.com/stretchr/testify/require"
)

type squareTest struct {
    x int
    y int
}

func sqr(x, y int) squareTest {
    return squareTest{x: x, y: y}
}

func TestShowBoards(t *testing.T) {
	cases := []struct {
		testName string
		setPieces map[squareTest]chess.Piece
		selected squareTest
	} {
		{
			testName: "RookInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(3, 2): chess.NewPiece(chess.PieceRook, chess.PlayerWhite),
			},
			selected: sqr(3, 2),
		},
		{
			testName: "RookInCorner",
			setPieces: map[squareTest]chess.Piece{
				sqr(7, 7): chess.NewPiece(chess.PieceRook, chess.PlayerWhite),
			},
			selected: sqr(7, 7),
		},
		{
			testName: "BishopInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(3, 2): chess.NewPiece(chess.PieceBishop, chess.PlayerWhite),
			},
			selected: sqr(3, 2),
		},
		{
			testName: "BishopInCorner",
			setPieces: map[squareTest]chess.Piece{
				sqr(7, 7): chess.NewPiece(chess.PieceBishop, chess.PlayerWhite),
			},
			selected: sqr(7, 7),
		},
		{
			testName: "QueenInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(4, 5): chess.NewPiece(chess.PieceQueen, chess.PlayerWhite),
			},
			selected: sqr(4, 5),
		},
		{
			testName: "QueenInCorner",
			setPieces: map[squareTest]chess.Piece{
				sqr(7, 0): chess.NewPiece(chess.PieceQueen, chess.PlayerWhite),
			},
			selected: sqr(7, 0),
		},
		{
			testName: "WhitePawnInStartingPos",
			setPieces: map[squareTest]chess.Piece{
				sqr(6, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(6, 2),
		},
		{
			testName: "WhitePawnInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(4, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(4, 2),
		},
		{
			testName: "WhitePawnInEnd",
			setPieces: map[squareTest]chess.Piece{
				sqr(1, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(1, 2),
		},
		{
			testName: "BlackPawnInStartingPos",
			setPieces: map[squareTest]chess.Piece{
				sqr(1, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerBlack),
			},
			selected: sqr(1, 2),
		},
		{
			testName: "BlackPawnInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(4, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerBlack),
			},
			selected: sqr(4, 2),
		},
		{
			testName: "BlackPawnInEnd",
			setPieces: map[squareTest]chess.Piece{
				sqr(6, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerBlack),
			},
			selected: sqr(6, 2),
		},

		{
			testName: "BlackKnightInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(3, 5): chess.NewPiece(chess.PieceKnight, chess.PlayerBlack),
			},
			selected: sqr(3, 5),
		},
		{
			testName: "BlackKnightInEdge",
			setPieces: map[squareTest]chess.Piece{
				sqr(0, 5): chess.NewPiece(chess.PieceKnight, chess.PlayerBlack),
			},
			selected: sqr(0, 5),
		},

		{
			testName: "WhiteKnightInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(3, 3): chess.NewPiece(chess.PieceKnight, chess.PlayerWhite),
			},
			selected: sqr(3, 3),
		},
		{
			testName: "WhiteKingtInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(3, 3): chess.NewPiece(chess.PieceKing, chess.PlayerWhite),
			},
			selected: sqr(3, 3),
		},
		{
			testName: "BlackKingtInCenter",
			setPieces: map[squareTest]chess.Piece{
				sqr(7, 0): chess.NewPiece(chess.PieceKing, chess.PlayerBlack),
			},
			selected: sqr(7, 0),
		},
		{
			testName: "BlackRookThreatsPawnI",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(2, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(2, 6),
		},
		{
			testName: "BlackRookThreatsPawnII",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(2, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
				sqr(3, 6): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(2, 6),
		},
		{
			testName: "BlackRookThreatsPawnIII",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(2, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
				sqr(2, 1): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(2, 6),
		},
		{
			testName: "BlackRookThreatsPawnIV",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(2, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerBlack),
				sqr(3, 6): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(2, 6),
		},
		{
			testName: "BlackRookThreatsPawnV",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(2, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerBlack),
				sqr(3, 6): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(3, 6),
		},
		{
			testName: "BlackRookThreatsPawnVI",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(2, 2): chess.NewPiece(chess.PiecePawn, chess.PlayerBlack),
				sqr(5, 6): chess.NewPiece(chess.PiecePawn, chess.PlayerWhite),
			},
			selected: sqr(5, 6),
		},
		{
			testName: "BlackRookCheckingI",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(7, 4): chess.NewPiece(chess.PieceRook, chess.PlayerWhite),
				sqr(2, 2): chess.NewPiece(chess.PieceKing, chess.PlayerWhite),
			},
			selected: sqr(2, 6),
		},
		{
			testName: "BlackRookCheckingII",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(7, 4): chess.NewPiece(chess.PieceRook, chess.PlayerWhite),
				sqr(2, 2): chess.NewPiece(chess.PieceKing, chess.PlayerWhite),
			},
			selected: sqr(2, 2),
		},
		{
			testName: "BlackRookCheckingIII",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(7, 4): chess.NewPiece(chess.PieceRook, chess.PlayerWhite),
				sqr(2, 2): chess.NewPiece(chess.PieceKing, chess.PlayerWhite),
			},
			selected: sqr(7, 4),
		},
		{
			testName: "BlackRookCheckingIV",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(7, 4): chess.NewPiece(chess.PieceRook, chess.PlayerWhite),
				sqr(2, 2): chess.NewPiece(chess.PieceKing, chess.PlayerWhite),

				sqr(7, 0): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(7, 7): chess.NewPiece(chess.PieceKing, chess.PlayerBlack),
			},
			selected: sqr(7, 4),
		},
		{
			testName: "BlackRookCheckingV",
			setPieces: map[squareTest]chess.Piece{
				sqr(2, 6): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(7, 4): chess.NewPiece(chess.PieceRook, chess.PlayerWhite),
				sqr(2, 2): chess.NewPiece(chess.PieceKing, chess.PlayerWhite),

				sqr(7, 0): chess.NewPiece(chess.PieceRook, chess.PlayerBlack),
				sqr(7, 7): chess.NewPiece(chess.PieceKing, chess.PlayerWhite),
			},
			selected: sqr(7, 4),
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
	        board := chess.NewChessBoard()
			for sq, piece := range c.setPieces {
	        	board.SetPiece(sq.x, sq.y, piece)
			}
	        sel, err := board.SelectPiece(c.selected.x, c.selected.y)
	        require.NoError(t, err)

	        printer.PrintSelect(&sel)
	    })
	}
}
