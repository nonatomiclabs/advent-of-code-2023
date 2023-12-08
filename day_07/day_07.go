package day_07

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	FiveOfAKind  HandType = 7
	FourOfAKind  HandType = 6
	FullHouse    HandType = 5
	ThreeOfAKind HandType = 4
	TwoPair      HandType = 3
	OnePair      HandType = 2
	HighCard     HandType = 1
)

var handTypeNames = map[HandType]string{
	FiveOfAKind:  "FiveOfAKind",
	FourOfAKind:  "FourOfAKind",
	FullHouse:    "FullHouse",
	ThreeOfAKind: "ThreeOfAKind",
	TwoPair:      "TwoPair",
	OnePair:      "OnePair",
	HighCard:     "HighCard",
}

type Hand struct {
	cards string
}

type CardCount struct {
	card  string
	count int
}

func (h *Hand) GetType(withJoker bool) HandType {

	cardsMapping := make(map[string]int)
	for _, card := range h.cards {
		cardsMapping[string(card)]++
	}

	cardCounts := []CardCount{}
	for cardName, cardCount := range cardsMapping {
		cardCounts = append(cardCounts, CardCount{card: cardName, count: cardCount})
	}

	sort.Slice(cardCounts, func(i int, j int) bool {
		leftCount := cardCounts[i].count
		rightCount := cardCounts[j].count

		if leftCount == rightCount {
			if cardCounts[i].card == "J" {
				return true
			}
		}

		return leftCount < rightCount
	})

	slices.Reverse(cardCounts)

	if cardCounts[0].card == "J" && len(cardCounts) >= 2 {
		swapF := reflect.Swapper(cardCounts)
		swapF(0, 1)
	}

	processedCards := []CardCount{}
	for index, cardCount := range cardCounts {
		if withJoker && index > 0 && cardCount.card == "J" {
			processedCards[0].count += cardCount.count
			continue
		}
		processedCards = append(processedCards, cardCount)
	}

	cardsCount := []int{}
	for _, c := range processedCards {
		cardsCount = append(cardsCount, c.count)
	}
	// fmt.Printf("\tcards count: %v\n", cardsCount)

	switch cardsCount[0] {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if cardsCount[1] == 2 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	case 2:
		if cardsCount[1] == 2 {
			return TwoPair
		} else {
			return OnePair
		}
	case 1:
		return HighCard
	default:
		return 0
	}
}

var cardValues = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

var cardValuesWithJoker = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func (thisHand *Hand) LowerThan(otherHand Hand, withJoker bool) bool {
	// fmt.Printf("Checking if hand %v is lower than hand %v\n", *thisHand, otherHand)
	leftHandType := thisHand.GetType(withJoker)
	rightHandType := otherHand.GetType(withJoker)
	// fmt.Printf("\tleft hand is a %s\tright hand is a %s\n", handTypeNames[leftHandType], handTypeNames[rightHandType])

	if leftHandType == rightHandType {
		// fmt.Printf("both hand have the same type, comparing their cards\n")
		var cardValuesToUse map[string]int
		if withJoker {
			cardValuesToUse = cardValuesWithJoker
		} else {
			cardValuesToUse = cardValues
		}

		for c := 0; c < 5; c++ {
			leftCardValue := cardValuesToUse[string(thisHand.cards[c])]
			rightCardValue := cardValuesToUse[string(otherHand.cards[c])]

			if leftCardValue == rightCardValue {
				continue
			}
			return leftCardValue < rightCardValue
		}
	}
	return leftHandType < rightHandType
}

func Solution(inputLines []string, partTow bool) int {

	hands := []Hand{}
	handsToBid := make(map[Hand]int)
	for _, line := range inputLines {
		splittedLine := strings.Split(line, " ")
		hand := Hand{cards: splittedLine[0]}
		bid, _ := strconv.Atoi(splittedLine[1])
		hands = append(hands, hand)
		handsToBid[hand] = bid
	}

	fmt.Printf("Hands: %v\n", hands)

	sort.Slice(hands, func(i int, j int) bool {
		return hands[i].LowerThan(hands[j], partTow)
	})

	fmt.Printf("Sorted hands: %v\n", hands)

	totalWinnings := 0

	bids := []int{}

	for rank, hand := range hands {
		bid := handsToBid[hand]
		bids = append(bids, bid)
		totalWinnings += bid * (rank + 1)
	}

	fmt.Printf("Sorted bids: %v\n", bids)

	fmt.Printf("Total winnings: %d\n", totalWinnings)
	return totalWinnings
}
