package main

import (
	"fmt"
	"log"
	"os"

	"github.com/oijazsh/go-sudoku/sudoku"
)

func main() {
	var grid sudoku.Grid
	err := grid.ReadInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	err = sudoku.Solve(&grid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(grid.String())
}
