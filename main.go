package main

import (
	"bufio"
	"github.com/jwils/connect4/game"
	"os"
)

func main() {
	p1 := game.RandomPlayer{}
	p2 := game.HumanPlayer{}
	p2.Reader = bufio.NewReader(os.Stdin)
	game := game.NewGame(p1, p2)
	game.Play()
}
