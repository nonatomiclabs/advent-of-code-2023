package day_04

import (
	"fmt"
	"regexp"
	"strings"
)

func Solution(inputLines []string, part2 bool) int {
	numbersRegex := regexp.MustCompile(`\d+`)
	cardResults := []int{}
	var totalResult int

	for index, line := range inputLines {
		fmt.Printf("Processing card %d: ", index)

		lineParts := strings.Split(line, ": ")
		lineWithoutHeader := lineParts[1]
		lineParts = strings.Split(lineWithoutHeader, " | ")
		winningNumbersSection := lineParts[0]
		userNumbersSection := lineParts[1]

		winningNumbers := numbersRegex.FindAllString(winningNumbersSection, -1)
		userNumbers := numbersRegex.FindAllString(userNumbersSection, -1)

		fmt.Printf("winning numbers: %v, user numbers: %v, ", winningNumbers, userNumbers)

		var cardResult int
		for _, userNumber := range userNumbers {
			for _, winningNumber := range winningNumbers {
				if userNumber == winningNumber {
					if cardResult == 0 || part2 {
						cardResult += 1
					} else {
						cardResult *= 2
					}
				}
			}
		}
		fmt.Printf("card result: %d\n", cardResult)
		cardResults = append(cardResults, cardResult)
		totalResult += cardResult
	}

	if part2 {
		var totalCardsCount int
		for cardIndex := range inputLines {
			totalCardsCount += getTotalCardsCountFromCard(cardIndex, cardResults)
		}

		fmt.Printf("total cards count: %d\n", totalCardsCount)
		return totalCardsCount
	} else {
		fmt.Printf("total result: %d\n", totalResult)
		return totalResult
	}
}

func getTotalCardsCountFromCard(cardIndex int, cardResults []int) int {
	currentCardResult := cardResults[cardIndex]

	totalCardsCount := 1 // the current card

	copiedCardIndices := []int{}
	for i := 0; i < currentCardResult; i++ {
		cardToCopyIndex := cardIndex + i + 1
		if cardToCopyIndex >= len(cardResults) {
			break
		}
		copiedCardIndices = append(copiedCardIndices, cardToCopyIndex)
	}

	for _, copiedCardIndex := range copiedCardIndices {
		totalCardsCount += getTotalCardsCountFromCard(copiedCardIndex, cardResults)
	}

	return totalCardsCount
}
