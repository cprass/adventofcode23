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

package utils

import (
	"fmt"
	"testing"
)

type Part int

const (
	Part1 Part = 1
	Part2 Part = 2
)

func (p Part) RunTest(t *testing.T, runner Runner, expected string) {
	input, err := LoadFile("test.txt")
	if err != nil {
		input, err = LoadFile(fmt.Sprintf("test%d.txt", p))
		if err != nil {
			t.Errorf("Test failed! Error: %s", err)
		}
	}

	isPartOne := p == Part(1)
	actual, err := runner.Run(input, isPartOne)
	if err != nil {
		t.Errorf("Test failed! Error: %s", err)
	}

	if actual != expected {
		t.Errorf("Test failed! Expected: %s, Actual: %s", expected, actual)
	}
}

func (p Part) Run(runner Runner, input []string) (string, error) {
	isPartOne := p == Part(1)
	return runner.Run(input, isPartOne)
}
