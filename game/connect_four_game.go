package game

import (
	"fmt"
)

const (
	Width  = 7
	Height = 6
)

const (
	Player1 = iota
	Player2
)

type ConnectFourBoard interface {
	MoveNumber() int
	WinningMove(move uint8) bool
	Move(move uint8)
	UndoMove(move uint8)
	ValidMove(move uint8) bool
	Reset()
	String()
	HasWon(player int) bool
}

func NewGame(p1, p2 Player) ConnectFourGame {
	return ConnectFourGame{&ConnectFourBitBoard{printer: ColorizedBoardPrinter{}, currentPlayer: Player1}, p1, p2}
}

type ConnectFourGame struct {
	board ConnectFourBoard
	p1    Player
	p2    Player
}

func (c *ConnectFourGame) GetPlayer(player int) Player {
	switch player {
	case Player1:
		return c.p1
	case Player2:
		return c.p2
	}
	panic("Unknown Player")
}

func (c *ConnectFourGame) Move(player int) {
	b := &c.board
	bCopy := b
	move := c.GetPlayer(player).Move(*bCopy)
	if !c.board.ValidMove(move) {
		panic(fmt.Sprintf("Invalid move: %v for player %v", move, player))
	}
	c.board.Move(move)
}

func (c *ConnectFourGame) Play() {
	c.board.Reset()
	for !c.board.HasWon((c.board.MoveNumber() + 1) % 2) {
		c.Move(c.board.MoveNumber() % 2)
	}
	c.board.String()
	fmt.Printf("Player %v has won", ((c.board.MoveNumber()+1)%2)+1)
}

func switchPlayer(p int) int {
	if p == Player1 {
		return Player2
	}
	return Player1
}
