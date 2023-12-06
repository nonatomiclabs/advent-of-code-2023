package day_06

import (
	"fmt"
	"regexp"
	"strconv"
)

type Race struct {
	Time         int
	BestDistance int
}

func ComputeBetterRacesCount(race Race) int {
	betterRacesCount := 0

	for i := 0; i < race.Time; i++ {
		distance := (race.Time - i) * i
		if distance > race.BestDistance {
			betterRacesCount++
		}
	}

	return betterRacesCount
}

func Solution(inputLines []string, part2 bool) int {
	numberRegex := regexp.MustCompile(`\d+`)
	raceTimes := []int{}
	raceBestDistances := []int{}

	for lineIndex, line := range inputLines {
		numbers := []int{}
		numberMatch := numberRegex.FindAllString(line, -1)
		for _, m := range numberMatch {
			number, _ := strconv.Atoi(m)
			numbers = append(numbers, number)
		}

		if part2 {
			totalNumberString := ""
			for _, m := range numberMatch {
				totalNumberString += m
			}
			totalNumber, _ := strconv.Atoi(totalNumberString)
			fmt.Printf("Parsed numbers without accounting for whitespace. Total number is %d\n", totalNumber)
			numbers = []int{totalNumber}
		}

		if lineIndex == 0 {
			raceTimes = numbers
		} else {
			raceBestDistances = numbers
		}

	}

	races := []Race{}
	for i := 0; i < len(raceTimes); i++ {
		race := Race{Time: raceTimes[i], BestDistance: raceBestDistances[i]}
		races = append(races, race)
	}

	totalBetterPossibleRaces := []int{}
	for raceIndex, race := range races {
		betterPossibleRaces := ComputeBetterRacesCount(race)
		fmt.Printf("For race %d: %d better races possible\n", raceIndex, betterPossibleRaces)
		totalBetterPossibleRaces = append(totalBetterPossibleRaces, betterPossibleRaces)
	}

	result := totalBetterPossibleRaces[0]
	for i := 1; i < len(totalBetterPossibleRaces); i++ {
		result *= totalBetterPossibleRaces[i]
	}

	fmt.Printf("Final result: %d\n", result)
	return result
}
