package main

import (
	"bufio"
	"expense-tracker/internal/model"
	"expense-tracker/internal/repository"
	"expense-tracker/internal/service"
	"fmt"
	"strconv"
	"time"
)

const dateLayout = "2006-01-02"

type transactionInput struct {
	transactionType model.TransactionType
	amount          int64
	category        string
	memo            string
	date            string
}

func handleAddTransaction(scanner *bufio.Scanner) bool {
	input, ok := readTransactionInput(scanner)
	if !ok {
		return true
	}

	data, ok := readAppData()
	if !ok {
		return true
	}

	updatedData, transaction, err := service.AddTransaction(
		data,
		input.transactionType,
		input.amount,
		input.category,
		input.memo,
		input.date,
	)
	if err != nil {
		fmt.Println("エラー:", err)
		return true
	}

	if err := repository.SaveData(updatedData); err != nil {
		fmt.Println("エラー:", err)
		return true
	}

	fmt.Printf("取引を登録しました: ID=%d\n", transaction.ID)
	return true
}

func readTransactionInput(scanner *bufio.Scanner) (transactionInput, bool) {
	fmt.Println("-------------------------------------------")
	fmt.Println("取引内容を入力してください")
	fmt.Println("-------------------------------------------")

	transactionType, ok := readLine(scanner, "取引種別（income: 収入 / expense: 支出）: ")
	if !ok {
		return transactionInput{}, false
	}

	amountText, ok := readLine(scanner, "金額（円）: ")
	if !ok {
		return transactionInput{}, false
	}
	amount, err := strconv.ParseInt(amountText, 10, 64)
	if err != nil {
		fmt.Println("エラー: 金額は整数で入力してください")
		return transactionInput{}, false
	}

	category, ok := readLine(scanner, "カテゴリ: ")
	if !ok {
		return transactionInput{}, false
	}
	memo, ok := readLine(scanner, "メモ（任意）: ")
	if !ok {
		return transactionInput{}, false
	}
	date, ok := readLine(scanner, "日付（任意、YYYY-MM-DD。未入力の場合は当日）: ")
	if !ok {
		return transactionInput{}, false
	}

	if date == "" {
		date = time.Now().Format(dateLayout)
	} else if _, err := time.Parse(dateLayout, date); err != nil {
		fmt.Println("エラー: 日付はYYYY-MM-DD形式で指定してください")
		return transactionInput{}, false
	}

	return transactionInput{
		transactionType: model.TransactionType(transactionType),
		amount:          amount,
		category:        category,
		memo:            memo,
		date:            date,
	}, true
}
