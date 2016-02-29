# A Sudoku solver written in Go

go-sudoku is a Sudoku solver which utilizes a [Dancing Links](https://en.wikipedia.org/wiki/Dancing_Links) technique.

## Usage

The command reads a puzzle from stdin and outputs the solution to stdout. Each line of the standard input is translated into one line of the Sudoku puzzle. Empty cells are denoted by `_`.

Example input:

```
2 _ _ 3 7 _ 6 _ _
7 _ _ _ _ 5 _ 9 _
4 _ _ _ _ 1 _ 3 _
_ 9 _ _ _ _ _ 7 _
3 _ _ _ 4 _ _ _ 1
_ 6 _ _ _ _ _ 2 _
_ 4 _ 6 _ _ _ _ 9
_ 2 _ 1 _ _ _ _ 7
_ _ 8 _ 2 7 _ _ 4
```
Output:
```
2 1 9 3 7 8 6 4 5
7 8 3 4 6 5 1 9 2
4 5 6 2 9 1 7 3 8
8 9 1 5 3 2 4 7 6
3 7 2 8 4 6 9 5 1
5 6 4 7 1 9 8 2 3
1 4 7 6 5 3 2 8 9
9 2 5 1 8 4 3 6 7
6 3 8 9 2 7 5 1 4
```
