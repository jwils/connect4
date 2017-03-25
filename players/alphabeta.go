package players

import (
	"github.com/jwils/connect4/game"
)

var (
	MAX_INT = int64(1) << 62
	MIN_INT = -1 * MAX_INT

	WIN_BASE  = int64(1)<<60 + 42
	LOSS_BASE = -1 * WIN_BASE

	MOVE_1_LOC = int64(1) << 7

	MOVE_ORDER = [...]uint8{3, 2, 4, 5, 1, 0, 6}
)

type AlphaBetaPlayer struct {
	EvalFuction func(game.ConnectFourBoard) float64
}

func (a AlphaBetaPlayer) Move(b game.ConnectFourBoard) uint8 {
	m := -1
	alpha := MIN_INT
	for i := uint8(0); i < game.Width; i++ {
		if b.ValidMove(MOVE_ORDER[i]) {
			b.Move(MOVE_ORDER[i])
			newV := a.AlphaBeta(b, 8, alpha, MAX_INT, false)
			b.UndoMove(MOVE_ORDER[i])
			if newV > alpha {
				m = int(MOVE_ORDER[i])
				alpha = newV
			}
		}
	}
	return uint8(m)
}

func (a *AlphaBetaPlayer) AlphaBeta(b game.ConnectFourBoard, depth int, alpha, beta int64, maxPlayer bool) int64 {
	if depth == 0 {
		return 0
	}
	if maxPlayer {
		if b.HasWon((b.MoveNumber() + 1) % 2) {
			// Want to lose as far as possible in the future
			return LOSS_BASE + int64(b.MoveNumber())
		}
		v := MIN_INT
		for i := uint8(0); i < game.Width; i++ {
			if b.ValidMove(MOVE_ORDER[i]) {
				b.Move(MOVE_ORDER[i])
				v = Max(a.AlphaBeta(b, depth-1, alpha, beta, false), v)
				b.UndoMove(MOVE_ORDER[i])
				alpha = Max(alpha, v)
				if beta <= alpha {
					break
				}
			}
		}
		return v
	} else {
		if b.HasWon((b.MoveNumber() + 1) % 2) {
			// Want to end as quickly as possible
			return WIN_BASE - int64(b.MoveNumber())
		}
		v := MAX_INT
		for i := uint8(0); i < game.Width; i++ {
			if b.ValidMove(MOVE_ORDER[i]) {
				b.Move(MOVE_ORDER[i])
				v = Min(a.AlphaBeta(b, depth-1, alpha, beta, true), v)
				b.UndoMove(MOVE_ORDER[i])
				v = Min(beta, v)
				if beta <= alpha {
					break
				}

			}
		}
		return v
	}
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
