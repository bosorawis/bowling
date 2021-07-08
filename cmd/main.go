package main

import (
	"bufio"
	"fmt"
	"github.com/dihmuzikien/bowling"
	"os"
	"strings"
)

type gameStateViewer interface {
	CurrentScore() int
	ScoreCard() []int
	CurrentFrame() int
}

func main() {
	g := bowling.NewGame()
	for !g.Finished() {
		currentFrame := g.CurrentFrame()
		play, err := input(currentFrame)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("scoring: '%s' for frame #%d\n", play, currentFrame)
		playErr := g.Play(play)
		if playErr != nil {
			fmt.Printf("failed to score: %v\n", playErr)
			continue
		}
		fmt.Printf("finished frame #%d\n", currentFrame)
		printGameState(g)
	}
	fmt.Printf("Total score is %d\n", g.FinalScore())
	fmt.Println("good bye!")
}

func printGameState(g gameStateViewer) {
	fmt.Printf("ScoreCard: %v | Current is %v\n", g.ScoreCard(), g.CurrentScore())
	fmt.Println("----------------------------")
}

func input(frame int) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">>> record your score for frame %d: ", frame)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	strTrim := strings.Trim(userInput, "\t \n")
	return strTrim, nil
}
