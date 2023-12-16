package day_05_test

import (
	"testing"

	"github.com/cprass/adventofcode23/internal/utils"
)

func TestPart1(t *testing.T) {
	input, err := utils.LoadFile("test.txt")
	if err != nil {
		t.Errorf("Test failed! Error: %s", err)
	}

	actual, err := utils.RunModule("05", true, input)

	if err != nil {
		t.Errorf("Test failed! Error: %s", err)
	}

	expected := "35"
	if actual != expected {
		t.Errorf("Test failed! Expected: %s, Actual: %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.LoadFile("test.txt")
	if err != nil {
		t.Errorf("Test failed! Error: %s", err)
	}

	actual, err := utils.RunModule("05", false, input)

	if err != nil {
		t.Errorf("Test failed! Error: %s", err)
	}

	expected := "46"
	if actual != expected {
		t.Errorf("Test failed! Expected: %s, Actual: %s", expected, actual)
	}
}
