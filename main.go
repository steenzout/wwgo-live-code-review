package main

import (
	"fmt"
)

func ElapsedDays(arg1, arg2, arg3 int) {
	elapsed := 0

	if arg1 == 1 {
		elapsed = arg2
	} else if arg1 == 2 {
		elapsed = 31 + arg2
	} else if arg1 == 3 {
		elapsed = 31 + 28 + arg2
	} else if arg1 == 4 {
		elapsed = 31 + 28 + 31 + arg2
	} else if arg1 == 5 {
		elapsed = 31 + 28 + 31 + 30 + arg2
	} else if arg1 == 6 {
		elapsed = 31 + 28 + 31 + 30 + 31 + arg2
	} else if arg1 == 7 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + arg2
	} else if arg1 == 8 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + arg2
	} else if arg1 == 9 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + arg2
	} else if arg1 == 10 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + arg2
	} else if arg1 == 11 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + arg2
	} else if arg1 == 12 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + arg2
	}

	fmt.Println(elapsed)
}

func main() {
	ElapsedDays(1, 1, 2001)
}
