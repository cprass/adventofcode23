// SPDX-License-Identifier: AGPL-3.0-or-later

package day_07

import (
	"fmt"
	"testing"

	"github.com/cprass/adventofcode23/internal/utils"
)

func TestPart1(t *testing.T) {
	utils.Part1.RunTest(t, Runner{}, "6440")
}

func TestPart2(t *testing.T) {
	utils.Part2.RunTest(t, Runner{}, "5905")
}

func TestFindRecurringCardPart1(t *testing.T) {
	var hand Hand

	err := hand.addCardsFromString("23345", true)
	if err != nil {
		t.Error(err)
	}
	hand.sortCards()

	idx, card := hand.findRecurringCard(hand.sortedCards, 2, true)
	if idx != 2 {
		t.Errorf("Expected index 1, got %d", idx)
		return
	}
	if card.name != "3" {
		t.Errorf("Expected card name 3, got %s", card.name)
	}
}

func TestFindRecurringCardPart2(t *testing.T) {
	var hand Hand

	err := hand.addCardsFromString("2334J", false)
	if err != nil {
		t.Error(err)
	}
	hand.sortCards()

	idx, card := hand.findRecurringCard(hand.sortedCards, 3, false)
	if idx != 1 {
		t.Errorf("Expected index 1, got %d", idx)
		return
	}
	if card.name != "3" {
		t.Errorf("Expected card name 3, got %s", card.name)
	}

	err = hand.addCardsFromString("JJJ24", false)
	if err != nil {
		t.Error(err)
	}
	hand.sortCards()

	idx, card = hand.findRecurringCard(hand.sortedCards, 4, false)
	if idx != 0 {
		t.Errorf("Expected index 0, got %d", idx)
		fmt.Println(card.name)
		return
	}
	if card.name != "4" {
		t.Errorf("Expected card name 4, got %s", card.name)
	}
}

func TestFiveOfAKind(t *testing.T) {
	var hand Hand

	err := hand.addCardsFromString("QAQQ7", false)
	if err != nil {
		t.Error(err)
	}
	hand.sortCards()

	match, _ := hand.findFiveOfAKind(false)
	if match {
		t.Errorf("Expected not five of a kind, got %v", match)
	}
}
