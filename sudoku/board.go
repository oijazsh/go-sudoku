package sudoku

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/oijazsh/go-sudoku/dlx"
)

// Grid represents the actual sudoku puzzle
type Grid [gridLen][gridLen]int

func (g *Grid) readPossibility(poss int) {
	row := poss / numCells
	col := (poss % numCells) / gridLen
	g[row][col] = poss%gridLen + 1
}

func (g *Grid) String() string {
	var buff bytes.Buffer
	for r := 0; r < gridLen; r++ {
		for c := 0; c < gridLen; c++ {
			buff.WriteString(strconv.Itoa(g[r][c]))
			buff.WriteString(" ")
		}
		buff.WriteString("\n")
	}
	return buff.String()
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
				row[3] = dlx.AddNode(poss, headers[blkConstraintsOff+(r/blockLen+c/blockLen*blockLen)*gridLen+v])

				dlx.BuildRow(row)
			}
		}
	}
}

func Solve(g *Grid) error {
	root := dlx.NewRoot()
	solution := make([]int, 0, numCells)
	genSparseMatrix(root)
	for r, row := range g {
		for c, v := range row {
			if v != 0 {
				n := dlx.Find(possibility(r, c, v-1), root)
				if n == nil {
					return fmt.Errorf("sudoku: no valid solution to given puzzle")
				}
				dlx.Cover(n)
			}
		}
	}
	dlx.Solve(root, &solution)
	for _, sol := range solution {
		g.readPossibility(sol)
	}
	return nil
}
