package sudoku

const blockLen = 3
const gridLen = 9
const constraintTypes = 4
const numCells = gridLen * gridLen
const maxCols = constraintTypes * numCells
const maxRows = numCells * gridLen

const rowConstraintsOff = numCells
const colConstraintsOff = numCells * 2
const blkConstraintsOff = numCells * 3
