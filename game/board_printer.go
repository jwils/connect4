package game

import "fmt"

type BoardPrinter interface {
	PrintHeader()
	PrintSquare(player int)
	BeginRow()
	EndRow()
	PrintFooter()
}

type ColorizedBoardPrinter struct {
}

func (c ColorizedBoardPrinter) PrintHeader() {
	fmt.Printf(" 0 1 2 3 4 5 6\n")
	fmt.Printf(" ______________\n")
}

func (c ColorizedBoardPrinter) changeColor(color string) {
	fmt.Printf("\x1B")
	fmt.Printf("[%sm", color)
}

func (c ColorizedBoardPrinter) clearColor() {
	fmt.Printf("\033")
	fmt.Printf("[0m")
}

func (c ColorizedBoardPrinter) PrintSquare(player int) {
	switch player {
	case 0:
		c.changeColor("31")
		fmt.Print("0 ")
		c.clearColor()
		break
	case 1:
		c.changeColor("33")
		fmt.Print("0 ")
		c.clearColor()
		break
	default:
		fmt.Printf("  ")
	}

}

func (c ColorizedBoardPrinter) BeginRow() {
	fmt.Print("|")
}

func (c ColorizedBoardPrinter) EndRow() {
	fmt.Print("|\n")
}

func (c ColorizedBoardPrinter) PrintFooter() {
	fmt.Printf(" ______________\n")
	fmt.Printf("\n")
}
