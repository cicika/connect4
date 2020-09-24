package app

import (
	"fmt"
)

type Token int
type GameResult int

var RedVictory = []Token{Red, Red, Red, Red}
var BlueVictory = []Token{Blue, Blue, Blue, Blue}

const (
	Draw             GameResult = 0
	PlayerOneVictory GameResult = 1
	PlayerTwoVictory GameResult = 2

	Empty Token = 0
	Red   Token = 1
	Blue  Token = 2

	defaultPlayerOne = "Red"
	defaultPlayerTwo = "Blue"

	defaultWidth  = 7
	defaultHeigth = 6
)

type Game struct {
	PlayerOne    string
	PlayerTwo    string
	Board        *GameBoard
	Result       GameResult
	GameOver     bool
	NextMoveOn   string
	TokensPlayed int
}

func StartGame(playerOne, playerTwo string, width, heigth int) *Game {
	if playerOne == "" {
		playerOne = defaultPlayerOne
	}
	if playerTwo == "" {
		playerTwo = defaultPlayerTwo
	}
	return &Game{
		PlayerOne:  playerOne,
		PlayerTwo:  playerTwo,
		Board:      NewBoard(width, heigth),
		NextMoveOn: playerOne,
	}
}

func (g *Game) MultipleMove(columns []int) (bool, int) {
	for i, c := range columns {
		res := g.SingleMove(c)
		if !res {
			return res, len(columns) - i + 1 //moves that didn't run
		}
	}
	return true, 0
}

func (g *Game) SingleMove(column int) bool {
	var t Token
	if g.NextMoveOn == g.PlayerOne {
		t = Blue
	} else {
		t = Red
	}
	res := g.Board.Add(column, t)
	if res {
		g.TokensPlayed++
	}
	if g.TokensPlayed >= 7 {
		g.Board.CheckForWinner()
	}

	if res && !g.Board.WinnerExists {
		if g.TokensPlayed%2 == 0 {
			g.NextMoveOn = g.PlayerOne
		} else {
			g.NextMoveOn = g.PlayerTwo
		}
	}
	draw := !g.Board.WinnerExists && g.TokensPlayed == g.Board.BoardSize
	if g.Board.WinnerExists || g.TokensPlayed == g.Board.BoardSize {
		g.GameOver = true
		if draw {
			g.Result = Draw
		} else {
			g.Result = PlayerOneVictory
			if g.NextMoveOn == g.PlayerTwo {
				g.Result = PlayerTwoVictory
			}
		}
	}
	return res
}

func (g *Game) End() {
	if g.Result == Draw {
		fmt.Println(DrawResult)
	} else {
		var player string
		if g.Result == PlayerOneVictory {
			player = g.PlayerOne
		} else {
			player = g.PlayerTwo
		}
		fmt.Printf(string(Victory)+"\n\r", player)
	}
}

type GameBoard struct {
	board        map[int][]Token
	rows         map[int][7]Token
	width        int
	heigth       int
	WinnerExists bool
	BoardSize    int
}

func NewBoard(width, heigth int) *GameBoard {
	if width == 0 {
		width = defaultWidth
	}
	if heigth == 0 {
		heigth = defaultHeigth
	}

	return &GameBoard{
		board:        make(map[int][]Token),
		rows:         make(map[int][7]Token),
		width:        width,
		heigth:       heigth,
		BoardSize:    width * heigth,
		WinnerExists: false,
	}
}

func (b *GameBoard) Add(column int, token Token) bool {
	if len(b.board[column]) == b.heigth {
		return false
	}
	b.board[column] = append(b.board[column], token)
	row := b.rows[len(b.board[column])-1]
	row[column] = token
	b.rows[len(b.board[column])-1] = row

	return true
}

func (b *GameBoard) AddMultiple(columns []int) (bool, int) {
	var token Token
	for p, c := range columns {
		if p%2 == 0 {
			token = Red
		} else {
			token = Blue
		}
		r := b.Add(c, token)
		if !r || b.WinnerExists {
			return false, p
		}
	}
	return true, len(columns)
}

func (b *GameBoard) CheckForWinner() {
	for _, column := range b.board {
		if len(column) >= 4 {
			b.checkSlice(column)
		}
	}
	if !b.WinnerExists {
		for _, row := range b.rows {
			var r []Token
			for _, x := range row {
				r = append(r, x)
			}
			b.checkSlice(r)
		}
	}

	if !b.WinnerExists {
		b.checkMatrices()
	}
}

func (b *GameBoard) checkSlice(in []Token) {
	for i, _ := range in {
		if len(in) < i+4+1 {
			return
		}
		s := in[i : i+4]
		if slicesMatch(s, RedVictory) || slicesMatch(s, BlueVictory) {
			b.WinnerExists = true
			return
		}
	}
}

func (b *GameBoard) checkMatrices() {
	for i, column := range b.board {
		for j, _ := range column {
			var s []Token
			// to the right
			_, ok1 := b.board[i+3]          // check if board column exists
			ok2 := len(b.board[i+3]) >= j+4 // check if element in the column exists
			if j <= 2 && ok1 && ok2 {
				s = []Token{
					b.board[i][j],
					b.board[i+1][j+1],
					b.board[i+2][j+2],
					b.board[i+3][j+3],
				}

			}
			// to the left
			_, ok1 = b.board[i-3]
			ok2 = len(b.board[i-3]) >= j+4
			if j >= 3 && ok1 && ok2 {
				s = []Token{
					b.board[i][j],
					b.board[i-1][j+1],
					b.board[i-2][j+2],
					b.board[i-3][j+3],
				}
			}
			if len(s) == 4 {
				if slicesMatch(s, RedVictory) || slicesMatch(s, BlueVictory) {
					b.WinnerExists = true
					return
				}
			}
		}
	}
}

func slicesMatch(s, cmp []Token) bool {
	if len(s) != len(cmp) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != cmp[i] {
			return false
		}
	}
	return true
}
