package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ROUND = 9
	X     = "X"
	O     = "O"
)

type Board struct {
	matrix *[3][3]string
}

func (b Board) String() string {
	var output strings.Builder

	for i := range 3 {
		output.WriteString(strings.Join(b.matrix[i][:], " | "))
		output.WriteString("\n")
	}

	return output.String()
}

func (b Board) markPosition(pos int, player string) error {
	line := int(math.Ceil(float64(pos)/3) - 1)
	col := int(math.Mod(float64(pos+2), 3))

	value := b.matrix[line][col]

	if value != "" {
		return fmt.Errorf("Position occuped")
	}

	b.matrix[line][col] = player

	return nil
}

func main() {
	currentRound := 1
	currentPlayer := X
	winner := ""

	fmt.Println("#### TIC-TAC-TOE ####")
	fmt.Printf("Let's start with the %s player\n", currentPlayer)

	board := Board{
		matrix: &[3][3]string{},
	}

	for currentRound <= 9 && winner == "" {
		fmt.Println("Round #", currentRound)
		pos, err := askPosition(currentPlayer)

		if err != nil {
			fmt.Println(err)
			continue
		}

		err = board.markPosition(pos, currentPlayer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(board)

		if currentPlayer == X {
			currentPlayer = O
		} else {
			currentPlayer = X
		}

		currentRound++
	}

}

func askPosition(currentPlayer string) (int, error) {
	var value string
	message := fmt.Sprintf("[%s player] Type a position: ", currentPlayer)
	fmt.Print(message)
	fmt.Scanf("%v\n", &value)
	nbmValue, err := strconv.Atoi(value)

	if err != nil {
		return 0, fmt.Errorf("%s is invalid number", value)
	}

	if nbmValue < 1 || nbmValue > 9 {
		return 0, fmt.Errorf("%s is not between 1 and 9", value)
	}

	return nbmValue, nil
}
