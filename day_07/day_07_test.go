package day_07

import "testing"

func TestHandType(t *testing.T) {
	var tests = []struct {
		hand      Hand
		withJoker bool
		want      HandType
	}{
		{
			Hand{cards: "32T3K"},
			false,
			OnePair,
		},
		{
			Hand{cards: "32T3K"},
			true,
			OnePair,
		},
		{
			Hand{cards: "KTJJT"},
			false,
			TwoPair,
		},
		{
			Hand{cards: "JJJJJ"},
			true,
			FiveOfAKind,
		},
		{
			Hand{cards: "KTJJT"},
			true,
			FourOfAKind,
		},
		{
			Hand{cards: "2233J"},
			true,
			FullHouse,
		},
		{
			Hand{cards: "2345J"},
			true,
			OnePair,
		},
		{
			Hand{cards: "Q2KJJ"},
			true,
			ThreeOfAKind,
		},
	}

	for _, tc := range tests {
		t.Run(tc.hand.cards, func(t *testing.T) {
			got := tc.hand.GetType(tc.withJoker)
			if got != tc.want {
				t.Errorf("got %s, want %s", handTypeNames[got], handTypeNames[tc.want])
			}
		})
	}
}

func TestLowerHandComparison(t *testing.T) {
	var tests = []struct {
		leftHand  Hand
		rightHand Hand
		withJoker bool
	}{
		{
			Hand{cards: "4K623"},
			Hand{cards: "746K8"},
			false,
		},
		{ // NOT SURE ABOUT THIS CASE
			Hand{cards: "JJJJJ"},
			Hand{cards: "2JJJJ"},
			false,
		},
		{
			Hand{cards: "J345A"},
			Hand{cards: "2345J"},
			true,
		},
		{
			Hand{cards: "2345A"},
			Hand{cards: "2345J"},
			true,
		},
		{
			Hand{cards: "T3Q33"},
			Hand{cards: "Q2KJJ"},
			true,
		},
	}

	for _, tc := range tests {
		t.Run("test", func(t *testing.T) {
			if !tc.leftHand.LowerThan(tc.rightHand, tc.withJoker) {
				t.Errorf("expected hand %v to be lower than hand %v, but it's not", tc.leftHand, tc.rightHand)
			}
		})
	}
}

func TestDay07(t *testing.T) {
	var tests = []struct {
		input   []string
		partTwo bool
		want    int
	}{
		{
			[]string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			false,
			6440,
		},
		{
			[]string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			true,
			5905,
		},
		{
			[]string{
				"2345A 1",
				"Q2KJJ 13",
				"Q2Q2Q 19",
				"T3T3J 17",
				"T3Q33 11",
				"2345J 3",
				"J345A 2",
				"32T3K 5",
				"T55J5 29",
				"KK677 7",
				"KTJJT 34",
				"QQQJA 31",
				"JJJJJ 37",
				"JAAAA 43",
				"AAAAJ 59",
				"AAAAA 61",
				"2AAAA 23",
				"2JJJJ 53",
				"JJJJ2 41",
			},
			false,
			6592,
		},
		{
			[]string{
				"2345A 1",
				"Q2KJJ 13",
				"Q2Q2Q 19",
				"T3T3J 17",
				"T3Q33 11",
				"2345J 3",
				"J345A 2",
				"32T3K 5",
				"T55J5 29",
				"KK677 7",
				"KTJJT 34",
				"QQQJA 31",
				"JJJJJ 37",
				"JAAAA 43",
				"AAAAJ 59",
				"AAAAA 61",
				"2AAAA 23",
				"2JJJJ 53",
				"JJJJ2 41",
			},
			true,
			6839,
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
