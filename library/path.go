package library

import "fmt"

const pathStart = "/Users/timothy/GolandProjects/aoc24/day"

func GetTestFilePath(day int) string {
	return fmt.Sprintf("%s%d/test.txt", pathStart, day)
}

func GetInputFilePath(day int) string {
	return fmt.Sprintf("%s%d/input.txt", pathStart, day)
}