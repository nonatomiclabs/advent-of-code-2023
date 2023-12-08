package day_08

import "testing"

func TestDay08(t *testing.T) {
	var tests = []struct {
		input   []string
		partTwo bool
		want    int
	}{
		{
			[]string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			false,
			2,
		},
		{
			[]string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},
			true,
			6,
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
