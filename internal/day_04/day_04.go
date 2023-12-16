// SPDX-License-Identifier: AGPL-3.0-or-later

package day_04

import (
	"math"
	"strconv"
	"strings"
)

type Runner struct{}

// calculates how many of the winning numbers are contained in the existing numbers
func nWinningNumbers(winningNumbers []int, numbers []int) int {
	matchingNumbers := 0
	// go through all winning numbers and check if it's contained in the numbers list
	for _, wn := range winningNumbers {
		for _, n := range numbers {
			if wn == n {
				matchingNumbers += 1
				// end the inner loop, since we only have to find each match once
				break
			}
		}
	}
	return matchingNumbers
}

func (r Runner) Run(input []string, isPartOne bool) (string, error) {
	// Part one is the sum of the card points
	// Part two is the sum of total scratchcards including copies
	sum := 0

	// Maps the card number (as array index) to the number of copies of that card
	var copies []int
	if !isPartOne {
		copies = make([]int, len(input))
	}

	for currIdx, line := range input {
		parts := strings.Split(line, " ")

		// Build list of winning numbers and actual numbers by looping through the parts of the string
		var winningNumbers []int
		var numbers []int
		isWinningNumbers := true

		for _, part := range parts {
			if part == "" || part == "Card" || part[len(part)-1:] == ":" {
				continue
			}
			if part == "|" {
				isWinningNumbers = false
				continue
			}

			strAsInt, err := strconv.Atoi(part)
			if err != nil {
				return "", err
			}
			if isWinningNumbers {
				winningNumbers = append(winningNumbers, strAsInt)
			} else {
				numbers = append(numbers, strAsInt)
			}
		}

		nWinning := nWinningNumbers(winningNumbers, numbers)

		if isPartOne {
			if nWinning == 0 {
				continue
			}
			sum += int(math.Pow(2, float64(nWinning-1)))
			continue
		}

		// We can only obtain copies of valid next cards, so we have to get rid of any overflowing numbers
		nCopies := min(len(input)-currIdx-1, nWinning)

		for i := 1; i <= nCopies; i++ {
			idx := currIdx + i
			// We add copies for each of the current copies
			copies[idx] += copies[currIdx] + 1
		}
	}

	if !isPartOne {
		// sum up all of the copies
		for _, nCopies := range copies {
			sum += nCopies
		}
		// add the number of base cards without the copies
		sum += len(input)
	}

	return strconv.Itoa(sum), nil
}
