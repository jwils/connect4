package main

import (
	"github.com/jwils/connect4/game"
	"github.com/jwils/connect4/players"
)

func main() {
	p1 := players.NegamaxPlayer{}
	p2 := players.AlphaBetaPlayer{}
	game := game.NewGame(p1, p2)
	game.Play()
}
