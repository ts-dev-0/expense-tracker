package main

import (
	"expense-tracker/internal/model"
	"expense-tracker/internal/repository"
	"fmt"
)

func readAppData() (model.Data, bool) {
	data, err := repository.ReadData()
	if err != nil {
		fmt.Println("エラー:", err)
		return model.Data{}, false
	}

	return data, true
}
