package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"aoc24/library/file_utils"
)

func main() {
	lines, _ := file_utils.CountLines("/Users/timothy/GolandProjects/aoc24/day1/input.txt")
	left := make([]int, 0, lines)
	right := make([]int, 0, lines)

	_ = file_utils.ForEachLineDo("/Users/timothy/GolandProjects/aoc24/day1/input.txt", func(line string) {
		parts := strings.Split(line, "   ")
		leftInt, _ := strconv.Atoi(parts[0])
		rightInt, _ := strconv.Atoi(parts[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	})

	slices.Sort(left)
	slices.Sort(right)

	sum1 := 0
	for i := range left {
		sum1 += int(math.Abs(float64(left[i] - right[i])))
	}

	sum2 := 0
	for i := range left {
		value := left[i]
		rightOccurrences := 0
		for _, rightValue := range right {
			if value == rightValue {
				rightOccurrences++
			}
		}

		sum2 += value * rightOccurrences
	}

	fmt.Printf("part 1: %d\n", sum1)
	fmt.Printf("part 2: %d\n", sum2)
}
