package main

type GameOutcome int

const (
	NotEnd GameOutcome = iota
	WinnerX
	WinnerO
	Tie
)

func (o GameOutcome) String() string {
	switch o {
	case NotEnd:
		return "\nGame is still in progress!\n"
	case WinnerX:
		return "\nPlayer X is the Winner!!!\n"
	case WinnerO:
		return "\nPlayer O is the Winner!!!\n"
	case Tie:
		return "\nThe game ended in Tie!\n"
	default:
		panic("not mapped")
	}
}
