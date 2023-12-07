# Advent of Code 2023

This is the code I wrote for solving the [Advent of Code 2023](https://adventofcode.com/2023) puzzles, released under a GNU Affero General Public License v3.0 or later.

## Usage

The entrypoing is a single Go application. There are several flags that can be used to control the output. The program reads the puzzle input in the form text files. The text input is expected at `./data.txt` or `./internal/day_{n}/test.txt` (for running the AoC test cases) per default, but the location can also be controlled via parameters.

This might look weird at first, but it is a good setup for actively working on puzzles. You can have your main puzzle input in `data.txt` while switching to the example data in `test.txt` without any issues.

## CLI Flags

`-p1` - Run part one of the puzzle (default is to run part two)

`-i myData.txt` - Load data from `myData.txt`, located at the root of the repository (data is loaded from `data.txt` if not set)

`-t` - Ignore data input and load data from `./internal/day_{n}/test.txt` instead, located at the root of the repository

`<int>`

## Parameters

The app takes a single positive integer as parameter, that determines what puzzle to load.

## Examples

```sh
go run cmd/adventofcode23/main.go 2
```

Executes **part two** of the **day two** puzzle using **`<repo-root>/data.txt`** as puzzle input.

---

```sh
go run cmd/adventofcode23/main.go -i ./foo/bar.txt -p1 1
```

Executes **part one** of the **day one** puzzle using **`<repo-root>/foo/bar.txt`** as puzzle input.

---

```sh
go run cmd/adventofcode23/main.go -t 2
```

Executes **part two** of the **day two** puzzle using the test puzzle input.

## Debugging

I find it most effective to come up with solutions using the debugger, as long as there isn't any TDD setup. There is a VSCode debugging setup, which has to be manually updated for test cases.

E.g. the `Launch tests` debug option uses `"args": ["-t", "2"],` which has to be manually updated to change the puzzle number or switch from part two to part one.

## Goals

1. Complete all AoC puzzles
2. Improve Golang skills

I certainly don't have time for every single puzzle before AoC is finished, but I might keep working on them afterwards (if it's possible to still continue at that point), eventually completing them all at some point.

## TODO

- add unit tests and solve puzzles using TDD, instead of using a command line parameter to switch to test data.
