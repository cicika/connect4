package main

import (
	"connect4/app"

	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var initialGameMoves string

func main() {
	flag.Parse()
	valid, columns := parseColumnsArray(initialGameMoves)
	var game *app.Game
	fmt.Println(app.Hello)
	fmt.Println(app.ExitTip)
	for {
		fmt.Print("> ")
		if !valid && len(columns) > 0 {
			fmt.Println(app.InitialInputInvalid)
		}
		blank := len(columns) == 0
		startNew := len(columns) > 0
		if blank {
			res, _ := app.StartNewPrompt()
			startNew = res == "y"
		}
		if !startNew {
			fmt.Println(app.SeeYaLater)
			os.Exit(0)
		}
		if startNew || !blank {
			fmt.Println(app.Default_GameStart)
			game = app.StartGame("", "", 0, 0)
		}

		for !game.GameOver {
			if !blank {
				fmt.Println(app.RunningInitialMoves)
				_, cnt := game.MultipleMove(columns)
				if cnt > 0 {
					fmt.Printf(string(app.MultipleMovesFailure)+"\n\r", columns[len(columns)-cnt:cnt])
				}
				if game.GameOver {
					game.End()
					columns = []int{}
					initialGameMoves = ""
					blank = true
					break
				}
			}
			res, err := app.NextMovePrompt(game.NextMoveOn)
			if err != nil {
				fmt.Println(err)
			} else {
				c, _ := strconv.Atoi(res)
				game.SingleMove(c)
				if game.GameOver {
					game.End()
					break
				}
			}
		}
	}
}

func init() {
	flag.StringVar(&initialGameMoves, "columns", "", "initial input")
}

func parseColumnsArray(in string) (bool, []int) {
	var out []int
	strArr := strings.Split(in, " ")
	if len(strArr) > 0 {
		for _, s := range strArr {
			i, err := strconv.Atoi(s)
			if err != nil {
				return false, out
			}
			out = append(out, i)
		}
		return true, out
	}
	return true, out
}
