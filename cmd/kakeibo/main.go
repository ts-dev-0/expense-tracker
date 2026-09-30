package main

import (
	"bufio"
	"os"
)

func main() {
	runMenu(bufio.NewScanner(os.Stdin))
}
