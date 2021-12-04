package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getOneCounts(bitmap [][]int) []int {
	oneCounts := make([]int, len(bitmap[0]))
	for i := range bitmap {
		for j := range bitmap[i] {
			if bitmap[i][j] == 1 {
				oneCounts[j]++
			}
		}
	}
	return oneCounts
}

func getRating(bitmap [][]int, commonBit int) int {
	reduced := append([][]int{}, bitmap...)
	index := 0

	for len(reduced) > 1 {
		tmp := make([][]int, 0)
		threshold := len(reduced)
		oneCounts := getOneCounts(reduced)

		for _, bits := range reduced {
			if oneCounts[index]*2 >= threshold && bits[index] == commonBit || oneCounts[index]*2 < threshold && bits[index] == 1-commonBit {
				tmp = append(tmp, bits)
			}
		}
		reduced = append([][]int{}, tmp...)
		index++
	}

	rating := 0
	for _, bit := range reduced[0] {
		rating <<= 1
		rating |= bit
	}
	return rating
}

func partOne(bitmap [][]int) {
	oneCounts := getOneCounts(bitmap)
	gammaRate, epsilonRate := 0, 0
	threshold := len(bitmap) / 2

	for _, count := range oneCounts {
		gammaRate <<= 1
		epsilonRate <<= 1
		if count > threshold {
			gammaRate |= 1
		} else {
			epsilonRate |= 1
		}
	}

	fmt.Println(gammaRate * epsilonRate)
}

func partTwo(bitmap [][]int) {
	oxygen := getRating(bitmap, 1)
	co2 := getRating(bitmap, 0)
	fmt.Println(oxygen * co2)
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lines := strings.Fields(string(bytes))
	bitmap := make([][]int, len(lines))
	for i, line := range lines {
		bitmap[i] = make([]int, len(line))

		chs := strings.Split(line, "")
		for j, ch := range chs {
			bit, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}
			bitmap[i][j] = bit
		}
	}

	partOne(bitmap)
	partTwo(bitmap)
}
