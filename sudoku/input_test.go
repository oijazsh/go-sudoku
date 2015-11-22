package sudoku

import (
	"strconv"
	"strings"
	"testing"
)

func TestCellValueGood(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		v, err := strCellValue(s)
		if err != nil {
			t.Error("error received for valid input")
		}
		if v != i {
			t.Errorf("incorrect return value for %v input", s)
		}
	}
	// Alse test for "_"
	s := "_"
	v, err := strCellValue(s)
	if err != nil {
		t.Error("error received for valid input")
	}
	if v != 0 {
		t.Errorf("incorrect return value for %v input", s)
	}
}

func TestCellValueBad(t *testing.T) {
	strs := []string{"01", "a", "-5", " "}
	for _, s := range strs {
		_, err := strCellValue(s)
		if err == nil {
			t.Errorf("expected error for input \"%v\" not received", s)
		}
	}
}

func TestSimpleInput(t *testing.T) {
	var grid Grid
	s := `1 _ 3 _ _ 6 _ 8 _
		_ 5 _ _ 8 _ 1 2 _
		7 _ 9 1 _ 3 _ 5 6
		_ 3 _ _ 6 7 _ 9 _
		5 _ 7 8 _ _ _ 3 _
		8 _ 1 _ 3 _ 5 _ 7
		_ 4 _ _ 7 8 _ 1 _
		6 _ 8 _ _ 2 _ 4 _
		_ 1 2 _ 4 5 _ 7 8`
	reader := strings.NewReader(s)
	err := grid.ReadInput(reader)
	if err != nil {
		t.Fail()
	}
}

func TestFewRows(t *testing.T) {
	var grid Grid
	s := `1 _ 3 _ _ 6 _ 8 _
		_ 5 _ _ 8 _ 1 2 _`
	reader := strings.NewReader(s)
	err := grid.ReadInput(reader)
	if err == nil {
		t.Fail()
	}
}

func TestFewColumns(t *testing.T) {
	var grid Grid
	s := `1 _ 3 8 _
		_ 5 _ _ 8
		7 _ 9 1 _
		_ 3 _ _ 6
		5 _ 7 8 _
		8 _ 1 _ 3
		_ 4 _ _ 7
		6 _ 8 _ _
		_ 1 2 _ 4`
	reader := strings.NewReader(s)
	err := grid.ReadInput(reader)
	if err == nil {
		t.Fail()
	}
}

func TestInvalidCell(t *testing.T) {
	var grid Grid
	s := `1 _ 3 _ _ 6 _  a _`
	reader := strings.NewReader(s)
	err := grid.ReadInput(reader)
	if err == nil {
		t.Fail()
	}
}
