package service

import (
	"cmp"
	"errors"
	"expense-tracker/internal/model"
	"slices"
)

func AddTransaction(
	data model.Data,
	transactionType model.TransactionType,
	amount int64,
	category string,
	memo string,
	date string,
) (model.Data, model.Transaction, error) {
	if transactionType != model.Income && transactionType != model.Expense {
		return model.Data{}, model.Transaction{}, errors.New(
			"種別はincome またはexpense を指定してください",
		)
	}

	if amount < 1 {
		return model.Data{}, model.Transaction{}, errors.New(
			"金額は1以上で指定してください",
		)
	}

	if category == "" {
		return model.Data{}, model.Transaction{}, errors.New(
			"カテゴリを指定してください",
		)
	}

	transaction := model.Transaction{
		ID:       data.NextID,
		Type:     transactionType,
		Amount:   amount,
		Category: category,
		Memo:     memo,
		Date:     date,
	}

	data.Transactions = append(data.Transactions, transaction)
	data.NextID++

	return data, transaction, nil
}

func ListTransactions(data model.Data) []model.Transaction {
	transactions := make([]model.Transaction, len(data.Transactions))

	copy(transactions, data.Transactions)

	slices.SortFunc(transactions, func(a, b model.Transaction) int {
		if result := cmp.Compare(b.Date, a.Date); result != 0 {
			return result
		}
		return cmp.Compare(b.ID, a.ID)
	})

	return transactions
}

type Summary struct {
	IncomeTotal  int64
	ExpenseTotal int64
	Balance      int64
}

func CalculateSummary(data model.Data) Summary {
	var summary Summary
	for _, transaction := range data.Transactions {
		switch transaction.Type {
		case model.Income:
			summary.IncomeTotal += transaction.Amount
		case model.Expense:
			summary.ExpenseTotal += transaction.Amount
		}
	}

	summary.Balance = summary.IncomeTotal - summary.ExpenseTotal
	return summary
}
