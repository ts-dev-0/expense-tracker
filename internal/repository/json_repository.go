package repository

import (
	"encoding/json"
	"errors"
	"expense-tracker/internal/model"
	"fmt"
	"os"
)

const dataFilePath = "data/kakeibo.json"

func ReadData() (model.Data, error) {
	contents, err := os.ReadFile(dataFilePath)
	if errors.Is(err, os.ErrNotExist) {
		return model.Data{
			Version:      1,
			NextID:       1,
			Transactions: []model.Transaction{},
		}, nil
	}
	if err != nil {
		return model.Data{}, fmt.Errorf("データファイルの読み込みに失敗しました: %w", err)
	}

	var data model.Data
	if err := json.Unmarshal(contents, &data); err != nil {
		return model.Data{}, fmt.Errorf("JSONの読み込みに失敗しました: %w", err)
	}

	return data, nil
}

func SaveData(data model.Data) error {
	if err := os.MkdirAll("data", 0755); err != nil {
		return fmt.Errorf("データディレクトリの作成に失敗しました: %w", err)
	}

	contents, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("JSONへの変換に失敗しました: %w", err)
	}

	if err := os.WriteFile(dataFilePath, contents, 0644); err != nil {
		return fmt.Errorf("データファイルの保存に失敗しました: %w", err)
	}

	return nil
}
