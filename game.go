package bowling

import (
	"fmt"
	"strconv"
	"strings"
)


const (
	totalPins = 10
	maxFrame  = 10
	lastFrame = 9
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
			result = append(result, totalPins)
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
	return result, nil
}

// Play records the scored played by human
func (g *game) Play(s string) error {
	roll, err := translateInput(s)
	if err != nil {
		return err
	}

	vErr := g.validateRoll(roll)
	if vErr != nil {
		return vErr
	}

	return g.record(roll)
}

func (g *game) validateRoll(rolls []int) error {
	if len(rolls) > 3 {
		return fmt.Errorf("only allow maximum of 3 throws")
	}
	if len(rolls) < 3 && sum(rolls) > 10 {
		return fmt.Errorf("invalid input: total score cannot be higher than 10 [got %d]", sum(rolls))
	}
	if g.currentFrame == lastFrame {
		if len(rolls) < 2 {
			return fmt.Errorf("you get at least 2 rolls in the last frame")
		}
		if rolls[0] + rolls[1] < totalPins && len(rolls) >= 3{
			return fmt.Errorf("only allow 3 throws if the pins are cleanedout")
		}
		if rolls[0] + rolls[1] >= totalPins && len(rolls) <= 2 {
			return fmt.Errorf("miss an extra throw for the last frame")
		}
	} else {
		if len(rolls) >= 3 {
			return fmt.Errorf("only last frame is allowed up to 3 rolls")
		}

	}
	return nil
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

func (g *game) FinalScore() int {
	total := 0
  	for _, f := range g.frames {
  		if f != nil{
			total += f.getScore()
		}
	}
	return total
}

func (g *game) record(rolls []int) error {
	if g.Finished() {
		return fmt.Errorf("game is already over")
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

// ScoreCard returns the score board
func (g *game) ScoreCard() []int {
	frameScores := []int{}
	for i := 0; i < g.currentFrame; i++ {
		if g.frames[i].completed() {
			frameScores = append(frameScores, g.frames[i].score)
		}
	}
	return frameScores
}

// Score total score at the current frame
func (g *game) CurrentScore() int {
	return sum(g.ScoreCard())
}
