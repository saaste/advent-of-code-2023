package day3

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/saaste/advent-of-code-2023/utils"
)

const (
	useTestInput bool   = false
	input        string = "day3/input.txt"
	testInput1   string = "day3/test.txt"
	testInput2   string = "day3/test.txt"
)

type Digit struct {
	value         int
	rowIndex      int
	colStartIndex int
	colEndIndex   int
}

func First() {
	// 539433 is the correct answer
	content, err := utils.ReadContent(input, testInput1, useTestInput)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(content, "\n")
	lines = lines[0 : len(lines)-1]
	digits := findDigits(lines)

	// validDigits := make([]Digit, 0)
	result := 0
	for _, digit := range digits {
		if isPartNumber(digit, lines) {
			// validDigits = append(validDigits, digit)
			result += digit.value
		}
	}

	fmt.Printf("Day 3 / Step 1: %d\n", result)
}

func findDigits(lines []string) []Digit {
	maxColIndex := len(lines[0]) - 1
	digits := make([]Digit, 0)

	for rowIndex, line := range lines {
		if len(line) == 0 {
			continue
		}
		colIndex := 0
		for colIndex <= maxColIndex {
			char := rune(line[colIndex])
			if unicode.IsDigit(char) {
				newDigit := string(char)
				startIndex := colIndex
				for colIndex+1 <= maxColIndex && unicode.IsDigit(rune(line[colIndex+1])) {
					newDigit += string(line[colIndex+1])
					colIndex++
				}
				endIndex := colIndex
				value, err := strconv.ParseInt(newDigit, 10, 32)
				if err != nil {
					log.Fatalf("failed to parse %s to integer: %v", newDigit, err)
				}
				digits = append(digits, Digit{value: int(value), rowIndex: rowIndex, colStartIndex: startIndex, colEndIndex: endIndex})
			}
			colIndex++
		}
	}

	return digits
}

func isPartNumber(d Digit, lines []string) bool {
	maxColIndex := len(lines[0]) - 1
	maxRowIndex := len(lines) - 1
	// Check left
	if d.colStartIndex > 0 {
		if getCharAt(lines, d.rowIndex, d.colStartIndex-1) != "." {
			return true
		}
	}

	// Check right
	if d.colEndIndex < maxColIndex {
		if getCharAt(lines, d.rowIndex, d.colEndIndex+1) != "." {
			return true
		}
	}

	// Check above
	if d.rowIndex > 0 {
		for i := d.colStartIndex; i <= d.colEndIndex; i++ {
			if getCharAt(lines, d.rowIndex-1, i) != "." {
				return true
			}
		}
	}

	// Check below
	if d.rowIndex < maxRowIndex {
		for i := d.colStartIndex; i <= d.colEndIndex; i++ {
			if getCharAt(lines, d.rowIndex+1, i) != "." {
				return true
			}
		}
	}

	// Check top left and right
	if d.rowIndex > 0 {
		if d.colStartIndex > 0 && getCharAt(lines, d.rowIndex-1, d.colStartIndex-1) != "." {
			return true
		}
		if d.colEndIndex < maxColIndex && getCharAt(lines, d.rowIndex-1, d.colEndIndex+1) != "." {
			return true
		}
	}

	// Check bottom left and right
	if d.rowIndex < maxRowIndex {
		if d.colStartIndex > 0 && getCharAt(lines, d.rowIndex+1, d.colStartIndex-1) != "." {
			return true
		}
		if d.colEndIndex < maxColIndex && getCharAt(lines, d.rowIndex+1, d.colEndIndex+1) != "." {
			return true
		}
	}

	return false
}

func getCharAt(lines []string, rowIndex int, colIndex int) string {
	return string(lines[rowIndex][colIndex])
}

func Second() {
	// 75847567 is correct answer
	content, err := utils.ReadContent(input, testInput1, useTestInput)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(content, "\n")
	lines = lines[0 : len(lines)-1]

	digits := findDigits(lines)
	gearRatioSum := 0
	for rowIndex, row := range lines {
		for colIndex, char := range row {
			if string(char) == "*" {
				adjacentDigits := findAdjacentDigits(rowIndex, colIndex, digits)
				if len(adjacentDigits) == 2 {
					// fmt.Printf("Gear on row %d and col %d is valid with digits %d and %d\n", rowIndex, colIndex, adjacentDigits[0].value, adjacentDigits[1].value)
					gearRatioSum += adjacentDigits[0].value * adjacentDigits[1].value
				}
			}
		}
	}

	fmt.Printf("Day 3 / Step 2: %d\n", gearRatioSum)
}

func findAdjacentDigits(rowIndex int, colIndex int, digits []Digit) []Digit {
	adjacentDigits := make([]Digit, 0)
	for _, digit := range digits {
		// Check left
		if digit.rowIndex == rowIndex && digit.colEndIndex == colIndex-1 {
			adjacentDigits = append(adjacentDigits, digit)
			continue
		}
		// Check right
		if digit.rowIndex == rowIndex && digit.colStartIndex == colIndex+1 {
			adjacentDigits = append(adjacentDigits, digit)
			continue
		}

		// Check top + adjacent
		if digit.rowIndex == rowIndex-1 {
			if digit.colStartIndex >= colIndex-1 && digit.colStartIndex <= colIndex+1 {
				adjacentDigits = append(adjacentDigits, digit)
				continue
			}
			if digit.colEndIndex >= colIndex-1 && digit.colEndIndex <= colIndex+1 {
				adjacentDigits = append(adjacentDigits, digit)
				continue
			}
		}

		// Check bottom + adjacent
		if digit.rowIndex == rowIndex+1 {
			if digit.colStartIndex >= colIndex-1 && digit.colStartIndex <= colIndex+1 {
				adjacentDigits = append(adjacentDigits, digit)
				continue
			}
			if digit.colEndIndex >= colIndex-1 && digit.colEndIndex <= colIndex+1 {
				adjacentDigits = append(adjacentDigits, digit)
				continue
			}
		}
	}
	return adjacentDigits
}
