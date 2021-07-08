package bowling

import (
	"fmt"
	"strconv"
	"strings"
)

var validInput = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '/', 'x', 'X'}

var (
	ErrValidation = fmt.Errorf("invalid input")
)

const (
	totalPins = 10
	maxFrame  = 10
	lastFrame = 9
	maxThrows = 21
)

type game struct {
	currentFrame int
	frames       [maxFrame]*frame
	scoringQueue []*frame
}

func NewGame() *game {
	return &game{
		currentFrame: 0,
		frames:       [maxFrame]*frame{},
	}
}

func (g *game) calculateBonus(rolls []int) {
	for i := 0; i < len(g.scoringQueue); i++ {
		scoreFrame := g.scoringQueue[i]
		for j := 0; j < len(rolls); j++ {
			if scoreFrame.remainingBonus <= 0 {
				break
			}
			scoreFrame.score += rolls[j]
			scoreFrame.remainingBonus -= 1
		}
		if scoreFrame.remainingBonus <= 0 {
			g.frames[scoreFrame.number] = scoreFrame
		}
	}
	temp := g.scoringQueue[:0]
	for _, v := range g.scoringQueue {
		if v.remainingBonus > 0 {
			temp = append(temp, v)
		}
	}
	g.scoringQueue = temp
}

func translateInput(s string) ([]int, error) {
	splited := strings.Split(s, ",")
	if len(splited) < 1 || len(splited) > 3 {
		return nil, fmt.Errorf("expect 1-3 plays separated by comma")
	}
	result := []int{}
	for i, c := range splited {
		if c == "X" {
			result = append(result, 10)
		} else if c == "/" {
			result = append(result, totalPins-result[i-1])
		} else if c == "-" {
			result = append(result, 0)
		} else {
			point, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			result = append(result, point)
		}
	}

	if len(result) < 3 && sum(result) > 10 {
		return nil, fmt.Errorf("invalid input: total score cannot be higher than 10 [got %d]", sum(result))
	}


	return result, nil
}

// Play records the scored played by human
func (g *game) Play(s string) error {
	roll, err := translateInput(s)
	if err != nil {
		return err
	}
	return g.record(roll)
}

// CurrentFrame prints out currently in-play frame (+1 for human readability)
func (g *game) CurrentFrame() int {
	if g.currentFrame + 1 > maxFrame {
		return maxFrame
	}
	return g.currentFrame + 1
}

// Finished checks if the game is completed
func (g *game) Finished() bool {
	if g.currentFrame >= maxFrame {
		return true
	}
	return false
}

func (g *game) record(rolls []int) error {
	if g.Finished() {
		return fmt.Errorf("game is already over")
	}
	if len(rolls) > 3 {
		return fmt.Errorf("cannot throw more than 3 balls")
	}
	g.calculateBonus(rolls)
	currentFrame := newFrame(rolls, g.currentFrame)
	if currentFrame.remainingBonus > 0 {
		g.scoringQueue = append(g.scoringQueue, currentFrame)
	}
	g.frames[g.currentFrame] = currentFrame
	g.currentFrame += 1
	return nil
}

// Score returns the score board
func (g *game) Score() []int {
	frameScores := []int{}
	for i := 0; i < g.currentFrame; i++ {
		if g.frames[i].completed() {
			frameScores = append(frameScores, g.frames[i].score)
		}
	}
	return frameScores
}
