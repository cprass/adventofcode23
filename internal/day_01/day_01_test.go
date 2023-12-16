package day_01

import (
	"testing"

	"github.com/cprass/adventofcode23/internal/utils"
)

func TestPart1(t *testing.T) {
	utils.Part1.RunTest(t, Runner{}, "142")
}

func TestPart2(t *testing.T) {
	utils.Part2.RunTest(t, Runner{}, "281")
}
