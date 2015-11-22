package sudoku

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// ReadInput parses the input from an io.Reader to construct the initial
// sudoku puzzle. Returns an error in case of malformed input
func (grid *Grid) ReadInput(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	i := 0
	for i < 9 && scanner.Scan() {
		line := scanner.Text()
		cells := strings.Fields(line)
		if len(cells) != 9 {
			return fmt.Errorf("sudoku input: row %v incorrect length", i+1)
		}
		for j, cell := range cells {
			val, err := strCellValue(cell)
			if err != nil {
				log.Printf("sudoku input: error in r%vc%v of input", i+1, j+1)
				return err
			}
			(*grid)[i][j] = val
		}
		i++
	}
	if i < 9 {
		return fmt.Errorf("sudoku input: too few rows in input")
	}
	return nil
}

func strCellValue(cell string) (int, error) {
	cellErr := fmt.Errorf("sudoku input: unacceptable cell value %v", cell)
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
