package main

type Symbol int

const (
	symbolX Symbol = iota
	symbolO
)

func (s Symbol) String() string {
	switch s {
	case symbolX:
		return "X"
	case symbolO:
		return "O"
	default:
		panic("unreachable")
	}
}

func (s Symbol) State() CellState {
	switch s {
	case symbolX:
		return X
	case symbolO:
		return O
	default:
		panic("unreachable")
	}
}
