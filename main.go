package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oijazsh/go-sudoku/sudoku"
)

var (
	blockCols = flag.Uint("bc", 3, "Number of columns in each block(region) of the puzzle.")
	blockRows = flag.Uint("br", 3, "Number of rows in each block(region) of the puzzle.")
)

func main() {
	flag.Parse()
	board, err := sudoku.Build(os.Stdin, int(*blockCols), int(*blockRows))
	if err != nil {
		log.Fatal(err)
	}

	ok := board.Solve()
	if !ok {
		log.Fatal("No solution exists for given sudoku puzzle")
	}
	fmt.Print(&board)
}
