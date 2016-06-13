package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Player interface {
	Move(ConnectFourBoard) uint8
}

type RandomPlayer struct {
}

func (r RandomPlayer) Move(b ConnectFourBoard) uint8 {
	m := rand.Intn(Height)
	for ; !b.validMove(uint8(m)); m = rand.Intn(Height) {

	}
	return uint8(m)
}

type HumanPlayer struct {
	Reader *bufio.Reader
}

func (h HumanPlayer) Move(b ConnectFourBoard) uint8 {
	b.String()

	fmt.Print("Enter move: ")
	move, e := h.Reader.ReadString('\n')
	if e != nil {
		panic(e)
	}

	imove, e2 := strconv.Atoi(strings.TrimSpace(move))
	if e2 != nil {
		panic(e2)
	}
	return uint8(imove)
}
