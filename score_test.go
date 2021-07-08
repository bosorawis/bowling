package bowling

import (
	"reflect"
	"testing"
)

func TestRecordAndScore(t *testing.T) {
	t.Run("Test recording and scoring my bowling game", func(t *testing.T) {
		testcases := []struct {
			input [][]int
			want  []int
		}{
			{
				input: [][]int{{1, 1}, {2, 2}, {3, 3}},
				want:  []int{2, 4, 6},
			},
			{
				input: [][]int{{10}, {10}, {10}, {5, 3}},
				want:  []int{30, 25, 18, 8},
			},
			{
				input: [][]int{{10}, {10}},
				want:  []int{},
			},
			{
				input: [][]int{{5, 5}, {5, 5}, {2, 3}},
				want:  []int{15, 12, 5},
			},

		}
		for _, tc := range testcases {
			g := game{}
			for _, v := range tc.input {
				g.Record(v)
			}
			got := g.Score()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("input %v. want %v got %v", tc.input, tc.want, got)
			}
		}
	})
}
