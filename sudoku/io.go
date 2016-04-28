package sudoku

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func (b *Board) setSize(blockCols, blockRows int) {
	b.blockCols = blockCols
	b.blockRows = blockRows

	nBlock := b.sideLen()
	b.grid = make([][]int, nBlock)

	// Allocate all rows as a single array
	cells := make([]int, nBlock*nBlock)
	for i := range b.grid {
		b.grid[i], cells = cells[:nBlock], cells[nBlock:]
	}
}

// Build parses the input from an io.Reader to construct the initial sudoku
// puzzle. Returns an error in case of malformed input.
func Build(r io.Reader, blockCols int, blockRows int) (Board, error) {
	scanner := bufio.NewScanner(r)
	board := Board{}
	board.setSize(blockCols, blockRows)

	sideLen := board.sideLen()

	i := 0
	for i < sideLen && scanner.Scan() {
		line := scanner.Text()
		cells := strings.Fields(line)

		if len(cells) != sideLen {
			return board, fmt.Errorf("sudoku: Input row %v incorrect length", i+1)
		}

		for j, cell := range cells {
			val, err := strCellValue(cell)
			if err != nil {
				return board, err
			}
			board.grid[i][j] = val
		}
		i++
	}
	return board, nil
}

// String returns a string representation of the sudoku puzzle
func (b *Board) String() string {
	var buff bytes.Buffer
	sideLen := b.sideLen()

	for r := 0; r < sideLen; r++ {
		for c := 0; c < sideLen; c++ {
			buff.WriteString(strconv.Itoa(b.grid[r][c]))
			if c != sideLen-1 {
				buff.WriteString(" ")
			}
		}
		buff.WriteString("\n")
	}
	return buff.String()
}

func strCellValue(cell string) (int, error) {
	cellErr := fmt.Errorf("sudoku: Unacceptable cell value %v", cell)
	if len(cell) > 1 {
		return 0, cellErr
	}
	if cell == "_" {
		return 0, nil
	} else if cell > "0" || cell < "9" {
		val, err := strconv.Atoi(cell)
		return val, err
	} else {
		return 0, cellErr
	}
}
