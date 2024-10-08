package main

import (
	"encoding/json"
	"os"
)

const dataFile = "expenses.json"

func LoadExpenses() ExpenseTracker {
	tracker := ExpenseTracker{
		Expenses: []Expense{},
		NextID:   1,
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return tracker
		}
		return tracker
	}

	err = json.Unmarshal(data, &tracker)
	if err != nil {
		return tracker
	}

	return tracker
}

func SaveExpenses(tracker ExpenseTracker) error {
	data, err := json.MarshalIndent(tracker, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}