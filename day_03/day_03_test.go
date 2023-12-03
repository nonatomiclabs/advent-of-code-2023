package day_03

import "testing"

func TestDay03(t *testing.T) {
	var tests = []struct {
		input []string
		part2 bool
		want  int
	}{
		// {
		// 	[]string{
		// 		"467..114..",
		// 		"...*......",
		// 		"..35..633.",
		// 		"......#...",
		// 		"617*......",
		// 		".....+.58.",
		// 		"..592.....",
		// 		"......755.",
		// 		"...$.*....",
		// 		".664.598..",
		// 	},
		// 	false,
		// 	4361,
		// },
		{
			[]string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			true,
			467835,
		},
	}

	for _, tc := range tests {
		t.Run("test", func(t *testing.T) {
			got := Solution(tc.input, tc.part2)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
