package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	// To solve part 1, use these regular expressions:
	// firstTokenRegex := regexp.MustCompile(`(\d)`)
	// lastTokenRegex := regexp.MustCompile(`.*(\d)`)

	// To solve part 2, use these instead:
	firstTokenRegex := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	lastTokenRegex := regexp.MustCompile(`.*(\d|one|two|three|four|five|six|seven|eight|nine)`)

	tokenValues := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	calibrationValues := 0

	for _, line := range lines {
		firstToken := firstTokenRegex.FindStringSubmatch(line)[1]
		lastToken := lastTokenRegex.FindStringSubmatch(line)[1]
		calibrationValues += 10*tokenValues[firstToken] + tokenValues[lastToken]
	}

	fmt.Println(calibrationValues)
}
