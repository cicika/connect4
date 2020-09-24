package app

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
)

type ConsoleMessage string

const (
	Hello                ConsoleMessage = "Hello! Welcome!"
	Default_GameStart    ConsoleMessage = "Starting a new game with default parameters..."
	RunningInitialMoves  ConsoleMessage = "Running initial moves..."
	MultipleMovesFailure ConsoleMessage = "Following moves didn't run: %+v"
	StartNew_Q           ConsoleMessage = "Would you like to start a new game? (y/n)"
	PlayerXMove          ConsoleMessage = "%s's move... Please, enter column"
	Victory              ConsoleMessage = "GAME OVER!! %s WINS!!!"
	DrawResult           ConsoleMessage = "GAME OVER!!! It's a DRAW!\n\rNo possible moves left. Nobody won, but nobody lost"
	InvalidColumn        ConsoleMessage = "Invalid column: %s. Column needs to be a number between %d and %d."
	InitialInputInvalid  ConsoleMessage = "Invalid input. Columns have to be numbers."
	SeeYaLater           ConsoleMessage = "Thanks for playing! Bye!"
	ExitTip              ConsoleMessage = "If you want to exit during the game, type 'exit' instead of a column"
)

func StartNewPrompt() (string, error) {
	validate := func(in string) error {
		if in != "y" && in != "n" {
			return errors.New("Invalid input")
		}
		return nil
	}
	p := promptui.Prompt{
		Label:    StartNew_Q,
		Validate: validate,
	}
	return p.Run()
}

func NextMovePrompt(player string) (string, error) {
	validate := func(in string) error {
		if in == "exit" {
			os.Exit(0)
		}
		_, err := strconv.Atoi(in)
		if err != nil {
			return errors.New(fmt.Sprintf(string(InvalidColumn), in, 0, 6))
		}
		return nil
	}
	p := promptui.Prompt{
		Label:    fmt.Sprintf(string(PlayerXMove), player),
		Validate: validate,
	}
	return p.Run()
}
