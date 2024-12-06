package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContents, _ := os.ReadFile("/Users/timothy/GolandProjects/aoc24/day3/input.txt")
	input := string(fileContents)

	allParts := []string{strings.Clone(input)}
	doParts := make([]string, 0)
	dontParts := strings.Split(input, "don't()")
	for i, part := range dontParts {
		if i == 0 {
			doParts = append(doParts, part)
			continue
		}

		_, after, found := strings.Cut(part, "do()")
		if !found {
			continue
		}

		doParts = append(doParts, after)
	}

	sum1 := 0
	sum2 := 0
	for _, allPart := range allParts {
		mulParts := strings.Split(allPart, "mul(")
		for _, part := range mulParts {
			subMulParts := strings.Split(part, ")")
			if isValid(subMulParts[0]) {
				numberStrings := strings.Split(subMulParts[0], ",")
				n1, _ := strconv.Atoi(numberStrings[0])
				n2, _ := strconv.Atoi(numberStrings[1])
				sum1 += n1 * n2
			}
		}
	}

	for _, doPart := range doParts {
		mulParts := strings.Split(doPart, "mul(")
		for _, part := range mulParts {
			subMulParts := strings.Split(part, ")")
			if isValid(subMulParts[0]) {
				numberStrings := strings.Split(subMulParts[0], ",")
				n1, _ := strconv.Atoi(numberStrings[0])
				n2, _ := strconv.Atoi(numberStrings[1])
				sum2 += n1 * n2
			}
		}
	}

	fmt.Printf("part 1: %d\n", sum1)
	fmt.Printf("part 2: %d\n", sum2)
}

func isValid(value string) bool {
	n1Length := 0
	n2Length := 0
	foundComma := false
	for _, char := range value {
		if char >= '0' && char <= '9' {
			if foundComma {
				n2Length++
			} else {
				n1Length++
			}

			continue
		}

		if char == ',' {
			foundComma = true
			continue
		}

		return false
	}

	if n1Length > 0 && n1Length <= 3 && n2Length > 0 && n2Length <= 3 {
		return true
	}

	return false
}
