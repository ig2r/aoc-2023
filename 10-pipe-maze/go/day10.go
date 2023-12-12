package main

import (
	"bytes"
	"fmt"
	"os"
)

type Maze = [][]byte

func findStartPos(maze Maze) (int, int) {
	width, height := len(maze[0]), len(maze)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if maze[i][j] == 'S' {
				return i, j
			}
		}
	}

	panic("No starting point found.")
}

const (
	North = 1 << iota
	South = 1 << iota
	East  = 1 << iota
	West  = 1 << iota
)

var connections = map[byte]int{
	'|': North | South,
	'-': East | West,
	'L': North | East,
	'J': North | West,
	'7': South | West,
	'F': South | East,
	'.': 0,
}

func main() {
	data, _ := os.ReadFile("sample.txt")
	maze := bytes.Split(data, []byte("\n"))

	i_0, j_0 := findStartPos(maze)
	fmt.Println(i_0, j_0, connections['-'])

	// TODO: Implementation pending
}
