package main

import (
	"fmt"

	"aoc24/library"
	"aoc24/library/file_utils"
)

type Point struct {
	X int
	Y int
}

type Guard struct {
	Position  Point
	Direction Point
}

func (this *Guard) TurnRight() {
	if this.Direction.X == 1 {
		this.Direction.X = 0
		this.Direction.Y = 1
		return
	}

	if this.Direction.Y == 1 {
		this.Direction.X = -1
		this.Direction.Y = 0
		return
	}

	if this.Direction.X == -1 {
		this.Direction.X = 0
		this.Direction.Y = -1
		return
	}

	this.Direction.X = 1
	this.Direction.Y = 0
}

func (this *Guard) MoveForward() {
	this.Position.X += this.Direction.X
	this.Position.Y += this.Direction.Y
}

func (this *Guard) CanMoveForward(thisMap [][]byte) bool {
	next := Point{this.Position.X + this.Direction.X, this.Position.Y + this.Direction.Y}
	if next.X < 0 || next.Y < 0 || next.X >= len(thisMap[0]) || next.Y >= len(thisMap) {
		return true
	}

	valueAtNext := thisMap[next.Y][next.X]
	return valueAtNext != '#'
}

func (this *Guard) InsideMap(theMap [][]byte) bool {
	return this.Position.X >= 0 && this.Position.Y >= 0 && this.Position.X < len(theMap[0]) && this.Position.Y < len(theMap)
}

func main() {
	number1 := 0
	number2 := 0

	guard := new(Guard)
	theMap := make([][]byte, 0)
	file_utils.ForEachLineDoIterator(library.GetInputFilePath(6), func(y int, line string) {
		theMap = append(theMap, []byte(line))

		for x, b := range line {
			if b == '^' {
				guard.Position = Point{X: x, Y: y}
				guard.Direction = Point{X: 0, Y: -1}
			}
		}
	})

	for guard.InsideMap(theMap) {
		for !guard.CanMoveForward(theMap) {
			guard.TurnRight()
		}

		valueHere := theMap[guard.Position.Y][guard.Position.X]
		if valueHere != 'X' {
			theMap[guard.Position.Y][guard.Position.X] = 'X'
			number1++
		}

		guard.MoveForward()
	}

	fmt.Printf("part 1: %d\n", number1)
	fmt.Printf("part 2: %d\n", number2)
}
