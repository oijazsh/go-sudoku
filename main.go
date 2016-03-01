package main

import (
	"fmt"
	"log"
	"os"

	"github.com/oijazsh/go-sudoku/sudoku"
)

func main() {
	var grid sudoku.Grid
	err := grid.Write(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	ok := grid.Solve()
	if !ok {
		log.Fatal("no solution exists for given sudoku puzzle")
	}
	fmt.Print(grid.String())
}
