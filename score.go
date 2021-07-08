package bowling

import (
	"fmt"
)

var validInput = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '/', 'x', 'X'}

var (
	ErrValidation = fmt.Errorf("invalid input")
)

const (
	totalPins = 10
	maxFrame  = 10
	maxThrows = 21
)

type frame struct {
	downPins       []int
	score          int
	remainingBonus int
	number         int
}

func (f *frame) completed() bool {
	return f.remainingBonus <= 0
}

func strikeFrame(rolls []int, num int) *frame {
	frame := &frame{
		downPins:       rolls,
		score:          sum(rolls),
		number:         num,
	}

	if len(rolls) == 1 && rolls[0] == 10 { // strike
		frame.remainingBonus = 2
	} else if len(rolls) == 2 && rolls[0] + rolls[1] == 10 {
		frame.remainingBonus = 1
	} else {
		frame.remainingBonus = 0
	}
	return frame
}
func newFrame(rolls []int, num int) *frame {
	return &frame{
		downPins:       []int{totalPins},
		score:          10,
		remainingBonus: 1,
		number:         num,
	}
}

type game struct {
	currentFrame int
	frames       [maxFrame]*frame
	scoringQueue []*frame
}

func newGame() *game {
	return &game{
		currentFrame: 0,
		frames:       [maxFrame]*frame{},
	}
}

func (g *game) calculateBonus(rolls []int){
	for i := 0 ; i < len(g.scoringQueue); i++ {
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

func (g *game) Record(rolls []int) error {
	if len(rolls) > 3 {
		return fmt.Errorf("cannot throw more than 3 balls")
	}

	var newFrame *frame
	g.calculateBonus(rolls)
	if len(rolls) == 1 && rolls[0] == 10 { // strike
		newFrame = strikeFrame(g.currentFrame)
		g.scoringQueue = append(g.scoringQueue, newFrame)
	} else if len(rolls) == 2 && rolls[0] + rolls[1] == 10{ // spare

	} else {
		newFrame = &frame{
			downPins: rolls,
			score:    sum(rolls),
		}
	}

	g.frames[g.currentFrame] = newFrame
	g.currentFrame += 1
	return nil
}

func sum(arr []int) (total int) {
	for _, v := range arr {
		total += v
	}
	return
}

func (g *game) Score() []int {
	frameScores := []int{}
	for i := 0; i < g.currentFrame; i++ {
		if g.frames[i].completed(){
			frameScores = append(frameScores, g.frames[i].score)
		}
	}
	return frameScores
}

func (g *game) Done() bool {
	if g.frames[maxFrame-1] != nil {
		return true
	}
	return false
}
