package day_02

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solution(inputLines []string, part2 bool) int {
	maxCubeCount := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var sumOfPossibleGames, sumOfPowers int

	gameIndexRegex := regexp.MustCompile("Game ([0-9]+)")
	gameSetResultRegex := regexp.MustCompile("([0-9]+) ([a-z]+)")

	for _, line := range inputLines {
		splitLine := strings.Split(line, ": ")
		if len(splitLine) != 2 {
			fmt.Fprint(os.Stderr, "could not split input line correctly")
			os.Exit(1)
		}
		result := splitLine[1]

		dismissGame := false
		highestCubeCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		gameSets := strings.Split(result, ";")
		for _, gameSet := range gameSets {

			gameSetResultMatch := gameSetResultRegex.FindAllStringSubmatch(gameSet, -1)

			for _, match := range gameSetResultMatch {
				cubeCount, _ := strconv.Atoi(match[1])
				cubeColor := match[2]

				if cubeCount > maxCubeCount[cubeColor] {
					dismissGame = true
				}

				currentHighest := highestCubeCount[cubeColor]
				if cubeCount > currentHighest {
					highestCubeCount[cubeColor] = cubeCount
				}
			}
		}

		if !dismissGame {
			// Determine the index of the game and add it to the sum of possible games
			prefix := splitLine[0]
			gameIndexMatch := gameIndexRegex.FindStringSubmatch(prefix)
			if len(gameIndexMatch) < 2 {
				fmt.Fprint(os.Stderr, "could not parse game index")
				os.Exit(1)
			}
			digits := gameIndexMatch[1]

			gameIndex, _ := strconv.Atoi(digits)
			sumOfPossibleGames += gameIndex
		}

		sumOfPowers += highestCubeCount["red"] * highestCubeCount["green"] * highestCubeCount["blue"]
	}

	var solution int
	if part2 {
		solution = sumOfPowers
	} else {
		solution = sumOfPossibleGames
	}

	slog.Info("Finished computation!", "solution", solution)
	return solution
}
