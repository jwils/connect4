package players

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/jwils/connect4/game"
)

type HumanPlayer struct {
	Reader *bufio.Reader
}

func (h HumanPlayer) Move(b game.ConnectFourBoard) uint8 {
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
		if imove < 0 || imove > game.Width-1 {
			fmt.Printf("\nMove out of range. Must be between 0 and %v\n", game.Width-1)
			fmt.Print("Enter move: ")
			continue

		}
		m = uint8(imove)
		if !b.ValidMove(m) {
			fmt.Print("Invalid move.")
			fmt.Print("Enter Move:")
			continue
		}
		break
	}
	return m
}
