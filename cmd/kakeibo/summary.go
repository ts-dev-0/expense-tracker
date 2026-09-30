package main

import (
	"bufio"
	"expense-tracker/internal/service"
	"fmt"
)

func handleShowSummary(_ *bufio.Scanner) bool {
	data, ok := readAppData()
	if !ok {
		return true
	}

	summary := service.CalculateSummary(data)
	fmt.Println("\n家計の集計")
	fmt.Printf("収入合計: %s\n", formatAmount(summary.IncomeTotal))
	fmt.Printf("支出合計: %s\n", formatAmount(summary.ExpenseTotal))
	fmt.Printf("収支:     %s\n\n", formatAmount(summary.Balance))

	return true
}
