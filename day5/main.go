package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rudeigerc/advent2021/utils"
)

func partOne() {
	n := 1000
	diagram := make([][]int, n)
	for i := range diagram {
		diagram[i] = make([]int, n)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		var (
			x1, y1, x2, y2 int
		)
		fmt.Sscanf(text, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if x1 == x2 {
			min, max := utils.MinMax(y1, y2)
			for y := min; y <= max; y++ {
				diagram[y][x1]++
			}
		} else if y1 == y2 {
			min, max := utils.MinMax(x1, x2)
			for x := min; x <= max; x++ {
				diagram[y1][x]++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	count := 0
	for y := range diagram {
		for x := range diagram[y] {
			if diagram[y][x] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func partTwo() {
	n := 1000
	diagram := make([][]int, n)
	for i := range diagram {
		diagram[i] = make([]int, n)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		var (
			x1, y1, x2, y2 int
		)
		fmt.Sscanf(text, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if x1 == x2 {
			min, max := utils.MinMax(y1, y2)
			for y := min; y <= max; y++ {
				diagram[y][x1]++
			}
		} else if y1 == y2 {
			min, max := utils.MinMax(x1, x2)
			for x := min; x <= max; x++ {
				diagram[y1][x]++
			}
		} else if utils.Abs(x1-x2) == utils.Abs(y1-y2) {
			minX, maxX := x1, x2
			minY, maxY := y1, y2

			if x1 > x2 {
				minX, maxX = x2, x1
				minY, maxY = y2, y1
			}

			direction := utils.Sgn((x1 - x2) * (y1 - y2))
			if direction >= 0 {
				for x, y := minX, minY; x <= maxX && y <= maxY; x, y = x+1, y+1 {
					diagram[y][x]++
				}
			} else {
				for x, y := minX, minY; x <= maxX && y >= maxY; x, y = x+1, y-1 {
					diagram[y][x]++
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	count := 0
	for y := range diagram {
		for x := range diagram[y] {
			if diagram[y][x] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	partOne()
	partTwo()
}
