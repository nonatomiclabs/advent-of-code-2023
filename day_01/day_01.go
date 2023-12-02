package day_01

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
)

func Solution(inputLines []string, countSpelledOutDigits bool) int {
	calibrationSum := 0
	lineIndex := 0

	for _, line := range inputLines {
		var firstDigit, lastDigit string

		firstDigitEncountered := false
		scannerStartIndex := 0
		for index, character := range line {
			var digit string
			if character >= 48 && character <= 57 {
				digit = string(character)
			}

			if countSpelledOutDigits {
				substrings := buildAllSubstringsInString(line[scannerStartIndex : index+1])

				for _, substring := range substrings {
					if d, err := spelledOutDigitToDigit(substring); err == nil {
						digit = d
						scannerStartIndex = index
					}
				}
			}

			if digit != "" {
				if !firstDigitEncountered {
					firstDigitEncountered = true
					firstDigit = digit
				}
				lastDigit = digit
			}
		}

		calibrationValue, _ := strconv.Atoi(firstDigit + lastDigit)
		slog.Debug(
			fmt.Sprintf("Processed line %d", lineIndex),
			"line", line,
			"first digit", firstDigit,
			"last digit", lastDigit,
			"calibration value", calibrationValue,
		)

		calibrationSum += calibrationValue
		lineIndex += 1
	}

	slog.Info("Finished computation!", "calibration sum", calibrationSum)
	return calibrationSum
}

func spelledOutDigitToDigit(s string) (string, error) {
	switch s {
	case "one":
		return "1", nil
	case "two":
		return "2", nil
	case "three":
		return "3", nil
	case "four":
		return "4", nil
	case "five":
		return "5", nil
	case "six":
		return "6", nil
	case "seven":
		return "7", nil
	case "eight":
		return "8", nil
	case "nine":
		return "9", nil
	}

	return "0", errors.New("the given string is not a spelled-out digit")
}

func buildAllSubstringsInString(s string) []string {
	substrings := []string{}

	for i := 0; i <= len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substrings = append(substrings, s[i:j])
		}
	}

	return substrings
}
