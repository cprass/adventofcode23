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

	"github.com/cprass/adventofcode23/internal/utils"
)

func printErrAndExit(message string, err error) {
	fmt.Printf("%s%s\n", message, err)
	os.Exit(1)
}

func main() {

	inputFile := flag.String("i", "data.txt", "Input file path relative to the current directory")
	isPartOne := flag.Bool("p1", false, "If set, runs code for part one, otherwise runs code for part 2")
	useTestInput := flag.Bool("t", false, "Run program with test input")

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("./adventofcode23 -i data.txt -p1 1")
	}

	flag.Parse()

	day := flag.Arg(0)
	if day == "" {
		printErrAndExit("please specify the day number as command-line argument", nil)
	}
	dayPadded := fmt.Sprintf("%02s", day)

	dataFilePath := *inputFile
	if *useTestInput {
		dataFilePath = fmt.Sprintf("internal/day_%s/test.txt", dayPadded)
	}

	input, err := utils.LoadFile(dataFilePath)
	if err != nil {
		printErrAndExit("error loading data: ", err)
	}

	result, err := utils.RunModule(dayPadded, *isPartOne, input)
	if err != nil {
		printErrAndExit("error running module: ", err)
	}

	fmt.Printf("Result:\n%s\n", result)

	os.Exit(0)
}
