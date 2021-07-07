package bowling

import (
	"fmt"
	"strconv"
)

var validInput = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '/', 'x', 'X'}
var (
	ErrValidation = fmt.Errorf("invalid input")
)

type game struct {
	rolls      []string
	fallenPins []int
}

func (g *game) Record(roll string) error {

	return nil
}

//func validateRoll(s string) error {
//	return nil
//}

func rollToFallenPins(roll string) ([]int, error) {
	scores := []int{}
	for i, c := range roll {
		if c == 'x' || c == 'X' {
			scores = append(scores, 10)
		} else if c == '/' {
			scores = append(scores, 10-scores[i-1])
		} else {
			sc, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, err
			}
			scores = append(scores, sc)
		}
	}
	return scores, nil
}

func isStrike(s string) bool {
	if s[0] == 'X' || s[0] == 'x' {
		return true
	}
	return false
}

func isSpare(s string) bool {
	if s[1] == '/' {
		return true
	}
	return false
}
