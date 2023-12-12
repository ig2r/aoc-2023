package main

import (
	"bytes"
	"fmt"
	"os"
)

type Maze = [][]byte

type Coords = struct {
	i int
	j int
}

const (
	North uint8 = 1 << iota
	South       = 1 << iota
	East        = 1 << iota
	West        = 1 << iota
)

var connections = map[byte]uint8{
	'|': North | South,
	'-': East | West,
	'L': North | East,
	'J': North | West,
	'7': South | West,
	'F': South | East,
	'.': 0,
}

func getDimensions(maze Maze) (width int, height int) {
	return len(maze[0]), len(maze)
}

func findStartPos(maze Maze) (int, int) {
	width, height := getDimensions(maze)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if maze[i][j] == 'S' {
				return i, j
			}
		}
	}

	panic("No starting point in maze.")
}

func getStartSymbol(maze Maze, i_0 int, j_0 int) byte {
	width, height := getDimensions(maze)

	startNeighbors := uint8(0)

	if i_0 > 0 && connections[maze[i_0-1][j_0]]&South == South {
		startNeighbors |= North
	}

	if j_0 > 0 && connections[maze[i_0][j_0-1]]&East == East {
		startNeighbors |= West
	}

	if i_0 < height-1 && connections[maze[i_0+1][j_0]]&North == North {
		startNeighbors |= South
	}

	if j_0 < width-1 && connections[maze[i_0][j_0+1]]&West == West {
		startNeighbors |= East
	}

	for k, v := range connections {
		if startNeighbors == v {
			return k
		}
	}

	panic("Unknown start symbol.")
}

// step moves by a single cell into the direction indicated by dir.
// Returns the new coordinates i', j' and the opposite of the direction travelled.
func step(i int, j int, dir uint8) (i_prime int, j_prime int, revDir uint8) {
	switch {
	case dir&North == North:
		return i - 1, j, South
	case dir&South == South:
		return i + 1, j, North
	case dir&West == West:
		return i, j - 1, East
	case dir&East == East:
		return i, j + 1, West
	default:
		panic("Invalid direction.")
	}
}

func traceLoop(maze Maze, i_0 int, j_0 int) []Coords {
	path := make([]Coords, 0)

	i, j := i_0, j_0
	revDir := uint8(0)

	for {
		path = append(path, Coords{i, j})

		dir := connections[maze[i][j]] & ^revDir
		i, j, revDir = step(i, j, dir)

		if i == i_0 && j == j_0 {
			return path
		}
	}
}

func main() {
	data, _ := os.ReadFile("sample.txt")
	maze := bytes.Split(data, []byte("\n"))

	// Find start position and actual pipe symbol at that position
	i_0, j_0 := findStartPos(maze)
	s := getStartSymbol(maze, i_0, j_0)
	maze[i_0][j_0] = s

	// Visit all points along the loop
	path := traceLoop(maze, i_0, j_0)

	// Answer part 1
	farthestPoint := len(path) / 2
	fmt.Printf("Farthest point is %d steps away from start\n", farthestPoint)

	// For part 2, mark cells occupied by loop in "visited" grid
	width, height := getDimensions(maze)
	visited := make([]byte, width*height)
	for _, c := range path {
		visited[c.i*width+c.j] = 1
	}

	// Now run Even-Odd Rule across maze, switching inside/outside every time
	// we cross a pipe segment that connects north (so we'll trace slightly "above"
	// the actual cells, ignoring any pipe segments that run east-west or never cross
	// that northern threshold).
	count := 0
	for i := 0; i < height; i++ {
		isInside := false
		for j := 0; j < width; j++ {
			if visited[i*width+j] > 0 {
				if connections[maze[i][j]]&North == North {
					isInside = !isInside
				}
			} else if isInside {
				count++
			}
		}
	}

	fmt.Printf("Cells inside loop: %d", count)
}
