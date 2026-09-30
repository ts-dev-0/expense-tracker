package main

import (
	"bufio"
	"expense-tracker/internal/model"
	"expense-tracker/internal/service"
	"fmt"
	"os"
	"text/tabwriter"
)

func handleListTransaction(_ *bufio.Scanner) bool {
	data, ok := readAppData()
	if !ok {
		return true
	}

	transactions := service.ListTransactions(data)
	if len(transactions) == 0 {
		fmt.Println("登録されている取引はありません")
		return true
	}

	fmt.Printf("\n取引履歴（%d件）\n", len(transactions))

	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\t日付\t種別\t金額\tカテゴリ\tメモ")
	fmt.Fprintln(writer, "--\t----------\t----\t--------\t--------\t----")
	for _, transaction := range transactions {
		fmt.Fprintf(
			writer,
			"%d\t%s\t%s\t%s\t%s\t%s\n",
			transaction.ID,
			transaction.Date,
			transactionTypeLabel(transaction.Type),
			formatAmount(transaction.Amount),
			transaction.Category,
			displayMemo(transaction.Memo),
		)
	}
	if err := writer.Flush(); err != nil {
		fmt.Println("履歴の表示中にエラーが発生しました:", err)
	}
	fmt.Println()

	return true
}

func transactionTypeLabel(transactionType model.TransactionType) string {
	switch transactionType {
	case model.Income:
		return "収入"
	case model.Expense:
		return "支出"
	default:
		return string(transactionType)
	}
}
