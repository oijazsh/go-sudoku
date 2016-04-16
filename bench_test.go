package main

import (
	"strings"
	"testing"

	"github.com/oijazsh/go-sudoku/sudoku"
)

func BenchmarkEasy(b *testing.B) {
	var grid sudoku.Grid
	s := "4 _ _ _ _ _ _ 5 _\n" +
		"_ _ 3 _ _ 7 4 2 _\n" +
		"_ 5 7 8 _ _ _ _ _\n" +
		"3 4 _ 5 7 _ 2 _ 6\n" +
		"_ _ 2 4 1 9 3 _ _\n" +
		"8 _ 5 _ 6 2 _ 7 4\n" +
		"_ _ _ _ _ 4 8 6 _\n" +
		"_ 7 4 1 _ _ 9 _ _\n" +
		"_ 2 _ _ _ _ _ _ 7\n"
	reader := strings.NewReader(s)
	grid.Write(reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := grid
		g.Solve()
	}
}

func BenchmarkMedium(b *testing.B) {
	var grid sudoku.Grid
	s := "_ 9 2 _ 8 _ _ _ _\n" +
		"3 _ 5 _ _ _ _ _ 8\n" +
		"6 _ _ 5 _ _ 9 3 _\n" +
		"_ _ _ _ 1 _ _ 2 _\n" +
		"_ _ 9 3 2 7 6 _ _\n" +
		"_ 4 _ _ 6 _ _ _ _\n" +
		"_ 3 8 _ _ 4 _ _ 2\n" +
		"9 _ _ _ _ _ 7 _ 5\n" +
		"_ _ _ _ 7 _ 4 8 _\n"
	reader := strings.NewReader(s)
	grid.Write(reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := grid
		g.Solve()
	}
}

func BenchmarkHard(b *testing.B) {
	var grid sudoku.Grid
	s := "2 _ _ 3 7 _ 6 _ _\n" +
		"7 _ _ _ _ 5 _ 9 _\n" +
		"4 _ _ _ _ 1 _ 3 _\n" +
		"_ 9 _ _ _ _ _ 7 _\n" +
		"3 _ _ _ 4 _ _ _ 1\n" +
		"_ 6 _ _ _ _ _ 2 _\n" +
		"_ 4 _ 6 _ _ _ _ 9\n" +
		"_ 2 _ 1 _ _ _ _ 7\n" +
		"_ _ 8 _ 2 7 _ _ 4\n"
	reader := strings.NewReader(s)
	grid.Write(reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := grid
		g.Solve()
	}
}

func BenchmarkEvil(b *testing.B) {
	var grid sudoku.Grid
	s := "9 _ _ _ _ 8 4 7 _\n" +
		"_ 1 _ _ _ _ _ _ 6\n" +
		"_ _ 5 _ _ _ _ 2 _\n" +
		"_ _ _ 9 _ _ 2 4 1\n" +
		"_ _ _ 3 _ 4 _ _ _\n" +
		"8 4 1 _ _ 7 _ _ _\n" +
		"_ 5 _ _ _ _ 3 _ _\n" +
		"1 _ _ _ _ _ _ 5 _\n" +
		"_ 3 6 5 _ _ _ _ 9\n"
	reader := strings.NewReader(s)
	grid.Write(reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := grid
		g.Solve()
	}
}
