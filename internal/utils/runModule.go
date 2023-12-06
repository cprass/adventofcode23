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
	"errors"
	"fmt"

	"github.com/cprass/adventofcode23/internal/day_01"
	"github.com/cprass/adventofcode23/internal/day_02"
)

func RunModule(day string, isPartOne bool, input []string) (string, error) {
	dayPadded := fmt.Sprintf("%02s", day)

	switch dayPadded {
	case "01":
		return day_01.Run(input, isPartOne)
	case "02":
		return day_02.Run(input, isPartOne)
	}

	return "", errors.New("requested module not found")
}
