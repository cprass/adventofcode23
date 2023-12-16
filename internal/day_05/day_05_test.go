// SPDX-License-Identifier: AGPL-3.0-or-later

package day_05

import (
	"testing"

	"github.com/cprass/adventofcode23/internal/utils"
)

func TestPart1(t *testing.T) {
	utils.Part1.RunTest(t, Runner{}, "35")
}

func TestPart2(t *testing.T) {
	utils.Part2.RunTest(t, Runner{}, "46")
}
