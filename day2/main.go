package main

import (
	"bufio"
	"fmt"
	"os"
)

type Command string

const (
	Forward Command = "forward"
	Down    Command = "down"
	Up      Command = "up"
)

func partOne() {
	horizontal, depth := 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		var (
			command string
			unit    int
		)
		fmt.Sscanln(text, &command, &unit)
		switch Command(command) {
		case Forward:
			horizontal += unit
		case Down:
			depth += unit
		case Up:
			depth -= unit
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(horizontal * depth)
}

func partTwo() {
	horizontal, depth, aim := 0, 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		var (
			command string
			unit    int
		)
		fmt.Sscanln(text, &command, &unit)
		switch Command(command) {
		case Forward:
			horizontal += unit
			depth += aim * unit
		case Down:
			aim += unit
		case Up:
			aim -= unit
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(horizontal * depth)
}

func main() {
	partOne()
	partTwo()
}
