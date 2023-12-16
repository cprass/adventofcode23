# Advent of Code 2023

This is the code I wrote for solving the [Advent of Code 2023](https://adventofcode.com/2023) puzzles, released under a GNU Affero General Public License v3.0 or later.

## Usage

The entrypoing is a single Go application. There are several flags that can be used to control the output. The program reads the puzzle input in the form text files. The puzzle input is expected at `./data.txt` or `./internal/day_{n}/test.txt` (for running the AoC test cases) per default, but the location can also be controlled via parameters.

Test inputs can also be split into `test1.txt` or `test2.txt` if there is different input for each test case.

You should always put your main puzzle input in `data.txt` at the root of the repo.

## CLI Flags

`-p1` - Run part one of the puzzle (default is to run part two)

`-i myData.txt` - Load data from `myData.txt`, located at the root of the repository (data is loaded from `data.txt` if not set)

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

## Debugging

I find it most effective to fix bugs using the debugger. There is a VSCode debugging setup, which has to be manually updated for test cases.
The `Launch` debug option uses arguments like `"args": ["-p1", "2"],` which have to be manually updated to change the puzzle number or switch from part two to part one.

## Testing

Tests provide a good way of doing TDD or just checking if everything is in order. You can use the following examples

```sh
# To run all tests
go test ./...

# To run only tests of a specific day
go test github.com/cprass/adventofcode23/internal/day_01

# To run only a specific test of a specific day
go test github.com/cprass/adventofcode23/internal/day_01 -run TestPart1
```

Typically there are two tests per module, `TestPart1` and `TestPart2`.

## Goals

1. Complete all AoC puzzles
2. Improve Golang skills

I certainly don't have time for every single puzzle before AoC is finished, but I might keep working on them afterwards (if it's possible to still continue at that point), eventually completing them all at some point.
