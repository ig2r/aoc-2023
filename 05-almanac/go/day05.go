package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start  int
	Length int
}

type Mapping struct {
	Dest   int
	Source int
	Length int
}

type Section = []Mapping

func getSeedNumbers(lines []string) []int {
	line := strings.TrimPrefix(lines[0], "seeds: ")
	components := strings.Split(line, " ")

	seedNumbers := make([]int, len(components))
	for i, c := range components {
		seedNumbers[i], _ = strconv.Atoi(c)
	}

	return seedNumbers
}

func getSeedRangesForPuzzlePart1(seedNumbers []int) []Range {
	ranges := make([]Range, len(seedNumbers))
	for i, seed := range seedNumbers {
		ranges[i] = Range{seed, 0}
	}

	return ranges
}

func getSeedRangesForPuzzlePart2(seedNumbers []int) []Range {
	ranges := make([]Range, len(seedNumbers)/2)
	for i := 0; i < len(ranges); i++ {
		ranges[i] = Range{seedNumbers[i*2], seedNumbers[i*2+1]}
	}

	return ranges
}

func parseMapping(line string) Mapping {
	components := strings.Split(line, " ")
	dest, _ := strconv.Atoi(components[0])
	src, _ := strconv.Atoi(components[1])
	size, _ := strconv.Atoi(components[2])
	return Mapping{dest, src, size}
}

func parseSections(lines []string) []Section {
	sections := make([]Section, 0)
	var section Section

	for _, line := range lines[2:] {
		if strings.Contains(line, ":") {
			section = Section{}
		} else if line == "" {
			sections = append(sections, section)
			section = nil
		} else {
			section = append(section, parseMapping(line))
		}
	}

	if section != nil {
		sections = append(sections, section)
	}

	return sections
}

func mapRange(r Range, section Section) Range {
	for _, mapping := range section {
		if r.Start >= mapping.Source && r.Start < mapping.Source+mapping.Length {
			return Range{mapping.Dest - mapping.Source + r.Start, r.Length}
		}
	}

	return r
}

func getSplitPoints(section Section) []int {
	results := make([]int, 0, len(section)*2)

	for _, m := range section {
		if !slices.Contains(results, m.Source) {
			results = append(results, m.Source)
		}

		if !slices.Contains(results, m.Source+m.Length) {
			results = append(results, m.Source+m.Length)
		}
	}

	slices.Sort(results)
	return results
}

func splitRanges(rs []Range, splitPoints []int) []Range {
	results := make([]Range, 0, len(rs))

	for _, r := range rs {
		for _, splitPoint := range splitPoints {
			if splitPoint >= r.Start && splitPoint < r.Start+r.Length {
				results = append(results, Range{r.Start, splitPoint - r.Start})
				r = Range{splitPoint, r.Length - (splitPoint - r.Start)}
			}
		}

		results = append(results, r)
	}

	return results
}

// Single translation step for a list of ranges using a single section
func transformRanges(rs []Range, section Section) []Range {
	splitPoints := getSplitPoints(section)
	rs = splitRanges(rs, splitPoints)
	results := []Range{}

	for _, r := range rs {
		results = append(results, mapRange(r, section))
	}

	return results
}

func getLowestLocationNumber(rs []Range) int {
	lowestNumber := math.MaxInt
	for _, r := range rs {
		lowestNumber = min(r.Start, lowestNumber)
	}

	return lowestNumber
}

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	sections := parseSections(lines)

	// Uncomment accordingly depending on which part of the puzzle you want to solve
	// rs := getSeedRangesForPuzzlePart1(getSeedNumbers(lines))
	rs := getSeedRangesForPuzzlePart2(getSeedNumbers(lines))

	for _, section := range sections {
		rs = transformRanges(rs, section)
	}

	fmt.Println(getLowestLocationNumber(rs))
}
