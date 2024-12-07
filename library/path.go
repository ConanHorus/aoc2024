package library

import "fmt"

const pathStart = "/Users/timothy/src/github/aoc2024/day"

func GetTestFilePath(day int) string {
	return fmt.Sprintf("%s%d/test.txt", pathStart, day)
}

func GetInputFilePath(day int) string {
	return fmt.Sprintf("%s%d/input.txt", pathStart, day)
}
