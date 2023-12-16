// SPDX-License-Identifier: AGPL-3.0-or-later

package day_02

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Runner struct{}

type GameSet struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id   int
	sets []GameSet
}

func parseLine(line string) (*Game, error) {
	reHead := regexp.MustCompile(`Game (\d+):\s+`)
	reColor := regexp.MustCompile(`(\d+)\s+(\w+)`)

	gameIdMatch := reHead.FindStringSubmatch(line)
	if len(gameIdMatch) != 2 {
		return nil, errors.New("unexpected match for head")
	}
	gameId, err := strconv.Atoi(gameIdMatch[1])
	if err != nil {
		return nil, err
	}

	var gameSets []GameSet
	game := Game{id: gameId}

	line2 := line[len(gameIdMatch)+6:]
	sets := strings.Split(line2, ";")
	for _, set := range sets {
		gameSet := GameSet{}
		colors := strings.Split(set, ",")
		for _, color := range colors {

			colorCountMatch := reColor.FindStringSubmatch(color)
			if len(colorCountMatch) != 3 {
				return nil, errors.New("unexpected match for color")
			}

			colorCount, err := strconv.Atoi(colorCountMatch[1])
			if err != nil {
				return nil, err
			}

			switch colorCountMatch[2] {
			case "red":
				gameSet.red = colorCount
			case "green":
				gameSet.green = colorCount
			case "blue":
				gameSet.blue = colorCount
			}
		}

		gameSets = append(gameSets, gameSet)
	}

	game.sets = gameSets

	return &game, nil
}

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func isValidGame(game *Game) bool {
	for _, set := range game.sets {
		if set.blue > MAX_BLUE || set.green > MAX_GREEN || set.red > MAX_RED {
			return false
		}
	}
	return true
}

func (r Runner) Run(input []string, isPartOne bool) (string, error) {
	var games []*Game
	for _, line := range input {
		game, err := parseLine(line)
		if err != nil {
			return "", err
		}

		if (isPartOne && isValidGame(game)) || !isPartOne {
			games = append(games, game)
		}
	}

	sum := 0
	sumPowers := 0
	for _, game := range games {
		sum += game.id

		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, set := range game.sets {
			if set.red > minRed {
				minRed = set.red
			}
			if set.green > minGreen {
				minGreen = set.green
			}
			if set.blue > minBlue {
				minBlue = set.blue
			}
		}

		power := minRed * minBlue * minGreen
		sumPowers += power
	}

	if isPartOne {
		return strconv.Itoa(sum), nil
	}

	return strconv.Itoa(sumPowers), nil
}
