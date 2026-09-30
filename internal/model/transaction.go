package model

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type Transaction struct {
	ID       int             `json:"id"`
	Type     TransactionType `json:"type"`
	Amount   int64           `json:"amount"`
	Category string          `json:"category"`
	Memo     string          `json:"memo"`
	Date     string          `json:"date"`
}

type Data struct {
	Version      int           `json:"version"`
	NextID       int           `json:"next_id"`
	Transactions []Transaction `json:"transactions"`
}
