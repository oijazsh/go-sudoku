package sudoku

import (
	"strings"
	"testing"
)

func BenchmarkEasy(b *testing.B) {
	var grid Grid
	s := `4 _ _ _ _ _ _ 5 _
  _ _ 3 _ _ 7 4 2 _
  _ 5 7 8 _ _ _ _ _
  3 4 _ 5 7 _ 2 _ 6
  _ _ 2 4 1 9 3 _ _
  8 _ 5 _ 6 2 _ 7 4
  _ _ _ _ _ 4 8 6 _
  _ 7 4 1 _ _ 9 _ _
  _ 2 _ _ _ _ _ _ 7`
	reader := strings.NewReader(s)
	grid.ReadInput(reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := grid
		Solve(&g)
	}
}

func BenchmarkMedium(b *testing.B) {
	var grid Grid
	s := `_ _ _ 4 1 7 _ 6 3
  _ 3 _ _ _ 8 7 _ _
  _ 4 _ _ _ _ _ 2 8
  _ _ _ 2 _ _ _ 8 _
  8 6 _ _ 3 _ _ 9 2
  _ 9 _ _ _ 5 _ _ _
  4 7 _ _ _ _ _ 5 _
  _ _ 5 8 _ _ _ 3 _
  3 8 _ 5 9 2 _ _ _`
	reader := strings.NewReader(s)
	grid.ReadInput(reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := grid
		Solve(&g)
	}
}

func BenchmarkHard(b *testing.B) {
	var grid Grid
	s := `_ _ _ _ 2 1 _ _ _
  _ _ 8 4 _ _ _ _ 1
  3 1 _ _ _ 8 7 _ _
  8 9 7 5 _ _ _ _ _
  _ _ 6 _ _ _ 1 _ _
  _ _ _ _ _ 4 8 9 7
  _ _ 5 1 _ _ _ 4 3
  1 _ _ _ _ 5 9 _ _
  _ _ _ 2 3 _ _ _ _`
	reader := strings.NewReader(s)
	grid.ReadInput(reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := grid
		Solve(&g)
	}
}

func BenchmarkEvil(b *testing.B) {
	var grid Grid
	s := `_ _ _ _ 8 _ 7 _ _
  8 _ _ _ _ 3 _ _ 2
  4 6 _ _ _ _ 3 _ _
  _ _ _ _ _ 5 2 1 _
  _ 7 _ _ _ _ _ 5 _
  _ 9 2 7 _ _ _ _ _
  _ _ 5 _ _ _ _ 3 6
  6 _ _ 4 _ _ _ _ 1
  _ _ 1 _ 6 _ _ _ _`
	reader := strings.NewReader(s)
	grid.ReadInput(reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := grid
		Solve(&g)
	}
}
