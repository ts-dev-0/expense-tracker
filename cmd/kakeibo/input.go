package main

import (
	"bufio"
	"fmt"
	"strings"
)

func readLine(scanner *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		reportScannerError(scanner)
		return "", false
	}

	return strings.TrimSpace(scanner.Text()), true
}

func reportScannerError(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		fmt.Println("入力の読み込み中にエラーが発生しました:", err)
	}
}
