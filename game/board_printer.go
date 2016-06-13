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

func (c ColorizedBoardPrinter) PrintSquare(player int) {
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
