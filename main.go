package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameState struct {
	slots [9]int
	turn  int
}

var (
	gameState    = GameState{}
	reader       = bufio.NewReader(os.Stdin)
	victoryLines = [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
		{1, 5, 9},
		{3, 5, 7},
	}
)

func main() {
	gameState.turn = 1

	for {
		fmt.Printf("Turn: Player %d\n", gameState.turn)
		printBoard()
		playTurn()
		if checkVictoryOrDraw() {
			break
		}
	}
}

func printBoard() {
	fmt.Println("_______")

	for i, slot := range gameState.slots {
		fmt.Print("|")
		switch slot {
		case 0:
			fmt.Print(" ")
		case 1:
			fmt.Print("X")
		case 2:
			fmt.Print("O")
		default:
			panic("Invalid slot state")
		}
		if i != 0 && (i+1)%3 == 0 {
			fmt.Print("|\n")
		}
	}

	fmt.Println("‾‾‾‾‾‾‾")
}

func playTurn() {
	var answer int
	var input string
	var err error

	fmt.Print("Choose a slot [1-9]: ")

	for {
		input, _ = reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		answer, err = strconv.Atoi(input)
		if err != nil || answer < 1 || answer > len(gameState.slots) || gameState.slots[answer-1] != 0 {
			fmt.Print("Invalid number, try again: ")
			continue
		}
		gameState.slots[answer-1] = gameState.turn
		break
	}

	if gameState.turn == 1 {
		gameState.turn = 2
	} else {
		gameState.turn = 1
	}
}

func checkVictoryOrDraw() bool {
	players := [2]int{1, 2}
	for _, line := range victoryLines {
		for _, player := range players {
			if gameState.slots[line[0]-1] == player && gameState.slots[line[1]-1] == player && gameState.slots[line[2]-1] == player {
				printBoard()
				fmt.Printf("Winner is Player %d!\n", player)
				return true
			}
		}
	}
	var emptyFound bool
	for _, slot := range gameState.slots {
		if slot == 0 {
			emptyFound = true
			break
		}
	}
	if !emptyFound {
		printBoard()
		fmt.Println("It is a draw!")
		return true
	}
	return false
}
