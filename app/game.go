package app

type Token int
type GameResult int

const (
	Draw             GameResult = 0
	PlayerOneVictory GameResult = 1
	PlayerTwoVictory GameResult = 2

	Red  Token = 1
	Blue Token = 2

	defaultPlayerOne = "Red"
	defaultPlayerTwo = "Blue"

	defaultWidth  = 7
	defaultHeigth = 6
)

type Game struct {
	PlayerOne string
	PlayerTwo string
	Board     *GameBoard
	Result    GameResult
}

func NewGame(playerOne, playerTwo string, width, heigth int) *Game {
	if playerOne == "" {
		playerOne = defaultPlayerOne
	}
	if playerTwo == "" {
		playerTwo = defaultPlayerTwo
	}
	return &Game{
		PlayerOne: playerOne,
		PlayerTwo: playerTwo,
		Board:     NewBoard(width, heigth),
	}
}

type GameBoard struct {
	board  map[int][]Token
	width  int
	heigth int
}

func NewBoard(width, heigth int) *GameBoard {
	if width == 0 {
		width = defaultWidth
	}
	if heigth == 0 {
		heigth = defaultHeigth
	}

	return &GameBoard{
		board:  make(map[int][]Token),
		width:  width,
		heigth: heigth,
	}
}

func (b *GameBoard) Add(column int, token Token) bool {
	if len(b.board[column]) == b.heigth {
		return false
	}
	b.board[column] = append(b.board[column], token)
	return true
}
