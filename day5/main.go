package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc24/library"
	"aoc24/library/file_utils"
)

var (
	instructions        = make(map[int][]*Entry)
	valueToInstructions = make(map[int][]int)
)

type Entry struct {
	Value int
	Seen  bool
}

func main() {
	updates := make([][]int, 0)

	secondHalf := false
	file_utils.ForEachLineDo(library.GetInputFilePath(5), func(line string) {
		if line == "" {
			secondHalf = true
			return
		}

		if secondHalf {
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))
			for i, part := range parts {
				update[i], _ = strconv.Atoi(part)
			}

			updates = append(updates, update)
			return
		}

		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		bucket, _ := instructions[before]
		bucket = append(bucket, &Entry{Value: after, Seen: false})
		instructions[before] = bucket

		befores, _ := valueToInstructions[after]
		befores = append(befores, before)
		valueToInstructions[after] = befores
	})

	count1 := 0
	count2 := 0

	for _, update := range updates {
		v := getMiddlePart1(update)
		count1 += v

		if v == 0 {
			count2 += getMiddlePart2(update)
		}
	}

	fmt.Printf("part 1: %d\n", count1)
	fmt.Printf("part 2: %d\n", count2)
}

func getMiddlePart1(update []int) int {
	if ok, _ := isGood(update); ok {
		return update[len(update)/2]
	}

	return 0
}

func getMiddlePart2(update []int) int {
	buffer := make([]int, len(update))
	copy(buffer, update)

	for {
		ok, badAt := isGood(buffer)
		if ok {
			break
		}

		buffer[badAt-1], buffer[badAt] = buffer[badAt], buffer[badAt-1]
	}

	return buffer[len(buffer)/2]
}

func isGood(update []int) (good bool, badAt int) {
	defer func() {
		for _, bucket := range instructions {
			for _, entry := range bucket {
				entry.Seen = false
			}
		}
	}()

	for iValue, value := range update {
		befores := valueToInstructions[value]
		for _, before := range befores {
			bucket := instructions[before]
			for _, entry := range bucket {
				if entry.Value == value {
					entry.Seen = true
				}
			}
		}

		bucket := instructions[value]
		for _, entry := range bucket {
			if entry.Seen {
				return false, iValue
			}
		}
	}

	return true, -1
}
