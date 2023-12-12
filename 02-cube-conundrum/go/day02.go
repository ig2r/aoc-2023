package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Round holds the numbers of red, green and blue cubes drawn during a single round of a game.
type Round struct {
	red   int
	green int
	blue  int
}

// A Game consists of an ID and one or more rounds of drawing cubes.
type Game struct {
	id     int
	rounds []Round
}

var gameRegexp, cubeRegexp *regexp.Regexp

func parseRound(roundStr string) Round {
	cubeMatches := cubeRegexp.FindAllStringSubmatch(roundStr, -1)

	var round Round
	for _, matches := range cubeMatches {
		n, _ := strconv.Atoi(matches[1])
		color := matches[2]

		switch color {
		case "red":
			round.red += n
		case "green":
			round.green += n
		case "blue":
			round.blue += n
		}
	}

	return round
}

func parseGame(line string) Game {
	matches := gameRegexp.FindStringSubmatch(line)
	gameId, _ := strconv.Atoi(matches[1])

	roundsStr := strings.Split(matches[2], `; `)
	rounds := make([]Round, 0, len(roundsStr))

	for _, roundStr := range roundsStr {
		rounds = append(rounds, parseRound((roundStr)))
	}

	return Game{gameId, rounds}
}

func loadGames(fileName string) []Game {
	data, _ := os.ReadFile(fileName)
	lines := strings.Split(string(data), "\n")
	games := make([]Game, 0, len(lines))

	for _, line := range lines {
		games = append(games, parseGame(line))
	}

	return games
}

// Functions for puzzle part 1

func isValidRound(round Round) bool {
	return round.red <= 12 && round.green <= 13 && round.blue <= 14
}

func isValidGame(game Game) bool {
	for _, round := range game.rounds {
		if !isValidRound(round) {
			return false
		}
	}

	return true
}

func validGameSum(games []Game) int {
	var sum int

	for _, game := range games {
		if isValidGame(game) {
			sum += game.id
		}
	}

	return sum
}

// Functions for puzzle part 2

// minCubesNeeded determines the minimum number of red, green and blue cubes
// needed for the provided game.
func minCubesNeeded(game Game) (red int, green int, blue int) {
	var r, g, b int

	for _, round := range game.rounds {
		r = max(r, round.red)
		g = max(g, round.green)
		b = max(b, round.blue)
	}

	return r, g, b
}

func powerSetSum(games []Game) int {
	var sum int

	for _, game := range games {
		r, g, b := minCubesNeeded(game)
		sum += r * g * b
	}

	return sum
}

func main() {
	gameRegexp = regexp.MustCompile(`Game (\d+): (.+)`)
	cubeRegexp = regexp.MustCompile(`(\d+) (red|green|blue)`)

	games := loadGames("input.txt")
	fmt.Println(validGameSum(games))
	fmt.Println(powerSetSum(games))
}
