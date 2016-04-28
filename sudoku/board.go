package sudoku

import (
	"log"

	"github.com/oijazsh/go-sudoku/dlx"
)

// Board contains all the information regarding the sudoku puzzle
type Board struct {
	grid      [][]int
	blockRows int
	blockCols int
}

func (b *Board) sideLen() int {
	return b.blockRows * b.blockCols
}

func (b *Board) togglePossibility(poss int) bool {
	row := b.possToRow(poss)
	col := b.possToCol(poss)
	val := b.possToValue(poss)
	if b.grid[row][col] == 0 {
		b.grid[row][col] = val
		return true
	}
	b.grid[row][col] = 0
	return false
}

func (b *Board) possToRow(possibility int) int {
	sideLen := b.sideLen()
	return possibility / (sideLen * sideLen)
}

func (b *Board) possToCol(possibility int) int {
	sideLen := b.sideLen()
	return (possibility % (sideLen * sideLen)) / sideLen
}

func (b *Board) possToValue(possibility int) int {
	sideLen := b.sideLen()
	return possibility%sideLen + 1
}

// possibility returns the row number of the possibility-constraint matrix
// representing the given sudoku row, column and value
func (b *Board) possibility(row, col, value int) int {
	sideLen := b.sideLen()
	return row*sideLen*sideLen + col*sideLen + value
}

func (b *Board) genSparseMatrix(root *dlx.Node) {
	sideLen := b.sideLen()
	nCells := sideLen * sideLen

	nConstraints := nCells * 4
	headers := make([]*dlx.Node, nConstraints)
	row := make([]*dlx.Node, 4)

	rowOffset := sideLen * sideLen
	colOffset := rowOffset * 2
	blkOffset := rowOffset * 3

	for i := range headers {
		headers[i] = dlx.AddHeader(root)
	}

	for r := 0; r < sideLen; r++ {
		for c := 0; c < sideLen; c++ {
			for v := 0; v < sideLen; v++ {
				poss := b.possibility(r, c, v)
				row[0] = dlx.AddNode(poss, headers[r*sideLen+c])
				row[1] = dlx.AddNode(poss, headers[rowOffset+r*sideLen+v])
				row[2] = dlx.AddNode(poss, headers[colOffset+c*sideLen+v])
				row[3] = dlx.AddNode(poss, headers[blkOffset+
					(r/b.blockRows*b.blockCols+c/b.blockCols)*sideLen+v])

				err := dlx.BuildRow(row)
				if err != nil {
					log.Fatal("sudoku: could not attach a row")
				}
			}
		}
	}
}

// Solve solves the given sudoku puzzle and returns whether it was successful
func (b *Board) Solve() bool {
	root := dlx.NewRoot()
	sideLen := b.sideLen()
	solution := make(chan int, sideLen*sideLen)
	solved := false

	b.genSparseMatrix(root)
	for r, row := range b.grid {
		for c, v := range row {
			if v != 0 {
				n := dlx.Find(b.possibility(r, c, v-1), root)
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
		b.togglePossibility(poss)
	}
	return solved
}
