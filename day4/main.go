package main

import (
	"fmt"

	"aoc24/library"
	"aoc24/library/file_utils"
)

func main() {
	lines, _ := file_utils.LoadAllLines(library.GetInputFilePath(4))

	count1 := 0
	count2 := 0
	for y, line := range lines {
		for x, v := range line {
			if v == 'X' {
				count1 += countUpHere(lines, x, y)
			}

			if v == 'A' {
				count2 += countXHere(lines, x, y)
			}
		}
	}

	fmt.Printf("part 1: %d\n", count1)
	fmt.Printf("part 2: %d\n", count2)
}

func countUpHere(lines []string, x int, y int) int {
	count := 0

	count += explore(lines, x, y, 1, 0)
	count += explore(lines, x, y, 1, 1)
	count += explore(lines, x, y, 0, 1)
	count += explore(lines, x, y, -1, 1)
	count += explore(lines, x, y, -1, 0)
	count += explore(lines, x, y, -1, -1)
	count += explore(lines, x, y, 0, -1)
	count += explore(lines, x, y, 1, -1)

	return count
}

func countXHere(lines []string, x int, y int) int {
	count := 0

	count += exploreX(lines, x, y, 1, 1)
	count += exploreX(lines, x, y, 1, -1)
	count += exploreX(lines, x, y, -1, -1)
	count += exploreX(lines, x, y, -1, 1)

	if count == 2 {
		return 1
	}

	return 0
}

func explore(lines []string, x int, y int, horizontal int, vertical int) int {
	requiredText := [4]byte{'X', 'M', 'A', 'S'}

	for i := 0; i < len(requiredText); i++ {
		xx := x + (i * horizontal)
		yy := y + (i * vertical)
		if xx < 0 || xx >= len(lines[0]) || yy < 0 || yy >= len(lines) {
			return 0
		}

		if lines[yy][xx] != requiredText[i] {
			return 0
		}
	}

	return 1
}

func exploreX(lines []string, x int, y int, horizontal int, vertical int) int {
	requiredText := [3]byte{'A', 'M', 'S'}
	multiplier := [3]int{0, 1, -1}

	for i := 0; i < len(requiredText); i++ {
		xx := x + (multiplier[i] * horizontal)
		yy := y + (multiplier[i] * vertical)
		if xx < 0 || xx >= len(lines[0]) || yy < 0 || yy >= len(lines) {
			return 0
		}

		if lines[yy][xx] != requiredText[i] {
			return 0
		}
	}

	return 1
}
