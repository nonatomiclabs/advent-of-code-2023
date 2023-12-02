package day_01

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay01(t *testing.T) {
	var tests = []struct {
		input                 []string
		countSpelledOutDigits bool
		want                  int
	}{
		{
			[]string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			false,
			142,
		},
		{
			[]string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			true,
			281,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Cound spelled-out digits: %t", tc.countSpelledOutDigits), func(t *testing.T) {
			got := Solution(tc.input, tc.countSpelledOutDigits)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestBuildAllSubstringsInString(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{"abc", []string{"a", "ab", "abc", "b", "bc", "c"}},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := buildAllSubstringsInString("abc")
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}

}
