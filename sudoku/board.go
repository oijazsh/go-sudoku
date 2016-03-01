package sudoku

import (
	"github.com/oijazsh/go-sudoku/dlx"
	"log"
)

// Grid represents the actual sudoku puzzle
type Grid [gridLen][gridLen]int

func (g *Grid) togglePossibility(poss int) bool {
	row := possToRow(poss)
	col := possToCol(poss)
	val := possToValue(poss)
	if g[row][col] == 0 {
		g[row][col] = val
		return true
	}
	g[row][col] = 0
	return false
}

func possToRow(possibility int) int {
	return possibility / numCells
}

func possToCol(possibility int) int {
	return (possibility % numCells) / gridLen
}

func possToValue(possibility int) int {
	return possibility%gridLen + 1
}

// possibility returns the row number of the possibility-constraint matrix
// representing the given sudoku row, column and value
func possibility(row, col, value int) int {
	return row*numCells + col*gridLen + value
}

func genSparseMatrix(root *dlx.Node) {
	var headers [maxCols]*dlx.Node
	row := make([]*dlx.Node, 4)

	for i := range headers {
		headers[i] = dlx.AddHeader(root)
	}

	for r := 0; r < gridLen; r++ {
		for c := 0; c < gridLen; c++ {
			for v := 0; v < gridLen; v++ {
				poss := possibility(r, c, v)
				row[0] = dlx.AddNode(poss, headers[r*gridLen+c])
				row[1] = dlx.AddNode(poss, headers[rowConstraintsOff+r*gridLen+v])
				row[2] = dlx.AddNode(poss, headers[colConstraintsOff+c*gridLen+v])
				row[3] = dlx.AddNode(poss, headers[blkConstraintsOff+
					(r/blockLen+c/blockLen*blockLen)*gridLen+v])

				err := dlx.BuildRow(row)
				if err != nil {
					log.Fatal("sudoku: could not attach a row")
				}
			}
		}
	}
}

// Solve solves the given sudoku puzzle and returns whether it was successful
func (g *Grid) Solve() bool {
	root := dlx.NewRoot()
	solution := make(chan int, numCells)
	solved := false

	genSparseMatrix(root)
	for r, row := range g {
		for c, v := range row {
			if v != 0 {
				n := dlx.Find(possibility(r, c, v-1), root)
				if n == nil {
					return false
				}
				dlx.Cover(n)
			}
		}
	}

	func() {
		defer close(solution)
		solved = dlx.Solve(root, solution)
	}()

	for poss := range solution {
		g.togglePossibility(poss)
	}
	return solved
}
