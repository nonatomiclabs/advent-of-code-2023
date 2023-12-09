package day_09

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func ComputeNextHistory(numbers []int, processedNumbers []int, partTwo bool) int {
	differences := []int{}
	for i := 1; i < len(numbers); i++ {
		difference := numbers[i] - numbers[i-1]
		differences = append(differences, difference)
	}
	fmt.Printf("\tDifferences: %v\n", differences)
	if partTwo {
		processedNumbers = append(processedNumbers, differences[0])
	} else {
		processedNumbers = append(processedNumbers, differences[len(differences)-1])
	}

	differenceAllZeroes := false
	for i, difference := range differences {
		if difference != 0 {
			break
		}
		if i == len(differences)-1 {
			differenceAllZeroes = true
		}
	}
	fmt.Printf("\tDifference is all zeroes: %t\n", differenceAllZeroes)

	result := 0
	if differenceAllZeroes {
		fmt.Printf("\tStopping. First/last numbers: %v\n", processedNumbers)

		if partTwo {
			slices.Reverse(processedNumbers)
			// processedNumbers = processedNumbers[1:]
		}

		i := 0
		if partTwo {
			i = 1
		}
		previousNumber := 0
		for ; i < len(processedNumbers); i++ {
			if partTwo {
				currentNumber := processedNumbers[i]
				difference := (currentNumber - previousNumber)
				previousNumber = difference
				if i == len(processedNumbers)-1 {
					result = difference
				}
			} else {
				result += processedNumbers[i]
			}
		}
		return result
	} else {
		result = ComputeNextHistory(differences, processedNumbers, partTwo)
	}
	return result
}

func Solution(inputLines []string, partTwo bool) int {
	var sum int

	numbersRegex := regexp.MustCompile(`-?\d+`)
	for lineIndex, line := range inputLines {
		numbers := []int{}
		match := numbersRegex.FindAllString(line, -1)

		for _, m := range match {
			digit, _ := strconv.Atoi(m)
			numbers = append(numbers, digit)
		}

		fmt.Printf("Processing line %d: %v\n", lineIndex, numbers)

		var initialNumber int
		if partTwo {
			initialNumber = numbers[0]
		} else {
			initialNumber = numbers[len(numbers)-1]
		}
		processedNumbers := []int{initialNumber}
		nextHistory := ComputeNextHistory(numbers, processedNumbers, partTwo)
		sum += nextHistory
	}

	fmt.Printf("total sum: %d\n", sum)
	return sum
}
