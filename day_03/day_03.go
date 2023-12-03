package day_03

import (
	"fmt"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

func Solution(inputLines []string, part2 bool) int {
	var specialCharactersCoordinates [][]int

	for i, line := range inputLines {
		for j, character := range line {
			if !((character >= 48 && character <= 57) || character == 46) {
				characterCoordinates := []int{j, i}
				if !part2 || (part2 && character == 42) {
					specialCharactersCoordinates = append(specialCharactersCoordinates, characterCoordinates)
				}
			}
		}
	}

	fmt.Printf("Found special characters at the following positions: %v\n", specialCharactersCoordinates)
	numbersAdjacentToSpecialCharacters := make(map[Point][]int)

	var sum int
	var gearRatioSum int
	numbersInLineRegexp := regexp.MustCompile(`\d+`)
	for index, line := range inputLines {
		numbersInLine := numbersInLineRegexp.FindAllStringSubmatchIndex(line, -1)
		fmt.Printf("Processing line %d, found numbers: %v\n", index, numbersInLine)

		for numberIndex, numberStartEndRange := range numbersInLine {
			// fmt.Printf("\tfor number range %v, ", numberStartEndRange)
			numberStart := numberStartEndRange[0]
			numberEnd := numberStartEndRange[1]
			coordinatesToConsider := [][]int{}

			startConsiderX := numberStart
			if numberStart > 0 {
				startConsiderX -= 1
				coordinatesToConsider = append(coordinatesToConsider, []int{startConsiderX, index})
			}

			endConsiderX := numberEnd - 1
			if numberEnd < len(line) {
				endConsiderX += 1
				coordinatesToConsider = append(coordinatesToConsider, []int{endConsiderX, index})
			}

			if index > 0 {
				// Check potential characters that are on the line above
				for i := startConsiderX; i <= endConsiderX; i++ {
					coordinatesToConsider = append(coordinatesToConsider, []int{i, index - 1})
				}
			}

			if index < len(inputLines) {
				// Check potential characters that are on the line below
				for i := startConsiderX; i <= endConsiderX; i++ {
					coordinatesToConsider = append(coordinatesToConsider, []int{i, index + 1})
				}
			}

			// fmt.Printf("coordinates to consider: %v\n", coordinatesToConsider)

			considerNumber := false
			adjacentSpecialCharater := Point{}
			for _, coordinatesSet := range coordinatesToConsider {
				for _, specialCharacterCoordinateSet := range specialCharactersCoordinates {
					if coordinatesSet[0] == specialCharacterCoordinateSet[0] && coordinatesSet[1] == specialCharacterCoordinateSet[1] {
						adjacentSpecialCharater.x = specialCharacterCoordinateSet[0]
						adjacentSpecialCharater.y = specialCharacterCoordinateSet[1]
						considerNumber = true
						break
					}
				}
			}

			actualNumberInLineString := numbersInLineRegexp.FindAllStringSubmatch(line, -1)[numberIndex][0]
			actualNumberInLine, _ := strconv.Atoi(actualNumberInLineString)
			if considerNumber {
				fmt.Printf("number %d is adjacent to a special character, adding it to the sum\n", actualNumberInLine)
				if !part2 {
					sum += actualNumberInLine
				}
				numbersAdjacentToSpecialCharacters[adjacentSpecialCharater] = append(numbersAdjacentToSpecialCharacters[adjacentSpecialCharater], actualNumberInLine)
			} else {
				fmt.Printf("number %d is not adjacent to a special character, dismissing it\n", actualNumberInLine)
			}
		}
	}

	if part2 {
		fmt.Printf("numbersAdjacentToSpecialCharacters: %v\n", numbersAdjacentToSpecialCharacters)
		for _, numbers := range numbersAdjacentToSpecialCharacters {
			if len(numbers) == 2 {
				gearRatio := numbers[0] * numbers[1]
				fmt.Printf("adding gear ratio composed of %d x %d: %d\n", numbers[0], numbers[1], gearRatio)
				gearRatioSum += gearRatio
				fmt.Printf("gear ratio sum is now: %d\n", gearRatioSum)
			}
		}
	}

	// fmt.Printf("specialCharactersCoordinates: %v\n", specialCharactersCoordinates)
	if part2 {
		fmt.Printf("gear ratio sum: %d\n", gearRatioSum)
		return gearRatioSum
	} else {
		fmt.Printf("sum: %d\n", sum)
		return sum
	}
}
