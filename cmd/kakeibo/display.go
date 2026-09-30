package main

import (
	"strconv"
	"strings"
)

func formatAmount(amount int64) string {
	number := strconv.FormatInt(amount, 10)
	start := 0
	if strings.HasPrefix(number, "-") {
		start = 1
	}

	for index := len(number) - 3; index > start; index -= 3 {
		number = number[:index] + "," + number[index:]
	}

	return number + "円"
}

func displayMemo(memo string) string {
	if memo == "" {
		return "-"
	}

	return memo
}
