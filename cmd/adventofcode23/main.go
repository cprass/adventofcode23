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

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/cprass/adventofcode23/internal/day_01"
	"github.com/cprass/adventofcode23/internal/day_02"
	"github.com/cprass/adventofcode23/internal/day_03"
	"github.com/cprass/adventofcode23/internal/day_04"
	"github.com/cprass/adventofcode23/internal/day_05"
	"github.com/cprass/adventofcode23/internal/utils"
)

// Create a map of runners
var runners = map[int]utils.Runner{
	1: day_01.Runner{},
	2: day_02.Runner{},
	3: day_03.Runner{},
	4: day_04.Runner{},
	5: day_05.Runner{},
}

func printErrAndExit(message string, err error) {
	m := message
	if err != nil {
		m = fmt.Sprintf("%s: %s", message, err)
	}
	fmt.Printf("%s\n", m)
	os.Exit(1)
}

func main() {
	var err error
	inputFile := flag.String("i", "data.txt", "Input file path relative to the current directory")
	isPartOne := flag.Bool("p1", false, "If set, runs code for part one, otherwise runs code for part 2")

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("./adventofcode23 -i data.txt -p1 1")
	}

	flag.Parse()

	dayArg := flag.Arg(0)
	if dayArg == "" {
		printErrAndExit("please specify the day number as command-line argument", nil)
	}
	dayNum, err := strconv.Atoi(dayArg)
	if err != nil {
		printErrAndExit("error parsing day number: ", err)
	}
	input, err := utils.LoadFile(*inputFile)
	if err != nil {
		printErrAndExit("error loading data: ", err)
	}

	runner, ok := runners[dayNum]
	if !ok {
		printErrAndExit(fmt.Sprintf("invalid day number: %d", dayNum), nil)
	}

	var part utils.Part
	if *isPartOne {
		part = utils.Part1
	} else {
		part = utils.Part2
	}

	result, err := part.Run(runner, input)

	if err != nil {
		printErrAndExit("error running module: ", err)
	}

	fmt.Printf("Running part %d of day %d\n", part, dayNum)
	fmt.Printf("https://adventofcode.com/2023/day/%d\n", dayNum)
	fmt.Printf("Result:\n%s\n", result)

	os.Exit(0)
}
