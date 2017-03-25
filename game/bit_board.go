package game

type ConnectFourBitBoard struct {
	playerBoard   [2]uint64
	height        [7]uint8
	moveNumber    int
	moves         [Width * Height]uint8
	printer       BoardPrinter
	currentPlayer int
}

func (c *ConnectFourBitBoard) String() {
	c.printer.PrintHeader()
	for i := 0; i < 6; i++ {
		c.printer.BeginRow()
		for col := 0; col <= 6; col++ {
			if c.height[col] > uint8(5-i) {
				player0 := (c.playerBoard[0] >> uint(col*7+(5-i))) & 1
				player1 := (c.playerBoard[1] >> uint(col*7+(5-i))) & 1
				playerMove := player0 + player1<<1
				c.printer.PrintSquare(int(playerMove) - 1)
			} else {
				c.printer.PrintSquare(-1)
			}
		}
		c.printer.EndRow()
	}
	c.printer.PrintFooter()
}

func (c *ConnectFourBitBoard) MoveNumber() int {
	return c.moveNumber
}

func (c *ConnectFourBitBoard) WinningMove(move uint8) bool {
	p := c.playerBoard[c.currentPlayer]
	p |= 1 << (7*move + c.height[move])

	horizontal := p & (p >> 7)
	vertical := p & (p >> 1)
	diagonal := p & (p >> 6)
	diagonal2 := p & (p >> 8)
	return ((horizontal & (horizontal >> 14)) | // check -
		(vertical & (vertical >> 2)) | // check |
		(diagonal & (diagonal >> 12)) | // check / diagonal
		(diagonal2 & (diagonal2 >> 16))) > 0 // check  \
}

//Copied from http://stackoverflow.com/questions/13327748/connect-four-bitboard
func (c *ConnectFourBitBoard) HasWon(player int) bool {
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

func (c *ConnectFourBitBoard) ValidMove(move uint8) bool {
	return move >= 0 && move < Width && c.height[move] < Height
}

func (b *ConnectFourBitBoard) Reset() {
	b.playerBoard[Player1] = 0
	b.playerBoard[Player2] = 0
	for i, _ := range b.height {
		b.height[i] = 0
	}
}

func (b *ConnectFourBitBoard) Move(move uint8) {
	b.playerBoard[b.currentPlayer] |= 1 << (7*move + b.height[move])
	b.height[move]++
	b.moves[b.moveNumber] = move
	b.moveNumber++
	b.currentPlayer = switchPlayer(b.currentPlayer)
}

func (b *ConnectFourBitBoard) UndoMove(move uint8) {
	b.currentPlayer = switchPlayer(b.currentPlayer)
	b.moveNumber--
	b.moves[b.moveNumber] = 0
	b.height[move]--
	b.playerBoard[b.currentPlayer] ^= 1 << (7*move + b.height[move])
}

func reversecols(num uint64) uint64 {
	col1 := num & 0x8F
	col2 := num & (0x8F << 7)
	col3 := num & (0x8F << 14)
	col4 := num & (0x8F << 21)
	col5 := num & (0x8F << 28)
	col6 := num & (0x8F << 35)
	col7 := num & (0x8F << 42)

	return (col1 << 42) | (col2 << 35) | (col3 << 28) | (col4 << 21) | (col5 << 14) | (col6 << 7) | col7
}
