package day4

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/saaste/advent-of-code-2023/utils"
)

const (
	useTestInput bool   = false
	input        string = "day4/input.txt"
	testInput1   string = "day4/test.txt"
	testInput2   string = "day4/test.txt"
)

func First() {
	// 23941 is the correct answer
	content, err := utils.ReadContent(input, testInput1, useTestInput)
	if err != nil {
		log.Fatalf("%v", err)
	}
	cards := parseCards(content)
	points := 0
	for _, card := range cards {
		intersect := intersection(card.myNumbers, card.correct)
		fmt.Printf("Intersection: %v\n", intersect)
		cardPoints := 0
		if len(intersect) > 1 {
			for i := range intersect {
				if i == 0 {
					cardPoints = 1
				} else {
					cardPoints = cardPoints * 2
				}
			}
		} else if len(intersect) == 1 {
			cardPoints = 1
		} else {
			cardPoints = 0
		}

		fmt.Printf("Card point %d\n", cardPoints)
		points += cardPoints
	}

	fmt.Printf("Day 4 / Step 1: %d\n", points)
}

func Second() {
	// content, err := utils.ReadContent(input, testInput1, useTestInput)
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }

	fmt.Printf("Day 4 / Step 2: \n")
}

type Card struct {
	correct   []string
	myNumbers []string
}

func parseCards(content string) []Card {
	lines := strings.Split(content, "\n")
	lines = lines[0 : len(lines)-1]

	cards := make([]Card, 0)

	for _, line := range lines {
		indexOfColon := strings.Index(line, ":")
		line = line[indexOfColon+2:]
		// fmt.Printf("%s\n", line)

		space := regexp.MustCompile(`\s+`)
		line := space.ReplaceAllString(line, " ")

		parts := strings.Split(line, " | ")
		// fmt.Printf("%v / %v\n", parts[0], parts[1])

		correctAnswers := strings.Split(parts[0], " ")
		myNumbers := strings.Split(parts[1], " ")

		cards = append(cards, Card{correct: correctAnswers, myNumbers: myNumbers})
	}
	return cards
}

func intersection(a []string, b []string) []string {
	result := make([]string, 0)

	for _, x := range a {
		for _, y := range b {
			if x == y {
				result = append(result, x)
				break
			}
		}
	}
	return result
}
