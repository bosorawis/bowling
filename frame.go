package bowling

type frame struct {
	downPins       []int
	score          int
	remainingBonus int
	number         int
}
func sum(arr []int) (total int) {
	for _, v := range arr {
		total += v
	}
	return
}

func newFrame(rolls []int, num int) *frame {
	frame := &frame{
		downPins: rolls,
		score:    sum(rolls),
		number:   num,
	}
	// special case for last play - no more bonus
	if num == 9 {
		return frame
	}

	if len(rolls) == 1 && rolls[0] == 10 { // strike
		frame.remainingBonus = 2
	} else if len(rolls) == 2 && rolls[0]+rolls[1] == 10 {
		frame.remainingBonus = 1
	} else {
		frame.remainingBonus = 0
	}
	return frame
}

func (f *frame) completed() bool {
	return f.remainingBonus <= 0
}

