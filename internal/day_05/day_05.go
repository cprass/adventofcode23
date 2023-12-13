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

package day_05

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type MapItem struct {
	destStart int
	srcStart  int
	length    int
}

type Map struct {
	srcType  string
	destType string
	items    []MapItem
}

type seedMaps = map[string]Map
type seedDefs = [][2]int

func genMap(lines []string) (*Map, error) {
	mapName := strings.Split(strings.Split(lines[0], " ")[0], "-to-")

	if len(mapName) != 2 {
		return nil, errors.New("could not parse map source and dest types")
	}

	srcType := mapName[0]
	destType := mapName[1]

	// create MapItems from the lines of strings
	var items []MapItem
	for _, line := range lines[1:] {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			return nil, errors.New("each map entry must have exactly three numbers divided by spaces")
		}
		destStart, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		srcStart, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		length, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		items = append(items, MapItem{
			destStart: destStart,
			srcStart:  srcStart,
			length:    length,
		})
	}

	// sort each of the item list by source-start value
	sort.Slice(items, func(i, j int) bool {
		return items[i].srcStart < items[j].srcStart
	})

	// add an item to the beginning, where src and dest matches
	if items[0].srcStart-1 >= 0 {
		firstItem := MapItem{
			destStart: 0,
			srcStart:  0,
			length:    items[0].srcStart,
		}
		var newItems []MapItem
		newItems = append(newItems, firstItem)
		items = append(newItems, items...)
	}

	result := Map{
		srcType:  srcType,
		destType: destType,
		items:    items,
	}

	return &result, nil
}

// recursive function that calculates the lowest location in the given range
func getLowestLocation(maps seedMaps, mapType string, seedStart int, seedRange int) (int, error) {
	currentMap := maps[mapType]
	lowestLocation := 0
	var err error

	// find a matching seed-map-item
	matchDestStart := 0
	matchDestRange := 0

	for _, item := range currentMap.items {
		if seedStart >= item.srcStart && seedStart < item.srcStart+item.length {
			diffStart := seedStart - item.srcStart
			matchDestStart = item.destStart + diffStart
			matchDestRange = min(item.length-diffStart, seedRange)
			break
		}
	}

	if matchDestRange == 0 {
		matchDestStart = seedStart
		matchDestRange = seedRange
	}

	if currentMap.destType == "location" {
		lowestLocation = matchDestStart
	} else {
		lowestLocation, err = getLowestLocation(maps, currentMap.destType, matchDestStart, matchDestRange)
		if err != nil {
			return 0, err
		}
	}

	remainingSeeds := seedRange - matchDestRange
	if remainingSeeds != 0 {
		// if not all of the seeds matched the current destination, return the min of the next lowest
		// location and the current lowest location
		nextLowestLocation, err := getLowestLocation(maps, mapType, seedStart+matchDestRange, remainingSeeds)
		if err != nil {
			return 0, err
		}
		return min(nextLowestLocation, lowestLocation), nil
	}

	return lowestLocation, nil
}

func parseSeeds(inputLine string, isPartOne bool) (seedDefs, error) {
	var seeds seedDefs
	values := strings.Split(inputLine[7:], " ")
	for i := 0; i < len(values)-1; i += 2 {
		valOne, err := strconv.Atoi(values[i])
		if err != nil {
			return nil, err
		}
		valTwo, err := strconv.Atoi(values[i+1])
		if err != nil {
			return nil, err
		}
		if isPartOne {
			// in part one we count each seed value as single seed definition
			seeds = append(seeds, [2]int{valOne, 1}, [2]int{valTwo, 1})
		} else {
			// in part two the values are tuples with start and range
			s := [2]int{valOne, valTwo}
			seeds = append(seeds, s)
		}
	}
	return seeds, nil
}

func Run(input []string, isPartOne bool) (string, error) {
	var blocks [][]string

	// Collect each line in different blocks, divided by empty lines
	i := 0
	for _, line := range input {
		if line == "" {
			i++
		} else {
			if len(blocks) < i+1 {
				blocks = append(blocks, []string{line})
			} else {
				blocks[i] = append(blocks[i], line)
			}
		}
	}

	maps := make(seedMaps)
	var seeds seedDefs
	var err error

	for _, block := range blocks {
		// Parse seed part
		if len(block) == 1 {

			seeds, err = parseSeeds(block[0], isPartOne)
			if err != nil {
				return "", err
			}
			continue
		}

		sMap, err := genMap(block)
		if err != nil {
			return "", err
		}
		maps[sMap.srcType] = *sMap
	}

	lowest := 0
	for _, seedVals := range seeds {
		newLowest, err := getLowestLocation(maps, "seed", seedVals[0], seedVals[1])
		if err != nil {
			return "", err
		}
		if lowest == 0 || newLowest < lowest {
			lowest = newLowest
		}
	}
	return strconv.Itoa(lowest), nil
}
