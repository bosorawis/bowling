package bowling

import (
	"reflect"
	"testing"
)

func TestRecordAndScore(t *testing.T) {
	t.Run("Test recording and scoring my bowling game", func(t *testing.T) {
		testcases := []struct {
			description string
			input       []string
			want        []int
			wantTotal   int
		}{
			{
				description: "strikes calculate ahead",
				input:       []string{"X", "X", "X", "5,3"},
				want:        []int{30, 25, 18, 8},
				wantTotal:   81,
			},
			{
				description: "incomplete strikes don't display scores",
				input:       []string{"X", "X"},
				want:        []int{},
				wantTotal:   0,
			},
			{
				description: "spares calculate ahead",
				input:       []string{"5,5", "5,5", "2,3"},
				want:        []int{15, 12, 5},
				wantTotal:   32,
			},
			{
				description: "perfect game",
				input:       []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "X,X,X"},
				want:        []int{30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
				wantTotal:   300,
			},
			{
				description: "spares last throw",
				input:       []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "-,/,X"},
				want:        []int{30, 30, 30, 30, 30, 30, 30, 20, 20, 20},
				wantTotal:   270,
			},
			{
				description: "rough game",
				input:       []string{"-,-", "-,-", "-,-", "-,-", "-,-", "-,-", "-,-", "-,-", "-,-", "-,-"},
				want:        []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				wantTotal:   0,
			},
			{
				description: "rough game",
				input:       []string{"-,-", "-,-", "-,-", "-,-", "-,-", "-,-", "-,-", "-,-", "-,-", "-,9"},
				want:        []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 9},
				wantTotal:   9,
			},
		}
		for _, tc := range testcases {
			g := game{}
			for _, v := range tc.input {
				err := g.Play(v)
				if err != nil {
					t.Fatalf("input %v: unexpected error %v", tc.input, err)
				}
			}
			got := g.Score()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("input %v. want %v got %v", tc.input, tc.want, got)
			}
			gotTotal := g.FinalScore()
			if gotTotal != tc.wantTotal {
				t.Errorf("input %v. want %d got %d", tc.input, tc.wantTotal, gotTotal)
			}
		}
	})
	t.Run("invalid rolls return error", func(t *testing.T) {
		testcases := []struct {
			description string
			input       []string
		}{
			{
				description: "roll invalid score",
				input:       []string{"X,2"},
			},
			{
				description: "roll invalid score",
				input:       []string{"6,5"},
			},
			{
				description: "try to roll 3 times without strikes/spares last frame",
				input:       []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "1,1,X"},
			},
			{
				description: "extra roll",
				input:       []string{"X,X,X"},
			},
			{
				description: "missing roll",
				input:       []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "X"},
			},
		}
		for _, tc := range testcases {
			g := game{}
			var err error
			for _, v := range tc.input {
				err = g.Play(v)
				if err != nil {
					break
				}
			}
			if err == nil {
				t.Errorf("%s, input: %v | expected an error but got nothing", tc.description, tc.input)
			}
		}

	})

	t.Run("10 rolls complete the game", func(t *testing.T) {
		input := [][]int{{10}, {10}, {10}, {10}, {10}, {10}, {10}, {10}, {10}, {10, 10, 10}}
		g := game{}
		for _, roll := range input {
			g.record(roll)
		}
		if g.Finished() != true {
			t.Fatal("expect game to be finished")
		}
	})

}

func TestTranslateThrow(t *testing.T) {
	t.Run("valid inputs", func(t *testing.T) {
		testcases := []struct {
			input string
			want  []int
		}{
			{
				input: "1,2",
				want:  []int{1, 2},
			},
			{
				input: "9,1",
				want:  []int{9, 1},
			},
			{
				input: "9,/",
				want:  []int{9, 1},
			},
			{
				input: "X",
				want:  []int{10},
			},
			{
				input: "7,/,X",
				want:  []int{7, 3, 10},
			},
		}
		for _, tc := range testcases {
			got, err := translateInput(tc.input)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("input %v want %v got %v", tc.input, tc.want, got)
			}
		}
	})
}
