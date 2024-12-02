package main

import (
	day_one "advent-of-code/day-one"
	solution "advent-of-code/utils"
	"fmt"
)

type Day struct {
	Solutions []solution.Solution
}

func main() {
	days := Day{
		day_one.Solution(),
	}

	for i, day := range days {
		fmt.Sprintf("Day %d:\n\nPart One: %v\n\nPart Two: %v", i, day.PartOne, day.PartTwo)
	}
}
