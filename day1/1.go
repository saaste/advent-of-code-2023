package day1

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2023/utils"
)

const (
	day1Test  bool   = false
	day1Input string = "day1/1_input.txt"
	day1Test1 string = "day1/1_test_1.txt"
	day1Test2 string = "day1/1_test_2.txt"
)

var writtenMap map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func First() {
	// 55017 is correct
	content, err := utils.ReadContent(day1Input, day1Test1, day1Test)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(content, "\n")
	r, _ := regexp.Compile(`\d{1}`)
	digits := make([]int64, 0)

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		results := r.FindAllString(line, -1)

		firstAndLast := results[0] + results[len(results)-1]
		digit, err := strconv.ParseInt(firstAndLast, 10, 64)
		if err != nil {
			log.Fatalf("Unable to parse string %s as int", firstAndLast)
		}
		digits = append(digits, digit)
	}

	var sum int64 = 0
	for _, digit := range digits {
		sum += digit
	}

	fmt.Printf("Day 1 / Step 1: %d\n", sum)
}

func Second() {
	content, err := utils.ReadContent(day1Input, day1Test2, day1Test)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(string(content), "\n")
	digits := make([]int64, 0)
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		matches := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		first := getFirst(line, matches)
		last := getLast(line, matches)
		firstAndLast := first + last

		digit, err := strconv.ParseInt(firstAndLast, 10, 64)
		if err != nil {
			log.Fatalf("Unable to parse string %s as int", firstAndLast)
		}
		digits = append(digits, digit)
	}

	var sum int64 = 0
	for _, digit := range digits {
		sum += digit
	}

	fmt.Printf("Day 1 / Step 2: %d\n", sum)
}

func getFirst(s string, matches []string) string {
	foundMatch := ""
	foundIndex := 999999999999
	for _, match := range matches {
		newIndex := strings.Index(s, match)
		if newIndex > -1 && newIndex < foundIndex {
			foundMatch = match
			foundIndex = newIndex
		}
	}

	digit, found := writtenMap[foundMatch]
	if found {
		return digit
	}

	return foundMatch
}

func getLast(s string, matches []string) string {
	foundMatch := ""
	foundIndex := -1
	for _, match := range matches {
		newIndex := strings.LastIndex(s, match)
		if newIndex > foundIndex {
			foundMatch = match
			foundIndex = newIndex
		}
	}

	digit, found := writtenMap[foundMatch]
	if found {
		return digit
	}

	return foundMatch
}
