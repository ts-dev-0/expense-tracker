package main

import (
	"bufio"
	"fmt"
)

type menuHandler func(*bufio.Scanner) bool

type menuItem struct {
	key     string
	label   string
	handler menuHandler
}

func runMenu(scanner *bufio.Scanner) {
	items := []menuItem{
		{key: "1", label: "取引を記録する", handler: handleAddTransaction},
		{key: "2", label: "取引の履歴を確認する", handler: handleListTransaction},
		{key: "3", label: "収入・支出・収支を確認する", handler: handleShowSummary},
		{key: "4", label: "アプリを終了する", handler: handleExit},
	}

	for {
		showMenu(items)

		choice, ok := readLine(scanner, "> ")
		if !ok {
			return
		}

		item, ok := findMenuItem(items, choice)
		if !ok {
			fmt.Println("メニューにない番号です")
			continue
		}

		if !item.handler(scanner) {
			return
		}
	}
}

func showMenu(items []menuItem) {
	fmt.Println("Kakeibo App")
	for _, item := range items {
		fmt.Printf("%s. %s\n", item.key, item.label)
	}
}

func findMenuItem(items []menuItem, key string) (menuItem, bool) {
	for _, item := range items {
		if item.key == key {
			return item, true
		}
	}

	return menuItem{}, false
}

func handleExit(_ *bufio.Scanner) bool {
	fmt.Println("アプリを終了します")
	return false
}
