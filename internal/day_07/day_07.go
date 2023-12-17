// SPDX-License-Identifier: AGPL-3.0-or-later

package day_07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Runner struct{}

var cardRank = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var handTypes = map[string]int{
	"One Pair":        1,
	"Two Pairs":       2,
	"Three of a Kind": 3,
	"Full House":      4,
	"Four of a Kind":  5,
	"Five of a Kind":  6,
}

var handTypeNames = map[int]string{
	1: "One Pair",
	2: "Two Pairs",
	3: "Three of a Kind",
	4: "Full House",
	5: "Four of a Kind",
	6: "Five of a Kind",
}

type Card struct {
	value int
	name  string
}

type Hand struct {
	cards       []Card
	sortedCards []Card
	handType    int
	bid         int
	jokers      int
}

// A function that returns the index and the struct of the first
// card that exists concurrently n times in a sorted slice of Cards
func (h *Hand) findRecurringCard(slice []Card, n int, isPartOne bool) (int, *Card) {
	if n == 5 && h.jokers == 5 {
		return 0, &slice[0]
	}

	endRange := max(0, n-h.jokers-1)

	for i := 0; i < len(slice)-n+1; i++ {
		if !isPartOne && slice[i].name == "J" {
			continue
		}
		if slice[i].value == slice[i+endRange].value {
			return i, &slice[i]
		}
	}
	return -1, nil
}

func (h *Hand) sortCards() {
	h.sortedCards = make([]Card, len(h.cards))
	copy(h.sortedCards, h.cards)

	sort.Slice(h.sortedCards, func(i, j int) bool {
		return h.sortedCards[i].value > h.sortedCards[j].value
	})
}

// Finds all five of a kind hand types
func (h *Hand) findFiveOfAKind(isPartOne bool) (bool, []Card) {
	idx, _ := h.findRecurringCard(h.sortedCards, 5, isPartOne)
	if idx == -1 {
		return false, nil
	}

	h.handType = handTypes["Five of a Kind"]
	return true, []Card{}
}

// Finds all four of a kind hand types
func (h *Hand) findFourOfAKind(isPartOne bool) (bool, []Card) {
	idx, _ := h.findRecurringCard(h.sortedCards, 4, isPartOne)
	if idx == -1 {
		return false, nil
	}
	h.handType = handTypes["Four of a Kind"]
	return true, append(h.sortedCards[:idx], h.sortedCards[idx+4:]...)
}

// Find all three of a kind hand types
func (h *Hand) findThreeOfAKindOrFullHouse(isPartOne bool) (bool, []Card) {
	idx, _ := h.findRecurringCard(h.sortedCards, 3, isPartOne)
	if idx == -1 {
		return false, nil
	}
	restCards := append(h.sortedCards[:idx], h.sortedCards[idx+3:]...)
	h.jokers = 0

	// Find another pair for a full house
	idx2, _ := h.findRecurringCard(restCards, 2, isPartOne)
	if idx2 != -1 {
		h.handType = handTypes["Full House"]
		return true, append(restCards[:idx2], restCards[idx2:]...)
	}

	h.handType = handTypes["Three of a Kind"]
	return true, append(h.sortedCards[:idx], h.sortedCards[idx+3:]...)
}

// Find one or two pairs of hand types
func (h *Hand) findPairs(isPartOne bool) (bool, []Card) {
	idx1, _ := h.findRecurringCard(h.sortedCards, 2, isPartOne)
	if idx1 == -1 {
		return false, nil
	}
	restCards := append(h.sortedCards[:idx1], h.sortedCards[idx1+2:]...)
	h.jokers = 0

	idx2, _ := h.findRecurringCard(restCards, 2, isPartOne)
	if idx2 == -1 {
		h.handType = handTypes["One Pair"]
	} else {
		h.handType = handTypes["Two Pairs"]
		restCards = append(restCards[:idx2], restCards[idx2+2:]...)
	}

	return true, restCards
}

func (h *Hand) findHandType(isPartOne bool) {
	var match bool
	match, _ = h.findFiveOfAKind(isPartOne)
	if match {
		return
	}
	match, _ = h.findFourOfAKind(isPartOne)
	if match {
		return
	}
	match, _ = h.findThreeOfAKindOrFullHouse(isPartOne)
	if match {
		return
	}
	_, _ = h.findPairs(isPartOne)
}

func (h *Hand) addCardsFromString(input string, isPartOne bool) error {
	cards := strings.Split(input, "")
	for _, card := range cards {
		value := cardRank[card]
		if !isPartOne && card == "J" {
			value = 1
			h.jokers++
		}
		if value == 0 {
			return fmt.Errorf("invalid card value: %s", card)
		}
		h.cards = append(h.cards, Card{value, card})
	}
	return nil
}

func createHandFromInput(line string, isPartOne bool) (*Hand, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid input: %s", line)
	}

	var hand Hand

	err := hand.addCardsFromString(parts[0], isPartOne)
	if err != nil {
		return nil, err
	}

	hand.sortCards()
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	hand.bid = bid

	hand.findHandType(isPartOne)

	return &hand, nil
}

func (r Runner) Run(input []string, isPartOne bool) (string, error) {
	hands := make([]*Hand, len(input))
	for i, line := range input {
		hand, err := createHandFromInput(line, isPartOne)
		if err != nil {
			return "", err
		}
		hands[i] = hand
	}

	// sort hands by rank
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			// If the hand types are equal, then the hand with the highest card value wins
			// But here we use the original card order, not the sorted card order
			for k := 0; k < len(hands[i].cards); k++ {
				if hands[i].cards[k].value == hands[j].cards[k].value {
					continue
				}
				return hands[i].cards[k].value < hands[j].cards[k].value
			}
		}

		return hands[i].handType < hands[j].handType
	})

	sum := 0
	for i, hand := range hands {
		cardsStr := ""
		for _, card := range hand.cards {
			cardsStr += card.name
		}
		// Keeping this line for debugging purposes
		fmt.Printf("Hand: %s, Rank: %d, Hand type: %s\n", cardsStr, i+1, handTypeNames[hand.handType])
		sum += hand.bid * (i + 1)
	}

	return strconv.Itoa(sum), nil
}
