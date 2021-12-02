package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func partOne() {
	count := 0
	prev := math.MaxInt64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if num > prev {
			count++
		}
		prev = num
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(count)
}

func partTwo() {
	count := 0
	i, j, k := math.MaxInt64, math.MaxInt64, math.MaxInt64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if num > i {
			count++
		}
		i, j, k = j, k, num
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(count)
}

func main() {
	partOne()
	partTwo()
}
