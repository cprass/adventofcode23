// SPDX-License-Identifier: AGPL-3.0-or-later

package day_01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Runner struct{}

func convertLiteral(s string) string {
	s = strings.ToLower(s)

	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}

	return s
}

func concatToInt(isPartOne bool, s1 string, s2 string) int {
	if !isPartOne {
		s1 = convertLiteral(s1)
		s2 = convertLiteral(s2)
	}
	i, err := strconv.Atoi(fmt.Sprintf("%s%s", s1, s2))
	if err != nil {
		return 0
	}
	return i
}

func (r Runner) Run(input []string, isPartOne bool) (string, error) {
	var re *regexp.Regexp
	if isPartOne {
		re = regexp.MustCompile(`(\d)`)
	} else {
		re = regexp.MustCompile(`(?i)(\d|one|two|three|four|five|six|seven|eight|nine)`)
	}

	sum := 0

	for _, line := range input {
		first := re.FindStringIndex(line)

		if first == nil {
			continue
		}

		firstStr := line[first[0]:first[1]]

		var last [2]int
		last[0] = first[0]
		last[1] = first[1]
		offset := first[0] + 1

		for offset < len(line) {
			next := re.FindStringIndex(line[offset:])
			if next == nil {
				break
			}
			last[0] = next[0] + offset
			last[1] = next[1] + offset
			offset += next[0] + 1
		}

		secondStr := line[last[0]:last[1]]

		res := concatToInt(isPartOne, firstStr, secondStr)
		sum += res
	}

	return strconv.Itoa(sum), nil
}
