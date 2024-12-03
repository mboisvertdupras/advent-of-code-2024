package day_one

import (
	"advent-of-code/day-one/part_one"
	solution "advent-of-code/utils"
)

func Solution() solution.Solution {
	solutions := solution.Solution{
		PartOne: part_one.Solve(),
	}

	return solutions
}
