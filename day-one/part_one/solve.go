package part_one

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFile = "day-one/input.txt"

type Columns [][]int

// Validate ensures all columns have the same length
func (c Columns) Validate() error {
	if len(c) == 0 {
		return fmt.Errorf("no columns present")
	}

	expectedLen := len(c[0])
	for i := 1; i < len(c); i++ {
		if len(c[i]) != expectedLen {
			return fmt.Errorf("column %d has length %d, expected %d",
				i, len(c[i]), expectedLen)
		}
	}
	return nil
}

// Sort sorts all columns
func (c Columns) Sort() {
	for i := range c {
		sort.Ints(c[i])
	}
}

// Helper function to count lines
func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("opening file for counting: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("counting lines: %w", err)
	}

	return lineCount, nil
}

func getDistance(a, b []int) int {
	totalDistance := 0
	for i := 0; i < len(a); i++ {
		totalDistance += int(math.Abs(float64(a[i] - b[i])))
	}
	return totalDistance
}

func readInput() (Columns, error) {
	// Open the file
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	// Count lines for pre-allocation
	lineCount, err := countLines(inputFile)
	if err != nil {
		return nil, fmt.Errorf("counting lines: %w", err)
	}

	// Initialize columns slice with 2 pre-allocated slices
	cols := make(Columns, 2)
	for i := range cols {
		cols[i] = make([]int, 0, lineCount)
	}

	// Create scanner for reading the file
	scanner := bufio.NewScanner(file)

	// Process each line
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		if len(values) != 2 {
			return nil, fmt.Errorf("invalid format: expected 2 values, got %d",
				len(values))
		}

		// Process each column value
		for i, val := range values {
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid number, column %d: %w",
					i+1, err)
			}
			cols[i] = append(cols[i], num)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning file: %w", err)
	}

	// Validate columns have equal length
	if err := cols.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Sort all columns
	cols.Sort()

	return cols, nil
}

func Solve() int {
	columns, err := readInput()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	distance := getDistance(columns[0], columns[1])
	return distance
}
