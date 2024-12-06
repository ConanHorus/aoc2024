package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc24/library"
	"aoc24/library/file_utils"
)

type reportObj struct {
	values []int
	safe1  *library.Lazy[bool]
	safe2  *library.Lazy[bool]
}

func newReport(line string) *reportObj {
	report := &reportObj{}
	report.safe1 = library.NewLazy[bool](func() bool { return isSafe(report.values) })
	report.safe2 = library.NewLazy[bool](func() bool {
		values := make([]int, 0, len(report.values))
		for i := 0; i < len(report.values); i++ {
			values = values[:0]
			values = append(values, report.values[:i]...)
			values = append(values, report.values[i+1:]...)
			if isSafe(values) {
				return true
			}
		}

		return false
	})

	parts := strings.Split(line, " ")
	for _, part := range parts {
		v, _ := strconv.Atoi(part)
		report.values = append(report.values, v)
	}

	return report
}

func isSafe(values []int) bool {
	low := math.MaxInt
	high := math.MinInt

	for i := range values {
		if i == 0 {
			continue
		}

		current := values[i]
		previous := values[i-1]
		diff := current - previous
		low = min(low, diff)
		high = max(high, diff)
	}

	return (low >= 1 && high <= 3) || (low >= -3 && high <= -1)
}

func main() {
	reports := make([]*reportObj, 0)
	file_utils.ForEachLineDo("/Users/timothy/GolandProjects/aoc24/day2/input.txt", func(line string) {
		reports = append(reports, newReport(line))
	})

	safeCount := 0
	safeCount2 := 0
	for _, report := range reports {
		if report.safe1.Value() {
			safeCount++
		}

		if report.safe1.Value() || report.safe2.Value() {
			safeCount2++
		}
	}

	fmt.Printf("part 1: %d\n", safeCount)
	fmt.Printf("part 2: %d\n", safeCount2)
}
