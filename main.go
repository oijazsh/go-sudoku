package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oijazsh/go-sudoku/sudoku"
)

func main() {
	var rank bool
	flag.BoolVar(&rank, "rank", false,
		"Rank the puzzle in addition to solving it. False by default.")

	flag.Parse()

	var grid sudoku.Grid
	err := grid.Write(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	ok := false
	if rank {
		ok = sudoku.SolveAndRank(&grid)
	} else {
		ok = sudoku.Solve(&grid)
	}
	if !ok {
		log.Fatal("no solution exists for given sudoku puzzle")
	}
	fmt.Print(grid.String())
}
