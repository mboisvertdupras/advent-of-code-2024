package main

import (
	day_one "advent-of-code/day-one"
	solution "advent-of-code/utils"
	"fmt"
)

func main() {
	days := []solution.Solution{
		day_one.Solution(),
	}

	for i, day := range days {
		fmt.Printf("Day %d:\n", i+1)
		fmt.Printf("  Part One: %d\n", day.PartOne)
		fmt.Printf("  Part Two: %d\n\n", day.PartTwo)
	}
}
