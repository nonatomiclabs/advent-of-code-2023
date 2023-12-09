package day_09

import "testing"

func TestDay09(t *testing.T) {
	var tests = []struct {
		input   []string
		partTwo bool
		want    int
	}{
		{
			[]string{
				"0 3 6 9 12 15",
			},
			false,
			18,
		},
		{
			[]string{
				"10 13 16 21 30 45",
			},
			false,
			68,
		},
		{
			[]string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			false,
			114,
		},
		{
			[]string{
				"10 13 16 21 30 45",
			},
			true,
			5,
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
