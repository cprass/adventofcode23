// SPDX-License-Identifier: AGPL-3.0-or-later

package day_06

import (
	"math"
	"strconv"
	"strings"
)

type Runner struct{}

type Game struct {
	recordTime int
	distance   int
}

func solveQuadratic(a, b, c int) (float64, float64) {
	sqrtPart := math.Sqrt(float64(b*b - 4*a*c))
	res1 := (-float64(b) + sqrtPart) / (2.0 * float64(a))
	res2 := (-float64(b) - sqrtPart) / (2.0 * float64(a))

	return res1, res2
}

func (g *Game) waysToWin() int {
	// The exact function to calculate the charge time to beat the record is
	// y = x * RECORD_TIME - x² - RECORD_DISTANCE
	// Solving that quadratic formula for y = 0 gives us all the points between which winning is possible

	res1, res2 := solveQuadratic(-1, g.recordTime, -g.distance)

	// We have to round up the lower bound and round down the upper bound
	// But if the result is a whole number, we have to add or subtract 1
	if math.Round(res1) == res1 {
		res1++
	}
	if math.Round(res2) == res2 {
		res2--
	}

	return int(math.Floor(res2)) - int(math.Ceil(res1)) + 1
}

func (r Runner) Run(input []string, isPartOne bool) (string, error) {
	var games []Game
	fieldsTime := strings.Fields(input[0])[1:]
	fieldsDistance := strings.Fields(input[1])[1:]

	if isPartOne {
		for i, field := range fieldsTime {
			time, err := strconv.Atoi(field)
			if err != nil {
				return "", err
			}
			distance, err := strconv.Atoi(fieldsDistance[i])
			if err != nil {
				return "", err
			}
			games = append(games, Game{time, distance})
		}
	} else {
		timeStr := strings.Join(fieldsTime, "")
		distanceStr := strings.Join(fieldsDistance, "")
		time, err := strconv.Atoi(timeStr)
		if err != nil {
			return "", err
		}
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			return "", err
		}
		games = append(games, Game{time, distance})
	}

	product := 1
	for _, game := range games {
		product *= game.waysToWin()
	}

	return strconv.Itoa(product), nil
}
