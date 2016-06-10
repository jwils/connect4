package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	Width  = 7
	Height = 6

	Player1 = 0
	Player2 = 1
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
	reader *bufio.Reader
}

func (h HumanPlayer) Move(b ConnectFourBoard) uint8 {
	b.Print()

	fmt.Print("Enter move: ")
	move, e := h.reader.ReadString('\n')
	if e != nil {
		panic(e)
	}

	imove, e2 := strconv.Atoi(strings.TrimSpace(move))
	if e2 != nil {
		panic(e2)
	}
	return uint8(imove)
}

type ConnectFourBoard struct {
	playerBoard [2]uint64
	height      [7]uint8
	moves       [Width * Height]uint8
}

func (c *ConnectFourBoard) Print() {
	fmt.Printf("0 1 2 3 4 5 6\n")
	fmt.Printf("--------------\n")
	for i := 0; i < 6; i++ {
		for col := 0; col <= 6; col++ {
			if c.height[col] > uint8(5-i) {
				playerMove := -1
				if (c.playerBoard[0]>>uint(col*7+(5-i)))&1 > 0 {
					playerMove = 0
				} else if (c.playerBoard[1]>>uint(col*7+(5-i)))&1 > 0 {
					playerMove = 1
				}
				c.printSquare(playerMove)
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}

func (c *ConnectFourBoard) printSquare(player int) {
	switch player {
	case 0:
		fmt.Printf("\x1B")
		fmt.Printf("[31m")
		fmt.Print("0 ")
		fmt.Printf("\033")
		fmt.Printf("[0m")
		break
	case 1:
		fmt.Printf("\x1B")
		fmt.Printf("[33m")
		fmt.Print("0 ")
		fmt.Printf("\033")
		fmt.Printf("[0m")
		break
	default:
		fmt.Printf("ER")
	}
}

type ConnectFourGame struct {
	board ConnectFourBoard
	p1    Player
	p2    Player
}

//Copied from http://stackoverflow.com/questions/13327748/connect-four-bitboard
func (c *ConnectFourBoard) hasWon(player int) bool {
	p := c.playerBoard[player]
	horizontal := p & (p >> 7)
	vertical := p & (p >> 1)
	diagonal := p & (p >> 6)
	diagonal2 := p & (p >> 8)
	return ((horizontal & (horizontal >> 14)) | // check -
		(vertical & (vertical >> 2)) | // check |
		(diagonal & (diagonal >> 12)) | // check / diagonal
		(diagonal2 & (diagonal2 >> 16))) > 0 // check  \
}

func (c *ConnectFourBoard) validMove(move uint8) bool {
	return move >= 0 && move < Width && c.height[move] < Height
}

func (c *ConnectFourGame) GetPlayer(player int) Player {
	if player == Player1 {
		return c.p1
	} else if player == Player2 {
		return c.p2
	}
	panic("Unknown Player")
}

func (c *ConnectFourGame) Move(player int) {
	boardCopy := c.board
	move := c.GetPlayer(player).Move(boardCopy)
	if !c.board.validMove(move) {
		panic("Invalid move:" + string(move))
	}

	c.board.playerBoard[player] |= 1 << (7*move + c.board.height[move])
	c.board.height[move]++
}

func (c *ConnectFourGame) Play() {

	c.board.playerBoard[0] = 0
	c.board.playerBoard[1] = 0
	for i, _ := range c.board.height {
		c.board.height[i] = 0

	}
	for {

		c.Move(Player1)
		if c.board.hasWon(Player1) {
			fmt.Printf("Player 1 Wins")
			break
		}
		c.Move(Player2)
		if c.board.hasWon(Player2) {
			fmt.Printf("Player 2 Wins")
			break
		}
	}
}

func main() {
	p1 := RandomPlayer{}
	p2 := HumanPlayer{}
	p2.reader = bufio.NewReader(os.Stdin)
	game := ConnectFourGame{ConnectFourBoard{}, p1, p2}
	game.Play()
}
