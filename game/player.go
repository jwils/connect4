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

	var m uint8
	fmt.Print("Enter move: ")
	for move, e := h.Reader.ReadString('\n'); true; move, e = h.Reader.ReadString('\n') {
		if e != nil {
			fmt.Print("\nError reading input. Please try again.\n")
			fmt.Print("Enter move: ")
			continue
		}

		imove, e2 := strconv.Atoi(strings.TrimSpace(move))
		if e2 != nil {
			fmt.Print("\nMove could not be parsed into an int. Please try agian.\n")
			fmt.Print("Enter move: ")
			continue
		}
		if imove < 0 || imove > Width - 1 {
			fmt.Printf("\nMove out of range. Must be between 0 and %v\n", Width - 1)
			fmt.Print("Enter move: ")
			continue

		}
		m = uint8(imove)
		break
	}
	return m
}
