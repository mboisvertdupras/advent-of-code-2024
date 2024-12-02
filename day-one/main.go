package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() ([]int, []int, error) {
	// Create two slices to store the values
	var firstColumn []int
	var secondColumn []int

	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return firstColumn, secondColumn, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line
	for scanner.Scan() {
		// Split the line by whitespace
		values := strings.Fields(scanner.Text())

		// Convert first value to int and append to first slice
		if val1, err := strconv.Atoi(values[0]); err == nil {
			firstColumn = append(firstColumn, val1)
		}

		// Convert second value to int and append to second slice
		if val2, err := strconv.Atoi(values[1]); err == nil {
			secondColumn = append(secondColumn, val2)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return firstColumn, secondColumn, err
	}

	// Print the results
	return firstColumn, secondColumn, err
}

func main() {
	leftColumn, rightColumn, err := readInput()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	fmt.Println("\nSorted first column:", leftColumn)
	fmt.Println("Sorted second column:", rightColumn)
}
