package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const size = 5

func partOne(nums []string, boards [][][]int, bingo [][][]int) {
	bingoBoard, bingoNum := drawBingo(nums, boards, bingo)

	bingoSum := 0
	for i := range bingo[bingoBoard] {
		for j := range bingo[bingoBoard][i] {
			if bingo[bingoBoard][i][j] == 0 {
				bingoSum += boards[bingoBoard][i][j]
			}
		}
	}

	fmt.Println(bingoSum * bingoNum)
}

func partTwo(nums []string, boards [][][]int, bingo [][][]int) {
	bingoBoard, bingoNum := drawBingoUntilLast(nums, boards, bingo)

	bingoSum := 0
	for i := range bingo[bingoBoard] {
		for j := range bingo[bingoBoard][i] {
			if bingo[bingoBoard][i][j] == 0 {
				bingoSum += boards[bingoBoard][i][j]
			}
		}
	}

	fmt.Println(bingoSum * bingoNum)
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(bytes), "\n\n")
	nums := strings.Split(inputs[0], ",")
	inputBoards := append([]string{}, inputs[1:]...)

	boards := make([][][]int, len(inputBoards))

	for index, inputBoard := range inputBoards {
		b := [][]int{}
		for _, row := range strings.Split(inputBoard, "\n") {
			r := []int{}
			for _, item := range strings.Fields(row) {
				num, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}
				r = append(r, num)
			}
			b = append(b, r)
		}
		boards[index] = b
	}

	bingo := make([][][]int, len(boards))

	for i := range bingo {
		board := make([][]int, size)
		for j := range board {
			board[j] = make([]int, size)
		}
		bingo[i] = board
	}

	// partOne(nums, boards, bingo)
	partTwo(nums, boards, bingo)
}

func drawBingo(nums []string, boards [][][]int, bingo [][][]int) (int, int) {
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		for n, board := range boards {
			for i, row := range board {
				for j, item := range row {
					if item == num {
						bingo[n][i][j]++
					}
				}
			}

			for i := 0; i < size; i++ {
				rowCount := 0
				for j := 0; j < size; j++ {
					rowCount += bingo[n][i][j]
				}

				if rowCount == size {
					return n, num
				}

				colCount := 0
				for k := 0; k < size; k++ {
					colCount += bingo[n][k][i]
				}

				if colCount == size {
					return n, num
				}
			}
		}
	}
	return 0, 0
}

func drawBingoUntilLast(nums []string, boards [][][]int, bingo [][][]int) (int, int) {
	boardFlags := make([]bool, len(boards))

	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		for n, board := range boards {
			for i, row := range board {
				for j, item := range row {
					if item == num {
						bingo[n][i][j]++
					}
				}
			}

			for i := 0; i < size; i++ {
				rowCount := 0
				for j := 0; j < size; j++ {
					rowCount += bingo[n][i][j]
				}

				if rowCount == size {
					boardFlags[n] = true
					flag := true
					for _, boardFlag := range boardFlags {
						flag = flag && boardFlag
					}
					if flag {
						return n, num
					}
				}

				colCount := 0
				for k := 0; k < size; k++ {
					colCount += bingo[n][k][i]
				}

				if colCount == size {
					boardFlags[n] = true
					flag := true
					for _, boardFlag := range boardFlags {
						flag = flag && boardFlag
					}
					if flag {
						return n, num
					}
				}
			}
		}
	}
	return 0, 0
}
