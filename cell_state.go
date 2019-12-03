package main

type CellState int

const (
	E CellState = iota // Empty cell
	X                  // Cell marked with X
	O                  // Cell marked with O
)

func toString(s CellState, i string) string {
	switch s {
	case E:
		return i
	case X:
		return "X"
	case O:
		return "O"
	default:
		panic("not mapped")
	}
}

func (s CellState) String() string {
	switch s {
	case E:
		return " "
	case X:
		return "X"
	case O:
		return "O"
	default:
		panic("not mapped")
	}
}
