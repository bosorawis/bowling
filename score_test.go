package bowling

import (
	"reflect"
	"testing"
)

func TestRollToFallenPin(t *testing.T) {
	t.Run("successful conversion", func(t *testing.T) {
		testcases := []struct {
			description string
			input       string
			want        []int
		}{
			{
				description: "no strike, no spare",
				input:       "12",
				want:        []int{1, 2},
			},
			{
				description: "strike",
				input:       "x",
				want:        []int{10},
			},
			{
				description: "spare",
				input:       "5/",
				want:        []int{5, 5},
			},
			{
				description: "three throws",
				input:       "1/9",
				want:        []int{1, 9, 9},
			},
			{
				description: "missed",
				input:       "50",
				want:        []int{5, 0},
			},
		}
		for _, tc := range testcases {
			got, err := rollToFallenPins(tc.input)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %v got %v", tc.want, got)
			}
		}
	})

}

func TestCalculateScore(t *testing.T) {

	//t.Run("Test Validate Roll", func(t *testing.T) {
	//	testcases := []struct {
	//		description   string
	//		input         string
	//		expectedError error
	//	}{
	//		{
	//			description:   "valid roll strike",
	//			input:         "x",
	//			expectedError: nil,
	//		},
	//		{
	//			description:   "valid roll spare",
	//			input:         "3/",
	//			expectedError: nil,
	//		},
	//		{
	//			description:   "valid roll points",
	//			input:         "35",
	//			expectedError: nil,
	//		},
	//		{
	//			description:   "valid roll - last play",
	//			input:         "3/5",
	//			expectedError: nil,
	//		},
	//		{
	//			description:   "invalid roll - invalid character",
	//			input:         "ab",
	//			expectedError: ErrValidation,
	//		},
	//		{
	//			description:   "invalid roll - strike on second throw",
	//			input:         "3x",
	//			expectedError: ErrValidation,
	//		},
	//		{
	//			description:   "invalid roll - 3 throws specified when shouldn't",
	//			input:         "345",
	//			expectedError: ErrValidation,
	//		},
	//	}
	//	for _, tc := range testcases {
	//		err := validateRoll(tc.input)
	//		if err != tc.expectedError{
	//			t.Errorf("input %s: expect %v got %v", tc.input, tc.expectedError, err)
	//		}
	//	}
	//})
	t.Run("Test Record and Get scores frame", func(t *testing.T) {
		testcases := []struct {
			description string
			throws      []string
			expected    [10]int
		}{
			{
				description: "no roll returns all 0's",
				expected:    [10]int{},
			},
			{
				description: "no strikes no spares",
				throws:      []string{"11", "22"},
				expected:    [10]int{2, 4},
			},
		}
		for _, tc := range testcases {
			g := game{}
			for _, play := range tc.throws {
				err := g.Record(play)
				if err != nil {
					t.Fatalf("unexpected error %v", err)
				}
			}
		}
	})

}
