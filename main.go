package main

import (
	"fmt"
)

// ElapsedDays prints the number of days since the 1st of January of that year.
func ElapsedDays(month, day, year int) {
	elapsed := 0

	if month == 1 {
		elapsed = day
	} else if month == 2 {
		elapsed = 31 + day
	} else if month == 3 {
		elapsed = 31 + 28 + day
	} else if month == 4 {
		elapsed = 31 + 28 + 31 + day
	} else if month == 5 {
		elapsed = 31 + 28 + 31 + 30 + day
	} else if month == 6 {
		elapsed = 31 + 28 + 31 + 30 + 31 + day
	} else if month == 7 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + day
	} else if month == 8 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + day
	} else if month == 9 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + day
	} else if month == 10 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + day
	} else if month == 11 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + day
	} else if month == 12 {
		elapsed = 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + day
	}

	fmt.Println(elapsed)
}

func main() {
	ElapsedDays(1, 1, 2001)
}
