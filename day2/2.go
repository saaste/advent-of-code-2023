package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2023/utils"
)

const (
	useTestInput bool   = false
	input        string = "day2/2_input.txt"
	testInput1   string = "day2/2_test_1.txt"
	testInput2   string = "day2/2_test_2.txt"
)

type round map[string]int64

type game struct {
	ID     int64
	Rounds []round
}

func First() {
	// 2720 is the correct answer
	content, err := utils.ReadContent(input, testInput1, useTestInput)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(content, "\n")
	games := parseGames(lines)

	maxRed := int64(12)
	maxGreen := int64(13)
	maxBlue := int64(14)

	possibleGames := make([]game, 0)

GAME:
	for _, game := range games {
		for _, round := range game.Rounds {
			if round["red"] > maxRed || round["green"] > maxGreen || round["blue"] > maxBlue {
				continue GAME
			}
		}
		possibleGames = append(possibleGames, game)
	}

	result := int64(0)
	for _, game := range possibleGames {
		result += game.ID
	}

	fmt.Printf("Day 2 / Step 1: %d\n", result)
}

func Second() {
	// 71535 is correct answer
	content, err := utils.ReadContent(input, testInput1, useTestInput)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(content, "\n")
	games := parseGames(lines)

	result := int64(0)
	for _, game := range games {
		minRed := int64(0)
		minGreen := int64(0)
		minBlue := int64(0)

		for _, round := range game.Rounds {
			if round["red"] > minRed {
				minRed = round["red"]
			}

			if round["green"] > minGreen {
				minGreen = round["green"]
			}

			if round["blue"] > minBlue {
				minBlue = round["blue"]
			}
		}
		result += minRed * minGreen * minBlue
	}

	fmt.Printf("Day 2 / Step 1: %d\n", result)
}

func parseGames(lines []string) []game {
	games := make([]game, 0)

	for _, line := range lines {
		if len(strings.TrimSpace(line)) < 1 {
			continue
		}
		parts := strings.Split(line, ": ")
		gameId, err := strconv.ParseInt(strings.Replace(parts[0], "Game ", "", -1), 10, 64)
		if err != nil {
			log.Fatalf("unable to parse game id: %+v, %v", parts, err)
		}

		rounds := make([]round, 0)

		parsedRounds := strings.Split(parts[1], "; ")
		for _, round := range parsedRounds {
			newRound := make(map[string]int64, 0)
			dice := strings.Split(round, ", ")
			for _, die := range dice {
				dieValues := strings.Split(die, " ")
				count, err := strconv.ParseInt(dieValues[0], 10, 64)
				if err != nil {
					log.Fatalf("unable to parse die count: %+v, %v", dieValues, err)
				}
				newRound[dieValues[1]] = count
			}
			rounds = append(rounds, newRound)
		}
		games = append(games, game{ID: gameId, Rounds: rounds})
	}
	return games
}
