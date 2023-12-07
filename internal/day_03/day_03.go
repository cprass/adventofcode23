// Copyright (C) 2023  C. Praß <https://github.com/cprass>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package day_03

import (
	"regexp"
	"strconv"
)

type GearRatios struct {
	rows int
	cols int

	lines []string

	rInts    *regexp.Regexp
	rSymbols *regexp.Regexp
	rGears   *regexp.Regexp
}

func (g *GearRatios) findNumberAt(row int, col int) (int, error) {
	intMatchesInLine := g.rInts.FindAllStringIndex(g.lines[row], -1)
	for _, match := range intMatchesInLine {
		if col >= match[0] && col < match[1] {
			num, err := strconv.Atoi(g.lines[row][match[0]:match[1]])
			if err != nil {
				return 0, err
			}
			return num, nil
		}
	}

	return 0, nil
}

// If the gear has two numbers next to it, return the gear ratio
func (g *GearRatios) getGearRatio(row int, col int) (int, error) {
	var list []int
	leftIdx := max(0, col-1)
	rightIdx := min(g.cols-1, col+1)

	if row > 0 {
		matches := g.rInts.FindAllStringIndex(g.lines[row-1][leftIdx:rightIdx+1], -1)
		for _, match := range matches {
			n, err := g.findNumberAt(row-1, match[0]+leftIdx)
			if err != nil {
				return 0, err
			}
			list = append(list, n)
		}
	}
	if row < g.rows-1 {
		matches := g.rInts.FindAllStringIndex(g.lines[row+1][leftIdx:rightIdx+1], -1)
		for _, match := range matches {
			n, err := g.findNumberAt(row+1, match[0]+leftIdx)
			if err != nil {
				return 0, err
			}
			list = append(list, n)
		}
	}
	matches := g.rInts.FindAllStringIndex(g.lines[row][leftIdx:rightIdx+1], -1)
	for _, match := range matches {
		n, err := g.findNumberAt(row, match[0]+leftIdx)
		if err != nil {
			return 0, err
		}
		list = append(list, n)
	}

	if len(list) == 2 {
		return list[0] * list[1], nil
	}

	return 0, nil
}

func hasNearbySymbol(r *regexp.Regexp, w []string) bool {
	for _, l := range w {
		if r.MatchString(l) {
			return true
		}
	}
	return false
}

func Run(input []string, isPartOne bool) (string, error) {
	g := GearRatios{
		rInts:    regexp.MustCompile(`\d+`),
		rSymbols: regexp.MustCompile(`[^\d\.]`),
		rGears:   regexp.MustCompile(`\*`),

		rows:  len(input),
		cols:  len(input[0]),
		lines: input,
	}

	sumP1 := 0
	sumP2 := 0

	for row, line := range input {

		if isPartOne {
			hasPrev := row > 0
			hasNext := row < len(input)-1

			allMatchesInLine := g.rInts.FindAllStringIndex(line, -1)

			for _, match := range allMatchesInLine {
				prevIdx := max(0, match[0]-1)
				nextIdx := min(len(line)-1, match[1]+1)

				var w []string
				w = append(w, line[prevIdx:nextIdx])
				if hasPrev {
					w = append(w, input[row-1][prevIdx:nextIdx])
				}
				if hasNext {
					w = append(w, input[row+1][prevIdx:nextIdx])
				}

				hasSymbol := hasNearbySymbol(g.rSymbols, w)
				if !hasSymbol {
					continue
				}

				num, err := strconv.Atoi(line[match[0]:match[1]])
				if err != nil {
					return "", err
				}

				sumP1 += num
			}
		}

		gearMatches := g.rGears.FindAllStringIndex(line, -1)
		for _, m := range gearMatches {
			col := m[0]
			gearRatio, err := g.getGearRatio(row, col)
			if err != nil {
				return "", err
			}
			sumP2 += gearRatio
		}
	}

	if isPartOne {
		return strconv.Itoa(sumP1), nil
	}

	return strconv.Itoa(sumP2), nil
}
