package utils

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinMax(x int, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sgn(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
