package players

import (
	"fmt"

	"github.com/jwils/connect4/game"
)

type NegamaxPlayer struct {
	EvalFuction func(game.ConnectFourBoard) float64
}

func (a NegamaxPlayer) Move(b game.ConnectFourBoard) uint8 {
	m := -1
	alpha := MIN_INT
	for i := uint8(0); i < game.Width; i++ {
		if b.ValidMove(MOVE_ORDER[i]) {
			b.Move(MOVE_ORDER[i])
			newV := -a.Negamax(b, 8, MIN_INT, -alpha)
			b.UndoMove(MOVE_ORDER[i])
			if newV > alpha {
				m = int(MOVE_ORDER[i])
				alpha = newV
			}
		}
	}
	if alpha < 0 {
		fmt.Printf("Score: %v\n", alpha)
		b.String()
	}
	return uint8(m)
}

func (a *NegamaxPlayer) Negamax(b game.ConnectFourBoard, depth int, alpha, beta int64) int64 {
	numValidMoves := 0
	for i := uint8(0); i < game.Width; i++ {
		numValidMoves++
		move := MOVE_ORDER[i]
		if b.ValidMove(move) && b.WinningMove(move) {
			return 42 - int64(b.MoveNumber())
		}
	}

	if b.MoveNumber() == game.Width*game.Height-1 {
		return 0
	}

	if depth == 0 && numValidMoves > 2 {
		// Todo
		return 0
	}

	for i := uint8(0); i < game.Width; i++ {
		move := MOVE_ORDER[i]
		if b.ValidMove(move) {
			b.Move(move)
			score := -a.Negamax(b, depth-1, -beta, -alpha)
			b.UndoMove(move)
			if score >= beta {
				return score
			}
			if score > alpha {
				alpha = score
			}
		}
	}
	return alpha
}
