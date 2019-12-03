package main

import (
	"fmt"
	"strings"
)

func main() {
	printWelcome() // print welcome message

	var play = true // manages multiple games

	for play == true {

		var gb = new(GameBoard)                            // initialize new game board
		gb.Cells = [9]CellState{E, E, E, E, E, E, E, E, E} // set cells of game board to empty

		playerX := choosePlayer(symbolX) // initialize player X
		playerO := choosePlayer(symbolO) // initialize player O

		gameOutcome := NotEnd       // game output initialized to NotEnd
		var currentPlayer = playerO // current player initialized to player O

		// the game continue until it is not ended
		for gameOutcome == NotEnd {
			currentPlayer = managePlayersTurn(currentPlayer, playerO, playerX) // choose of current player
			printGameBoard(gb)                                                 // print intermediate game board
			assignCellValue(currentPlayer, gb)                                 // assigne the symbol to the chosen cell
			gameOutcome = evaluateGameOutcome(gb)                              // evaluate the game output
		}

		fmt.Println(gameOutcome.String()) // print the game outcome
		printFinalGameBoard(gb)           // print the final game board

		play = evaluateNewGame() // evaluate if the players want to play again
	}
}

func choosePlayer(s Symbol) Player {
	fmt.Print("Insert the name of the player [" + s.String() + "]: ")
	var namePlayer string
	fmt.Scanln(&namePlayer)
	fmt.Println("Hello " + namePlayer + "! You have symbol: " + s.String() + "\n")
	return Player{namePlayer, s}

}

func managePlayersTurn(currentPlayer Player, playerO Player, playerX Player) Player {
	if currentPlayer == playerO {
		return playerX
	} else {
		return playerO
	}
}

func evaluateGameOutcome(gb *GameBoard) GameOutcome {
	var x = (gb.Cells[0] == X && gb.Cells[1] == X && gb.Cells[2] == X) || // evaluate rows for X
		(gb.Cells[3] == X && gb.Cells[4] == X && gb.Cells[5] == X) ||
		(gb.Cells[6] == X && gb.Cells[7] == X && gb.Cells[8] == X) ||

		(gb.Cells[0] == X && gb.Cells[3] == X && gb.Cells[6] == X) || // evaluate columns for X
		(gb.Cells[1] == X && gb.Cells[4] == X && gb.Cells[7] == X) ||
		(gb.Cells[2] == X && gb.Cells[5] == X && gb.Cells[8] == X) ||

		(gb.Cells[0] == X && gb.Cells[4] == X && gb.Cells[8] == X) || // evaluate diagonals for X
		(gb.Cells[2] == X && gb.Cells[4] == X && gb.Cells[6] == X)

	var o = (gb.Cells[0] == O && gb.Cells[1] == O && gb.Cells[2] == O) || // evaluate rows for O
		(gb.Cells[3] == O && gb.Cells[4] == O && gb.Cells[5] == O) ||
		(gb.Cells[6] == O && gb.Cells[7] == O && gb.Cells[8] == O) ||

		(gb.Cells[0] == O && gb.Cells[3] == O && gb.Cells[6] == O) || // evaluate columns for O
		(gb.Cells[1] == O && gb.Cells[4] == O && gb.Cells[7] == O) ||
		(gb.Cells[2] == O && gb.Cells[5] == O && gb.Cells[8] == O) ||

		(gb.Cells[0] == O && gb.Cells[4] == O && gb.Cells[8] == O) || // evaluate diagonals for O
		(gb.Cells[2] == O && gb.Cells[4] == O && gb.Cells[6] == O)

	var emptyCells = gb.Cells[0] == E || gb.Cells[1] == E || gb.Cells[2] == E || // evaluate empty cells
		gb.Cells[3] == E || gb.Cells[4] == E || gb.Cells[5] == E ||
		gb.Cells[6] == E || gb.Cells[7] == E || gb.Cells[8] == E

	switch {
	case x && !o:
		return WinnerX
	case o && !x:
		return WinnerO
	case !emptyCells:
		return Tie
	default:
		return NotEnd
	}
}

func assignCellValue(p Player, gb *GameBoard) {
	fmt.Print("\n" + p.name + " [" + p.symbol.String() + "] " + "it's your turn, write the number of the cell [1-9]: ")
	var input int
	fmt.Scanln(&input)
	if input > 0 && input < 10 && gb.Cells[input-1] == E {
		gb.Cells[input-1] = p.symbol.State()
	} else {
		fmt.Println("Not valid cell number!")
		assignCellValue(p, gb)
	}
}

func evaluateNewGame() bool {
	fmt.Print("Do you want to play again? [ y / n ]: ")
	var input string
	fmt.Scanln(&input)

	return strings.EqualFold(input, "y")
}

func printGameBoard(gb *GameBoard) {
	fmt.Println()
	fmt.Printf("%v│%v│%v\n", toString(gb.Cells[0], "1"), toString(gb.Cells[1], "2"), toString(gb.Cells[2], "3"))
	fmt.Println("─┼─┼─")
	fmt.Printf("%v│%v│%v\n", toString(gb.Cells[3], "4"), toString(gb.Cells[4], "5"), toString(gb.Cells[5], "6"))
	fmt.Println("─┼─┼─")
	fmt.Printf("%v│%v│%v\n", toString(gb.Cells[6], "7"), toString(gb.Cells[7], "8"), toString(gb.Cells[8], "9"))
	fmt.Println("─┼─┼─")
}

func printFinalGameBoard(gb *GameBoard) {
	fmt.Printf("%v│%v│%v\n", gb.Cells[0].String(), gb.Cells[1].String(), gb.Cells[2].String())
	fmt.Println("─┼─┼─")
	fmt.Printf("%v│%v│%v\n", gb.Cells[3].String(), gb.Cells[4].String(), gb.Cells[5].String())
	fmt.Println("─┼─┼─")
	fmt.Printf("%v│%v│%v\n", gb.Cells[6].String(), gb.Cells[7].String(), gb.Cells[8].String())
	fmt.Println("─┼─┼─")
	fmt.Println()
}
func printWelcome() {
	fmt.Println()
	fmt.Println(" ------------------------")
	fmt.Println("| WELCOME TO TIC-TAC-TOE |")
	fmt.Println(" ------------------------")
	fmt.Println()
}
