package day_06

import "testing"

func TestDay06(t *testing.T) {
	var tests = []struct {
		input   []string
		partTwo bool
		want    int
	}{
		{
			[]string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			false,
			288,
		},
		{
			[]string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			true,
			71503,
		},
	}

	for _, tc := range tests {
		t.Run("test", func(t *testing.T) {
			got := Solution(tc.input, tc.partTwo)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
