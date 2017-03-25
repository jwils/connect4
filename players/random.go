package players

import (
	"math/rand"

	"github.com/jwils/connect4/game"
)

type RandomPlayer struct {
}

func (r RandomPlayer) Move(b game.ConnectFourBoard) uint8 {
	m := rand.Intn(game.Height)
	for ; !b.ValidMove(uint8(m)); m = rand.Intn(game.Height) {

	}
	return uint8(m)
}
